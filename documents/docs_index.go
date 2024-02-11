package documents

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pzierahn/chatbot_services/account"
	"github.com/pzierahn/chatbot_services/llm"
	"github.com/pzierahn/chatbot_services/pdf"
	pb "github.com/pzierahn/chatbot_services/proto"
	"github.com/pzierahn/chatbot_services/vectordb"
	"io"
	"strings"
	"time"
)

type embeddingsBatch struct {
	userId string
	pages  []string
	stream pb.DocumentService_IndexServer
}

type document struct {
	id           string
	userId       string
	collectionId string
	filename     string
	path         string
	embeddings   []*embedding
}

type embedding struct {
	Page      uint32
	Text      string
	Embedding []float32
	Tokens    uint32
	Error     error
}

func (service *Service) getDocPages(ctx context.Context, path string) ([]string, error) {

	obj := service.storage.Object(path)
	read, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = read.Close() }()

	raw, err := io.ReadAll(read)
	if err != nil {
		return nil, err
	}

	return pdf.GetPagesFromBytes(ctx, raw)
}

func (service *Service) processEmbeddings(ctx context.Context, batch *embeddingsBatch) ([]*embedding, error) {
	totalPages := len(batch.pages)

	_ = batch.stream.Send(&pb.IndexProgress{
		TotalPages:     uint32(totalPages),
		ProcessedPages: 0,
	})

	var embeddings []*embedding
	var processed uint32

	results := make(chan *embedding, 1)
	defer close(results)

	queue := make(chan int, totalPages)
	for page := 0; page < totalPages; page++ {
		queue <- page
	}
	defer close(queue)

	for agent := 0; agent < 50; agent++ {
		go func(agent int) {
			for {
				select {
				case page, ok := <-queue:
					if !ok {
						return
					}

					text := strings.TrimSpace(batch.pages[page])
					if len(text) <= 0 {
						results <- &embedding{
							Page:  uint32(page),
							Error: nil,
						}
						continue
					}

					// TODO: Pool embeddings tracking
					ctx, cnl := context.WithTimeout(ctx, time.Second*5)
					resp, err := service.embeddings.CreateEmbeddings(ctx, &llm.EmbeddingRequest{
						Input:  text,
						UserId: batch.userId,
					})
					cnl()

					if err != nil {
						results <- &embedding{
							Page:  uint32(page),
							Error: err,
						}

						break
					} else {
						results <- &embedding{
							Page:      uint32(page),
							Text:      text,
							Embedding: resp.Data,
							Tokens:    uint32(resp.Tokens),
							Error:     err,
						}
					}

				case <-ctx.Done():
					return
				}
			}
		}(agent)
	}

	var errorCount int
	for result := range results {
		if result.Error != nil {
			if errorCount > 100 {
				return nil, result.Error
			} else {
				queue <- int(result.Page)
				errorCount++
				continue
			}
		}

		if result.Embedding != nil {
			embeddings = append(embeddings, result)
		}

		_ = batch.stream.Send(&pb.IndexProgress{
			TotalPages:     uint32(totalPages),
			ProcessedPages: processed + 1,
		})
		processed++

		if processed >= uint32(totalPages) {
			break
		}
	}

	return embeddings, nil
}

func (service *Service) insertEmbeddings(ctx context.Context, doc *document) error {
	tx, err := service.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()
	_, err = tx.Exec(
		ctx,
		`insert into documents (id, user_id, filename, path, collection_id)
			values ($1, $2, $3, $4, $5)`,
		doc.id,
		doc.userId,
		doc.filename,
		doc.path,
		doc.collectionId)
	if err != nil {
		return err
	}

	var vectors []*vectordb.Vector

	for _, fragment := range doc.embeddings {
		chunkId := uuid.NewString()

		_, err = tx.Exec(ctx,
			`insert into document_chunks (id, document_id, page, text)
				values ($1, $2, $3, $4)`,
			chunkId,
			doc.id,
			fragment.Page,
			fragment.Text,
		)
		if err != nil {
			return err
		}

		vectors = append(vectors, &vectordb.Vector{
			Id:           chunkId,
			DocumentId:   doc.id,
			CollectionId: doc.collectionId,
			UserId:       doc.userId,
			Filename:     doc.filename,
			Text:         fragment.Text,
			Page:         fragment.Page,
			Vector:       fragment.Embedding,
		})
	}

	err = service.vectorDB.Upsert(vectors)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (service *Service) Index(doc *pb.Document, stream pb.DocumentService_IndexServer) error {

	ctx := stream.Context()
	userId, err := service.auth.Verify(ctx)
	if err != nil {
		return err
	}

	funding, err := service.account.HasFunding(ctx)
	if err != nil {
		return err
	}

	if !funding {
		return account.NoFundingError()
	}

	pages, err := service.getDocPages(ctx, doc.Path)
	if err != nil {
		return err
	}

	obj := service.storage.Object(doc.Path)

	embeddings, err := service.processEmbeddings(ctx, &embeddingsBatch{
		userId: userId,
		pages:  pages,
		stream: stream,
	})
	if err != nil {
		_ = obj.Delete(ctx)
		return err
	}

	err = service.insertEmbeddings(ctx, &document{
		id:           doc.Id,
		userId:       userId,
		collectionId: doc.CollectionId,
		filename:     doc.Filename,
		path:         doc.Path,
		embeddings:   embeddings,
	})
	if err != nil {
		_ = obj.Delete(ctx)
		return err
	}

	return err
}

package chat

import (
	"context"
	pb "github.com/pzierahn/brainboost/proto"
	"sort"
	"strings"
)

type documentPages struct {
	id           string
	collectionId string
	userId       string
	pages        []uint32
}

type documentChunk struct {
	chunkId string
	page    uint32
	text    string
}

func (service *Service) getDocumentChunks(ctx context.Context, query documentPages) ([]documentChunk, error) {
	rows, err := service.db.Query(ctx,
		`SELECT chunk.id, chunk.page, chunk.text
		FROM document_chunks as chunk, documents as doc
		WHERE
		    document_id = $1 AND
		    doc.id = chunk.document_id AND
		    doc.collection_id = $2 AND
		    user_id = $3 AND
		    page = ANY($4)
		ORDER BY page`,
		query.id, query.collectionId, query.userId, query.pages)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chunks []documentChunk

	for rows.Next() {
		var chunk documentChunk

		err = rows.Scan(
			&chunk.chunkId,
			&chunk.page,
			&chunk.text,
		)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

func (service *Service) getDocumentsContext(ctx context.Context, userId string, prompt *pb.Prompt) ([]string, []string, error) {
	sort.Slice(prompt.Documents, func(i, j int) bool {
		return prompt.Documents[i].Id < prompt.Documents[j].Id
	})

	for _, doc := range prompt.Documents {
		sort.Slice(doc.Pages, func(i, j int) bool {
			return doc.Pages[i] < doc.Pages[j]
		})
	}

	var chunkIds []string
	var contextTexts []string

	for _, doc := range prompt.Documents {
		chunks, err := service.getDocumentChunks(ctx, documentPages{
			id:           doc.Id,
			collectionId: prompt.CollectionId,
			userId:       userId,
			pages:        doc.Pages,
		})
		if err != nil {
			return nil, nil, err
		}

		var texts []string
		for _, chunk := range chunks {
			chunkIds = append(chunkIds, chunk.chunkId)
			texts = append(texts, chunk.text)
		}

		contextTexts = append(contextTexts, strings.Join(texts, "\n"))
	}

	return chunkIds, contextTexts, nil
}

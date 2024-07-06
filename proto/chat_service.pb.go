// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: chat_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId   string        `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Prompt       string        `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	ModelOptions *ModelOptions `protobuf:"bytes,3,opt,name=model_options,json=modelOptions,proto3" json:"model_options,omitempty"`
}

func (x *CompletionRequest) Reset() {
	*x = CompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionRequest) ProtoMessage() {}

func (x *CompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionRequest.ProtoReflect.Descriptor instead.
func (*CompletionRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompletionRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *CompletionRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *CompletionRequest) GetModelOptions() *ModelOptions {
	if x != nil {
		return x.ModelOptions
	}
	return nil
}

type CompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completion string `protobuf:"bytes,1,opt,name=completion,proto3" json:"completion,omitempty"`
}

func (x *CompletionResponse) Reset() {
	*x = CompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionResponse) ProtoMessage() {}

func (x *CompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionResponse.ProtoReflect.Descriptor instead.
func (*CompletionResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompletionResponse) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

type BatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentIds  []string      `protobuf:"bytes,1,rep,name=document_ids,json=documentIds,proto3" json:"document_ids,omitempty"`
	Prompts      []string      `protobuf:"bytes,2,rep,name=prompts,proto3" json:"prompts,omitempty"`
	ModelOptions *ModelOptions `protobuf:"bytes,3,opt,name=model_options,json=modelOptions,proto3" json:"model_options,omitempty"`
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchRequest) GetDocumentIds() []string {
	if x != nil {
		return x.DocumentIds
	}
	return nil
}

func (x *BatchRequest) GetPrompts() []string {
	if x != nil {
		return x.Prompts
	}
	return nil
}

func (x *BatchRequest) GetModelOptions() *ModelOptions {
	if x != nil {
		return x.ModelOptions
	}
	return nil
}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentIds []string                    `protobuf:"bytes,1,rep,name=document_ids,json=documentIds,proto3" json:"document_ids,omitempty"`
	Prompts     []string                    `protobuf:"bytes,2,rep,name=prompts,proto3" json:"prompts,omitempty"`
	PromptTitle []string                    `protobuf:"bytes,3,rep,name=prompt_title,json=promptTitle,proto3" json:"prompt_title,omitempty"`
	Items       []*BatchResponse_Completion `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchResponse) GetDocumentIds() []string {
	if x != nil {
		return x.DocumentIds
	}
	return nil
}

func (x *BatchResponse) GetPrompts() []string {
	if x != nil {
		return x.Prompts
	}
	return nil
}

func (x *BatchResponse) GetPromptTitle() []string {
	if x != nil {
		return x.PromptTitle
	}
	return nil
}

func (x *BatchResponse) GetItems() []*BatchResponse_Completion {
	if x != nil {
		return x.Items
	}
	return nil
}

type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Thread ID to post the message to
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// Prompt to generate completion
	Prompt string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Model options
	ModelOptions *ModelOptions `protobuf:"bytes,3,opt,name=model_options,json=modelOptions,proto3" json:"model_options,omitempty"`
	// Search options
	Threshold float32 `protobuf:"fixed32,4,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Limit     uint32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// Attached documents
	DocumentIds []string `protobuf:"bytes,6,rep,name=document_ids,json=documentIds,proto3" json:"document_ids,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{4}
}

func (x *Prompt) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *Prompt) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *Prompt) GetModelOptions() *ModelOptions {
	if x != nil {
		return x.ModelOptions
	}
	return nil
}

func (x *Prompt) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Prompt) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Prompt) GetDocumentIds() []string {
	if x != nil {
		return x.DocumentIds
	}
	return nil
}

type ModelOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model       string  `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Temperature float32 `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	MaxTokens   uint32  `protobuf:"varint,3,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	TopP        float32 `protobuf:"fixed32,4,opt,name=top_p,json=topP,proto3" json:"top_p,omitempty"`
}

func (x *ModelOptions) Reset() {
	*x = ModelOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelOptions) ProtoMessage() {}

func (x *ModelOptions) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelOptions.ProtoReflect.Descriptor instead.
func (*ModelOptions) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{5}
}

func (x *ModelOptions) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ModelOptions) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ModelOptions) GetMaxTokens() uint32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *ModelOptions) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the message
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Prompt used to generate the message
	Prompt string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Generated completion
	Completion string `protobuf:"bytes,3,opt,name=completion,proto3" json:"completion,omitempty"`
	// ReferenceIDs and their scores
	References map[string]float32 `protobuf:"bytes,4,rep,name=references,proto3" json:"references,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{6}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *Message) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

func (x *Message) GetReferences() map[string]float32 {
	if x != nil {
		return x.References
	}
	return nil
}

type Thread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Thread) Reset() {
	*x = Thread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thread) ProtoMessage() {}

func (x *Thread) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thread.ProtoReflect.Descriptor instead.
func (*Thread) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{7}
}

func (x *Thread) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thread) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ThreadID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ThreadID) Reset() {
	*x = ThreadID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadID) ProtoMessage() {}

func (x *ThreadID) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadID.ProtoReflect.Descriptor instead.
func (*ThreadID) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{8}
}

func (x *ThreadID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MessageID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ThreadId string `protobuf:"bytes,2,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
}

func (x *MessageID) Reset() {
	*x = MessageID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageID) ProtoMessage() {}

func (x *MessageID) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageID.ProtoReflect.Descriptor instead.
func (*MessageID) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{9}
}

func (x *MessageID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageID) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

type ThreadIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ThreadIDs) Reset() {
	*x = ThreadIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadIDs) ProtoMessage() {}

func (x *ThreadIDs) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadIDs.ProtoReflect.Descriptor instead.
func (*ThreadIDs) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{10}
}

func (x *ThreadIDs) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BatchResponse_Completion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId    uint32 `protobuf:"varint,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	DocumentTitle string `protobuf:"bytes,2,opt,name=document_title,json=documentTitle,proto3" json:"document_title,omitempty"`
	Prompt        uint32 `protobuf:"varint,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Completion    string `protobuf:"bytes,4,opt,name=completion,proto3" json:"completion,omitempty"`
}

func (x *BatchResponse_Completion) Reset() {
	*x = BatchResponse_Completion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse_Completion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse_Completion) ProtoMessage() {}

func (x *BatchResponse_Completion) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse_Completion.ProtoReflect.Descriptor instead.
func (*BatchResponse_Completion) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *BatchResponse_Completion) GetDocumentId() uint32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *BatchResponse_Completion) GetDocumentTitle() string {
	if x != nil {
		return x.DocumentTitle
	}
	return ""
}

func (x *BatchResponse_Completion) GetPrompt() uint32 {
	if x != nil {
		return x.Prompt
	}
	return 0
}

func (x *BatchResponse_Completion) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

var File_chat_service_proto protoreflect.FileDescriptor

var file_chat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x35, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a,
	0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x42, 0x0a, 0x0d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x34, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbf, 0x02, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x8c, 0x01, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd8, 0x01, 0x0a, 0x06, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x42, 0x0a, 0x0d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x76, 0x35, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x22, 0x7a, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x5f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50,
	0x22, 0xda, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x3d,
	0x0a, 0x0f, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4e, 0x0a,
	0x06, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x1a, 0x0a,
	0x08, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x09, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x32, 0xca, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x35, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x35, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x4f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x49, 0x44, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x49, 0x44, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x35, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_chat_service_proto_rawDescOnce sync.Once
	file_chat_service_proto_rawDescData = file_chat_service_proto_rawDesc
)

func file_chat_service_proto_rawDescGZIP() []byte {
	file_chat_service_proto_rawDescOnce.Do(func() {
		file_chat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_service_proto_rawDescData)
	})
	return file_chat_service_proto_rawDescData
}

var file_chat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chat_service_proto_goTypes = []any{
	(*CompletionRequest)(nil),        // 0: chatbot.chat.v5.CompletionRequest
	(*CompletionResponse)(nil),       // 1: chatbot.chat.v5.CompletionResponse
	(*BatchRequest)(nil),             // 2: chatbot.chat.v5.BatchRequest
	(*BatchResponse)(nil),            // 3: chatbot.chat.v5.BatchResponse
	(*Prompt)(nil),                   // 4: chatbot.chat.v5.Prompt
	(*ModelOptions)(nil),             // 5: chatbot.chat.v5.ModelOptions
	(*Message)(nil),                  // 6: chatbot.chat.v5.Message
	(*Thread)(nil),                   // 7: chatbot.chat.v5.Thread
	(*ThreadID)(nil),                 // 8: chatbot.chat.v5.ThreadID
	(*MessageID)(nil),                // 9: chatbot.chat.v5.MessageID
	(*ThreadIDs)(nil),                // 10: chatbot.chat.v5.ThreadIDs
	(*BatchResponse_Completion)(nil), // 11: chatbot.chat.v5.BatchResponse.Completion
	nil,                              // 12: chatbot.chat.v5.Message.ReferencesEntry
	(*Collection)(nil),               // 13: chatbot.collections.v2.Collection
	(*emptypb.Empty)(nil),            // 14: google.protobuf.Empty
}
var file_chat_service_proto_depIdxs = []int32{
	5,  // 0: chatbot.chat.v5.CompletionRequest.model_options:type_name -> chatbot.chat.v5.ModelOptions
	5,  // 1: chatbot.chat.v5.BatchRequest.model_options:type_name -> chatbot.chat.v5.ModelOptions
	11, // 2: chatbot.chat.v5.BatchResponse.items:type_name -> chatbot.chat.v5.BatchResponse.Completion
	5,  // 3: chatbot.chat.v5.Prompt.model_options:type_name -> chatbot.chat.v5.ModelOptions
	12, // 4: chatbot.chat.v5.Message.references:type_name -> chatbot.chat.v5.Message.ReferencesEntry
	6,  // 5: chatbot.chat.v5.Thread.messages:type_name -> chatbot.chat.v5.Message
	4,  // 6: chatbot.chat.v5.ChatService.PostMessage:input_type -> chatbot.chat.v5.Prompt
	8,  // 7: chatbot.chat.v5.ChatService.GetThread:input_type -> chatbot.chat.v5.ThreadID
	13, // 8: chatbot.chat.v5.ChatService.ListThreadIDs:input_type -> chatbot.collections.v2.Collection
	8,  // 9: chatbot.chat.v5.ChatService.DeleteThread:input_type -> chatbot.chat.v5.ThreadID
	9,  // 10: chatbot.chat.v5.ChatService.DeleteMessageFromThread:input_type -> chatbot.chat.v5.MessageID
	0,  // 11: chatbot.chat.v5.ChatService.Completion:input_type -> chatbot.chat.v5.CompletionRequest
	6,  // 12: chatbot.chat.v5.ChatService.PostMessage:output_type -> chatbot.chat.v5.Message
	7,  // 13: chatbot.chat.v5.ChatService.GetThread:output_type -> chatbot.chat.v5.Thread
	10, // 14: chatbot.chat.v5.ChatService.ListThreadIDs:output_type -> chatbot.chat.v5.ThreadIDs
	14, // 15: chatbot.chat.v5.ChatService.DeleteThread:output_type -> google.protobuf.Empty
	14, // 16: chatbot.chat.v5.ChatService.DeleteMessageFromThread:output_type -> google.protobuf.Empty
	1,  // 17: chatbot.chat.v5.ChatService.Completion:output_type -> chatbot.chat.v5.CompletionResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chat_service_proto_init() }
func file_chat_service_proto_init() {
	if File_chat_service_proto != nil {
		return
	}
	file_collection_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CompletionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompletionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Prompt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ModelOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Thread); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ThreadID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*MessageID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ThreadIDs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_service_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*BatchResponse_Completion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_service_proto_goTypes,
		DependencyIndexes: file_chat_service_proto_depIdxs,
		MessageInfos:      file_chat_service_proto_msgTypes,
	}.Build()
	File_chat_service_proto = out.File
	file_chat_service_proto_rawDesc = nil
	file_chat_service_proto_goTypes = nil
	file_chat_service_proto_depIdxs = nil
}

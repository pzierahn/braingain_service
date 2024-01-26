// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: chat_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatService_Chat_FullMethodName              = "/endpoint.brainboost.chat.v2.ChatService/Chat"
	ChatService_GetChatMessages_FullMethodName   = "/endpoint.brainboost.chat.v2.ChatService/GetChatMessages"
	ChatService_GetChatMessage_FullMethodName    = "/endpoint.brainboost.chat.v2.ChatService/GetChatMessage"
	ChatService_DeleteChatMessage_FullMethodName = "/endpoint.brainboost.chat.v2.ChatService/DeleteChatMessage"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Chat(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*ChatMessage, error)
	GetChatMessages(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*ChatMessages, error)
	GetChatMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*ChatMessage, error)
	DeleteChatMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*MessageID, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Chat(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*ChatMessage, error) {
	out := new(ChatMessage)
	err := c.cc.Invoke(ctx, ChatService_Chat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatMessages(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*ChatMessages, error) {
	out := new(ChatMessages)
	err := c.cc.Invoke(ctx, ChatService_GetChatMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*ChatMessage, error) {
	out := new(ChatMessage)
	err := c.cc.Invoke(ctx, ChatService_GetChatMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteChatMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*MessageID, error) {
	out := new(MessageID)
	err := c.cc.Invoke(ctx, ChatService_DeleteChatMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Chat(context.Context, *Prompt) (*ChatMessage, error)
	GetChatMessages(context.Context, *Collection) (*ChatMessages, error)
	GetChatMessage(context.Context, *MessageID) (*ChatMessage, error)
	DeleteChatMessage(context.Context, *MessageID) (*MessageID, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Chat(context.Context, *Prompt) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServiceServer) GetChatMessages(context.Context, *Collection) (*ChatMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMessages not implemented")
}
func (UnimplementedChatServiceServer) GetChatMessage(context.Context, *MessageID) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMessage not implemented")
}
func (UnimplementedChatServiceServer) DeleteChatMessage(context.Context, *MessageID) (*MessageID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Chat(ctx, req.(*Prompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatMessages(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatMessage(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteChatMessage(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.brainboost.chat.v2.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat",
			Handler:    _ChatService_Chat_Handler,
		},
		{
			MethodName: "GetChatMessages",
			Handler:    _ChatService_GetChatMessages_Handler,
		},
		{
			MethodName: "GetChatMessage",
			Handler:    _ChatService_GetChatMessage_Handler,
		},
		{
			MethodName: "DeleteChatMessage",
			Handler:    _ChatService_DeleteChatMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat_service.proto",
}
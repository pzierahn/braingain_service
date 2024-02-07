// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: collection_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CollectionService_GetAll_FullMethodName = "/endpoint.brainboost.collections.v1.CollectionService/GetAll"
	CollectionService_Create_FullMethodName = "/endpoint.brainboost.collections.v1.CollectionService/Create"
	CollectionService_Update_FullMethodName = "/endpoint.brainboost.collections.v1.CollectionService/Update"
	CollectionService_Delete_FullMethodName = "/endpoint.brainboost.collections.v1.CollectionService/Delete"
)

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionServiceClient interface {
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Collections, error)
	Create(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error)
	Update(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error)
	Delete(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Collections, error) {
	out := new(Collections)
	err := c.cc.Invoke(ctx, CollectionService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) Create(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, CollectionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) Update(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, CollectionService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) Delete(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CollectionService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServiceServer is the server API for CollectionService service.
// All implementations must embed UnimplementedCollectionServiceServer
// for forward compatibility
type CollectionServiceServer interface {
	GetAll(context.Context, *emptypb.Empty) (*Collections, error)
	Create(context.Context, *Collection) (*Collection, error)
	Update(context.Context, *Collection) (*Collection, error)
	Delete(context.Context, *Collection) (*emptypb.Empty, error)
	mustEmbedUnimplementedCollectionServiceServer()
}

// UnimplementedCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionServiceServer struct {
}

func (UnimplementedCollectionServiceServer) GetAll(context.Context, *emptypb.Empty) (*Collections, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCollectionServiceServer) Create(context.Context, *Collection) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCollectionServiceServer) Update(context.Context, *Collection) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCollectionServiceServer) Delete(context.Context, *Collection) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCollectionServiceServer) mustEmbedUnimplementedCollectionServiceServer() {}

// UnsafeCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServiceServer will
// result in compilation errors.
type UnsafeCollectionServiceServer interface {
	mustEmbedUnimplementedCollectionServiceServer()
}

func RegisterCollectionServiceServer(s grpc.ServiceRegistrar, srv CollectionServiceServer) {
	s.RegisterService(&CollectionService_ServiceDesc, srv)
}

func _CollectionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).Create(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).Update(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).Delete(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionService_ServiceDesc is the grpc.ServiceDesc for CollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.brainboost.collections.v1.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _CollectionService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CollectionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CollectionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CollectionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection_service.proto",
}

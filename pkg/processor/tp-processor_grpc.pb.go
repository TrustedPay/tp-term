// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: tp-processor.proto

package processor

import (
	context "context"
	proto "github.com/TrustedPay/tp-term/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TPProcessor_AuthorizeTransaction_FullMethodName = "/TPProcessor/AuthorizeTransaction"
)

// TPProcessorClient is the client API for TPProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TPProcessorClient interface {
	AuthorizeTransaction(ctx context.Context, in *proto.Transaction, opts ...grpc.CallOption) (*AuthorizeTransactionReply, error)
}

type tPProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewTPProcessorClient(cc grpc.ClientConnInterface) TPProcessorClient {
	return &tPProcessorClient{cc}
}

func (c *tPProcessorClient) AuthorizeTransaction(ctx context.Context, in *proto.Transaction, opts ...grpc.CallOption) (*AuthorizeTransactionReply, error) {
	out := new(AuthorizeTransactionReply)
	err := c.cc.Invoke(ctx, TPProcessor_AuthorizeTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TPProcessorServer is the server API for TPProcessor service.
// All implementations must embed UnimplementedTPProcessorServer
// for forward compatibility
type TPProcessorServer interface {
	AuthorizeTransaction(context.Context, *proto.Transaction) (*AuthorizeTransactionReply, error)
	mustEmbedUnimplementedTPProcessorServer()
}

// UnimplementedTPProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedTPProcessorServer struct {
}

func (UnimplementedTPProcessorServer) AuthorizeTransaction(context.Context, *proto.Transaction) (*AuthorizeTransactionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeTransaction not implemented")
}
func (UnimplementedTPProcessorServer) mustEmbedUnimplementedTPProcessorServer() {}

// UnsafeTPProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TPProcessorServer will
// result in compilation errors.
type UnsafeTPProcessorServer interface {
	mustEmbedUnimplementedTPProcessorServer()
}

func RegisterTPProcessorServer(s grpc.ServiceRegistrar, srv TPProcessorServer) {
	s.RegisterService(&TPProcessor_ServiceDesc, srv)
}

func _TPProcessor_AuthorizeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TPProcessorServer).AuthorizeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TPProcessor_AuthorizeTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TPProcessorServer).AuthorizeTransaction(ctx, req.(*proto.Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// TPProcessor_ServiceDesc is the grpc.ServiceDesc for TPProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TPProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TPProcessor",
	HandlerType: (*TPProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeTransaction",
			Handler:    _TPProcessor_AuthorizeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tp-processor.proto",
}
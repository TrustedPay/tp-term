// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: tp-term.proto

package tpterm

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
	TPTerm_SignRequest_FullMethodName = "/TPTerm/SignRequest"
)

// TPTermClient is the client API for TPTerm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TPTermClient interface {
	SignRequest(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionSignature, error)
}

type tPTermClient struct {
	cc grpc.ClientConnInterface
}

func NewTPTermClient(cc grpc.ClientConnInterface) TPTermClient {
	return &tPTermClient{cc}
}

func (c *tPTermClient) SignRequest(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionSignature, error) {
	out := new(TransactionSignature)
	err := c.cc.Invoke(ctx, TPTerm_SignRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TPTermServer is the server API for TPTerm service.
// All implementations must embed UnimplementedTPTermServer
// for forward compatibility
type TPTermServer interface {
	SignRequest(context.Context, *Transaction) (*TransactionSignature, error)
	mustEmbedUnimplementedTPTermServer()
}

// UnimplementedTPTermServer must be embedded to have forward compatible implementations.
type UnimplementedTPTermServer struct {
}

func (UnimplementedTPTermServer) SignRequest(context.Context, *Transaction) (*TransactionSignature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignRequest not implemented")
}
func (UnimplementedTPTermServer) mustEmbedUnimplementedTPTermServer() {}

// UnsafeTPTermServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TPTermServer will
// result in compilation errors.
type UnsafeTPTermServer interface {
	mustEmbedUnimplementedTPTermServer()
}

func RegisterTPTermServer(s grpc.ServiceRegistrar, srv TPTermServer) {
	s.RegisterService(&TPTerm_ServiceDesc, srv)
}

func _TPTerm_SignRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TPTermServer).SignRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TPTerm_SignRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TPTermServer).SignRequest(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// TPTerm_ServiceDesc is the grpc.ServiceDesc for TPTerm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TPTerm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TPTerm",
	HandlerType: (*TPTermServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignRequest",
			Handler:    _TPTerm_SignRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tp-term.proto",
}
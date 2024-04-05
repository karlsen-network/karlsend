// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: messages.proto

package protowire

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
	P2P_MessageStream_FullMethodName = "/protowire.P2P/MessageStream"
)

// P2PClient is the client API for P2P service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PClient interface {
	MessageStream(ctx context.Context, opts ...grpc.CallOption) (P2P_MessageStreamClient, error)
}

type p2PClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PClient(cc grpc.ClientConnInterface) P2PClient {
	return &p2PClient{cc}
}

func (c *p2PClient) MessageStream(ctx context.Context, opts ...grpc.CallOption) (P2P_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &P2P_ServiceDesc.Streams[0], P2P_MessageStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PMessageStreamClient{stream}
	return x, nil
}

type P2P_MessageStreamClient interface {
	Send(*KarlsendMessage) error
	Recv() (*KarlsendMessage, error)
	grpc.ClientStream
}

type p2PMessageStreamClient struct {
	grpc.ClientStream
}

func (x *p2PMessageStreamClient) Send(m *KarlsendMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *p2PMessageStreamClient) Recv() (*KarlsendMessage, error) {
	m := new(KarlsendMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2PServer is the server API for P2P service.
// All implementations must embed UnimplementedP2PServer
// for forward compatibility
type P2PServer interface {
	MessageStream(P2P_MessageStreamServer) error
	mustEmbedUnimplementedP2PServer()
}

// UnimplementedP2PServer must be embedded to have forward compatible implementations.
type UnimplementedP2PServer struct {
}

func (UnimplementedP2PServer) MessageStream(P2P_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}
func (UnimplementedP2PServer) mustEmbedUnimplementedP2PServer() {}

// UnsafeP2PServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PServer will
// result in compilation errors.
type UnsafeP2PServer interface {
	mustEmbedUnimplementedP2PServer()
}

func RegisterP2PServer(s grpc.ServiceRegistrar, srv P2PServer) {
	s.RegisterService(&P2P_ServiceDesc, srv)
}

func _P2P_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(P2PServer).MessageStream(&p2PMessageStreamServer{stream})
}

type P2P_MessageStreamServer interface {
	Send(*KarlsendMessage) error
	Recv() (*KarlsendMessage, error)
	grpc.ServerStream
}

type p2PMessageStreamServer struct {
	grpc.ServerStream
}

func (x *p2PMessageStreamServer) Send(m *KarlsendMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *p2PMessageStreamServer) Recv() (*KarlsendMessage, error) {
	m := new(KarlsendMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2P_ServiceDesc is the grpc.ServiceDesc for P2P service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2P_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protowire.P2P",
	HandlerType: (*P2PServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageStream",
			Handler:       _P2P_MessageStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

const (
	RPC_MessageStream_FullMethodName = "/protowire.RPC/MessageStream"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RPC_ServiceDesc.Streams[0], RPC_MessageStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCMessageStreamClient{stream}
	return x, nil
}

type RPC_MessageStreamClient interface {
	Send(*KarlsendMessage) error
	Recv() (*KarlsendMessage, error)
	grpc.ClientStream
}

type rPCMessageStreamClient struct {
	grpc.ClientStream
}

func (x *rPCMessageStreamClient) Send(m *KarlsendMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCMessageStreamClient) Recv() (*KarlsendMessage, error) {
	m := new(KarlsendMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	MessageStream(RPC_MessageStreamServer) error
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) MessageStream(RPC_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCServer).MessageStream(&rPCMessageStreamServer{stream})
}

type RPC_MessageStreamServer interface {
	Send(*KarlsendMessage) error
	Recv() (*KarlsendMessage, error)
	grpc.ServerStream
}

type rPCMessageStreamServer struct {
	grpc.ServerStream
}

func (x *rPCMessageStreamServer) Send(m *KarlsendMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCMessageStreamServer) Recv() (*KarlsendMessage, error) {
	m := new(KarlsendMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protowire.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageStream",
			Handler:       _RPC_MessageStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

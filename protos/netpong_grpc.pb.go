// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: protos/netpong.proto

package netpong

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

// NetPongClient is the client API for NetPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetPongClient interface {
	// Initiates a game of pong as a bi-directional stream of paddle moves from each player.
	Play(ctx context.Context, opts ...grpc.CallOption) (NetPong_PlayClient, error)
}

type netPongClient struct {
	cc grpc.ClientConnInterface
}

func NewNetPongClient(cc grpc.ClientConnInterface) NetPongClient {
	return &netPongClient{cc}
}

func (c *netPongClient) Play(ctx context.Context, opts ...grpc.CallOption) (NetPong_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &NetPong_ServiceDesc.Streams[0], "/netpong.NetPong/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &netPongPlayClient{stream}
	return x, nil
}

type NetPong_PlayClient interface {
	Send(*MovePaddle) error
	Recv() (*MovePaddle, error)
	grpc.ClientStream
}

type netPongPlayClient struct {
	grpc.ClientStream
}

func (x *netPongPlayClient) Send(m *MovePaddle) error {
	return x.ClientStream.SendMsg(m)
}

func (x *netPongPlayClient) Recv() (*MovePaddle, error) {
	m := new(MovePaddle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetPongServer is the server API for NetPong service.
// All implementations must embed UnimplementedNetPongServer
// for forward compatibility
type NetPongServer interface {
	// Initiates a game of pong as a bi-directional stream of paddle moves from each player.
	Play(NetPong_PlayServer) error
	mustEmbedUnimplementedNetPongServer()
}

// UnimplementedNetPongServer must be embedded to have forward compatible implementations.
type UnimplementedNetPongServer struct {
}

func (UnimplementedNetPongServer) Play(NetPong_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedNetPongServer) mustEmbedUnimplementedNetPongServer() {}

// UnsafeNetPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetPongServer will
// result in compilation errors.
type UnsafeNetPongServer interface {
	mustEmbedUnimplementedNetPongServer()
}

func RegisterNetPongServer(s grpc.ServiceRegistrar, srv NetPongServer) {
	s.RegisterService(&NetPong_ServiceDesc, srv)
}

func _NetPong_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetPongServer).Play(&netPongPlayServer{stream})
}

type NetPong_PlayServer interface {
	Send(*MovePaddle) error
	Recv() (*MovePaddle, error)
	grpc.ServerStream
}

type netPongPlayServer struct {
	grpc.ServerStream
}

func (x *netPongPlayServer) Send(m *MovePaddle) error {
	return x.ServerStream.SendMsg(m)
}

func (x *netPongPlayServer) Recv() (*MovePaddle, error) {
	m := new(MovePaddle)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetPong_ServiceDesc is the grpc.ServiceDesc for NetPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netpong.NetPong",
	HandlerType: (*NetPongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _NetPong_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/netpong.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// HELLOGRPCClient is the client API for HELLOGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HELLOGRPCClient interface {
	SayHi(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type hELLOGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewHELLOGRPCClient(cc grpc.ClientConnInterface) HELLOGRPCClient {
	return &hELLOGRPCClient{cc}
}

func (c *hELLOGRPCClient) SayHi(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/hello_grpc.HELLOGRPC/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HELLOGRPCServer is the server API for HELLOGRPC service.
// All implementations must embed UnimplementedHELLOGRPCServer
// for forward compatibility
type HELLOGRPCServer interface {
	SayHi(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedHELLOGRPCServer()
}

// UnimplementedHELLOGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedHELLOGRPCServer struct {
}

func (UnimplementedHELLOGRPCServer) SayHi(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedHELLOGRPCServer) mustEmbedUnimplementedHELLOGRPCServer() {}

// UnsafeHELLOGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HELLOGRPCServer will
// result in compilation errors.
type UnsafeHELLOGRPCServer interface {
	mustEmbedUnimplementedHELLOGRPCServer()
}

func RegisterHELLOGRPCServer(s grpc.ServiceRegistrar, srv HELLOGRPCServer) {
	s.RegisterService(&HELLOGRPC_ServiceDesc, srv)
}

func _HELLOGRPC_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HELLOGRPCServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello_grpc.HELLOGRPC/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HELLOGRPCServer).SayHi(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// HELLOGRPC_ServiceDesc is the grpc.ServiceDesc for HELLOGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HELLOGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello_grpc.HELLOGRPC",
	HandlerType: (*HELLOGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _HELLOGRPC_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_grpc.proto",
}

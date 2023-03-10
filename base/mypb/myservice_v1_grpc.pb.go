// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: myservice_v1.proto

package mypb

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

// App1Client is the client API for App1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type App1Client interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type app1Client struct {
	cc grpc.ClientConnInterface
}

func NewApp1Client(cc grpc.ClientConnInterface) App1Client {
	return &app1Client{cc}
}

func (c *app1Client) SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/mypb.app1/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// App1Server is the server API for App1 service.
// All implementations should embed UnimplementedApp1Server
// for forward compatibility
type App1Server interface {
	// Sends a greeting
	SayHello(context.Context, *Hello) (*Hello, error)
}

// UnimplementedApp1Server should be embedded to have forward compatible implementations.
type UnimplementedApp1Server struct {
}

func (UnimplementedApp1Server) SayHello(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

// UnsafeApp1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to App1Server will
// result in compilation errors.
type UnsafeApp1Server interface {
	mustEmbedUnimplementedApp1Server()
}

func RegisterApp1Server(s grpc.ServiceRegistrar, srv App1Server) {
	s.RegisterService(&App1_ServiceDesc, srv)
}

func _App1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(App1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mypb.app1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(App1Server).SayHello(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

// App1_ServiceDesc is the grpc.ServiceDesc for App1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mypb.app1",
	HandlerType: (*App1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _App1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myservice_v1.proto",
}

// App2Client is the client API for App2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type App2Client interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type app2Client struct {
	cc grpc.ClientConnInterface
}

func NewApp2Client(cc grpc.ClientConnInterface) App2Client {
	return &app2Client{cc}
}

func (c *app2Client) SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/mypb.app2/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// App2Server is the server API for App2 service.
// All implementations should embed UnimplementedApp2Server
// for forward compatibility
type App2Server interface {
	// Sends a greeting
	SayHello(context.Context, *Hello) (*Hello, error)
}

// UnimplementedApp2Server should be embedded to have forward compatible implementations.
type UnimplementedApp2Server struct {
}

func (UnimplementedApp2Server) SayHello(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

// UnsafeApp2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to App2Server will
// result in compilation errors.
type UnsafeApp2Server interface {
	mustEmbedUnimplementedApp2Server()
}

func RegisterApp2Server(s grpc.ServiceRegistrar, srv App2Server) {
	s.RegisterService(&App2_ServiceDesc, srv)
}

func _App2_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(App2Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mypb.app2/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(App2Server).SayHello(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

// App2_ServiceDesc is the grpc.ServiceDesc for App2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mypb.app2",
	HandlerType: (*App2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _App2_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myservice_v1.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: account.proto

package messager

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

// MessageReceiverClient is the client API for MessageReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageReceiverClient interface {
	EnableAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error)
}

type messageReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageReceiverClient(cc grpc.ClientConnInterface) MessageReceiverClient {
	return &messageReceiverClient{cc}
}

func (c *messageReceiverClient) EnableAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/messager.MessageReceiver/EnableAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageReceiverServer is the server API for MessageReceiver service.
// All implementations must embed UnimplementedMessageReceiverServer
// for forward compatibility
type MessageReceiverServer interface {
	EnableAccount(context.Context, *Account) (*Response, error)
	mustEmbedUnimplementedMessageReceiverServer()
}

// UnimplementedMessageReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedMessageReceiverServer struct {
}

func (UnimplementedMessageReceiverServer) EnableAccount(context.Context, *Account) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableAccount not implemented")
}
func (UnimplementedMessageReceiverServer) mustEmbedUnimplementedMessageReceiverServer() {}

// UnsafeMessageReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageReceiverServer will
// result in compilation errors.
type UnsafeMessageReceiverServer interface {
	mustEmbedUnimplementedMessageReceiverServer()
}

func RegisterMessageReceiverServer(s grpc.ServiceRegistrar, srv MessageReceiverServer) {
	s.RegisterService(&MessageReceiver_ServiceDesc, srv)
}

func _MessageReceiver_EnableAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageReceiverServer).EnableAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messager.MessageReceiver/EnableAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageReceiverServer).EnableAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageReceiver_ServiceDesc is the grpc.ServiceDesc for MessageReceiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageReceiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messager.MessageReceiver",
	HandlerType: (*MessageReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableAccount",
			Handler:    _MessageReceiver_EnableAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

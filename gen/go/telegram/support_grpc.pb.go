// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: telegram/support.proto

package supv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TelegramSupportBot_SendMessage_FullMethodName = "/telegram.TelegramSupportBot/SendMessage"
)

// TelegramSupportBotClient is the client API for TelegramSupportBot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramSupportBotClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type telegramSupportBotClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramSupportBotClient(cc grpc.ClientConnInterface) TelegramSupportBotClient {
	return &telegramSupportBotClient{cc}
}

func (c *telegramSupportBotClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, TelegramSupportBot_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramSupportBotServer is the server API for TelegramSupportBot service.
// All implementations must embed UnimplementedTelegramSupportBotServer
// for forward compatibility.
type TelegramSupportBotServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedTelegramSupportBotServer()
}

// UnimplementedTelegramSupportBotServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTelegramSupportBotServer struct{}

func (UnimplementedTelegramSupportBotServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTelegramSupportBotServer) mustEmbedUnimplementedTelegramSupportBotServer() {}
func (UnimplementedTelegramSupportBotServer) testEmbeddedByValue()                            {}

// UnsafeTelegramSupportBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramSupportBotServer will
// result in compilation errors.
type UnsafeTelegramSupportBotServer interface {
	mustEmbedUnimplementedTelegramSupportBotServer()
}

func RegisterTelegramSupportBotServer(s grpc.ServiceRegistrar, srv TelegramSupportBotServer) {
	// If the following call pancis, it indicates UnimplementedTelegramSupportBotServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TelegramSupportBot_ServiceDesc, srv)
}

func _TelegramSupportBot_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramSupportBotServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelegramSupportBot_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramSupportBotServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelegramSupportBot_ServiceDesc is the grpc.ServiceDesc for TelegramSupportBot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelegramSupportBot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telegram.TelegramSupportBot",
	HandlerType: (*TelegramSupportBotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _TelegramSupportBot_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telegram/support.proto",
}

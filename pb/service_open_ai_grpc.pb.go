// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service_open_ai.proto

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

// OpenAiClient is the client API for OpenAi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenAiClient interface {
	ChatGpt(ctx context.Context, in *ChatGptRequest, opts ...grpc.CallOption) (*ChatGptReply, error)
}

type openAiClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenAiClient(cc grpc.ClientConnInterface) OpenAiClient {
	return &openAiClient{cc}
}

func (c *openAiClient) ChatGpt(ctx context.Context, in *ChatGptRequest, opts ...grpc.CallOption) (*ChatGptReply, error) {
	out := new(ChatGptReply)
	err := c.cc.Invoke(ctx, "/pb.OpenAi/ChatGpt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenAiServer is the server API for OpenAi service.
// All implementations must embed UnimplementedOpenAiServer
// for forward compatibility
type OpenAiServer interface {
	ChatGpt(context.Context, *ChatGptRequest) (*ChatGptReply, error)
	mustEmbedUnimplementedOpenAiServer()
}

// UnimplementedOpenAiServer must be embedded to have forward compatible implementations.
type UnimplementedOpenAiServer struct {
}

func (UnimplementedOpenAiServer) ChatGpt(context.Context, *ChatGptRequest) (*ChatGptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatGpt not implemented")
}
func (UnimplementedOpenAiServer) mustEmbedUnimplementedOpenAiServer() {}

// UnsafeOpenAiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenAiServer will
// result in compilation errors.
type UnsafeOpenAiServer interface {
	mustEmbedUnimplementedOpenAiServer()
}

func RegisterOpenAiServer(s grpc.ServiceRegistrar, srv OpenAiServer) {
	s.RegisterService(&OpenAi_ServiceDesc, srv)
}

func _OpenAi_ChatGpt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatGptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenAiServer).ChatGpt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OpenAi/ChatGpt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenAiServer).ChatGpt(ctx, req.(*ChatGptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenAi_ServiceDesc is the grpc.ServiceDesc for OpenAi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenAi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OpenAi",
	HandlerType: (*OpenAiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatGpt",
			Handler:    _OpenAi_ChatGpt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_open_ai.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: gpt.proto

package proto

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

// GptServiceClient is the client API for GptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GptServiceClient interface {
	DetectLanguage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error)
	ExtractSection(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ExtractSectionResponse, error)
	ExtractDiploma(ctx context.Context, in *ExtractDiplomaRequest, opts ...grpc.CallOption) (*ExtractDiplomaResponse, error)
	FilterKeyWords(ctx context.Context, in *RepeatedTextRequest, opts ...grpc.CallOption) (*FilterKeyWordsResponse, error)
	FormatReferences(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ReformattedReferenceResponse, error)
}

type gptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGptServiceClient(cc grpc.ClientConnInterface) GptServiceClient {
	return &gptServiceClient{cc}
}

func (c *gptServiceClient) DetectLanguage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error) {
	out := new(DetectLanguageResponse)
	err := c.cc.Invoke(ctx, "/pp.gpt.GptService/DetectLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) ExtractSection(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ExtractSectionResponse, error) {
	out := new(ExtractSectionResponse)
	err := c.cc.Invoke(ctx, "/pp.gpt.GptService/ExtractSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) ExtractDiploma(ctx context.Context, in *ExtractDiplomaRequest, opts ...grpc.CallOption) (*ExtractDiplomaResponse, error) {
	out := new(ExtractDiplomaResponse)
	err := c.cc.Invoke(ctx, "/pp.gpt.GptService/ExtractDiploma", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) FilterKeyWords(ctx context.Context, in *RepeatedTextRequest, opts ...grpc.CallOption) (*FilterKeyWordsResponse, error) {
	out := new(FilterKeyWordsResponse)
	err := c.cc.Invoke(ctx, "/pp.gpt.GptService/FilterKeyWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) FormatReferences(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ReformattedReferenceResponse, error) {
	out := new(ReformattedReferenceResponse)
	err := c.cc.Invoke(ctx, "/pp.gpt.GptService/FormatReferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GptServiceServer is the server API for GptService service.
// All implementations must embed UnimplementedGptServiceServer
// for forward compatibility
type GptServiceServer interface {
	DetectLanguage(context.Context, *TextRequest) (*DetectLanguageResponse, error)
	ExtractSection(context.Context, *TextRequest) (*ExtractSectionResponse, error)
	ExtractDiploma(context.Context, *ExtractDiplomaRequest) (*ExtractDiplomaResponse, error)
	FilterKeyWords(context.Context, *RepeatedTextRequest) (*FilterKeyWordsResponse, error)
	FormatReferences(context.Context, *TextRequest) (*ReformattedReferenceResponse, error)
	mustEmbedUnimplementedGptServiceServer()
}

// UnimplementedGptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGptServiceServer struct {
}

func (UnimplementedGptServiceServer) DetectLanguage(context.Context, *TextRequest) (*DetectLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectLanguage not implemented")
}
func (UnimplementedGptServiceServer) ExtractSection(context.Context, *TextRequest) (*ExtractSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractSection not implemented")
}
func (UnimplementedGptServiceServer) ExtractDiploma(context.Context, *ExtractDiplomaRequest) (*ExtractDiplomaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractDiploma not implemented")
}
func (UnimplementedGptServiceServer) FilterKeyWords(context.Context, *RepeatedTextRequest) (*FilterKeyWordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterKeyWords not implemented")
}
func (UnimplementedGptServiceServer) FormatReferences(context.Context, *TextRequest) (*ReformattedReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormatReferences not implemented")
}
func (UnimplementedGptServiceServer) mustEmbedUnimplementedGptServiceServer() {}

// UnsafeGptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GptServiceServer will
// result in compilation errors.
type UnsafeGptServiceServer interface {
	mustEmbedUnimplementedGptServiceServer()
}

func RegisterGptServiceServer(s grpc.ServiceRegistrar, srv GptServiceServer) {
	s.RegisterService(&GptService_ServiceDesc, srv)
}

func _GptService_DetectLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).DetectLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.gpt.GptService/DetectLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).DetectLanguage(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_ExtractSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).ExtractSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.gpt.GptService/ExtractSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).ExtractSection(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_ExtractDiploma_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractDiplomaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).ExtractDiploma(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.gpt.GptService/ExtractDiploma",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).ExtractDiploma(ctx, req.(*ExtractDiplomaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_FilterKeyWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepeatedTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).FilterKeyWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.gpt.GptService/FilterKeyWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).FilterKeyWords(ctx, req.(*RepeatedTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_FormatReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).FormatReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.gpt.GptService/FormatReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).FormatReferences(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GptService_ServiceDesc is the grpc.ServiceDesc for GptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pp.gpt.GptService",
	HandlerType: (*GptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectLanguage",
			Handler:    _GptService_DetectLanguage_Handler,
		},
		{
			MethodName: "ExtractSection",
			Handler:    _GptService_ExtractSection_Handler,
		},
		{
			MethodName: "ExtractDiploma",
			Handler:    _GptService_ExtractDiploma_Handler,
		},
		{
			MethodName: "FilterKeyWords",
			Handler:    _GptService_FilterKeyWords_Handler,
		},
		{
			MethodName: "FormatReferences",
			Handler:    _GptService_FormatReferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gpt.proto",
}

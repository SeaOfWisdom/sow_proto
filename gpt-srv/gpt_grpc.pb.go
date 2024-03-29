// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
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

const (
	GptService_DetectLanguage_FullMethodName           = "/pp.gpt.GptService/DetectLanguage"
	GptService_ExtractSection_FullMethodName           = "/pp.gpt.GptService/ExtractSection"
	GptService_ExtractDiploma_FullMethodName           = "/pp.gpt.GptService/ExtractDiploma"
	GptService_FilterKeyWords_FullMethodName           = "/pp.gpt.GptService/FilterKeyWords"
	GptService_FormatReferences_FullMethodName         = "/pp.gpt.GptService/FormatReferences"
	GptService_ReviewPaper_FullMethodName              = "/pp.gpt.GptService/ReviewPaper"
	GptService_CensorNickname_FullMethodName           = "/pp.gpt.GptService/CensorNickname"
	GptService_TransliterateNameSurname_FullMethodName = "/pp.gpt.GptService/TransliterateNameSurname"
)

// GptServiceClient is the client API for GptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GptServiceClient interface {
	DetectLanguage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error)
	ExtractSection(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ExtractSectionResponse, error)
	ExtractDiploma(ctx context.Context, in *ExtractDiplomaRequest, opts ...grpc.CallOption) (*ExtractDiplomaResponse, error)
	FilterKeyWords(ctx context.Context, in *RepeatedTextRequest, opts ...grpc.CallOption) (*FilterKeyWordsResponse, error)
	FormatReferences(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ReformattedReferenceResponse, error)
	ReviewPaper(ctx context.Context, in *ScientificCriteriaRequest, opts ...grpc.CallOption) (*ScientificReviewResponse, error)
	CensorNickname(ctx context.Context, in *NicknameRequest, opts ...grpc.CallOption) (*NicknameResponse, error)
	TransliterateNameSurname(ctx context.Context, in *NameSurnameRequest, opts ...grpc.CallOption) (*NameSurnameResponse, error)
}

type gptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGptServiceClient(cc grpc.ClientConnInterface) GptServiceClient {
	return &gptServiceClient{cc}
}

func (c *gptServiceClient) DetectLanguage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error) {
	out := new(DetectLanguageResponse)
	err := c.cc.Invoke(ctx, GptService_DetectLanguage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) ExtractSection(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ExtractSectionResponse, error) {
	out := new(ExtractSectionResponse)
	err := c.cc.Invoke(ctx, GptService_ExtractSection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) ExtractDiploma(ctx context.Context, in *ExtractDiplomaRequest, opts ...grpc.CallOption) (*ExtractDiplomaResponse, error) {
	out := new(ExtractDiplomaResponse)
	err := c.cc.Invoke(ctx, GptService_ExtractDiploma_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) FilterKeyWords(ctx context.Context, in *RepeatedTextRequest, opts ...grpc.CallOption) (*FilterKeyWordsResponse, error) {
	out := new(FilterKeyWordsResponse)
	err := c.cc.Invoke(ctx, GptService_FilterKeyWords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) FormatReferences(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ReformattedReferenceResponse, error) {
	out := new(ReformattedReferenceResponse)
	err := c.cc.Invoke(ctx, GptService_FormatReferences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) ReviewPaper(ctx context.Context, in *ScientificCriteriaRequest, opts ...grpc.CallOption) (*ScientificReviewResponse, error) {
	out := new(ScientificReviewResponse)
	err := c.cc.Invoke(ctx, GptService_ReviewPaper_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) CensorNickname(ctx context.Context, in *NicknameRequest, opts ...grpc.CallOption) (*NicknameResponse, error) {
	out := new(NicknameResponse)
	err := c.cc.Invoke(ctx, GptService_CensorNickname_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gptServiceClient) TransliterateNameSurname(ctx context.Context, in *NameSurnameRequest, opts ...grpc.CallOption) (*NameSurnameResponse, error) {
	out := new(NameSurnameResponse)
	err := c.cc.Invoke(ctx, GptService_TransliterateNameSurname_FullMethodName, in, out, opts...)
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
	ReviewPaper(context.Context, *ScientificCriteriaRequest) (*ScientificReviewResponse, error)
	CensorNickname(context.Context, *NicknameRequest) (*NicknameResponse, error)
	TransliterateNameSurname(context.Context, *NameSurnameRequest) (*NameSurnameResponse, error)
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
func (UnimplementedGptServiceServer) ReviewPaper(context.Context, *ScientificCriteriaRequest) (*ScientificReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewPaper not implemented")
}
func (UnimplementedGptServiceServer) CensorNickname(context.Context, *NicknameRequest) (*NicknameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CensorNickname not implemented")
}
func (UnimplementedGptServiceServer) TransliterateNameSurname(context.Context, *NameSurnameRequest) (*NameSurnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransliterateNameSurname not implemented")
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
		FullMethod: GptService_DetectLanguage_FullMethodName,
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
		FullMethod: GptService_ExtractSection_FullMethodName,
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
		FullMethod: GptService_ExtractDiploma_FullMethodName,
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
		FullMethod: GptService_FilterKeyWords_FullMethodName,
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
		FullMethod: GptService_FormatReferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).FormatReferences(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_ReviewPaper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScientificCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).ReviewPaper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GptService_ReviewPaper_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).ReviewPaper(ctx, req.(*ScientificCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_CensorNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).CensorNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GptService_CensorNickname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).CensorNickname(ctx, req.(*NicknameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GptService_TransliterateNameSurname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameSurnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GptServiceServer).TransliterateNameSurname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GptService_TransliterateNameSurname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GptServiceServer).TransliterateNameSurname(ctx, req.(*NameSurnameRequest))
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
		{
			MethodName: "ReviewPaper",
			Handler:    _GptService_ReviewPaper_Handler,
		},
		{
			MethodName: "CensorNickname",
			Handler:    _GptService_CensorNickname_Handler,
		},
		{
			MethodName: "TransliterateNameSurname",
			Handler:    _GptService_TransliterateNameSurname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gpt.proto",
}

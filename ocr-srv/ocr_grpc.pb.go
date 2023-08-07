// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: ocr-srv/ocr.proto

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

// OCRClient is the client API for OCR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCRClient interface {
	ExtractText(ctx context.Context, in *ExtractTextRequest, opts ...grpc.CallOption) (*ExtractTextResponse, error)
	ExtractValidatorText(ctx context.Context, in *ExtractValidatorRequest, opts ...grpc.CallOption) (*ExtractValidatorResponse, error)
	ExtractLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageResponse, error)
}

type oCRClient struct {
	cc grpc.ClientConnInterface
}

func NewOCRClient(cc grpc.ClientConnInterface) OCRClient {
	return &oCRClient{cc}
}

func (c *oCRClient) ExtractText(ctx context.Context, in *ExtractTextRequest, opts ...grpc.CallOption) (*ExtractTextResponse, error) {
	out := new(ExtractTextResponse)
	err := c.cc.Invoke(ctx, "/ocr.OCR/ExtractText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCRClient) ExtractValidatorText(ctx context.Context, in *ExtractValidatorRequest, opts ...grpc.CallOption) (*ExtractValidatorResponse, error) {
	out := new(ExtractValidatorResponse)
	err := c.cc.Invoke(ctx, "/ocr.OCR/ExtractValidatorText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCRClient) ExtractLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageResponse, error) {
	out := new(LanguageResponse)
	err := c.cc.Invoke(ctx, "/ocr.OCR/ExtractLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCRServer is the server API for OCR service.
// All implementations must embed UnimplementedOCRServer
// for forward compatibility
type OCRServer interface {
	ExtractText(context.Context, *ExtractTextRequest) (*ExtractTextResponse, error)
	ExtractValidatorText(context.Context, *ExtractValidatorRequest) (*ExtractValidatorResponse, error)
	ExtractLanguage(context.Context, *LanguageRequest) (*LanguageResponse, error)
	mustEmbedUnimplementedOCRServer()
}

// UnimplementedOCRServer must be embedded to have forward compatible implementations.
type UnimplementedOCRServer struct {
}

func (UnimplementedOCRServer) ExtractText(context.Context, *ExtractTextRequest) (*ExtractTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractText not implemented")
}
func (UnimplementedOCRServer) ExtractValidatorText(context.Context, *ExtractValidatorRequest) (*ExtractValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractValidatorText not implemented")
}
func (UnimplementedOCRServer) ExtractLanguage(context.Context, *LanguageRequest) (*LanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractLanguage not implemented")
}
func (UnimplementedOCRServer) mustEmbedUnimplementedOCRServer() {}

// UnsafeOCRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCRServer will
// result in compilation errors.
type UnsafeOCRServer interface {
	mustEmbedUnimplementedOCRServer()
}

func RegisterOCRServer(s grpc.ServiceRegistrar, srv OCRServer) {
	s.RegisterService(&OCR_ServiceDesc, srv)
}

func _OCR_ExtractText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServer).ExtractText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocr.OCR/ExtractText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServer).ExtractText(ctx, req.(*ExtractTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCR_ExtractValidatorText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServer).ExtractValidatorText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocr.OCR/ExtractValidatorText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServer).ExtractValidatorText(ctx, req.(*ExtractValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCR_ExtractLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServer).ExtractLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocr.OCR/ExtractLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServer).ExtractLanguage(ctx, req.(*LanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OCR_ServiceDesc is the grpc.ServiceDesc for OCR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocr.OCR",
	HandlerType: (*OCRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExtractText",
			Handler:    _OCR_ExtractText_Handler,
		},
		{
			MethodName: "ExtractValidatorText",
			Handler:    _OCR_ExtractValidatorText_Handler,
		},
		{
			MethodName: "ExtractLanguage",
			Handler:    _OCR_ExtractLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocr-srv/ocr.proto",
}

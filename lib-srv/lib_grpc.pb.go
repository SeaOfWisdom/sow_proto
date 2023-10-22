// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: lib-srv/lib.proto

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

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	MakeAsPurchased(ctx context.Context, in *MakeAsPurchasedRequest, opts ...grpc.CallOption) (*Null, error)
	SearchSubmittedReviews(ctx context.Context, in *SearchWorkReviewsRequest, opts ...grpc.CallOption) (*SearchReviewsResponse, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) MakeAsPurchased(ctx context.Context, in *MakeAsPurchasedRequest, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/pp.contractor.LibraryService/MakeAsPurchased", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) SearchSubmittedReviews(ctx context.Context, in *SearchWorkReviewsRequest, opts ...grpc.CallOption) (*SearchReviewsResponse, error) {
	out := new(SearchReviewsResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.LibraryService/SearchSubmittedReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations must embed UnimplementedLibraryServiceServer
// for forward compatibility
type LibraryServiceServer interface {
	MakeAsPurchased(context.Context, *MakeAsPurchasedRequest) (*Null, error)
	SearchSubmittedReviews(context.Context, *SearchWorkReviewsRequest) (*SearchReviewsResponse, error)
	mustEmbedUnimplementedLibraryServiceServer()
}

// UnimplementedLibraryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServiceServer struct {
}

func (UnimplementedLibraryServiceServer) MakeAsPurchased(context.Context, *MakeAsPurchasedRequest) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAsPurchased not implemented")
}
func (UnimplementedLibraryServiceServer) SearchSubmittedReviews(context.Context, *SearchWorkReviewsRequest) (*SearchReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSubmittedReviews not implemented")
}
func (UnimplementedLibraryServiceServer) mustEmbedUnimplementedLibraryServiceServer() {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_MakeAsPurchased_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeAsPurchasedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).MakeAsPurchased(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.LibraryService/MakeAsPurchased",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).MakeAsPurchased(ctx, req.(*MakeAsPurchasedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_SearchSubmittedReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWorkReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).SearchSubmittedReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.LibraryService/SearchSubmittedReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).SearchSubmittedReviews(ctx, req.(*SearchWorkReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pp.contractor.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeAsPurchased",
			Handler:    _LibraryService_MakeAsPurchased_Handler,
		},
		{
			MethodName: "SearchSubmittedReviews",
			Handler:    _LibraryService_SearchSubmittedReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib-srv/lib.proto",
}

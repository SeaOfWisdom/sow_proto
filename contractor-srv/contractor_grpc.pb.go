// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: contractor-srv/contractor.proto

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

// ContractorServiceClient is the client API for ContractorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractorServiceClient interface {
	PublishWork(ctx context.Context, in *PublishWorkRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	PurchaseWork(ctx context.Context, in *PurchaseWorkRequest, opts ...grpc.CallOption) (*PurchaseWorkResponse, error)
	MakeAuthor(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	MakeAdmin(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	MakeReviewer(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	Faucet(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	AddParticipant(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	GetStatus(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*TxStatusResponse, error)
	GetPaperById(ctx context.Context, in *PaperByIdRequest, opts ...grpc.CallOption) (*PaperByIdResponse, error)
	GetParticipantRole(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RoleResponse, error)
}

type contractorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractorServiceClient(cc grpc.ClientConnInterface) ContractorServiceClient {
	return &contractorServiceClient{cc}
}

func (c *contractorServiceClient) PublishWork(ctx context.Context, in *PublishWorkRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/PublishWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) PurchaseWork(ctx context.Context, in *PurchaseWorkRequest, opts ...grpc.CallOption) (*PurchaseWorkResponse, error) {
	out := new(PurchaseWorkResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/PurchaseWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeAuthor(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/MakeAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeAdmin(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/MakeAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeReviewer(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/MakeReviewer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) Faucet(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/Faucet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) AddParticipant(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/AddParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetStatus(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*TxStatusResponse, error) {
	out := new(TxStatusResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetPaperById(ctx context.Context, in *PaperByIdRequest, opts ...grpc.CallOption) (*PaperByIdResponse, error) {
	out := new(PaperByIdResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/GetPaperById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetParticipantRole(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/pp.contractor.ContractorService/GetParticipantRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractorServiceServer is the server API for ContractorService service.
// All implementations must embed UnimplementedContractorServiceServer
// for forward compatibility
type ContractorServiceServer interface {
	PublishWork(context.Context, *PublishWorkRequest) (*TxHashResponse, error)
	PurchaseWork(context.Context, *PurchaseWorkRequest) (*PurchaseWorkResponse, error)
	MakeAuthor(context.Context, *AccountRequest) (*TxHashResponse, error)
	MakeAdmin(context.Context, *AccountRequest) (*TxHashResponse, error)
	MakeReviewer(context.Context, *AccountRequest) (*TxHashResponse, error)
	Faucet(context.Context, *FaucetRequest) (*TxHashResponse, error)
	AddParticipant(context.Context, *AccountRequest) (*TxHashResponse, error)
	GetStatus(context.Context, *TxStatusRequest) (*TxStatusResponse, error)
	GetPaperById(context.Context, *PaperByIdRequest) (*PaperByIdResponse, error)
	GetParticipantRole(context.Context, *AccountRequest) (*RoleResponse, error)
	mustEmbedUnimplementedContractorServiceServer()
}

// UnimplementedContractorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractorServiceServer struct {
}

func (UnimplementedContractorServiceServer) PublishWork(context.Context, *PublishWorkRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishWork not implemented")
}
func (UnimplementedContractorServiceServer) PurchaseWork(context.Context, *PurchaseWorkRequest) (*PurchaseWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseWork not implemented")
}
func (UnimplementedContractorServiceServer) MakeAuthor(context.Context, *AccountRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAuthor not implemented")
}
func (UnimplementedContractorServiceServer) MakeAdmin(context.Context, *AccountRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAdmin not implemented")
}
func (UnimplementedContractorServiceServer) MakeReviewer(context.Context, *AccountRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeReviewer not implemented")
}
func (UnimplementedContractorServiceServer) Faucet(context.Context, *FaucetRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Faucet not implemented")
}
func (UnimplementedContractorServiceServer) AddParticipant(context.Context, *AccountRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipant not implemented")
}
func (UnimplementedContractorServiceServer) GetStatus(context.Context, *TxStatusRequest) (*TxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedContractorServiceServer) GetPaperById(context.Context, *PaperByIdRequest) (*PaperByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaperById not implemented")
}
func (UnimplementedContractorServiceServer) GetParticipantRole(context.Context, *AccountRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipantRole not implemented")
}
func (UnimplementedContractorServiceServer) mustEmbedUnimplementedContractorServiceServer() {}

// UnsafeContractorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractorServiceServer will
// result in compilation errors.
type UnsafeContractorServiceServer interface {
	mustEmbedUnimplementedContractorServiceServer()
}

func RegisterContractorServiceServer(s grpc.ServiceRegistrar, srv ContractorServiceServer) {
	s.RegisterService(&ContractorService_ServiceDesc, srv)
}

func _ContractorService_PublishWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).PublishWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/PublishWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).PublishWork(ctx, req.(*PublishWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_PurchaseWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).PurchaseWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/PurchaseWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).PurchaseWork(ctx, req.(*PurchaseWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_MakeAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).MakeAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/MakeAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).MakeAuthor(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_MakeAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).MakeAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/MakeAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).MakeAdmin(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_MakeReviewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).MakeReviewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/MakeReviewer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).MakeReviewer(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_Faucet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaucetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).Faucet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/Faucet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).Faucet(ctx, req.(*FaucetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_AddParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).AddParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/AddParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).AddParticipant(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).GetStatus(ctx, req.(*TxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_GetPaperById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaperByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).GetPaperById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/GetPaperById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).GetPaperById(ctx, req.(*PaperByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_GetParticipantRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).GetParticipantRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pp.contractor.ContractorService/GetParticipantRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).GetParticipantRole(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractorService_ServiceDesc is the grpc.ServiceDesc for ContractorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pp.contractor.ContractorService",
	HandlerType: (*ContractorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishWork",
			Handler:    _ContractorService_PublishWork_Handler,
		},
		{
			MethodName: "PurchaseWork",
			Handler:    _ContractorService_PurchaseWork_Handler,
		},
		{
			MethodName: "MakeAuthor",
			Handler:    _ContractorService_MakeAuthor_Handler,
		},
		{
			MethodName: "MakeAdmin",
			Handler:    _ContractorService_MakeAdmin_Handler,
		},
		{
			MethodName: "MakeReviewer",
			Handler:    _ContractorService_MakeReviewer_Handler,
		},
		{
			MethodName: "Faucet",
			Handler:    _ContractorService_Faucet_Handler,
		},
		{
			MethodName: "AddParticipant",
			Handler:    _ContractorService_AddParticipant_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _ContractorService_GetStatus_Handler,
		},
		{
			MethodName: "GetPaperById",
			Handler:    _ContractorService_GetPaperById_Handler,
		},
		{
			MethodName: "GetParticipantRole",
			Handler:    _ContractorService_GetParticipantRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contractor-srv/contractor.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: contractor-srv/contractor.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContractorService_PublishWork_FullMethodName                = "/pp.contractor.ContractorService/PublishWork"
	ContractorService_PurchaseWork_FullMethodName               = "/pp.contractor.ContractorService/PurchaseWork"
	ContractorService_MakeAuthor_FullMethodName                 = "/pp.contractor.ContractorService/MakeAuthor"
	ContractorService_MakeAdmin_FullMethodName                  = "/pp.contractor.ContractorService/MakeAdmin"
	ContractorService_MakeReviewer_FullMethodName               = "/pp.contractor.ContractorService/MakeReviewer"
	ContractorService_Faucet_FullMethodName                     = "/pp.contractor.ContractorService/Faucet"
	ContractorService_AddParticipant_FullMethodName             = "/pp.contractor.ContractorService/AddParticipant"
	ContractorService_GetStatus_FullMethodName                  = "/pp.contractor.ContractorService/GetStatus"
	ContractorService_GetPaperById_FullMethodName               = "/pp.contractor.ContractorService/GetPaperById"
	ContractorService_GetParticipantRole_FullMethodName         = "/pp.contractor.ContractorService/GetParticipantRole"
	ContractorService_AddReviewsBatch_FullMethodName            = "/pp.contractor.ContractorService/AddReviewsBatch"
	ContractorService_IsPaperReadableFor_FullMethodName         = "/pp.contractor.ContractorService/IsPaperReadableFor"
	ContractorService_PaperExpiresFor_FullMethodName            = "/pp.contractor.ContractorService/PaperExpiresFor"
	ContractorService_CheckNewParticipantAddress_FullMethodName = "/pp.contractor.ContractorService/CheckNewParticipantAddress"
	ContractorService_VerifySignature_FullMethodName            = "/pp.contractor.ContractorService/VerifySignature"
)

// ContractorServiceClient is the client API for ContractorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractorServiceClient interface {
	PublishWork(ctx context.Context, in *PublishWorkRequest, opts ...grpc.CallOption) (*PublishWorkResponse, error)
	PurchaseWork(ctx context.Context, in *PurchaseWorkRequest, opts ...grpc.CallOption) (*PurchaseWorkResponse, error)
	MakeAuthor(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	MakeAdmin(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	MakeReviewer(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	Faucet(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	AddParticipant(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	GetStatus(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*TxStatusResponse, error)
	GetPaperById(ctx context.Context, in *PaperByIdRequest, opts ...grpc.CallOption) (*PaperByIdResponse, error)
	GetParticipantRole(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	AddReviewsBatch(ctx context.Context, in *AddReviewsRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	IsPaperReadableFor(ctx context.Context, in *IsPaperReadableRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
	PaperExpiresFor(ctx context.Context, in *PaperExpiresForRequest, opts ...grpc.CallOption) (*PaperExpiresForResponse, error)
	CheckNewParticipantAddress(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VerifySignature(ctx context.Context, in *VerifySignatureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contractorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractorServiceClient(cc grpc.ClientConnInterface) ContractorServiceClient {
	return &contractorServiceClient{cc}
}

func (c *contractorServiceClient) PublishWork(ctx context.Context, in *PublishWorkRequest, opts ...grpc.CallOption) (*PublishWorkResponse, error) {
	out := new(PublishWorkResponse)
	err := c.cc.Invoke(ctx, ContractorService_PublishWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) PurchaseWork(ctx context.Context, in *PurchaseWorkRequest, opts ...grpc.CallOption) (*PurchaseWorkResponse, error) {
	out := new(PurchaseWorkResponse)
	err := c.cc.Invoke(ctx, ContractorService_PurchaseWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeAuthor(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_MakeAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeAdmin(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_MakeAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) MakeReviewer(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_MakeReviewer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) Faucet(ctx context.Context, in *FaucetRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_Faucet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) AddParticipant(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_AddParticipant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetStatus(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*TxStatusResponse, error) {
	out := new(TxStatusResponse)
	err := c.cc.Invoke(ctx, ContractorService_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetPaperById(ctx context.Context, in *PaperByIdRequest, opts ...grpc.CallOption) (*PaperByIdResponse, error) {
	out := new(PaperByIdResponse)
	err := c.cc.Invoke(ctx, ContractorService_GetPaperById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) GetParticipantRole(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, ContractorService_GetParticipantRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) AddReviewsBatch(ctx context.Context, in *AddReviewsRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, ContractorService_AddReviewsBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) IsPaperReadableFor(ctx context.Context, in *IsPaperReadableRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, ContractorService_IsPaperReadableFor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) PaperExpiresFor(ctx context.Context, in *PaperExpiresForRequest, opts ...grpc.CallOption) (*PaperExpiresForResponse, error) {
	out := new(PaperExpiresForResponse)
	err := c.cc.Invoke(ctx, ContractorService_PaperExpiresFor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) CheckNewParticipantAddress(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContractorService_CheckNewParticipantAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractorServiceClient) VerifySignature(ctx context.Context, in *VerifySignatureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContractorService_VerifySignature_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractorServiceServer is the server API for ContractorService service.
// All implementations must embed UnimplementedContractorServiceServer
// for forward compatibility
type ContractorServiceServer interface {
	PublishWork(context.Context, *PublishWorkRequest) (*PublishWorkResponse, error)
	PurchaseWork(context.Context, *PurchaseWorkRequest) (*PurchaseWorkResponse, error)
	MakeAuthor(context.Context, *AccountRequest) (*TxHashResponse, error)
	MakeAdmin(context.Context, *AccountRequest) (*TxHashResponse, error)
	MakeReviewer(context.Context, *AccountRequest) (*TxHashResponse, error)
	Faucet(context.Context, *FaucetRequest) (*TxHashResponse, error)
	AddParticipant(context.Context, *AccountRequest) (*TxHashResponse, error)
	GetStatus(context.Context, *TxStatusRequest) (*TxStatusResponse, error)
	GetPaperById(context.Context, *PaperByIdRequest) (*PaperByIdResponse, error)
	GetParticipantRole(context.Context, *AccountRequest) (*RoleResponse, error)
	AddReviewsBatch(context.Context, *AddReviewsRequest) (*TxHashResponse, error)
	IsPaperReadableFor(context.Context, *IsPaperReadableRequest) (*BooleanResponse, error)
	PaperExpiresFor(context.Context, *PaperExpiresForRequest) (*PaperExpiresForResponse, error)
	CheckNewParticipantAddress(context.Context, *AccountRequest) (*emptypb.Empty, error)
	VerifySignature(context.Context, *VerifySignatureRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedContractorServiceServer()
}

// UnimplementedContractorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractorServiceServer struct {
}

func (UnimplementedContractorServiceServer) PublishWork(context.Context, *PublishWorkRequest) (*PublishWorkResponse, error) {
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
func (UnimplementedContractorServiceServer) AddReviewsBatch(context.Context, *AddReviewsRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReviewsBatch not implemented")
}
func (UnimplementedContractorServiceServer) IsPaperReadableFor(context.Context, *IsPaperReadableRequest) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPaperReadableFor not implemented")
}
func (UnimplementedContractorServiceServer) PaperExpiresFor(context.Context, *PaperExpiresForRequest) (*PaperExpiresForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaperExpiresFor not implemented")
}
func (UnimplementedContractorServiceServer) CheckNewParticipantAddress(context.Context, *AccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNewParticipantAddress not implemented")
}
func (UnimplementedContractorServiceServer) VerifySignature(context.Context, *VerifySignatureRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySignature not implemented")
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
		FullMethod: ContractorService_PublishWork_FullMethodName,
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
		FullMethod: ContractorService_PurchaseWork_FullMethodName,
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
		FullMethod: ContractorService_MakeAuthor_FullMethodName,
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
		FullMethod: ContractorService_MakeAdmin_FullMethodName,
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
		FullMethod: ContractorService_MakeReviewer_FullMethodName,
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
		FullMethod: ContractorService_Faucet_FullMethodName,
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
		FullMethod: ContractorService_AddParticipant_FullMethodName,
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
		FullMethod: ContractorService_GetStatus_FullMethodName,
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
		FullMethod: ContractorService_GetPaperById_FullMethodName,
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
		FullMethod: ContractorService_GetParticipantRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).GetParticipantRole(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_AddReviewsBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).AddReviewsBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractorService_AddReviewsBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).AddReviewsBatch(ctx, req.(*AddReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_IsPaperReadableFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPaperReadableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).IsPaperReadableFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractorService_IsPaperReadableFor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).IsPaperReadableFor(ctx, req.(*IsPaperReadableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_PaperExpiresFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaperExpiresForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).PaperExpiresFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractorService_PaperExpiresFor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).PaperExpiresFor(ctx, req.(*PaperExpiresForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_CheckNewParticipantAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).CheckNewParticipantAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractorService_CheckNewParticipantAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).CheckNewParticipantAddress(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractorService_VerifySignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServiceServer).VerifySignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContractorService_VerifySignature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServiceServer).VerifySignature(ctx, req.(*VerifySignatureRequest))
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
		{
			MethodName: "AddReviewsBatch",
			Handler:    _ContractorService_AddReviewsBatch_Handler,
		},
		{
			MethodName: "IsPaperReadableFor",
			Handler:    _ContractorService_IsPaperReadableFor_Handler,
		},
		{
			MethodName: "PaperExpiresFor",
			Handler:    _ContractorService_PaperExpiresFor_Handler,
		},
		{
			MethodName: "CheckNewParticipantAddress",
			Handler:    _ContractorService_CheckNewParticipantAddress_Handler,
		},
		{
			MethodName: "VerifySignature",
			Handler:    _ContractorService_VerifySignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contractor-srv/contractor.proto",
}

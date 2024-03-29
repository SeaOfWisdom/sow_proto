// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: lib-srv/lib.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidatorModeration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorId string `protobuf:"bytes,1,opt,name=validator_id,json=validatorId,proto3" json:"validator_id,omitempty"`
	IsSuccess   bool   `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *ValidatorModeration) Reset() {
	*x = ValidatorModeration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorModeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorModeration) ProtoMessage() {}

func (x *ValidatorModeration) ProtoReflect() protoreflect.Message {
	mi := &file_lib_srv_lib_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorModeration.ProtoReflect.Descriptor instead.
func (*ValidatorModeration) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorModeration) GetValidatorId() string {
	if x != nil {
		return x.ValidatorId
	}
	return ""
}

func (x *ValidatorModeration) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type MakeAsPurchasedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReaderAddress string `protobuf:"bytes,1,opt,name=reader_address,json=readerAddress,proto3" json:"reader_address,omitempty"`
	WorkId        string `protobuf:"bytes,2,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
}

func (x *MakeAsPurchasedRequest) Reset() {
	*x = MakeAsPurchasedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeAsPurchasedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeAsPurchasedRequest) ProtoMessage() {}

func (x *MakeAsPurchasedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_srv_lib_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeAsPurchasedRequest.ProtoReflect.Descriptor instead.
func (*MakeAsPurchasedRequest) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{1}
}

func (x *MakeAsPurchasedRequest) GetReaderAddress() string {
	if x != nil {
		return x.ReaderAddress
	}
	return ""
}

func (x *MakeAsPurchasedRequest) GetWorkId() string {
	if x != nil {
		return x.WorkId
	}
	return ""
}

type SearchWorkReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Since   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Until   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *SearchWorkReviewsRequest) Reset() {
	*x = SearchWorkReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchWorkReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchWorkReviewsRequest) ProtoMessage() {}

func (x *SearchWorkReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_srv_lib_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchWorkReviewsRequest.ProtoReflect.Descriptor instead.
func (*SearchWorkReviewsRequest) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{2}
}

func (x *SearchWorkReviewsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SearchWorkReviewsRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *SearchWorkReviewsRequest) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

type SearchReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkReviews map[string]*timestamppb.Timestamp `protobuf:"bytes,1,rep,name=workReviews,proto3" json:"workReviews,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SearchReviewsResponse) Reset() {
	*x = SearchReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReviewsResponse) ProtoMessage() {}

func (x *SearchReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_srv_lib_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReviewsResponse.ProtoReflect.Descriptor instead.
func (*SearchReviewsResponse) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{3}
}

func (x *SearchReviewsResponse) GetWorkReviews() map[string]*timestamppb.Timestamp {
	if x != nil {
		return x.WorkReviews
	}
	return nil
}

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_lib_srv_lib_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{4}
}

var File_lib_srv_lib_proto protoreflect.FileDescriptor

var file_lib_srv_lib_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x69, 0x62, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x58, 0x0a, 0x16,
	0x4d, 0x61, 0x6b, 0x65, 0x41, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69,
	0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x1a, 0x5a, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0x9c, 0x02, 0x0a, 0x0e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x4d,
	0x61, 0x6b, 0x65, 0x41, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x25,
	0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d,
	0x61, 0x6b, 0x65, 0x41, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x67, 0x0a, 0x16, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x12, 0x27, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x4d, 0x61, 0x6b, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_srv_lib_proto_rawDescOnce sync.Once
	file_lib_srv_lib_proto_rawDescData = file_lib_srv_lib_proto_rawDesc
)

func file_lib_srv_lib_proto_rawDescGZIP() []byte {
	file_lib_srv_lib_proto_rawDescOnce.Do(func() {
		file_lib_srv_lib_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_srv_lib_proto_rawDescData)
	})
	return file_lib_srv_lib_proto_rawDescData
}

var file_lib_srv_lib_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lib_srv_lib_proto_goTypes = []interface{}{
	(*ValidatorModeration)(nil),      // 0: pp.contractor.ValidatorModeration
	(*MakeAsPurchasedRequest)(nil),   // 1: pp.contractor.MakeAsPurchasedRequest
	(*SearchWorkReviewsRequest)(nil), // 2: pp.contractor.SearchWorkReviewsRequest
	(*SearchReviewsResponse)(nil),    // 3: pp.contractor.SearchReviewsResponse
	(*Null)(nil),                     // 4: pp.contractor.Null
	nil,                              // 5: pp.contractor.SearchReviewsResponse.WorkReviewsEntry
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_lib_srv_lib_proto_depIdxs = []int32{
	6, // 0: pp.contractor.SearchWorkReviewsRequest.since:type_name -> google.protobuf.Timestamp
	6, // 1: pp.contractor.SearchWorkReviewsRequest.until:type_name -> google.protobuf.Timestamp
	5, // 2: pp.contractor.SearchReviewsResponse.workReviews:type_name -> pp.contractor.SearchReviewsResponse.WorkReviewsEntry
	6, // 3: pp.contractor.SearchReviewsResponse.WorkReviewsEntry.value:type_name -> google.protobuf.Timestamp
	1, // 4: pp.contractor.LibraryService.MakeAsPurchased:input_type -> pp.contractor.MakeAsPurchasedRequest
	2, // 5: pp.contractor.LibraryService.SearchSubmittedReviews:input_type -> pp.contractor.SearchWorkReviewsRequest
	0, // 6: pp.contractor.LibraryService.MakeValidatorModeration:input_type -> pp.contractor.ValidatorModeration
	4, // 7: pp.contractor.LibraryService.MakeAsPurchased:output_type -> pp.contractor.Null
	3, // 8: pp.contractor.LibraryService.SearchSubmittedReviews:output_type -> pp.contractor.SearchReviewsResponse
	4, // 9: pp.contractor.LibraryService.MakeValidatorModeration:output_type -> pp.contractor.Null
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_lib_srv_lib_proto_init() }
func file_lib_srv_lib_proto_init() {
	if File_lib_srv_lib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_srv_lib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorModeration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lib_srv_lib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeAsPurchasedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lib_srv_lib_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchWorkReviewsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lib_srv_lib_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReviewsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lib_srv_lib_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lib_srv_lib_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_srv_lib_proto_goTypes,
		DependencyIndexes: file_lib_srv_lib_proto_depIdxs,
		MessageInfos:      file_lib_srv_lib_proto_msgTypes,
	}.Build()
	File_lib_srv_lib_proto = out.File
	file_lib_srv_lib_proto_rawDesc = nil
	file_lib_srv_lib_proto_goTypes = nil
	file_lib_srv_lib_proto_depIdxs = nil
}

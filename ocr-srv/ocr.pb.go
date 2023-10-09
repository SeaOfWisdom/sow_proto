// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: ocr-srv/ocr.proto

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

type ExtractTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	IsPaper bool   `protobuf:"varint,2,opt,name=isPaper,proto3" json:"isPaper,omitempty"`
}

func (x *ExtractTextRequest) Reset() {
	*x = ExtractTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractTextRequest) ProtoMessage() {}

func (x *ExtractTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractTextRequest.ProtoReflect.Descriptor instead.
func (*ExtractTextRequest) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractTextRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ExtractTextRequest) GetIsPaper() bool {
	if x != nil {
		return x.IsPaper
	}
	return false
}

type LanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *LanguageRequest) Reset() {
	*x = LanguageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageRequest) ProtoMessage() {}

func (x *LanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageRequest.ProtoReflect.Descriptor instead.
func (*LanguageRequest) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{1}
}

func (x *LanguageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type LanguageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LanguageResponse) Reset() {
	*x = LanguageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageResponse) ProtoMessage() {}

func (x *LanguageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageResponse.ProtoReflect.Descriptor instead.
func (*LanguageResponse) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{2}
}

func (x *LanguageResponse) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *LanguageResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ExtractValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontSideImage []byte `protobuf:"bytes,1,opt,name=front_side_image,json=frontSideImage,proto3" json:"front_side_image,omitempty"`
	BackSideImage  []byte `protobuf:"bytes,2,opt,name=back_side_image,json=backSideImage,proto3" json:"back_side_image,omitempty"`
}

func (x *ExtractValidatorRequest) Reset() {
	*x = ExtractValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractValidatorRequest) ProtoMessage() {}

func (x *ExtractValidatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractValidatorRequest.ProtoReflect.Descriptor instead.
func (*ExtractValidatorRequest) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{3}
}

func (x *ExtractValidatorRequest) GetFrontSideImage() []byte {
	if x != nil {
		return x.FrontSideImage
	}
	return nil
}

func (x *ExtractValidatorRequest) GetBackSideImage() []byte {
	if x != nil {
		return x.BackSideImage
	}
	return nil
}

type ExtractValidatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontSideInfo *ValidatorFrontSideInfo `protobuf:"bytes,1,opt,name=front_side_info,json=frontSideInfo,proto3" json:"front_side_info,omitempty"`
	BackSideInfo  *ValidatorBackSideInfo  `protobuf:"bytes,2,opt,name=back_side_info,json=backSideInfo,proto3" json:"back_side_info,omitempty"`
}

func (x *ExtractValidatorResponse) Reset() {
	*x = ExtractValidatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractValidatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractValidatorResponse) ProtoMessage() {}

func (x *ExtractValidatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractValidatorResponse.ProtoReflect.Descriptor instead.
func (*ExtractValidatorResponse) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{4}
}

func (x *ExtractValidatorResponse) GetFrontSideInfo() *ValidatorFrontSideInfo {
	if x != nil {
		return x.FrontSideInfo
	}
	return nil
}

func (x *ExtractValidatorResponse) GetBackSideInfo() *ValidatorBackSideInfo {
	if x != nil {
		return x.BackSideInfo
	}
	return nil
}

type ValidatorBackSideInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Series string                 `protobuf:"bytes,2,opt,name=series,proto3" json:"series,omitempty"`
	Date   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ValidatorBackSideInfo) Reset() {
	*x = ValidatorBackSideInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorBackSideInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorBackSideInfo) ProtoMessage() {}

func (x *ValidatorBackSideInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorBackSideInfo.ProtoReflect.Descriptor instead.
func (*ValidatorBackSideInfo) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{5}
}

func (x *ValidatorBackSideInfo) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *ValidatorBackSideInfo) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *ValidatorBackSideInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ValidatorFrontSideInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number     uint64                 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	LastName   string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName  string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string                 `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Date       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ValidatorFrontSideInfo) Reset() {
	*x = ValidatorFrontSideInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorFrontSideInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorFrontSideInfo) ProtoMessage() {}

func (x *ValidatorFrontSideInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorFrontSideInfo.ProtoReflect.Descriptor instead.
func (*ValidatorFrontSideInfo) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatorFrontSideInfo) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ValidatorFrontSideInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ValidatorFrontSideInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ValidatorFrontSideInfo) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *ValidatorFrontSideInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ExtractTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Authors  []string          `protobuf:"bytes,2,rep,name=authors,proto3" json:"authors,omitempty"`
	Abstract string            `protobuf:"bytes,3,opt,name=abstract,proto3" json:"abstract,omitempty"`
	Keywords []string          `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Main     map[string]string `protobuf:"bytes,5,rep,name=main,proto3" json:"main,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExtractTextResponse) Reset() {
	*x = ExtractTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocr_srv_ocr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractTextResponse) ProtoMessage() {}

func (x *ExtractTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocr_srv_ocr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractTextResponse.ProtoReflect.Descriptor instead.
func (*ExtractTextResponse) Descriptor() ([]byte, []int) {
	return file_ocr_srv_ocr_proto_rawDescGZIP(), []int{7}
}

func (x *ExtractTextResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ExtractTextResponse) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *ExtractTextResponse) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *ExtractTextResponse) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *ExtractTextResponse) GetMain() map[string]string {
	if x != nil {
		return x.Main
	}
	return nil
}

var File_ocr_srv_ocr_proto protoreflect.FileDescriptor

var file_ocr_srv_ocr_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x63, 0x72, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x6f, 0x63, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x63, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x12, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x22,
	0x25, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x6b, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x69,
	0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0xa1, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x63, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x40, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x72, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x77, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xbd, 0x01, 0x0a,
	0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53,
	0x69, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xee, 0x01, 0x0a,
	0x13, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x63, 0x72,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe0, 0x01,
	0x0a, 0x03, 0x4f, 0x43, 0x52, 0x12, 0x42, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x6f, 0x63, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6f, 0x63, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x14, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x1c, 0x2e, 0x6f, 0x63, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6f, 0x63, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x6f, 0x63, 0x72, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x63, 0x72, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ocr_srv_ocr_proto_rawDescOnce sync.Once
	file_ocr_srv_ocr_proto_rawDescData = file_ocr_srv_ocr_proto_rawDesc
)

func file_ocr_srv_ocr_proto_rawDescGZIP() []byte {
	file_ocr_srv_ocr_proto_rawDescOnce.Do(func() {
		file_ocr_srv_ocr_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocr_srv_ocr_proto_rawDescData)
	})
	return file_ocr_srv_ocr_proto_rawDescData
}

var file_ocr_srv_ocr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ocr_srv_ocr_proto_goTypes = []interface{}{
	(*ExtractTextRequest)(nil),       // 0: ocr.ExtractTextRequest
	(*LanguageRequest)(nil),          // 1: ocr.LanguageRequest
	(*LanguageResponse)(nil),         // 2: ocr.LanguageResponse
	(*ExtractValidatorRequest)(nil),  // 3: ocr.ExtractValidatorRequest
	(*ExtractValidatorResponse)(nil), // 4: ocr.ExtractValidatorResponse
	(*ValidatorBackSideInfo)(nil),    // 5: ocr.ValidatorBackSideInfo
	(*ValidatorFrontSideInfo)(nil),   // 6: ocr.ValidatorFrontSideInfo
	(*ExtractTextResponse)(nil),      // 7: ocr.ExtractTextResponse
	nil,                              // 8: ocr.ExtractTextResponse.MainEntry
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
}
var file_ocr_srv_ocr_proto_depIdxs = []int32{
	6, // 0: ocr.ExtractValidatorResponse.front_side_info:type_name -> ocr.ValidatorFrontSideInfo
	5, // 1: ocr.ExtractValidatorResponse.back_side_info:type_name -> ocr.ValidatorBackSideInfo
	9, // 2: ocr.ValidatorBackSideInfo.date:type_name -> google.protobuf.Timestamp
	9, // 3: ocr.ValidatorFrontSideInfo.date:type_name -> google.protobuf.Timestamp
	8, // 4: ocr.ExtractTextResponse.main:type_name -> ocr.ExtractTextResponse.MainEntry
	0, // 5: ocr.OCR.ExtractText:input_type -> ocr.ExtractTextRequest
	3, // 6: ocr.OCR.ExtractValidatorText:input_type -> ocr.ExtractValidatorRequest
	1, // 7: ocr.OCR.ExtractLanguage:input_type -> ocr.LanguageRequest
	7, // 8: ocr.OCR.ExtractText:output_type -> ocr.ExtractTextResponse
	4, // 9: ocr.OCR.ExtractValidatorText:output_type -> ocr.ExtractValidatorResponse
	2, // 10: ocr.OCR.ExtractLanguage:output_type -> ocr.LanguageResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ocr_srv_ocr_proto_init() }
func file_ocr_srv_ocr_proto_init() {
	if File_ocr_srv_ocr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocr_srv_ocr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractTextRequest); i {
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
		file_ocr_srv_ocr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageRequest); i {
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
		file_ocr_srv_ocr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageResponse); i {
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
		file_ocr_srv_ocr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractValidatorRequest); i {
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
		file_ocr_srv_ocr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractValidatorResponse); i {
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
		file_ocr_srv_ocr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorBackSideInfo); i {
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
		file_ocr_srv_ocr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorFrontSideInfo); i {
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
		file_ocr_srv_ocr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractTextResponse); i {
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
			RawDescriptor: file_ocr_srv_ocr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocr_srv_ocr_proto_goTypes,
		DependencyIndexes: file_ocr_srv_ocr_proto_depIdxs,
		MessageInfos:      file_ocr_srv_ocr_proto_msgTypes,
	}.Build()
	File_ocr_srv_ocr_proto = out.File
	file_ocr_srv_ocr_proto_rawDesc = nil
	file_ocr_srv_ocr_proto_goTypes = nil
	file_ocr_srv_ocr_proto_depIdxs = nil
}

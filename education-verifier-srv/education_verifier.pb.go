// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: education-verifier-srv/education_verifier.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName      string `protobuf:"bytes,1,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName     string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName    string `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	DiplomaSeries string `protobuf:"bytes,4,opt,name=diploma_series,json=diplomaSeries,proto3" json:"diploma_series,omitempty"`
	DiplomaNumber string `protobuf:"bytes,5,opt,name=diploma_number,json=diplomaNumber,proto3" json:"diploma_number,omitempty"`
	Country       string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_education_verifier_srv_education_verifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_education_verifier_srv_education_verifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_education_verifier_srv_education_verifier_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *VerifyRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *VerifyRequest) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *VerifyRequest) GetDiplomaSeries() string {
	if x != nil {
		return x.DiplomaSeries
	}
	return ""
}

func (x *VerifyRequest) GetDiplomaNumber() string {
	if x != nil {
		return x.DiplomaNumber
	}
	return ""
}

func (x *VerifyRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool   `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Science string `protobuf:"bytes,3,opt,name=science,proto3" json:"science,omitempty"`
	Code    string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Link    string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_education_verifier_srv_education_verifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_education_verifier_srv_education_verifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_education_verifier_srv_education_verifier_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *VerifyResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VerifyResponse) GetScience() string {
	if x != nil {
		return x.Science
	}
	return ""
}

func (x *VerifyResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_education_verifier_srv_education_verifier_proto protoreflect.FileDescriptor

var file_education_verifier_srv_education_verifier_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x22, 0xd4, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d,
	0x61, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x61, 0x0a, 0x18, 0x45,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x12, 0x1c, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_education_verifier_srv_education_verifier_proto_rawDescOnce sync.Once
	file_education_verifier_srv_education_verifier_proto_rawDescData = file_education_verifier_srv_education_verifier_proto_rawDesc
)

func file_education_verifier_srv_education_verifier_proto_rawDescGZIP() []byte {
	file_education_verifier_srv_education_verifier_proto_rawDescOnce.Do(func() {
		file_education_verifier_srv_education_verifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_education_verifier_srv_education_verifier_proto_rawDescData)
	})
	return file_education_verifier_srv_education_verifier_proto_rawDescData
}

var file_education_verifier_srv_education_verifier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_education_verifier_srv_education_verifier_proto_goTypes = []interface{}{
	(*VerifyRequest)(nil),  // 0: pp.contractor.VerifyRequest
	(*VerifyResponse)(nil), // 1: pp.contractor.VerifyResponse
}
var file_education_verifier_srv_education_verifier_proto_depIdxs = []int32{
	0, // 0: pp.contractor.EducationVerifierService.Verify:input_type -> pp.contractor.VerifyRequest
	1, // 1: pp.contractor.EducationVerifierService.Verify:output_type -> pp.contractor.VerifyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_education_verifier_srv_education_verifier_proto_init() }
func file_education_verifier_srv_education_verifier_proto_init() {
	if File_education_verifier_srv_education_verifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_education_verifier_srv_education_verifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_education_verifier_srv_education_verifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
			RawDescriptor: file_education_verifier_srv_education_verifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_education_verifier_srv_education_verifier_proto_goTypes,
		DependencyIndexes: file_education_verifier_srv_education_verifier_proto_depIdxs,
		MessageInfos:      file_education_verifier_srv_education_verifier_proto_msgTypes,
	}.Build()
	File_education_verifier_srv_education_verifier_proto = out.File
	file_education_verifier_srv_education_verifier_proto_rawDesc = nil
	file_education_verifier_srv_education_verifier_proto_goTypes = nil
	file_education_verifier_srv_education_verifier_proto_depIdxs = nil
}
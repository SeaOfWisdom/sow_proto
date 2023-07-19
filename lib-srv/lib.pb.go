// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: lib-srv/lib.proto

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

type PutWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReaderAddress string `protobuf:"bytes,1,opt,name=reader_address,json=readerAddress,proto3" json:"reader_address,omitempty"`
	WorkId        string `protobuf:"bytes,2,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
}

func (x *PutWorkRequest) Reset() {
	*x = PutWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutWorkRequest) ProtoMessage() {}

func (x *PutWorkRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PutWorkRequest.ProtoReflect.Descriptor instead.
func (*PutWorkRequest) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{0}
}

func (x *PutWorkRequest) GetReaderAddress() string {
	if x != nil {
		return x.ReaderAddress
	}
	return ""
}

func (x *PutWorkRequest) GetWorkId() string {
	if x != nil {
		return x.WorkId
	}
	return ""
}

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_srv_lib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_lib_srv_lib_proto_rawDescGZIP(), []int{1}
}

var File_lib_srv_lib_proto protoreflect.FileDescriptor

var file_lib_srv_lib_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x69, 0x62, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x50, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0x4f, 0x0a, 0x0e,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x07, 0x50, 0x75, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x2e, 0x70, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_lib_srv_lib_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lib_srv_lib_proto_goTypes = []interface{}{
	(*PutWorkRequest)(nil), // 0: pp.contractor.PutWorkRequest
	(*Null)(nil),           // 1: pp.contractor.Null
}
var file_lib_srv_lib_proto_depIdxs = []int32{
	0, // 0: pp.contractor.LibraryService.PutWork:input_type -> pp.contractor.PutWorkRequest
	1, // 1: pp.contractor.LibraryService.PutWork:output_type -> pp.contractor.Null
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lib_srv_lib_proto_init() }
func file_lib_srv_lib_proto_init() {
	if File_lib_srv_lib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_srv_lib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutWorkRequest); i {
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
			NumMessages:   2,
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: jwt-srv/jwt.proto

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

type TokenBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iss         string `protobuf:"bytes,1,opt,name=iss,proto3" json:"iss,omitempty"`
	Sub         string `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Role        int64  `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
	Web3Address string `protobuf:"bytes,4,opt,name=web3_address,json=web3Address,proto3" json:"web3_address,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Language    string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *TokenBody) Reset() {
	*x = TokenBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_srv_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBody) ProtoMessage() {}

func (x *TokenBody) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_srv_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBody.ProtoReflect.Descriptor instead.
func (*TokenBody) Descriptor() ([]byte, []int) {
	return file_jwt_srv_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *TokenBody) GetIss() string {
	if x != nil {
		return x.Iss
	}
	return ""
}

func (x *TokenBody) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *TokenBody) GetRole() int64 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *TokenBody) GetWeb3Address() string {
	if x != nil {
		return x.Web3Address
	}
	return ""
}

func (x *TokenBody) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TokenBody) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_srv_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_srv_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_jwt_srv_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DecodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool       `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Token string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Body  *TokenBody `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DecodeResp) Reset() {
	*x = DecodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_srv_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResp) ProtoMessage() {}

func (x *DecodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_srv_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResp.ProtoReflect.Descriptor instead.
func (*DecodeResp) Descriptor() ([]byte, []int) {
	return file_jwt_srv_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeResp) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *DecodeResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DecodeResp) GetBody() *TokenBody {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_jwt_srv_jwt_proto protoreflect.FileDescriptor

var file_jwt_srv_jwt_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x77, 0x74, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x09,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x33, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x65, 0x62, 0x33, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x6d, 0x0a, 0x0a, 0x4a, 0x77, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4a, 0x57, 0x54, 0x12, 0x11, 0x2e, 0x70, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x70, 0x2e, 0x6a, 0x77,
	0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x4a, 0x57, 0x54, 0x12, 0x0d, 0x2e, 0x70, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x12, 0x2e, 0x70, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jwt_srv_jwt_proto_rawDescOnce sync.Once
	file_jwt_srv_jwt_proto_rawDescData = file_jwt_srv_jwt_proto_rawDesc
)

func file_jwt_srv_jwt_proto_rawDescGZIP() []byte {
	file_jwt_srv_jwt_proto_rawDescOnce.Do(func() {
		file_jwt_srv_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_jwt_srv_jwt_proto_rawDescData)
	})
	return file_jwt_srv_jwt_proto_rawDescData
}

var file_jwt_srv_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_jwt_srv_jwt_proto_goTypes = []interface{}{
	(*TokenBody)(nil),  // 0: pp.jwt.TokenBody
	(*Token)(nil),      // 1: pp.jwt.Token
	(*DecodeResp)(nil), // 2: pp.jwt.DecodeResp
}
var file_jwt_srv_jwt_proto_depIdxs = []int32{
	0, // 0: pp.jwt.DecodeResp.body:type_name -> pp.jwt.TokenBody
	0, // 1: pp.jwt.JwtService.GenerateJWT:input_type -> pp.jwt.TokenBody
	1, // 2: pp.jwt.JwtService.DecodeJWT:input_type -> pp.jwt.Token
	1, // 3: pp.jwt.JwtService.GenerateJWT:output_type -> pp.jwt.Token
	2, // 4: pp.jwt.JwtService.DecodeJWT:output_type -> pp.jwt.DecodeResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_jwt_srv_jwt_proto_init() }
func file_jwt_srv_jwt_proto_init() {
	if File_jwt_srv_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jwt_srv_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBody); i {
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
		file_jwt_srv_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_jwt_srv_jwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResp); i {
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
			RawDescriptor: file_jwt_srv_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jwt_srv_jwt_proto_goTypes,
		DependencyIndexes: file_jwt_srv_jwt_proto_depIdxs,
		MessageInfos:      file_jwt_srv_jwt_proto_msgTypes,
	}.Build()
	File_jwt_srv_jwt_proto = out.File
	file_jwt_srv_jwt_proto_rawDesc = nil
	file_jwt_srv_jwt_proto_goTypes = nil
	file_jwt_srv_jwt_proto_depIdxs = nil
}

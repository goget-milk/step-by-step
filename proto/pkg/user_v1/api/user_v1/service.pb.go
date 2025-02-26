// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: api/user_v1/service.proto

package user_v1

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsHuman bool   `protobuf:"varint,3,opt,name=is_human,json=isHuman,proto3" json:"is_human,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_api_user_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetIsHuman() bool {
	if x != nil {
		return x.IsHuman
	}
	return false
}

var File_api_user_v1_service_proto protoreflect.FileDescriptor

var file_api_user_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x67, 0x65, 0x74, 0x2d, 0x6d, 0x69, 0x6c, 0x6b, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x2d, 0x62, 0x79,
	0x2d, 0x73, 0x74, 0x65, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_service_proto_rawDescOnce sync.Once
	file_api_user_v1_service_proto_rawDescData = file_api_user_v1_service_proto_rawDesc
)

func file_api_user_v1_service_proto_rawDescGZIP() []byte {
	file_api_user_v1_service_proto_rawDescOnce.Do(func() {
		file_api_user_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_service_proto_rawDescData)
	})
	return file_api_user_v1_service_proto_rawDescData
}

var file_api_user_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_user_v1_service_proto_goTypes = []any{
	(*UserInfo)(nil), // 0: user_v1.UserInfo
}
var file_api_user_v1_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_service_proto_init() }
func file_api_user_v1_service_proto_init() {
	if File_api_user_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_service_proto_goTypes,
		DependencyIndexes: file_api_user_v1_service_proto_depIdxs,
		MessageInfos:      file_api_user_v1_service_proto_msgTypes,
	}.Build()
	File_api_user_v1_service_proto = out.File
	file_api_user_v1_service_proto_rawDesc = nil
	file_api_user_v1_service_proto_goTypes = nil
	file_api_user_v1_service_proto_depIdxs = nil
}

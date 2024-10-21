// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/v1/echo.proto

package protov1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	mi := &file_proto_v1_echo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_echo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_proto_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Description) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	mi := &file_proto_v1_echo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_echo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	mi := &file_proto_v1_echo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_echo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_echo_proto_rawDescGZIP(), []int{2}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var file_proto_v1_echo_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Description)(nil),
		Field:         50000,
		Name:          "echo.description",
		Tag:           "bytes,50000,opt,name=description",
		Filename:      "proto/v1/echo.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional echo.Description description = 50000;
	E_Description = &file_proto_v1_echo_proto_extTypes[0]
)

var File_proto_v1_echo_proto protoreflect.FileDescriptor

var file_proto_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x82, 0xb5,
	0x18, 0x11, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x65,
	0x63, 0x68, 0x6f, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0c,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x82,
	0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x45, 0x63, 0x68, 0x6f, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x51, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a,
	0x54, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x76, 0x63, 0x68, 0x61, 0x72, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_echo_proto_rawDescOnce sync.Once
	file_proto_v1_echo_proto_rawDescData = file_proto_v1_echo_proto_rawDesc
)

func file_proto_v1_echo_proto_rawDescGZIP() []byte {
	file_proto_v1_echo_proto_rawDescOnce.Do(func() {
		file_proto_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_echo_proto_rawDescData)
	})
	return file_proto_v1_echo_proto_rawDescData
}

var file_proto_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_v1_echo_proto_goTypes = []any{
	(*Description)(nil),               // 0: echo.Description
	(*EchoRequest)(nil),               // 1: echo.EchoRequest
	(*EchoResponse)(nil),              // 2: echo.EchoResponse
	(*descriptorpb.FieldOptions)(nil), // 3: google.protobuf.FieldOptions
}
var file_proto_v1_echo_proto_depIdxs = []int32{
	3, // 0: echo.description:extendee -> google.protobuf.FieldOptions
	0, // 1: echo.description:type_name -> echo.Description
	1, // 2: echo.EchoService.Echo:input_type -> echo.EchoRequest
	2, // 3: echo.EchoService.Echo:output_type -> echo.EchoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_echo_proto_init() }
func file_proto_v1_echo_proto_init() {
	if File_proto_v1_echo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_echo_proto_goTypes,
		DependencyIndexes: file_proto_v1_echo_proto_depIdxs,
		MessageInfos:      file_proto_v1_echo_proto_msgTypes,
		ExtensionInfos:    file_proto_v1_echo_proto_extTypes,
	}.Build()
	File_proto_v1_echo_proto = out.File
	file_proto_v1_echo_proto_rawDesc = nil
	file_proto_v1_echo_proto_goTypes = nil
	file_proto_v1_echo_proto_depIdxs = nil
}

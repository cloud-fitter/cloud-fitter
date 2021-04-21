// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: idl/ecs/ecs.proto

package pbecs

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ECSInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName  string   `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	PublicIps []string `protobuf:"bytes,2,rep,name=public_ips,json=publicIps,proto3" json:"public_ips,omitempty"`
}

func (x *ECSInstance) Reset() {
	*x = ECSInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ecs_ecs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECSInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECSInstance) ProtoMessage() {}

func (x *ECSInstance) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ecs_ecs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECSInstance.ProtoReflect.Descriptor instead.
func (*ECSInstance) Descriptor() ([]byte, []int) {
	return file_idl_ecs_ecs_proto_rawDescGZIP(), []int{0}
}

func (x *ECSInstance) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *ECSInstance) GetPublicIps() []string {
	if x != nil {
		return x.PublicIps
	}
	return nil
}

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ecs_ecs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ecs_ecs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_idl_ecs_ecs_proto_rawDescGZIP(), []int{1}
}

func (x *StringMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_idl_ecs_ecs_proto protoreflect.FileDescriptor

var file_idl_ecs_ecs_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x65, 0x63, 0x73, 0x2f, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x65, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0b, 0x45, 0x43, 0x53, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x70, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x55, 0x0a, 0x0a, 0x45, 0x43,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x65, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x65, 0x63, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x3a, 0x01,
	0x2a, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x62, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_ecs_ecs_proto_rawDescOnce sync.Once
	file_idl_ecs_ecs_proto_rawDescData = file_idl_ecs_ecs_proto_rawDesc
)

func file_idl_ecs_ecs_proto_rawDescGZIP() []byte {
	file_idl_ecs_ecs_proto_rawDescOnce.Do(func() {
		file_idl_ecs_ecs_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_ecs_ecs_proto_rawDescData)
	})
	return file_idl_ecs_ecs_proto_rawDescData
}

var file_idl_ecs_ecs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idl_ecs_ecs_proto_goTypes = []interface{}{
	(*ECSInstance)(nil),   // 0: pbecs.ECSInstance
	(*StringMessage)(nil), // 1: pbecs.StringMessage
}
var file_idl_ecs_ecs_proto_depIdxs = []int32{
	1, // 0: pbecs.ECSService.Echo:input_type -> pbecs.StringMessage
	1, // 1: pbecs.ECSService.Echo:output_type -> pbecs.StringMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_ecs_ecs_proto_init() }
func file_idl_ecs_ecs_proto_init() {
	if File_idl_ecs_ecs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_ecs_ecs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECSInstance); i {
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
		file_idl_ecs_ecs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
			RawDescriptor: file_idl_ecs_ecs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_ecs_ecs_proto_goTypes,
		DependencyIndexes: file_idl_ecs_ecs_proto_depIdxs,
		MessageInfos:      file_idl_ecs_ecs_proto_msgTypes,
	}.Build()
	File_idl_ecs_ecs_proto = out.File
	file_idl_ecs_ecs_proto_rawDesc = nil
	file_idl_ecs_ecs_proto_goTypes = nil
	file_idl_ecs_ecs_proto_depIdxs = nil
}

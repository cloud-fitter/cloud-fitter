// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: idl/pboss/oss.proto

package pboss

import (
	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
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

type OssInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云类型
	Provider pbtenant.CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
	// 账号名称
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// bucket名称
	BucketName string `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	// 区域
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *OssInstance) Reset() {
	*x = OssInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pboss_oss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OssInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OssInstance) ProtoMessage() {}

func (x *OssInstance) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pboss_oss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OssInstance.ProtoReflect.Descriptor instead.
func (*OssInstance) Descriptor() ([]byte, []int) {
	return file_idl_pboss_oss_proto_rawDescGZIP(), []int{0}
}

func (x *OssInstance) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali
}

func (x *OssInstance) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *OssInstance) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *OssInstance) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ListDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云名称
	Provider pbtenant.CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
	// 账户名称，根据config.yaml中的配置，默认为第一个配置的账户
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// 分页相关参数，页码
	PageNumber int32 `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// 分页相关参数，每页数量
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 分页相关参数，下一页的token
	NextToken string `protobuf:"bytes,5,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *ListDetailReq) Reset() {
	*x = ListDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pboss_oss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailReq) ProtoMessage() {}

func (x *ListDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pboss_oss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDetailReq.ProtoReflect.Descriptor instead.
func (*ListDetailReq) Descriptor() ([]byte, []int) {
	return file_idl_pboss_oss_proto_rawDescGZIP(), []int{1}
}

func (x *ListDetailReq) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali
}

func (x *ListDetailReq) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *ListDetailReq) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListDetailReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDetailReq) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type ListDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Oss 机器集合
	Osses []*OssInstance `protobuf:"bytes,1,rep,name=osses,proto3" json:"osses,omitempty"`
	// 查询是否完成，如果为否-false，则可以将下面三个分页参数填入到请求中，继续查询
	Finished bool `protobuf:"varint,2,opt,name=finished,proto3" json:"finished,omitempty"`
	// 分页相关参数，页码
	PageNumber int32 `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// 分页相关参数，每页数量
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 分页相关参数，下一页的token
	NextToken string `protobuf:"bytes,5,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	// 请求id，出现问题后提供给云厂商，排查问题
	RequestId string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *ListDetailResp) Reset() {
	*x = ListDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pboss_oss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailResp) ProtoMessage() {}

func (x *ListDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pboss_oss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDetailResp.ProtoReflect.Descriptor instead.
func (*ListDetailResp) Descriptor() ([]byte, []int) {
	return file_idl_pboss_oss_proto_rawDescGZIP(), []int{2}
}

func (x *ListDetailResp) GetOsses() []*OssInstance {
	if x != nil {
		return x.Osses
	}
	return nil
}

func (x *ListDetailResp) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *ListDetailResp) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListDetailResp) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDetailResp) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *ListDetailResp) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云名称
	Provider pbtenant.CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pboss_oss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pboss_oss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_idl_pboss_oss_proto_rawDescGZIP(), []int{3}
}

func (x *ListReq) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Oss 机器集合
	Osses []*OssInstance `protobuf:"bytes,1,rep,name=osses,proto3" json:"osses,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pboss_oss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pboss_oss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_idl_pboss_oss_proto_rawDescGZIP(), []int{4}
}

func (x *ListResp) GetOsses() []*OssInstance {
	if x != nil {
		return x.Osses
	}
	return nil
}

var File_idl_pboss_oss_proto protoreflect.FileDescriptor

var file_idl_pboss_oss_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x1a, 0x19, 0x69, 0x64,
	0x6c, 0x2f, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x4f, 0x73, 0x73, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4f, 0x73, 0x73, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4f, 0x73, 0x73, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x32, 0xa9, 0x01, 0x0a,
	0x0a, 0x4f, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x40, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x73,
	0x73, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6f, 0x73, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x6f, 0x73, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_pboss_oss_proto_rawDescOnce sync.Once
	file_idl_pboss_oss_proto_rawDescData = file_idl_pboss_oss_proto_rawDesc
)

func file_idl_pboss_oss_proto_rawDescGZIP() []byte {
	file_idl_pboss_oss_proto_rawDescOnce.Do(func() {
		file_idl_pboss_oss_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_pboss_oss_proto_rawDescData)
	})
	return file_idl_pboss_oss_proto_rawDescData
}

var file_idl_pboss_oss_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_pboss_oss_proto_goTypes = []interface{}{
	(*OssInstance)(nil),         // 0: pboss.OssInstance
	(*ListDetailReq)(nil),       // 1: pboss.ListDetailReq
	(*ListDetailResp)(nil),      // 2: pboss.ListDetailResp
	(*ListReq)(nil),             // 3: pboss.ListReq
	(*ListResp)(nil),            // 4: pboss.ListResp
	(pbtenant.CloudProvider)(0), // 5: pbtenant.CloudProvider
}
var file_idl_pboss_oss_proto_depIdxs = []int32{
	5, // 0: pboss.OssInstance.provider:type_name -> pbtenant.CloudProvider
	5, // 1: pboss.ListDetailReq.provider:type_name -> pbtenant.CloudProvider
	0, // 2: pboss.ListDetailResp.osses:type_name -> pboss.OssInstance
	5, // 3: pboss.ListReq.provider:type_name -> pbtenant.CloudProvider
	0, // 4: pboss.ListResp.osses:type_name -> pboss.OssInstance
	1, // 5: pboss.OssService.ListOssDetail:input_type -> pboss.ListDetailReq
	3, // 6: pboss.OssService.ListOss:input_type -> pboss.ListReq
	2, // 7: pboss.OssService.ListOssDetail:output_type -> pboss.ListDetailResp
	4, // 8: pboss.OssService.ListOss:output_type -> pboss.ListResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_idl_pboss_oss_proto_init() }
func file_idl_pboss_oss_proto_init() {
	if File_idl_pboss_oss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_pboss_oss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OssInstance); i {
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
		file_idl_pboss_oss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDetailReq); i {
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
		file_idl_pboss_oss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDetailResp); i {
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
		file_idl_pboss_oss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_idl_pboss_oss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
			RawDescriptor: file_idl_pboss_oss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_pboss_oss_proto_goTypes,
		DependencyIndexes: file_idl_pboss_oss_proto_depIdxs,
		MessageInfos:      file_idl_pboss_oss_proto_msgTypes,
	}.Build()
	File_idl_pboss_oss_proto = out.File
	file_idl_pboss_oss_proto_rawDesc = nil
	file_idl_pboss_oss_proto_goTypes = nil
	file_idl_pboss_oss_proto_depIdxs = nil
}

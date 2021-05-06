// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: idl/pbrds/rds.proto

package pbrds

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

type RDSInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云类型
	Provider pbtenant.CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
	// 账号名称
	AccoutName string `protobuf:"bytes,2,opt,name=accout_name,json=accoutName,proto3" json:"accout_name,omitempty"`
	// 实例id
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// 实例名称
	InstanceName string `protobuf:"bytes,4,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// 地域，数据中心
	RegionName string `protobuf:"bytes,5,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	// 实例类型
	InstanceType string `protobuf:"bytes,6,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// 数据库类型
	Engine string `protobuf:"bytes,7,opt,name=engine,proto3" json:"engine,omitempty"`
	// 数据库类型的版本
	EngineVersion string `protobuf:"bytes,8,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	// 数据库实例规格
	InstanceClass string `protobuf:"bytes,9,opt,name=instance_class,json=instanceClass,proto3" json:"instance_class,omitempty"`
	// 状态
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// 创建时间，ISO8601
	CreationTime string `protobuf:"bytes,11,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	// 过期时间
	ExpireTime string `protobuf:"bytes,12,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *RDSInstance) Reset() {
	*x = RDSInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbrds_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RDSInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RDSInstance) ProtoMessage() {}

func (x *RDSInstance) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbrds_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RDSInstance.ProtoReflect.Descriptor instead.
func (*RDSInstance) Descriptor() ([]byte, []int) {
	return file_idl_pbrds_rds_proto_rawDescGZIP(), []int{0}
}

func (x *RDSInstance) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali_cloud
}

func (x *RDSInstance) GetAccoutName() string {
	if x != nil {
		return x.AccoutName
	}
	return ""
}

func (x *RDSInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *RDSInstance) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *RDSInstance) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *RDSInstance) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *RDSInstance) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *RDSInstance) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *RDSInstance) GetInstanceClass() string {
	if x != nil {
		return x.InstanceClass
	}
	return ""
}

func (x *RDSInstance) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RDSInstance) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *RDSInstance) GetExpireTime() string {
	if x != nil {
		return x.ExpireTime
	}
	return ""
}

type ListDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云名称
	Provider pbtenant.CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
	// 账户名称，根据统计接口返回值，或直接根据config.yaml中的配置
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// 区域Id，参考 tenant.proto 中的各个云的区域
	RegionId int32 `protobuf:"varint,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// 分页相关参数，页码
	PageNumber int32 `protobuf:"varint,4,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// 分页相关参数，每页数量
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 分页相关参数，下一页的token
	NextToken string `protobuf:"bytes,6,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *ListDetailReq) Reset() {
	*x = ListDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbrds_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailReq) ProtoMessage() {}

func (x *ListDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbrds_rds_proto_msgTypes[1]
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
	return file_idl_pbrds_rds_proto_rawDescGZIP(), []int{1}
}

func (x *ListDetailReq) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali_cloud
}

func (x *ListDetailReq) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *ListDetailReq) GetRegionId() int32 {
	if x != nil {
		return x.RegionId
	}
	return 0
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

	// RDS 机器集合
	Rdses []*RDSInstance `protobuf:"bytes,1,rep,name=rdses,proto3" json:"rdses,omitempty"`
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
		mi := &file_idl_pbrds_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailResp) ProtoMessage() {}

func (x *ListDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbrds_rds_proto_msgTypes[2]
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
	return file_idl_pbrds_rds_proto_rawDescGZIP(), []int{2}
}

func (x *ListDetailResp) GetRdses() []*RDSInstance {
	if x != nil {
		return x.Rdses
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
		mi := &file_idl_pbrds_rds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbrds_rds_proto_msgTypes[3]
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
	return file_idl_pbrds_rds_proto_rawDescGZIP(), []int{3}
}

func (x *ListReq) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali_cloud
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RDS 机器集合
	Rdses []*RDSInstance `protobuf:"bytes,1,rep,name=rdses,proto3" json:"rdses,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbrds_rds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbrds_rds_proto_msgTypes[4]
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
	return file_idl_pbrds_rds_proto_rawDescGZIP(), []int{4}
}

func (x *ListResp) GetRdses() []*RDSInstance {
	if x != nil {
		return x.Rdses
	}
	return nil
}

var File_idl_pbrds_rds_proto protoreflect.FileDescriptor

var file_idl_pbrds_rds_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x72, 0x64, 0x73, 0x2f, 0x72, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x72, 0x64, 0x73, 0x1a, 0x19, 0x69, 0x64,
	0x6c, 0x2f, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x03, 0x0a, 0x0b, 0x52, 0x44, 0x53, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xd2, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x64, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x72, 0x64, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x28, 0x0a, 0x05, 0x72, 0x64, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x05, 0x72, 0x64, 0x73, 0x65, 0x73, 0x32, 0x9f, 0x01, 0x0a, 0x0a, 0x52,
	0x44, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x44, 0x53, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x72,
	0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x72, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22,
	0x0b, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12,
	0x3b, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x44, 0x53, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x72,
	0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x72,
	0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x72, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x72,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_pbrds_rds_proto_rawDescOnce sync.Once
	file_idl_pbrds_rds_proto_rawDescData = file_idl_pbrds_rds_proto_rawDesc
)

func file_idl_pbrds_rds_proto_rawDescGZIP() []byte {
	file_idl_pbrds_rds_proto_rawDescOnce.Do(func() {
		file_idl_pbrds_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_pbrds_rds_proto_rawDescData)
	})
	return file_idl_pbrds_rds_proto_rawDescData
}

var file_idl_pbrds_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_pbrds_rds_proto_goTypes = []interface{}{
	(*RDSInstance)(nil),         // 0: pbrds.RDSInstance
	(*ListDetailReq)(nil),       // 1: pbrds.ListDetailReq
	(*ListDetailResp)(nil),      // 2: pbrds.ListDetailResp
	(*ListReq)(nil),             // 3: pbrds.ListReq
	(*ListResp)(nil),            // 4: pbrds.ListResp
	(pbtenant.CloudProvider)(0), // 5: pbtenant.CloudProvider
}
var file_idl_pbrds_rds_proto_depIdxs = []int32{
	5, // 0: pbrds.RDSInstance.provider:type_name -> pbtenant.CloudProvider
	5, // 1: pbrds.ListDetailReq.provider:type_name -> pbtenant.CloudProvider
	0, // 2: pbrds.ListDetailResp.rdses:type_name -> pbrds.RDSInstance
	5, // 3: pbrds.ListReq.provider:type_name -> pbtenant.CloudProvider
	0, // 4: pbrds.ListResp.rdses:type_name -> pbrds.RDSInstance
	1, // 5: pbrds.RDSService.ListRDSDetail:input_type -> pbrds.ListDetailReq
	3, // 6: pbrds.RDSService.ListRDS:input_type -> pbrds.ListReq
	2, // 7: pbrds.RDSService.ListRDSDetail:output_type -> pbrds.ListDetailResp
	4, // 8: pbrds.RDSService.ListRDS:output_type -> pbrds.ListResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_idl_pbrds_rds_proto_init() }
func file_idl_pbrds_rds_proto_init() {
	if File_idl_pbrds_rds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_pbrds_rds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RDSInstance); i {
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
		file_idl_pbrds_rds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbrds_rds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbrds_rds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbrds_rds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_idl_pbrds_rds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_pbrds_rds_proto_goTypes,
		DependencyIndexes: file_idl_pbrds_rds_proto_depIdxs,
		MessageInfos:      file_idl_pbrds_rds_proto_msgTypes,
	}.Build()
	File_idl_pbrds_rds_proto = out.File
	file_idl_pbrds_rds_proto_rawDesc = nil
	file_idl_pbrds_rds_proto_goTypes = nil
	file_idl_pbrds_rds_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: idl/pbkafka/kafka.proto

package pbkafka

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

type KafkaInstance struct {
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
	// 默认接入点
	EndPoint string `protobuf:"bytes,6,opt,name=end_point,json=endPoint,proto3" json:"end_point,omitempty"`
	// Topic最大数量
	TopicNumLimit int32 `protobuf:"varint,7,opt,name=topic_num_limit,json=topicNumLimit,proto3" json:"topic_num_limit,omitempty"`
	// 磁盘容量
	DistSize int32 `protobuf:"varint,8,opt,name=dist_size,json=distSize,proto3" json:"dist_size,omitempty"`
	// 状态
	Status string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// 创建时间，ISO8601
	CreateTime string `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 过期时间
	ExpiredTime string `protobuf:"bytes,11,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time,omitempty"`
}

func (x *KafkaInstance) Reset() {
	*x = KafkaInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbkafka_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaInstance) ProtoMessage() {}

func (x *KafkaInstance) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbkafka_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaInstance.ProtoReflect.Descriptor instead.
func (*KafkaInstance) Descriptor() ([]byte, []int) {
	return file_idl_pbkafka_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *KafkaInstance) GetProvider() pbtenant.CloudProvider {
	if x != nil {
		return x.Provider
	}
	return pbtenant.CloudProvider_ali
}

func (x *KafkaInstance) GetAccoutName() string {
	if x != nil {
		return x.AccoutName
	}
	return ""
}

func (x *KafkaInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *KafkaInstance) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *KafkaInstance) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *KafkaInstance) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *KafkaInstance) GetTopicNumLimit() int32 {
	if x != nil {
		return x.TopicNumLimit
	}
	return 0
}

func (x *KafkaInstance) GetDistSize() int32 {
	if x != nil {
		return x.DistSize
	}
	return 0
}

func (x *KafkaInstance) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *KafkaInstance) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *KafkaInstance) GetExpiredTime() string {
	if x != nil {
		return x.ExpiredTime
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
		mi := &file_idl_pbkafka_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailReq) ProtoMessage() {}

func (x *ListDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbkafka_kafka_proto_msgTypes[1]
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
	return file_idl_pbkafka_kafka_proto_rawDescGZIP(), []int{1}
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

	// kafka 机器集合
	Kafkas []*KafkaInstance `protobuf:"bytes,1,rep,name=kafkas,proto3" json:"kafkas,omitempty"`
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
		mi := &file_idl_pbkafka_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetailResp) ProtoMessage() {}

func (x *ListDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbkafka_kafka_proto_msgTypes[2]
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
	return file_idl_pbkafka_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *ListDetailResp) GetKafkas() []*KafkaInstance {
	if x != nil {
		return x.Kafkas
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
		mi := &file_idl_pbkafka_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbkafka_kafka_proto_msgTypes[3]
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
	return file_idl_pbkafka_kafka_proto_rawDescGZIP(), []int{3}
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

	// Kafka 机器集合
	Kafkas []*KafkaInstance `protobuf:"bytes,1,rep,name=kafkas,proto3" json:"kafkas,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbkafka_kafka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbkafka_kafka_proto_msgTypes[4]
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
	return file_idl_pbkafka_kafka_proto_rawDescGZIP(), []int{4}
}

func (x *ListResp) GetKafkas() []*KafkaInstance {
	if x != nil {
		return x.Kafkas
	}
	return nil
}

var File_idl_pbkafka_kafka_proto protoreflect.FileDescriptor

var file_idl_pbkafka_kafka_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x1a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0d,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd8, 0x01, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2e, 0x0a, 0x06, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x06, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x73, 0x32, 0xbb, 0x01, 0x0a, 0x0c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x48, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x22, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x3a, 0x01,
	0x2a, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64,
	0x6c, 0x2f, 0x70, 0x62, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_idl_pbkafka_kafka_proto_rawDescOnce sync.Once
	file_idl_pbkafka_kafka_proto_rawDescData = file_idl_pbkafka_kafka_proto_rawDesc
)

func file_idl_pbkafka_kafka_proto_rawDescGZIP() []byte {
	file_idl_pbkafka_kafka_proto_rawDescOnce.Do(func() {
		file_idl_pbkafka_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_pbkafka_kafka_proto_rawDescData)
	})
	return file_idl_pbkafka_kafka_proto_rawDescData
}

var file_idl_pbkafka_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_pbkafka_kafka_proto_goTypes = []interface{}{
	(*KafkaInstance)(nil),       // 0: pbkafka.KafkaInstance
	(*ListDetailReq)(nil),       // 1: pbkafka.ListDetailReq
	(*ListDetailResp)(nil),      // 2: pbkafka.ListDetailResp
	(*ListReq)(nil),             // 3: pbkafka.ListReq
	(*ListResp)(nil),            // 4: pbkafka.ListResp
	(pbtenant.CloudProvider)(0), // 5: pbtenant.CloudProvider
}
var file_idl_pbkafka_kafka_proto_depIdxs = []int32{
	5, // 0: pbkafka.KafkaInstance.provider:type_name -> pbtenant.CloudProvider
	5, // 1: pbkafka.ListDetailReq.provider:type_name -> pbtenant.CloudProvider
	0, // 2: pbkafka.ListDetailResp.kafkas:type_name -> pbkafka.KafkaInstance
	5, // 3: pbkafka.ListReq.provider:type_name -> pbtenant.CloudProvider
	0, // 4: pbkafka.ListResp.kafkas:type_name -> pbkafka.KafkaInstance
	1, // 5: pbkafka.KafkaService.ListKafkaDetail:input_type -> pbkafka.ListDetailReq
	3, // 6: pbkafka.KafkaService.ListKafka:input_type -> pbkafka.ListReq
	2, // 7: pbkafka.KafkaService.ListKafkaDetail:output_type -> pbkafka.ListDetailResp
	4, // 8: pbkafka.KafkaService.ListKafka:output_type -> pbkafka.ListResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_idl_pbkafka_kafka_proto_init() }
func file_idl_pbkafka_kafka_proto_init() {
	if File_idl_pbkafka_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_pbkafka_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaInstance); i {
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
		file_idl_pbkafka_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbkafka_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbkafka_kafka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_pbkafka_kafka_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_idl_pbkafka_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_pbkafka_kafka_proto_goTypes,
		DependencyIndexes: file_idl_pbkafka_kafka_proto_depIdxs,
		MessageInfos:      file_idl_pbkafka_kafka_proto_msgTypes,
	}.Build()
	File_idl_pbkafka_kafka_proto = out.File
	file_idl_pbkafka_kafka_proto_rawDesc = nil
	file_idl_pbkafka_kafka_proto_goTypes = nil
	file_idl_pbkafka_kafka_proto_depIdxs = nil
}

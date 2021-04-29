// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: idl/pbtenant/tenant.proto

package pbtenant

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type CloudProvider int32

const (
	CloudProvider_ali_cloud     CloudProvider = 0
	CloudProvider_tencent_cloud CloudProvider = 1
	CloudProvider_huawei_cloud  CloudProvider = 2
	CloudProvider_aws_cloud     CloudProvider = 3
	// 暂不支持
	CloudProvider_google_cloud CloudProvider = 4
)

// Enum value maps for CloudProvider.
var (
	CloudProvider_name = map[int32]string{
		0: "ali_cloud",
		1: "tencent_cloud",
		2: "huawei_cloud",
		3: "aws_cloud",
		4: "google_cloud",
	}
	CloudProvider_value = map[string]int32{
		"ali_cloud":     0,
		"tencent_cloud": 1,
		"huawei_cloud":  2,
		"aws_cloud":     3,
		"google_cloud":  4,
	}
)

func (x CloudProvider) Enum() *CloudProvider {
	p := new(CloudProvider)
	*p = x
	return p
}

func (x CloudProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[0].Descriptor()
}

func (CloudProvider) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[0]
}

func (x CloudProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudProvider.Descriptor instead.
func (CloudProvider) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{0}
}

type CloudProduct int32

const (
	// 所有产品
	CloudProduct_product_all CloudProduct = 0
	// ECS类产品：阿里云ECS,腾讯云CVM，华为云ECS，亚马逊EC2
	CloudProduct_product_ecs CloudProduct = 1
)

// Enum value maps for CloudProduct.
var (
	CloudProduct_name = map[int32]string{
		0: "product_all",
		1: "product_ecs",
	}
	CloudProduct_value = map[string]int32{
		"product_all": 0,
		"product_ecs": 1,
	}
)

func (x CloudProduct) Enum() *CloudProduct {
	p := new(CloudProduct)
	*p = x
	return p
}

func (x CloudProduct) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudProduct) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[1].Descriptor()
}

func (CloudProduct) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[1]
}

func (x CloudProduct) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudProduct.Descriptor instead.
func (CloudProduct) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{1}
}

// 阿里云区域，需要将对应的 _ 转化为 -
type AliRegionId int32

const (
	AliRegionId_ali_all            AliRegionId = 0
	AliRegionId_ali_cn_qingdao     AliRegionId = 1  // 青岛
	AliRegionId_ali_cn_beijing     AliRegionId = 2  // 北京
	AliRegionId_ali_cn_zhangjiakou AliRegionId = 3  // 张家口
	AliRegionId_ali_cn_huhehaote   AliRegionId = 4  // 呼和浩特
	AliRegionId_ali_cn_wulanchabu  AliRegionId = 5  // 乌兰察布
	AliRegionId_ali_cn_hangzhou    AliRegionId = 6  // 杭州
	AliRegionId_ali_cn_shanghai    AliRegionId = 7  // 上海
	AliRegionId_ali_cn_shenzhen    AliRegionId = 8  // 深圳
	AliRegionId_ali_cn_heyuan      AliRegionId = 9  // 河源
	AliRegionId_ali_cn_guangzhou   AliRegionId = 10 // 广州
	AliRegionId_ali_cn_chengdu     AliRegionId = 11 // 成都
	AliRegionId_ali_cn_hongkong    AliRegionId = 12 // 中国香港-香港
	AliRegionId_ali_ap_southeast_1 AliRegionId = 13 // 亚太东南1-新加坡
	AliRegionId_ali_ap_southeast_2 AliRegionId = 14 // 亚太东南2-悉尼
	AliRegionId_ali_ap_southeast_3 AliRegionId = 15 // 亚太东南3-吉隆坡
	AliRegionId_ali_ap_southeast_5 AliRegionId = 16 // 亚太东南5-雅加达
	AliRegionId_ali_ap_south_1     AliRegionId = 17 // 亚太南部1-孟买
	AliRegionId_ali_ap_northeast_1 AliRegionId = 18 // 亚太东北1-东京
	AliRegionId_ali_us_west_1      AliRegionId = 19 // 美国西部1-硅谷
	AliRegionId_ali_us_east_1      AliRegionId = 20 // 美国东部1-弗吉尼亚
	AliRegionId_ali_eu_central_1   AliRegionId = 21 // 欧洲中部1-法兰克福
	AliRegionId_ali_eu_west_1      AliRegionId = 22 // 英国（伦敦）-伦敦
	AliRegionId_ali_me_east_1      AliRegionId = 23 // 中东东部1-迪拜
)

// Enum value maps for AliRegionId.
var (
	AliRegionId_name = map[int32]string{
		0:  "ali_all",
		1:  "ali_cn_qingdao",
		2:  "ali_cn_beijing",
		3:  "ali_cn_zhangjiakou",
		4:  "ali_cn_huhehaote",
		5:  "ali_cn_wulanchabu",
		6:  "ali_cn_hangzhou",
		7:  "ali_cn_shanghai",
		8:  "ali_cn_shenzhen",
		9:  "ali_cn_heyuan",
		10: "ali_cn_guangzhou",
		11: "ali_cn_chengdu",
		12: "ali_cn_hongkong",
		13: "ali_ap_southeast_1",
		14: "ali_ap_southeast_2",
		15: "ali_ap_southeast_3",
		16: "ali_ap_southeast_5",
		17: "ali_ap_south_1",
		18: "ali_ap_northeast_1",
		19: "ali_us_west_1",
		20: "ali_us_east_1",
		21: "ali_eu_central_1",
		22: "ali_eu_west_1",
		23: "ali_me_east_1",
	}
	AliRegionId_value = map[string]int32{
		"ali_all":            0,
		"ali_cn_qingdao":     1,
		"ali_cn_beijing":     2,
		"ali_cn_zhangjiakou": 3,
		"ali_cn_huhehaote":   4,
		"ali_cn_wulanchabu":  5,
		"ali_cn_hangzhou":    6,
		"ali_cn_shanghai":    7,
		"ali_cn_shenzhen":    8,
		"ali_cn_heyuan":      9,
		"ali_cn_guangzhou":   10,
		"ali_cn_chengdu":     11,
		"ali_cn_hongkong":    12,
		"ali_ap_southeast_1": 13,
		"ali_ap_southeast_2": 14,
		"ali_ap_southeast_3": 15,
		"ali_ap_southeast_5": 16,
		"ali_ap_south_1":     17,
		"ali_ap_northeast_1": 18,
		"ali_us_west_1":      19,
		"ali_us_east_1":      20,
		"ali_eu_central_1":   21,
		"ali_eu_west_1":      22,
		"ali_me_east_1":      23,
	}
)

func (x AliRegionId) Enum() *AliRegionId {
	p := new(AliRegionId)
	*p = x
	return p
}

func (x AliRegionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AliRegionId) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[2].Descriptor()
}

func (AliRegionId) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[2]
}

func (x AliRegionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AliRegionId.Descriptor instead.
func (AliRegionId) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{2}
}

// 阿里云区域，需要将对应的 _ 转化为 -
type TencentRegionId int32

const (
	TencentRegionId_tc_all               TencentRegionId = 0
	TencentRegionId_tc_ap_bangkok        TencentRegionId = 1  // 曼谷
	TencentRegionId_tc_ap_beijing        TencentRegionId = 2  // 北京
	TencentRegionId_tc_ap_chengdu        TencentRegionId = 3  // 成都
	TencentRegionId_tc_ap_chongqing      TencentRegionId = 4  // 重庆
	TencentRegionId_tc_ap_guangzhou      TencentRegionId = 5  // 广州
	TencentRegionId_tc_ap_guangzhou_open TencentRegionId = 6  // 广州Open
	TencentRegionId_tc_ap_hongkong       TencentRegionId = 7  // 中国香港
	TencentRegionId_tc_ap_mumbai         TencentRegionId = 8  // 孟买
	TencentRegionId_tc_ap_seoul          TencentRegionId = 9  // 首尔
	TencentRegionId_tc_ap_shanghai       TencentRegionId = 10 // 上海
	TencentRegionId_tc_ap_shanghai_fsi   TencentRegionId = 11 // 上海金融
	TencentRegionId_tc_ap_shenzhen_fsi   TencentRegionId = 12 // 深圳金融
	TencentRegionId_tc_ap_singapore      TencentRegionId = 13 // 新加坡
	TencentRegionId_tc_ap_tokyo          TencentRegionId = 14 // 东京
	TencentRegionId_tc_eu_frankfurt      TencentRegionId = 15 // 法兰克福
	TencentRegionId_tc_eu_moscow         TencentRegionId = 16 // 莫斯科
	TencentRegionId_tc_na_ashburn        TencentRegionId = 17 // 阿什本
	TencentRegionId_tc_na_siliconvalley  TencentRegionId = 18 // 硅谷
	TencentRegionId_tc_na_toronto        TencentRegionId = 19 // 多伦多
)

// Enum value maps for TencentRegionId.
var (
	TencentRegionId_name = map[int32]string{
		0:  "tc_all",
		1:  "tc_ap_bangkok",
		2:  "tc_ap_beijing",
		3:  "tc_ap_chengdu",
		4:  "tc_ap_chongqing",
		5:  "tc_ap_guangzhou",
		6:  "tc_ap_guangzhou_open",
		7:  "tc_ap_hongkong",
		8:  "tc_ap_mumbai",
		9:  "tc_ap_seoul",
		10: "tc_ap_shanghai",
		11: "tc_ap_shanghai_fsi",
		12: "tc_ap_shenzhen_fsi",
		13: "tc_ap_singapore",
		14: "tc_ap_tokyo",
		15: "tc_eu_frankfurt",
		16: "tc_eu_moscow",
		17: "tc_na_ashburn",
		18: "tc_na_siliconvalley",
		19: "tc_na_toronto",
	}
	TencentRegionId_value = map[string]int32{
		"tc_all":               0,
		"tc_ap_bangkok":        1,
		"tc_ap_beijing":        2,
		"tc_ap_chengdu":        3,
		"tc_ap_chongqing":      4,
		"tc_ap_guangzhou":      5,
		"tc_ap_guangzhou_open": 6,
		"tc_ap_hongkong":       7,
		"tc_ap_mumbai":         8,
		"tc_ap_seoul":          9,
		"tc_ap_shanghai":       10,
		"tc_ap_shanghai_fsi":   11,
		"tc_ap_shenzhen_fsi":   12,
		"tc_ap_singapore":      13,
		"tc_ap_tokyo":          14,
		"tc_eu_frankfurt":      15,
		"tc_eu_moscow":         16,
		"tc_na_ashburn":        17,
		"tc_na_siliconvalley":  18,
		"tc_na_toronto":        19,
	}
)

func (x TencentRegionId) Enum() *TencentRegionId {
	p := new(TencentRegionId)
	*p = x
	return p
}

func (x TencentRegionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TencentRegionId) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[3].Descriptor()
}

func (TencentRegionId) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[3]
}

func (x TencentRegionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TencentRegionId.Descriptor instead.
func (TencentRegionId) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{3}
}

// 华为云区域，需要将对应的 _ 转化为 -
type HuaweiRegionId int32

const (
	HuaweiRegionId_hw_all            HuaweiRegionId = 0
	HuaweiRegionId_hw_cn_north_1     HuaweiRegionId = 1
	HuaweiRegionId_hw_cn_north_4     HuaweiRegionId = 2
	HuaweiRegionId_hw_cn_south_1     HuaweiRegionId = 3
	HuaweiRegionId_hw_cn_east_2      HuaweiRegionId = 4
	HuaweiRegionId_hw_cn_east_3      HuaweiRegionId = 5
	HuaweiRegionId_hw_cn_southwest_2 HuaweiRegionId = 6
	HuaweiRegionId_hw_ap_southeast_1 HuaweiRegionId = 7
	HuaweiRegionId_hw_ap_southeast_2 HuaweiRegionId = 8
	HuaweiRegionId_hw_ap_southeast_3 HuaweiRegionId = 9
	HuaweiRegionId_hw_af_south_1     HuaweiRegionId = 10
)

// Enum value maps for HuaweiRegionId.
var (
	HuaweiRegionId_name = map[int32]string{
		0:  "hw_all",
		1:  "hw_cn_north_1",
		2:  "hw_cn_north_4",
		3:  "hw_cn_south_1",
		4:  "hw_cn_east_2",
		5:  "hw_cn_east_3",
		6:  "hw_cn_southwest_2",
		7:  "hw_ap_southeast_1",
		8:  "hw_ap_southeast_2",
		9:  "hw_ap_southeast_3",
		10: "hw_af_south_1",
	}
	HuaweiRegionId_value = map[string]int32{
		"hw_all":            0,
		"hw_cn_north_1":     1,
		"hw_cn_north_4":     2,
		"hw_cn_south_1":     3,
		"hw_cn_east_2":      4,
		"hw_cn_east_3":      5,
		"hw_cn_southwest_2": 6,
		"hw_ap_southeast_1": 7,
		"hw_ap_southeast_2": 8,
		"hw_ap_southeast_3": 9,
		"hw_af_south_1":     10,
	}
)

func (x HuaweiRegionId) Enum() *HuaweiRegionId {
	p := new(HuaweiRegionId)
	*p = x
	return p
}

func (x HuaweiRegionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HuaweiRegionId) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[4].Descriptor()
}

func (HuaweiRegionId) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[4]
}

func (x HuaweiRegionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HuaweiRegionId.Descriptor instead.
func (HuaweiRegionId) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{4}
}

// 亚马逊云区域，需要将对应的 _ 转化为 -
type AwsRegionId int32

const (
	AwsRegionId_aws_all            AwsRegionId = 0
	AwsRegionId_aws_us_east_2      AwsRegionId = 1  // US East (Ohio)
	AwsRegionId_aws_us_east_1      AwsRegionId = 2  // US East (N. Virginia)
	AwsRegionId_aws_us_west_1      AwsRegionId = 3  // US West (N. California)
	AwsRegionId_aws_us_west_2      AwsRegionId = 4  // US West (Oregon)
	AwsRegionId_aws_af_south_1     AwsRegionId = 5  // Africa (Cape Town)
	AwsRegionId_aws_ap_east_1      AwsRegionId = 6  // Asia Pacific (Hong Kong)
	AwsRegionId_aws_ap_south_1     AwsRegionId = 7  // Asia Pacific (Mumbai)
	AwsRegionId_aws_ap_northeast_3 AwsRegionId = 8  // Asia Pacific (Osaka)
	AwsRegionId_aws_ap_northeast_2 AwsRegionId = 9  // Asia Pacific (Seoul)
	AwsRegionId_aws_ap_northeast_1 AwsRegionId = 10 // Asia Pacific (Singapore)
	AwsRegionId_aws_ap_southeast_2 AwsRegionId = 11 // Asia Pacific (Sydney)
	AwsRegionId_aws_ap_southeast_1 AwsRegionId = 12 // Asia Pacific (Tokyo)
	AwsRegionId_aws_ca_central_1   AwsRegionId = 13 // Canada (Central)
	AwsRegionId_aws_eu_central_1   AwsRegionId = 14 // Europe (Frankfurt)
	AwsRegionId_aws_eu_west_1      AwsRegionId = 15 // Europe (Ireland)
	AwsRegionId_aws_eu_west_2      AwsRegionId = 16 // Europe (London)
	AwsRegionId_aws_eu_south_1     AwsRegionId = 17 // Europe (Milan)
	AwsRegionId_aws_eu_west_3      AwsRegionId = 18 // Europe (Paris)
	AwsRegionId_aws_eu_north_1     AwsRegionId = 19 // Europe (Stockholm)
	AwsRegionId_aws_me_south_1     AwsRegionId = 20 // Middle East (Bahrain)
	AwsRegionId_aws_sa_east_1      AwsRegionId = 21 // South America (São Paulo)
)

// Enum value maps for AwsRegionId.
var (
	AwsRegionId_name = map[int32]string{
		0:  "aws_all",
		1:  "aws_us_east_2",
		2:  "aws_us_east_1",
		3:  "aws_us_west_1",
		4:  "aws_us_west_2",
		5:  "aws_af_south_1",
		6:  "aws_ap_east_1",
		7:  "aws_ap_south_1",
		8:  "aws_ap_northeast_3",
		9:  "aws_ap_northeast_2",
		10: "aws_ap_northeast_1",
		11: "aws_ap_southeast_2",
		12: "aws_ap_southeast_1",
		13: "aws_ca_central_1",
		14: "aws_eu_central_1",
		15: "aws_eu_west_1",
		16: "aws_eu_west_2",
		17: "aws_eu_south_1",
		18: "aws_eu_west_3",
		19: "aws_eu_north_1",
		20: "aws_me_south_1",
		21: "aws_sa_east_1",
	}
	AwsRegionId_value = map[string]int32{
		"aws_all":            0,
		"aws_us_east_2":      1,
		"aws_us_east_1":      2,
		"aws_us_west_1":      3,
		"aws_us_west_2":      4,
		"aws_af_south_1":     5,
		"aws_ap_east_1":      6,
		"aws_ap_south_1":     7,
		"aws_ap_northeast_3": 8,
		"aws_ap_northeast_2": 9,
		"aws_ap_northeast_1": 10,
		"aws_ap_southeast_2": 11,
		"aws_ap_southeast_1": 12,
		"aws_ca_central_1":   13,
		"aws_eu_central_1":   14,
		"aws_eu_west_1":      15,
		"aws_eu_west_2":      16,
		"aws_eu_south_1":     17,
		"aws_eu_west_3":      18,
		"aws_eu_north_1":     19,
		"aws_me_south_1":     20,
		"aws_sa_east_1":      21,
	}
)

func (x AwsRegionId) Enum() *AwsRegionId {
	p := new(AwsRegionId)
	*p = x
	return p
}

func (x AwsRegionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AwsRegionId) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_pbtenant_tenant_proto_enumTypes[5].Descriptor()
}

func (AwsRegionId) Type() protoreflect.EnumType {
	return &file_idl_pbtenant_tenant_proto_enumTypes[5]
}

func (x AwsRegionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AwsRegionId.Descriptor instead.
func (AwsRegionId) EnumDescriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{5}
}

// 云配置信息
type CloudConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云配置
	Configs []*CloudConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *CloudConfigs) Reset() {
	*x = CloudConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbtenant_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudConfigs) ProtoMessage() {}

func (x *CloudConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbtenant_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudConfigs.ProtoReflect.Descriptor instead.
func (*CloudConfigs) Descriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *CloudConfigs) GetConfigs() []*CloudConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

type CloudConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 云服务提供商，具体参考 CloudProvider 的定义
	Provider CloudProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=pbtenant.CloudProvider" json:"provider,omitempty"`
	// 账户名称，由用户自定义，必须全局唯一，方便多个系统之间的维护
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 认证方式1：与 access_secret 结合使用，两者均非空时生效
	AccessId string `protobuf:"bytes,3,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	// 认证方式1：与 access_id 结合使用，两者均非空时生效
	AccessSecret string `protobuf:"bytes,4,opt,name=access_secret,json=accessSecret,proto3" json:"access_secret,omitempty"`
}

func (x *CloudConfig) Reset() {
	*x = CloudConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pbtenant_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudConfig) ProtoMessage() {}

func (x *CloudConfig) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pbtenant_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudConfig.ProtoReflect.Descriptor instead.
func (*CloudConfig) Descriptor() ([]byte, []int) {
	return file_idl_pbtenant_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *CloudConfig) GetProvider() CloudProvider {
	if x != nil {
		return x.Provider
	}
	return CloudProvider_ali_cloud
}

func (x *CloudConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudConfig) GetAccessId() string {
	if x != nil {
		return x.AccessId
	}
	return ""
}

func (x *CloudConfig) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

var File_idl_pbtenant_tenant_proto protoreflect.FileDescriptor

var file_idl_pbtenant_tenant_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2a,
	0x64, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x0d, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x61, 0x77, 0x73, 0x5f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x10, 0x04, 0x2a, 0x30, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x61, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x65, 0x63, 0x73, 0x10, 0x01, 0x2a, 0x86, 0x04, 0x0a, 0x0b, 0x41, 0x6c, 0x69, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x5f, 0x61,
	0x6c, 0x6c, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x71,
	0x69, 0x6e, 0x67, 0x64, 0x61, 0x6f, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x6c, 0x69, 0x5f,
	0x63, 0x6e, 0x5f, 0x62, 0x65, 0x69, 0x6a, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x6a, 0x69, 0x61, 0x6b,
	0x6f, 0x75, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x68,
	0x75, 0x68, 0x65, 0x68, 0x61, 0x6f, 0x74, 0x65, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x61, 0x6c,
	0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x61, 0x62, 0x75, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x68, 0x61, 0x6e, 0x67,
	0x7a, 0x68, 0x6f, 0x75, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e,
	0x5f, 0x73, 0x68, 0x61, 0x6e, 0x67, 0x68, 0x61, 0x69, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x61,
	0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x73, 0x68, 0x65, 0x6e, 0x7a, 0x68, 0x65, 0x6e, 0x10, 0x08,
	0x12, 0x11, 0x0a, 0x0d, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x68, 0x65, 0x79, 0x75, 0x61,
	0x6e, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x67, 0x75,
	0x61, 0x6e, 0x67, 0x7a, 0x68, 0x6f, 0x75, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x6c, 0x69,
	0x5f, 0x63, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x64, 0x75, 0x10, 0x0b, 0x12, 0x13, 0x0a,
	0x0f, 0x61, 0x6c, 0x69, 0x5f, 0x63, 0x6e, 0x5f, 0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67,
	0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x6c, 0x69, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75,
	0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x0d, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x6c,
	0x69, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x32,
	0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x6c, 0x69, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75,
	0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x33, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x6c,
	0x69, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x35,
	0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x6c, 0x69, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75,
	0x74, 0x68, 0x5f, 0x31, 0x10, 0x11, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x6c, 0x69, 0x5f, 0x61, 0x70,
	0x5f, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x12, 0x12, 0x11,
	0x0a, 0x0d, 0x61, 0x6c, 0x69, 0x5f, 0x75, 0x73, 0x5f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x31, 0x10,
	0x13, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x6c, 0x69, 0x5f, 0x75, 0x73, 0x5f, 0x65, 0x61, 0x73, 0x74,
	0x5f, 0x31, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x61, 0x6c, 0x69, 0x5f, 0x65, 0x75, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x31, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x6c,
	0x69, 0x5f, 0x65, 0x75, 0x5f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x16, 0x12, 0x11, 0x0a,
	0x0d, 0x61, 0x6c, 0x69, 0x5f, 0x6d, 0x65, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x17,
	0x2a, 0xa1, 0x03, 0x0a, 0x0f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x63, 0x5f, 0x61, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x62, 0x61, 0x6e, 0x67, 0x6b, 0x6f,
	0x6b, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x62, 0x65, 0x69,
	0x6a, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f,
	0x63, 0x68, 0x65, 0x6e, 0x67, 0x64, 0x75, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x74, 0x63, 0x5f,
	0x61, 0x70, 0x5f, 0x63, 0x68, 0x6f, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x7a, 0x68, 0x6f,
	0x75, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x67, 0x75, 0x61,
	0x6e, 0x67, 0x7a, 0x68, 0x6f, 0x75, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x10, 0x06, 0x12, 0x12, 0x0a,
	0x0e, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x10,
	0x07, 0x12, 0x10, 0x0a, 0x0c, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x6d, 0x75, 0x6d, 0x62, 0x61,
	0x69, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x6f,
	0x75, 0x6c, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x68,
	0x61, 0x6e, 0x67, 0x68, 0x61, 0x69, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x74, 0x63, 0x5f, 0x61,
	0x70, 0x5f, 0x73, 0x68, 0x61, 0x6e, 0x67, 0x68, 0x61, 0x69, 0x5f, 0x66, 0x73, 0x69, 0x10, 0x0b,
	0x12, 0x16, 0x0a, 0x12, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x68, 0x65, 0x6e, 0x7a, 0x68,
	0x65, 0x6e, 0x5f, 0x66, 0x73, 0x69, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x74, 0x63, 0x5f, 0x61,
	0x70, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65, 0x10, 0x0d, 0x12, 0x0f, 0x0a,
	0x0b, 0x74, 0x63, 0x5f, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x79, 0x6f, 0x10, 0x0e, 0x12, 0x13,
	0x0a, 0x0f, 0x74, 0x63, 0x5f, 0x65, 0x75, 0x5f, 0x66, 0x72, 0x61, 0x6e, 0x6b, 0x66, 0x75, 0x72,
	0x74, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x74, 0x63, 0x5f, 0x65, 0x75, 0x5f, 0x6d, 0x6f, 0x73,
	0x63, 0x6f, 0x77, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x74, 0x63, 0x5f, 0x6e, 0x61, 0x5f, 0x61,
	0x73, 0x68, 0x62, 0x75, 0x72, 0x6e, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x74, 0x63, 0x5f, 0x6e,
	0x61, 0x5f, 0x73, 0x69, 0x6c, 0x69, 0x63, 0x6f, 0x6e, 0x76, 0x61, 0x6c, 0x6c, 0x65, 0x79, 0x10,
	0x12, 0x12, 0x11, 0x0a, 0x0d, 0x74, 0x63, 0x5f, 0x6e, 0x61, 0x5f, 0x74, 0x6f, 0x72, 0x6f, 0x6e,
	0x74, 0x6f, 0x10, 0x13, 0x2a, 0xe8, 0x01, 0x0a, 0x0e, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x68, 0x77, 0x5f, 0x61, 0x6c,
	0x6c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x68, 0x77, 0x5f, 0x63, 0x6e, 0x5f, 0x6e, 0x6f, 0x72,
	0x74, 0x68, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x68, 0x77, 0x5f, 0x63, 0x6e, 0x5f,
	0x6e, 0x6f, 0x72, 0x74, 0x68, 0x5f, 0x34, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x68, 0x77, 0x5f,
	0x63, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x68, 0x77, 0x5f, 0x63, 0x6e, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x32, 0x10, 0x04, 0x12, 0x10,
	0x0a, 0x0c, 0x68, 0x77, 0x5f, 0x63, 0x6e, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x33, 0x10, 0x05,
	0x12, 0x15, 0x0a, 0x11, 0x68, 0x77, 0x5f, 0x63, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x77,
	0x65, 0x73, 0x74, 0x5f, 0x32, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x68, 0x77, 0x5f, 0x61, 0x70,
	0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x07, 0x12, 0x15,
	0x0a, 0x11, 0x68, 0x77, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73,
	0x74, 0x5f, 0x32, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x68, 0x77, 0x5f, 0x61, 0x70, 0x5f, 0x73,
	0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x33, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d,
	0x68, 0x77, 0x5f, 0x61, 0x66, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x10, 0x0a, 0x2a,
	0xcd, 0x03, 0x0a, 0x0b, 0x41, 0x77, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x0b, 0x0a, 0x07, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x61, 0x77, 0x73, 0x5f, 0x75, 0x73, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x32, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x61, 0x77, 0x73, 0x5f, 0x75, 0x73, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x77, 0x73, 0x5f, 0x75, 0x73, 0x5f, 0x77, 0x65, 0x73,
	0x74, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x77, 0x73, 0x5f, 0x75, 0x73, 0x5f,
	0x77, 0x65, 0x73, 0x74, 0x5f, 0x32, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f,
	0x61, 0x66, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d,
	0x61, 0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x06, 0x12,
	0x12, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f,
	0x31, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x6e, 0x6f,
	0x72, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x33, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x61,
	0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f,
	0x32, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x6e, 0x6f,
	0x72, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x61,
	0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f,
	0x32, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x70, 0x5f, 0x73, 0x6f,
	0x75, 0x74, 0x68, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x61,
	0x77, 0x73, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x31, 0x10,
	0x0d, 0x12, 0x14, 0x0a, 0x10, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x75, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x31, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x77, 0x73, 0x5f, 0x65,
	0x75, 0x5f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x77,
	0x73, 0x5f, 0x65, 0x75, 0x5f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x32, 0x10, 0x10, 0x12, 0x12, 0x0a,
	0x0e, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x75, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x10,
	0x11, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x75, 0x5f, 0x77, 0x65, 0x73, 0x74,
	0x5f, 0x33, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x75, 0x5f, 0x6e,
	0x6f, 0x72, 0x74, 0x68, 0x5f, 0x31, 0x10, 0x13, 0x12, 0x12, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f,
	0x6d, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d,
	0x61, 0x77, 0x73, 0x5f, 0x73, 0x61, 0x5f, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x31, 0x10, 0x15, 0x32,
	0x85, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x74, 0x92, 0x41, 0x71, 0x12, 0x1e, 0xe6, 0x89, 0x80, 0xe6, 0x9c, 0x89, 0xe4, 0xba,
	0x91, 0xe7, 0xa7, 0x9f, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe8, 0xae, 0xa4, 0xe8, 0xaf, 0x81,
	0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x1a, 0x4f, 0x0a, 0x1f, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f,
	0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x46, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x66, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x62, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_pbtenant_tenant_proto_rawDescOnce sync.Once
	file_idl_pbtenant_tenant_proto_rawDescData = file_idl_pbtenant_tenant_proto_rawDesc
)

func file_idl_pbtenant_tenant_proto_rawDescGZIP() []byte {
	file_idl_pbtenant_tenant_proto_rawDescOnce.Do(func() {
		file_idl_pbtenant_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_pbtenant_tenant_proto_rawDescData)
	})
	return file_idl_pbtenant_tenant_proto_rawDescData
}

var file_idl_pbtenant_tenant_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_idl_pbtenant_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idl_pbtenant_tenant_proto_goTypes = []interface{}{
	(CloudProvider)(0),   // 0: pbtenant.CloudProvider
	(CloudProduct)(0),    // 1: pbtenant.CloudProduct
	(AliRegionId)(0),     // 2: pbtenant.AliRegionId
	(TencentRegionId)(0), // 3: pbtenant.TencentRegionId
	(HuaweiRegionId)(0),  // 4: pbtenant.HuaweiRegionId
	(AwsRegionId)(0),     // 5: pbtenant.AwsRegionId
	(*CloudConfigs)(nil), // 6: pbtenant.CloudConfigs
	(*CloudConfig)(nil),  // 7: pbtenant.CloudConfig
}
var file_idl_pbtenant_tenant_proto_depIdxs = []int32{
	7, // 0: pbtenant.CloudConfigs.configs:type_name -> pbtenant.CloudConfig
	0, // 1: pbtenant.CloudConfig.provider:type_name -> pbtenant.CloudProvider
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_idl_pbtenant_tenant_proto_init() }
func file_idl_pbtenant_tenant_proto_init() {
	if File_idl_pbtenant_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_pbtenant_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudConfigs); i {
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
		file_idl_pbtenant_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudConfig); i {
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
			RawDescriptor: file_idl_pbtenant_tenant_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_pbtenant_tenant_proto_goTypes,
		DependencyIndexes: file_idl_pbtenant_tenant_proto_depIdxs,
		EnumInfos:         file_idl_pbtenant_tenant_proto_enumTypes,
		MessageInfos:      file_idl_pbtenant_tenant_proto_msgTypes,
	}.Build()
	File_idl_pbtenant_tenant_proto = out.File
	file_idl_pbtenant_tenant_proto_rawDesc = nil
	file_idl_pbtenant_tenant_proto_goTypes = nil
	file_idl_pbtenant_tenant_proto_depIdxs = nil
}

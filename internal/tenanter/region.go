package tenanter

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

var (
	ErrNoExistAliRegionId     = errors.New("region id not exist in ali")
	ErrNoExistTencentRegionId = errors.New("region id not exist in tencent")
	ErrNoExistHuaweiRegionId  = errors.New("region id not exist in huawei")
	ErrNoExistAwsRegionId     = errors.New("region id not exist in aws")
)

type Region interface {
	GetId() int32
	GetName() string
}

type region struct {
	provider   pbtenant.CloudProvider
	regionId   int32
	regionName string
}

func NewRegion(provider pbtenant.CloudProvider, regionId int32) (Region, error) {
	r := &region{
		provider: provider,
		regionId: regionId,
	}
	var err error

	switch provider {
	case pbtenant.CloudProvider_ali_cloud:
		r.regionName, err = getAliRegionName(regionId)
	case pbtenant.CloudProvider_tencent_cloud:
		r.regionName, err = getTencentRegionName(regionId)
	case pbtenant.CloudProvider_huawei_cloud:
		r.regionName, err = getHuaweiRegionName(regionId)
	case pbtenant.CloudProvider_aws_cloud:
		r.regionName, err = getAwsRegionName(regionId)
	}

	return r, err
}

func (r *region) GetName() string {
	return r.regionName
}

func (r *region) GetId() int32 {
	return r.regionId
}

func GetAllRegionIds(provider pbtenant.CloudProvider) (regions []Region) {
	switch provider {
	case pbtenant.CloudProvider_ali_cloud:
		for rId := range pbtenant.AliRegionId_name {
			if rId != int32(pbtenant.AliRegionId_ali_all) {
				region, _ := NewRegion(provider, rId)
				regions = append(regions, region)
			}
		}
	case pbtenant.CloudProvider_tencent_cloud:
		for rId := range pbtenant.TencentRegionId_name {
			if rId != int32(pbtenant.TencentRegionId_tc_all) {
				region, _ := NewRegion(provider, rId)
				regions = append(regions, region)
			}
		}
	case pbtenant.CloudProvider_huawei_cloud:
		for rId := range pbtenant.HuaweiRegionId_name {
			if rId != int32(pbtenant.HuaweiRegionId_hw_all) {
				region, _ := NewRegion(provider, rId)
				regions = append(regions, region)
			}
		}
	case pbtenant.CloudProvider_aws_cloud:
		for rId := range pbtenant.AwsRegionId_name {
			if rId != int32(pbtenant.AwsRegionId_aws_all) {
				region, _ := NewRegion(provider, rId)
				regions = append(regions, region)
			}
		}
	}

	return
}

// prefix ali_
func getAliRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.AliRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.AliRegionId_ali_all) {
		return "", errors.WithMessagef(ErrNoExistAliRegionId, "input region id is %d", regionId)
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[4:], nil
}

// prefix tc_
func getTencentRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.TencentRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.TencentRegionId_tc_all) {
		return "", errors.WithMessagef(ErrNoExistTencentRegionId, "input region id is %d", regionId)
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[3:], nil
}

// prefix hw_
func getHuaweiRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.HuaweiRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.HuaweiRegionId_hw_all) {
		return "", errors.WithMessagef(ErrNoExistHuaweiRegionId, "input region id is %d", regionId)
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[3:], nil
}

// prefix aws_
func getAwsRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.AwsRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.AwsRegionId_aws_all) {
		return "", errors.WithMessagef(ErrNoExistAwsRegionId, "input region id is %d", regionId)
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[4:], nil
}

package tenanter

import (
	"fmt"
	"strings"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/pkg/errors"
)

var (
	ErrNoExistAliRegionId     = errors.New("region id not exist in ali")
	ErrNoExistTencentRegionId = errors.New("region id not exist in tencent")
	ErrNoExistHuaweiRegionId  = errors.New("region id not exist in huawei")
	ErrNoExistAwsRegionId     = errors.New("region id not exist in aws")
)

// prefix ali_
func GetAliRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.AliRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.AliRegionId_ali_all) {
		return "", errors.WithMessage(ErrNoExistAliRegionId, fmt.Sprintf("input region id is %d", regionId))
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[4:], nil
}

// prefix tc_
func GetTencentRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.TencentRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.TencentRegionId_tc_all) {
		return "", errors.WithMessage(ErrNoExistTencentRegionId, fmt.Sprintf("input region id is %d", regionId))
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[3:], nil
}

// prefix hw_
func GetHuaweiRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.HuaweiRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.HuaweiRegionId_hw_all) {
		return "", errors.WithMessage(ErrNoExistHuaweiRegionId, fmt.Sprintf("input region id is %d", regionId))
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[3:], nil
}

// prefix aws_
func GetAwsRegionName(regionId int32) (string, error) {
	name, ok := pbtenant.AwsRegionId_name[regionId]
	if !ok || regionId == int32(pbtenant.AwsRegionId_aws_all) {
		return "", errors.WithMessage(ErrNoExistAwsRegionId, fmt.Sprintf("input region id is %d", regionId))
	}
	region := strings.ReplaceAll(name, "_", "-")
	return region[4:], nil
}

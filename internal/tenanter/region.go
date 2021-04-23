package tenanter

import (
	"strings"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/pkg/errors"
)

var ErrNoExistAliyunRegionId = errors.New("region id not exist in aliyun")

// prefix ali_
func GetAliRegionId(regionId pbtenant.AliRegionId) string {
	region := strings.ReplaceAll(regionId.String(), "_", "-")
	return region[4:]
}

// prefix tc_
func GetTencentRegionId(regionId pbtenant.TencentRegionId) string {
	region := strings.ReplaceAll(regionId.String(), "_", "-")
	return region[3:]
}

// prefix hw_
func GetHuaweiRegionId(regionId pbtenant.HuaweiRegionId) string {
	region := strings.ReplaceAll(regionId.String(), "_", "-")
	return region[3:]
}

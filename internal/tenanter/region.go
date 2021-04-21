package tenanter

import (
	"strings"

	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/tenant"
	"github.com/pkg/errors"
)

var ErrNoExistAliyunRegionId = errors.New("region id not exist in aliyun")

func GetAliRegionId(regionId pbtenant.AliRegionId) string {
	return strings.ReplaceAll(regionId.String(), "_", "-")
}

func GetTencentRegionId(regionId pbtenant.TencentRegionId) string {
	return strings.ReplaceAll(regionId.String(), "_", "-")
}

package tenant

import (
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/pkg/errors"

	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/tenant"
)

var ErrNoExistAliyunRegionId = errors.New("region id not exist in aliyun")

type AliyunTenant struct {
	*sdk.Client
}

func NewTenant(regionId pbtenant.AliyunRegionId, accessKeyId, accessKeySecret string) (*AliyunTenant, error) {
	if _, ok := pbtenant.AliyunRegionId_name[int32(regionId)]; !ok {
		return nil, errors.WithMessage(ErrNoExistAliyunRegionId, fmt.Sprintf("region id is %d", regionId))
	}
	client, err := sdk.NewClientWithAccessKey(getRegionId(regionId), accessKeyId, accessKeySecret)
	if err != nil {
		return nil, errors.Wrap(err, "init aliyun client error")
	}
	return &AliyunTenant{client}, nil
}

func getRegionId(regionId pbtenant.AliyunRegionId) string {
	return strings.ReplaceAll(regionId.String(), "_", "-")
}

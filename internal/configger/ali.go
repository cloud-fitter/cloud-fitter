package configger

import (
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	alicfg "github.com/aliyun/alibaba-cloud-sdk-go/services/config"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

var aliClientMutex sync.Mutex

type AliCfg struct {
	cli        *alicfg.Client
	regionId   pbtenant.AliRegionId
	regionName string

	tenanter.Tenanter
}

func NewAliConfigClient(regionId int32, tenant tenanter.Tenanter) (Configger, error) {
	var client *alicfg.Client

	rName, err := tenanter.GetAliRegionName(regionId)
	if err != nil {
		return nil, errors.WithMessage(err, "get region name")
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		// 配置资源必须连接上海
		client, err = alicfg.NewClientWithAccessKey("cn-shanghai", t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali cfg client error")
	}

	return &AliCfg{
		cli:        client,
		regionId:   pbtenant.AliRegionId(regionId),
		regionName: rName,
		Tenanter:   tenant.Clone(),
	}, nil
}

func (cfg *AliCfg) Statistic() (*pbcfg.StatisticRespList, error) {
	req := alicfg.CreateListDiscoveredResourcesRequest()
	req.PageNumber = requests.NewInteger(1)
	req.PageSize = requests.NewInteger(1)
	req.ResourceTypes = "ACS::ECS::Instance"
	req.Regions = cfg.regionName

	resp, err := cfg.cli.ListDiscoveredResources(req)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDiscoveredResources error")
	}

	return &pbcfg.StatisticRespList{
		Provider:    pbtenant.CloudProvider_ali_cloud,
		AccountName: cfg.AccountName(),
		Product:     pbtenant.CloudProduct_product_ecs,
		RegionId:    int32(cfg.regionId),
		RegionName:  cfg.regionName,
		Count:       int64(resp.DiscoveredResourceProfiles.TotalCount),
	}, nil
}

package configger

import "github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"

type Configger interface {
	// 统计接口
	Statistic() (*pbcfg.StatisticRespList, error)
}

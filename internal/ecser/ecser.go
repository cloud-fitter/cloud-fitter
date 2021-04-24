package ecser

import "github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"

type Ecser interface {
	// 统计接口
	ECSStatistic() (*pbecs.ECSStatisticRespList, error)
	DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error)
}

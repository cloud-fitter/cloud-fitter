package ecser

import pbecs "github.com/cloud-fitter/cloud-fitter/gen/idl/ecs"

type Ecser interface {
	DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error)
}

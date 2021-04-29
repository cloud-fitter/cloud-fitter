package ecser

import "github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"

type Ecser interface {
	DescribeInstances(pageNumber, pageSize int32) (*pbecs.ListResp, error)
}

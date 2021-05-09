package tenanter

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

func TestGetAllRegionIds(t *testing.T) {
	type args struct {
		provider pbtenant.CloudProvider
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "ali", args: args{provider: pbtenant.CloudProvider_ali}},
		{name: "tencent", args: args{provider: pbtenant.CloudProvider_tencent}},
		{name: "huawei", args: args{provider: pbtenant.CloudProvider_huawei}},
		{name: "aws", args: args{provider: pbtenant.CloudProvider_aws}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRegions := GetAllRegionIds(tt.args.provider); len(gotRegions) == 0 {
				t.Errorf("GetAllRegionIds() = %v, want >0", gotRegions)
			}
		})
	}
}

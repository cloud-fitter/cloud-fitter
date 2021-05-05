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
		{name: "ali", args: args{provider: pbtenant.CloudProvider_ali_cloud}},
		{name: "tencent", args: args{provider: pbtenant.CloudProvider_tencent_cloud}},
		{name: "huawei", args: args{provider: pbtenant.CloudProvider_huawei_cloud}},
		{name: "aws", args: args{provider: pbtenant.CloudProvider_aws_cloud}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRegions := GetAllRegionIds(tt.args.provider); len(gotRegions) == 0 {
				t.Errorf("GetAllRegionIds() = %v, want >0", gotRegions)
			}
		})
	}
}

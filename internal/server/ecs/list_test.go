package ecs

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestList(t *testing.T) {
	type args struct {
		req       *pbecs.ListReq
		tenanters []tenanter.Tenanter
	}
	tests := []struct {
		name    string
		args    args
		want    []*pbcfg.StatisticRespList
		wantErr bool
	}{
		{name: "ali", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_ali_cloud, RegionId: int32(pbtenant.AliRegionId_ali_cn_hangzhou), PageNumber: 1, PageSize: 10}, tenanters: aliTenant}, wantErr: false},
		{name: "tencent", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_tencent_cloud, RegionId: int32(pbtenant.TencentRegionId_tc_ap_beijing), PageNumber: 1, PageSize: 10}, tenanters: tcTenant}, wantErr: false},
		{name: "aws", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_aws_cloud, RegionId: int32(pbtenant.AwsRegionId_aws_us_east_2), PageNumber: 1, PageSize: 10}, tenanters: awsTenant}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.req, tt.args.tenanters)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
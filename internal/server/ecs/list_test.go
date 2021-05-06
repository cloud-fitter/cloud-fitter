package ecs

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

func TestListDetail(t *testing.T) {
	type args struct {
		req *pbecs.ListDetailReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ali", args: args{req: &pbecs.ListDetailReq{Provider: pbtenant.CloudProvider_ali_cloud, RegionId: int32(pbtenant.AliRegionId_ali_cn_hangzhou), PageNumber: 1, PageSize: 10}}, wantErr: false},
		{name: "tencent", args: args{req: &pbecs.ListDetailReq{Provider: pbtenant.CloudProvider_tencent_cloud, RegionId: int32(pbtenant.TencentRegionId_tc_ap_beijing), PageNumber: 1, PageSize: 10}}, wantErr: false},
		{name: "aws", args: args{req: &pbecs.ListDetailReq{Provider: pbtenant.CloudProvider_aws_cloud, RegionId: int32(pbtenant.AwsRegionId_aws_us_east_2), PageNumber: 1, PageSize: 10}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListDetail(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDetail() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		req *pbecs.ListReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ali", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_ali_cloud}}, wantErr: false},
		{name: "tencent", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_tencent_cloud}}, wantErr: false},
		{name: "huawei", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_huawei_cloud}}, wantErr: false},
		{name: "aws", args: args{req: &pbecs.ListReq{Provider: pbtenant.CloudProvider_aws_cloud}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

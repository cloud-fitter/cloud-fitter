package ecser

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestEcser_ListDetail(t *testing.T) {
	region, _ := tenanter.NewRegion(pbtenant.CloudProvider_ali, int32(pbtenant.AliRegionId_ali_cn_hangzhou))
	ali, _ := NewEcsClient(pbtenant.CloudProvider_ali, region, aliTenant[0])
	aliFailed, _ := NewEcsClient(pbtenant.CloudProvider_ali, region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_tencent, int32(pbtenant.TencentRegionId_tc_ap_beijing))
	tc, _ := NewEcsClient(pbtenant.CloudProvider_tencent, region, tcTenant[0])
	tcFailed, _ := NewEcsClient(pbtenant.CloudProvider_tencent, region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_huawei, int32(pbtenant.HuaweiRegionId_hw_cn_southwest_2))
	hw, _ := NewEcsClient(pbtenant.CloudProvider_huawei, region, hwTenant[0])
	// hwFailed, _ := newHuaweiEcsClient(int32(pbtenant.HuaweiRegionId_hw_cn_north_1), tenanter.NewTenantWithAccessKey("empty", "", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_aws, int32(pbtenant.AwsRegionId_aws_us_east_2))
	aws, _ := NewEcsClient(pbtenant.CloudProvider_aws, region, awsTenant[0])

	// google, _ := NewGoogleEcsClient(tenanter.NewTenantWithAccessKey("", ""))

	type args struct {
		req *pbecs.ListDetailReq
	}
	tests := []struct {
		name    string
		fields  Ecser
		args    args
		wantErr bool
	}{
		{name: "ali wrong cli", fields: aliFailed, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		{name: "ali wrong page number", fields: ali, args: args{&pbecs.ListDetailReq{PageNumber: 0, PageSize: 1}}, wantErr: true},
		{name: "ali wrong page size", fields: ali, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 0}}, wantErr: true},
		{name: "ali right cli", fields: ali, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		{name: "tc wrong cli", fields: tcFailed, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		{name: "tc wrong page number", fields: tc, args: args{&pbecs.ListDetailReq{PageNumber: 0, PageSize: 1}}, wantErr: true},
		{name: "tc wrong page size", fields: tc, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 0}}, wantErr: true},
		{name: "tc right cli", fields: tc, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "hw wrong cli", fields: hwFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "hw right cli", fields: hw, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		{name: "aws right cli", fields: aws, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "right cli", fields: google, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.ListDetail(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDetail() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", err)
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

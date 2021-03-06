package rediser

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbredis"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestRediser_ListDetail(t *testing.T) {
	// region, _ := tenanter.NewRegion(pbtenant.CloudProvider_ali, int32(pbtenant.AliRegionId_ali_cn_hangzhou))
	// ali, _ := NewRedisClient(pbtenant.CloudProvider_ali, region, aliTenant[0])
	// aliFailed, _ := NewRedisClient(pbtenant.CloudProvider_ali, region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ := tenanter.NewRegion(pbtenant.CloudProvider_tencent, int32(pbtenant.TencentRegionId_tc_ap_beijing))
	tc, _ := NewRedisClient(pbtenant.CloudProvider_tencent, region, tcTenant[0])
	tcFailed, _ := NewRedisClient(pbtenant.CloudProvider_tencent, region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_huawei, int32(pbtenant.HuaweiRegionId_hw_ap_southeast_1))
	hw, _ := NewRedisClient(pbtenant.CloudProvider_huawei, region, hwTenant[0])
	// hwFailed, _ := NewRedisClient(pbtenant.CloudProvider_huawei, region, tenanter.NewTenantWithAccessKey("empty", "", ""))
	//
	// region, _ = tenanter.NewRegion(pbtenant.CloudProvider_aws, int32(pbtenant.AwsRegionId_aws_us_east_2))
	// aws, _ := NewRedisClient(pbtenant.CloudProvider_aws, region, awsTenant[0])

	// google, _ := NewGoogleEcsClient(tenanter.NewTenantWithAccessKey("", ""))

	type args struct {
		req *pbredis.ListDetailReq
	}
	tests := []struct {
		name    string
		fields  Rediser
		args    args
		wantErr bool
	}{
		// {name: "ali wrong cli", fields: aliFailed, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		// {name: "ali wrong page size", fields: ali, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 0}}, wantErr: true},
		// {name: "ali right cli", fields: ali, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		{name: "tc wrong cli", fields: tcFailed, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		{name: "tc wrong page number", fields: tc, args: args{&pbredis.ListDetailReq{PageNumber: 0, PageSize: 1}}, wantErr: true},
		{name: "tc wrong page size", fields: tc, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 0}}, wantErr: true},
		{name: "tc right cli", fields: tc, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "hw wrong cli", fields: hwFailed, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: true},
		{name: "hw right cli", fields: hw, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "aws right cli", fields: aws, args: args{&pbredis.ListDetailReq{PageNumber: 1, PageSize: 30}}, wantErr: false},

		// {name: "right cli", fields: google, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.fields)
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

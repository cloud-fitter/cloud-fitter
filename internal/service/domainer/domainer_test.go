package domainer

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestDomainer_ListDetail(t *testing.T) {
	ali, _ := NewDomainClient(pbtenant.CloudProvider_ali, aliTenant[0])
	aliFailed, _ := NewDomainClient(pbtenant.CloudProvider_ali, tenanter.NewTenantWithAccessKey("empty", "", ""))

	tc, _ := NewDomainClient(pbtenant.CloudProvider_tencent, tcTenant[0])
	tcFailed, _ := NewDomainClient(pbtenant.CloudProvider_tencent, tenanter.NewTenantWithAccessKey("empty", "", ""))

	// region, _ = tenanter.NewRegion(pbtenant.CloudProvider_huawei_cloud, int32(pbtenant.HuaweiRegionId_hw_cn_southwest_2))
	// hw, _ := NewHuaweiEcsClient(region, hwTenant[0])
	// // hwFailed, _ := NewHuaweiEcsClient(int32(pbtenant.HuaweiRegionId_hw_cn_north_1), tenanter.NewTenantWithAccessKey("empty", "", "", ""))
	//
	// region, _ = tenanter.NewRegion(pbtenant.CloudProvider_aws_cloud, int32(pbtenant.AwsRegionId_aws_us_east_2))
	// aws, _ := NewAwsEcsClient(region, awsTenant[0])

	// google, _ := NewGoogleEcsClient(tenanter.NewTenantWithAccessKey("", ""))

	type args struct {
		req *pbdomain.ListDetailReq
	}
	tests := []struct {
		name    string
		fields  Rdser
		args    args
		wantErr bool
	}{
		{name: "ali wrong cli", fields: aliFailed, args: args{&pbdomain.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		{name: "ali right cli", fields: ali, args: args{&pbdomain.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		{name: "tc wrong cli", fields: tcFailed, args: args{&pbdomain.ListDetailReq{PageNumber: 1, PageSize: 1}}, wantErr: true},
		{name: "tc wrong page size", fields: tc, args: args{&pbdomain.ListDetailReq{PageNumber: 1, PageSize: 0}}, wantErr: true},
		{name: "tc right cli", fields: tc, args: args{&pbdomain.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "hw wrong cli", fields: hwFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		// {name: "hw right cli", fields: hw, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

		// {name: "aws right cli", fields: aws, args: args{&pbecs.ListDetailReq{PageNumber: 1, PageSize: 10}}, wantErr: false},

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

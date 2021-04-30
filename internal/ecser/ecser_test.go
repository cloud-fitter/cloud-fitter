package ecser

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestEcser_DescribeInstances(t *testing.T) {
	ali, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), aliTenant[0])
	aliFailed, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), tenanter.NewTenantWithAccessKey("empty", "", ""))

	tc, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tcTenant[0])
	tcFailed, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tenanter.NewTenantWithAccessKey("empty", "", ""))

	hw, _ := NewHuaweiEcsClient(int32(pbtenant.HuaweiRegionId_hw_cn_southwest_2), hwTenant[0])
	// hwFailed, _ := NewHuaweiEcsClient(int32(pbtenant.HuaweiRegionId_hw_cn_north_1), tenanter.NewTenantWithAccessKey("empty", "", "", ""))

	aws, _ := NewAwsEcsClient(int32(pbtenant.AwsRegionId_aws_us_east_2), awsTenant[0])

	// google, _ := NewGoogleEcsClient(tenanter.NewTenantWithAccessKey("", ""))

	type args struct {
		pageNumber int32
		pageSize   int32
	}
	tests := []struct {
		name    string
		fields  Ecser
		args    args
		wantErr bool
	}{
		{name: "ali wrong cli", fields: aliFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "ali wrong page number", fields: ali, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		{name: "ali wrong page size", fields: ali, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		{name: "ali right cli", fields: ali, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		{name: "tc wrong cli", fields: tcFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "tc wrong page number", fields: tc, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		{name: "tc wrong page size", fields: tc, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		{name: "tc right cli", fields: tc, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		// {name: "hw wrong cli", fields: hwFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "hw right cli", fields: hw, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		{name: "aws right cli", fields: aws, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		// {name: "right cli", fields: google, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.DescribeInstances(tt.args.pageNumber, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeInstances() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", err)
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

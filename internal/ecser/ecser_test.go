package ecser

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestEcser_DescribeInstances(t *testing.T) {
	ali, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), aliTenant.Clone())
	aliFailed, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), tenanter.NewTenantWithAccessKey("", ""))

	tc, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tcTenant.Clone())
	tcFailed, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tenanter.NewTenantWithAccessKey("", ""))

	// hw, _ := NewHuaweiEcsClient(pbtenant.HuaweiRegionId_hw_cn_north_1, hwTenant)
	// hwFailed, _ := NewHuaweiEcsClient(pbtenant.HuaweiRegionId_hw_cn_north_1, tenanter.NewTenantWithAccessKey("", ""))

	type args struct {
		pageNumber int
		pageSize   int
	}
	tests := []struct {
		name    string
		fields  Ecser
		args    args
		wantErr bool
	}{
		{name: "wrong cli", fields: aliFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "wrong page number", fields: ali, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		{name: "wrong page size", fields: ali, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		{name: "right cli", fields: ali, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		{name: "wrong cli", fields: tcFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		{name: "wrong page number", fields: tc, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		{name: "wrong page size", fields: tc, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		{name: "right cli", fields: tc, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		// {name: "wrong cli", fields: hwFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		// {name: "wrong page number", fields: hw, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		// {name: "wrong page size", fields: hw, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		// {name: "right cli", fields: hw, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.DescribeInstances(tt.args.pageNumber, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", err)
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

func TestEcser_ECSStatistic(t *testing.T) {
	ali, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), aliTenant.Clone())
	aliFailed, _ := NewAliEcsClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), tenanter.NewTenantWithAccessKey("", ""))

	tc, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tcTenant.Clone())
	tcFailed, _ := NewTencentCvmClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tenanter.NewTenantWithAccessKey("", ""))

	// hw, _ := NewHuaweiEcsClient(pbtenant.HuaweiRegionId_hw_cn_north_1, hwTenant)
	// hwFailed, _ := NewHuaweiEcsClient(pbtenant.HuaweiRegionId_hw_cn_north_1, tenanter.NewTenantWithAccessKey("", ""))

	aws, _ := NewAwsEcsClient(awsTenant.Clone())

	type args struct {
	}
	tests := []struct {
		name    string
		fields  Ecser
		args    args
		wantErr bool
	}{
		{name: "wrong cli", fields: aliFailed, args: args{}, wantErr: true},
		{name: "right cli", fields: ali, args: args{}, wantErr: false},

		{name: "wrong cli", fields: tcFailed, args: args{}, wantErr: true},
		{name: "right cli", fields: tc, args: args{}, wantErr: false},

		// {name: "wrong cli", fields: hwFailed, args: args{pageNumber: 1, pageSize: 1}, wantErr: true},
		// {name: "wrong page number", fields: hw, args: args{pageNumber: 0, pageSize: 1}, wantErr: true},
		// {name: "wrong page size", fields: hw, args: args{pageNumber: 1, pageSize: 0}, wantErr: true},
		// {name: "right cli", fields: hw, args: args{pageNumber: 1, pageSize: 10}, wantErr: false},

		{name: "right cli", fields: aws, args: args{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.ECSStatistic()
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", err)
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

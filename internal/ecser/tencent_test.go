package ecser

import (
	"testing"

	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/tenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestNewTencentCvmClient(t *testing.T) {
	type args struct {
		regionId pbtenant.TencentRegionId
		tenanter tenanter.Tenanter
	}
	tests := []struct {
		name    string
		args    args
		want    *AliEcs
		wantErr bool
	}{
		{
			name: "wrong",
			args: args{
				regionId: pbtenant.TencentRegionId_ap_beijing,
				tenanter: tenanter.NewTenantWithAccessKey("", ""),
			},
			wantErr: false,
		},
		{
			name: "right",
			args: args{
				regionId: pbtenant.TencentRegionId_ap_beijing,
				tenanter: demoTenant,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTencentCvmClient(tt.args.regionId, tt.args.tenanter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAliEcsClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTencentCvm_DescribeInstances(t *testing.T) {
	cliWrong, _ := NewTencentCvmClient(pbtenant.TencentRegionId_ap_beijing, tenanter.NewTenantWithAccessKey("", ""))
	cliRight, _ := NewTencentCvmClient(pbtenant.TencentRegionId_ap_beijing, demo2Tenant)
	cliRight2, _ := NewTencentCvmClient(pbtenant.TencentRegionId_ap_guangzhou, demo2Tenant)
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
		{
			name:   "wrong cli",
			fields: cliWrong,
			args: args{
				pageNumber: 1,
				pageSize:   1,
			},
			wantErr: true,
		},
		{
			name:   "wrong page number",
			fields: cliRight,
			args: args{
				pageNumber: 0,
				pageSize:   1,
			},
			wantErr: true,
		},
		{
			name:   "wrong page size",
			fields: cliRight,
			args: args{
				pageNumber: 1,
				pageSize:   0,
			},
			wantErr: true,
		},
		{
			name:   "right cli",
			fields: cliRight,
			args: args{
				pageNumber: 1,
				pageSize:   10,
			},
			wantErr: false,
		},
		{
			name:   "right cli2",
			fields: cliRight2,
			args: args{
				pageNumber: 1,
				pageSize:   10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.DescribeInstances(tt.args.pageNumber, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DescribeInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

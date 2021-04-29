package configger

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestConfigger_Statistic(t *testing.T) {
	ali, _ := NewAliConfigClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), aliTenant[0])
	aliFailed, _ := NewAliConfigClient(int32(pbtenant.AliRegionId_ali_cn_hangzhou), tenanter.NewTenantWithAccessKey("empty", "", ""))

	tc, _ := NewTencentCfgClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tcTenant[0])
	tcFailed, _ := NewAliConfigClient(int32(pbtenant.TencentRegionId_tc_ap_beijing), tenanter.NewTenantWithAccessKey("empty", "", ""))

	aws, _ := NewAwsCfgClient(int32(pbtenant.AwsRegionId_aws_us_east_2), awsTenant[0])
	awsFailed, _ := NewAwsCfgClient(int32(pbtenant.AwsRegionId_aws_us_east_2), tenanter.NewTenantWithAccessKey("empty", "", ""))

	type args struct {
	}
	tests := []struct {
		name    string
		fields  Configger
		args    args
		wantErr bool
	}{
		{name: "ali wrong cli", fields: aliFailed, args: args{}, wantErr: true},
		{name: "ali right cli", fields: ali, args: args{}, wantErr: false},

		{name: "tc wrong cli", fields: tcFailed, args: args{}, wantErr: true},
		{name: "tc right cli", fields: tc, args: args{}, wantErr: false},

		{name: "aws wrong cli", fields: awsFailed, args: args{}, wantErr: true},
		{name: "aws right cli", fields: aws, args: args{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.Statistic()
			if (err != nil) != tt.wantErr {
				t.Errorf("Statistic() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", err)
			if err == nil {
				t.Log(resp)
			}
		})
	}
}

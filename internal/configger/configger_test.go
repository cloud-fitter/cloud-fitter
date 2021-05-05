package configger

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestConfigger_Statistic(t *testing.T) {
	region, _ := tenanter.NewRegion(pbtenant.CloudProvider_ali_cloud, int32(pbtenant.AliRegionId_ali_cn_hangzhou))
	ali, _ := NewAliCfgClient(region, aliTenant[0])
	aliFailed, _ := NewAliCfgClient(region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_tencent_cloud, int32(pbtenant.TencentRegionId_tc_ap_beijing))
	tc, _ := NewTencentCfgClient(region, tcTenant[0])
	tcFailed, _ := NewTencentCfgClient(region, tenanter.NewTenantWithAccessKey("empty", "", ""))

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_huawei_cloud, int32(pbtenant.HuaweiRegionId_hw_cn_southwest_2))
	hw, _ := NewHuaweiCfgClient(region, hwTenant[0])

	region, _ = tenanter.NewRegion(pbtenant.CloudProvider_aws_cloud, int32(pbtenant.AwsRegionId_aws_us_east_2))
	aws, _ := NewAwsCfgClient(region, awsTenant[0])
	awsFailed, _ := NewAwsCfgClient(region, tenanter.NewTenantWithAccessKey("empty", "", ""))

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

		{name: "hw right cli", fields: hw, args: args{}, wantErr: false},

		{name: "aws wrong cli", fields: awsFailed, args: args{}, wantErr: true},
		{name: "aws right cli", fields: aws, args: args{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fields.Statistic(context.Background())
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

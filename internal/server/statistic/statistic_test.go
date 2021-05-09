package statistic

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestStatistic(t *testing.T) {
	type args struct {
		provider  pbtenant.CloudProvider
		tenanters []tenanter.Tenanter
	}
	tests := []struct {
		name    string
		args    args
		want    []*pbstatistic.StatisticInfo
		wantErr bool
	}{
		{name: "ali", args: args{provider: pbtenant.CloudProvider_ali, tenanters: aliTenant}, wantErr: false},
		{name: "tencent", args: args{provider: pbtenant.CloudProvider_tencent, tenanters: tcTenant}, wantErr: false},
		{name: "huawei", args: args{provider: pbtenant.CloudProvider_huawei, tenanters: hwTenant}, wantErr: false},
		{name: "aws", args: args{provider: pbtenant.CloudProvider_aws, tenanters: awsTenant}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Statistic(context.Background(), tt.args.provider, tt.args.tenanters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Statistic() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

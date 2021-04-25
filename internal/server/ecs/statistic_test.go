package ecs

import (
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

func TestStatistic(t *testing.T) {
	type args struct {
		provider pbtenant.CloudProvider
	}
	tests := []struct {
		name    string
		args    args
		want    []*pbecs.ECSStatisticRespList
		wantErr bool
	}{
		{
			name:    "ali",
			args:    args{provider: pbtenant.CloudProvider_ali_cloud},
			wantErr: false,
		},
		{
			name:    "tencent",
			args:    args{provider: pbtenant.CloudProvider_tencent_cloud},
			wantErr: false,
		},
		{
			name:    "aws",
			args:    args{provider: pbtenant.CloudProvider_aws_cloud},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Statistic(tt.args.provider)
			if (err != nil) != tt.wantErr {
				t.Errorf("Statistic() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

package tenant

import (
	"errors"
	"testing"

	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/tenant"
)

func TestNewTenant(t *testing.T) {
	type args struct {
		regionId        pbtenant.AliyunRegionId
		accessKeyId     string
		accessKeySecret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "empty",
			args: args{
				regionId:        -1,
				accessKeyId:     "wrong",
				accessKeySecret: "wrong",
			},
			wantErr: ErrNoExistAliyunRegionId,
		},
		{
			name: "right",
			args: args{
				regionId:        pbtenant.AliyunRegionId_cn_beijing,
				accessKeyId:     "wrong",
				accessKeySecret: "wrong",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTenant(tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret)
			if err != tt.wantErr && !errors.Is(err, tt.wantErr) {
				t.Errorf("NewTenant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

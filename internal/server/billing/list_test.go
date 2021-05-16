package billing

import (
	"context"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

func TestListDetail(t *testing.T) {
	type args struct {
		req *pbbilling.ListDetailReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ali", args: args{req: &pbbilling.ListDetailReq{Provider: pbtenant.CloudProvider_ali, BillingCycle: "2021-05"}}, wantErr: false},
		// {name: "tencent", args: args{req: &pbbilling.ListDetailReq{Provider: pbtenant.CloudProvider_tencent}}, wantErr: false},
		// {name: "aws", args: args{req: &pbbilling.ListDetailReq{Provider: pbtenant.CloudProvider_aws}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListDetail(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDetail() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

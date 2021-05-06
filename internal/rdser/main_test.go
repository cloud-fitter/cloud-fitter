package rdser

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	aliTenant, tcTenant, hwTenant, awsTenant []tenanter.Tenanter
)

func TestMain(m *testing.M) {
	err := tenanter.LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}
	if aliTenant, err = tenanter.GetTenanters(pbtenant.CloudProvider_ali_cloud); err != nil {
		panic("get aliTenant failed")
	}
	if tcTenant, err = tenanter.GetTenanters(pbtenant.CloudProvider_tencent_cloud); err != nil {
		panic("get tcTenant failed")
	}
	if hwTenant, err = tenanter.GetTenanters(pbtenant.CloudProvider_huawei_cloud); err != nil {
		panic("get hwTenant failed")
	}
	if awsTenant, err = tenanter.GetTenanters(pbtenant.CloudProvider_aws_cloud); err != nil {
		panic("get awsTenant failed")
	}
	os.Exit(m.Run())
}

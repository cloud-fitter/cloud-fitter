package tenanter

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

var (
	aliTenant Tenanter
	tcTenant  Tenanter
	hwTenant  Tenanter
)

func TestMain(m *testing.M) {
	err := LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}

	var ok bool
	if aliTenant, ok = GetTenanter(pbtenant.CloudProvider_ali_cloud, "demo"); !ok {
		panic("get aliTenant failed")
	}
	if tcTenant, ok = GetTenanter(pbtenant.CloudProvider_tencent_cloud, "demo2"); !ok {
		panic("get tcTenantr failed")
	}
	if hwTenant, ok = GetTenanter(pbtenant.CloudProvider_huawei_cloud, "demo3"); !ok {
		panic("get hwTenant failed")
	}
	os.Exit(m.Run())
}

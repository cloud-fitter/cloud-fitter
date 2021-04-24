package ecser

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	aliTenant tenanter.Tenanter
	tcTenant  tenanter.Tenanter
	hwTenant  tenanter.Tenanter
)

func TestMain(m *testing.M) {
	err := tenanter.LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}

	var ok bool
	if aliTenant, ok = tenanter.GetTenanter(pbtenant.CloudProvider_ali_cloud, "demo"); !ok {
		panic("get aliTenant failed")
	}
	if tcTenant, ok = tenanter.GetTenanter(pbtenant.CloudProvider_tencent_cloud, "demo2"); !ok {
		panic("get tcTenant failed")
	}
	if hwTenant, ok = tenanter.GetTenanter(pbtenant.CloudProvider_huawei_cloud, "demo3"); !ok {
		panic("get hwTenant failed")
	}
	os.Exit(m.Run())
}

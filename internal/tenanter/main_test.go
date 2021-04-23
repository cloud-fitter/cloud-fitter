package tenanter

import (
	"os"
	"testing"
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
	if aliTenant, ok = GetTenanter("demo"); !ok {
		panic("get aliTenant failed")
	}
	if tcTenant, ok = GetTenanter("demo2"); !ok {
		panic("get tcTenantr failed")
	}
	if hwTenant, ok = GetTenanter("demo3"); !ok {
		panic("get hwTenant failed")
	}
	os.Exit(m.Run())
}

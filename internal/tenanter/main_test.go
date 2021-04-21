package tenanter

import (
	"os"
	"testing"
)

var (
	demoTenant  Tenanter
	demo2Tenant Tenanter
)

func TestMain(m *testing.M) {
	err := LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}

	var ok bool
	if demoTenant, ok = GetTenanter("demo"); !ok {
		panic("get demo tenanter failed")
	}
	if demo2Tenant, ok = GetTenanter("demo2"); !ok {
		panic("get demo2 tenanter failed")
	}
	os.Exit(m.Run())
}

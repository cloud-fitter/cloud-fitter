package ecser

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	demoTenant  tenanter.Tenanter
	demo2Tenant tenanter.Tenanter
)

func TestMain(m *testing.M) {
	err := tenanter.LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}

	var ok bool
	if demoTenant, ok = tenanter.GetTenanter("demo"); !ok {
		panic("get demo tenanter failed")
	}
	if demo2Tenant, ok = tenanter.GetTenanter("demo2"); !ok {
		panic("get demo2 tenanter failed")
	}
	os.Exit(m.Run())
}

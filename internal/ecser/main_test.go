package ecser

import (
	"os"
	"testing"

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
	if aliTenant, ok = tenanter.GetTenanter("demo"); !ok {
		panic("get aliTenant failed")
	}
	if tcTenant, ok = tenanter.GetTenanter("demo2"); !ok {
		panic("get tcTenant failed")
	}
	if hwTenant, ok = tenanter.GetTenanter("demo3"); !ok {
		panic("get hwTenant failed")
	}
	os.Exit(m.Run())
}

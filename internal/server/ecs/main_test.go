package ecs

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func TestMain(m *testing.M) {
	err := tenanter.LoadCloudConfigs("../../../config.yaml")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

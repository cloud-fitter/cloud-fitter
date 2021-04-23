package tenanter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

const osEnvKey = "CLOUD_FITTER_CONFIGS"

var (
	ErrLoadTenanterFromFile  = errors.New("load tenanter from file failed")
	ErrLoadTenanterFromOsEnv = errors.New("load tenanter from os env failed")
	ErrLoadTenanterFileEmpty = errors.New("load tenanter from file failed")
)

type Tenanter interface {
	Clone() Tenanter
}

var gStore = globalStore{stores: make(map[pbtenant.CloudProvider]map[string]Tenanter)}

type globalStore struct {
	sync.Mutex
	stores map[pbtenant.CloudProvider]map[string]Tenanter
}

func LoadCloudConfigs(configFile string) error {
	if err := LoadCloudConfigsFromFile(configFile); errors.Is(err, ErrLoadTenanterFileEmpty) {
		return LoadCloudConfigsFromOsEnv()
	}
	return nil
}

func LoadCloudConfigsFromFile(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return ErrLoadTenanterFileEmpty
	}

	var configs = new(pbtenant.CloudConfigs)
	if err = yaml.Unmarshal(b, configs); err != nil {
		return errors.WithMessage(ErrLoadTenanterFromFile, err.Error())
	}

	return load(configs)
}

func LoadCloudConfigsFromOsEnv() error {
	data := os.Getenv(osEnvKey)
	var configs = new(pbtenant.CloudConfigs)
	if err := json.Unmarshal([]byte(data), configs); err != nil {
		return errors.WithMessage(ErrLoadTenanterFromOsEnv, err.Error())
	}

	return load(configs)
}

func load(configs *pbtenant.CloudConfigs) error {
	gStore.Lock()
	defer gStore.Unlock()

	for k := range pbtenant.CloudProvider_name {
		gStore.stores[pbtenant.CloudProvider(k)] = make(map[string]Tenanter)
	}

	// 这里 name 会覆盖
	for _, c := range configs.Configs {
		if c.AccessId != "" && c.AccessSecret != "" {
			gStore.stores[c.Provider][c.Name] = NewTenantWithAccessKey(c.AccessId, c.AccessSecret)
		}
	}
	return nil
}

func GetTenantersMap(provider pbtenant.CloudProvider) map[string]Tenanter {
	gStore.Lock()
	defer gStore.Unlock()

	tmap := make(map[string]Tenanter, len(gStore.stores[provider]))
	for k := range gStore.stores[provider] {
		tmap[k] = gStore.stores[provider][k].Clone()
	}
	return tmap
}

func GetTenanter(provider pbtenant.CloudProvider, name string) (Tenanter, bool) {
	gStore.Lock()
	defer gStore.Unlock()

	tenanter, ok := gStore.stores[provider][name]
	if !ok {
		return nil, false
	}
	return tenanter.Clone(), true
}

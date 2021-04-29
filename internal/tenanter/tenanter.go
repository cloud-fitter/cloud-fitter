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
	ErrNoTenanters           = errors.New("no tenanters for the cloud")
)

type Tenanter interface {
	AccountName() string
	Clone() Tenanter
}

var gStore = globalStore{stores: make(map[pbtenant.CloudProvider][]Tenanter)}

type globalStore struct {
	sync.Mutex
	stores map[pbtenant.CloudProvider][]Tenanter
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

	for _, c := range configs.Configs {
		if c.AccessId != "" && c.AccessSecret != "" {
			gStore.stores[c.Provider] = append(gStore.stores[c.Provider], NewTenantWithAccessKey(c.Name, c.AccessId, c.AccessSecret))
		}
	}
	return nil
}

func GetTenanters(provider pbtenant.CloudProvider) ([]Tenanter, error) {
	gStore.Lock()
	defer gStore.Unlock()

	if len(gStore.stores[provider]) == 0 {
		return nil, errors.WithMessagef(ErrNoTenanters, "cloud is %v", provider)
	}

	var tenanters = make([]Tenanter, len(gStore.stores[provider]))
	for k := range gStore.stores[provider] {
		tenanters[k] = gStore.stores[provider][k].Clone()
	}
	return tenanters, nil
}

package tenanter

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

var (
	ErrTenantNameExist       = errors.New("tenant name already exist")
	ErrLoadTenanterFromFile  = errors.New("load tenanter from file failed")
	ErrLoadTenanterFileEmpty = errors.New("load tenanter from file failed")
)

type Tenanter interface {
}

var gStore = globalStore{stores: make(map[string]Tenanter)}

type globalStore struct {
	sync.RWMutex
	stores map[string]Tenanter
}

func LoadCloudConfigs(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return ErrLoadTenanterFileEmpty
	}

	var configs = new(pbtenant.CloudConfigs)
	if err = yaml.Unmarshal(b, configs); err != nil {
		return errors.WithMessage(ErrLoadTenanterFromFile, err.Error())
	}

	gStore.Lock()
	defer gStore.Unlock()

	// 保证原子性，先全部检查无重复，再插入
	for _, c := range configs.Ali {
		if _, ok := gStore.stores[c.Name]; ok {
			return errors.WithMessage(ErrTenantNameExist, fmt.Sprintf("name is %s", c.Name))
		}
	}
	for _, c := range configs.Tencent {
		if _, ok := gStore.stores[c.Name]; ok {
			return errors.WithMessage(ErrTenantNameExist, fmt.Sprintf("name is %s", c.Name))
		}
	}
	for _, c := range configs.Huawei {
		if _, ok := gStore.stores[c.Name]; ok {
			return errors.WithMessage(ErrTenantNameExist, fmt.Sprintf("name is %s", c.Name))
		}
	}

	for _, c := range configs.Ali {
		if c.AccessId != "" && c.AccessSecret != "" {
			gStore.stores[c.Name] = NewTenantWithAccessKey(c.AccessId, c.AccessSecret)
		}
	}
	for _, c := range configs.Tencent {
		if c.AccessId != "" && c.AccessSecret != "" {
			gStore.stores[c.Name] = NewTenantWithAccessKey(c.AccessId, c.AccessSecret)
		}
	}
	for _, c := range configs.Huawei {
		if c.AccessId != "" && c.AccessSecret != "" {
			gStore.stores[c.Name] = NewTenantWithAccessKey(c.AccessId, c.AccessSecret)
		}
	}
	return nil
}

func GetTenanter(name string) (Tenanter, bool) {
	gStore.RLock()
	defer gStore.RUnlock()

	if gStore.stores == nil {
		return nil, false
	}
	tenanter, ok := gStore.stores[name]
	return tenanter, ok
}

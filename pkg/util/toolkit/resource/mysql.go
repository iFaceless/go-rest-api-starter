package resource

import (
	"fmt"
	"strings"
)

type MySQLConfig struct {
	*Config
}

func (mc *MySQLConfig) Addr() string {
	return fmt.Sprintf("%s:%s@%s:%s/%s", mc.User(), mc.Password(), mc.Host(), mc.Port(), mc.Database())
}

func (mc *MySQLConfig) Database() string {
	return strings.TrimLeft(mc.Path(), "/")
}

type MySQLResource struct {
	MasterInstances []*MySQLConfig
	SlaveInstances  []*MySQLConfig
}

func DiscoverMySQL(name string) (*MySQLResource, error) {
	masterConfigs := findResourceConfigs("mysql", fmt.Sprintf("%s_master", name))
	if len(masterConfigs) == 0 {
		return nil, errResourceNotFound("mysql", name)
	}

	slaveConfigs := findResourceConfigs("mysql", fmt.Sprintf("%s_slave", name))

	res := &MySQLResource{
		MasterInstances: make([]*MySQLConfig, 0, len(masterConfigs)),
		SlaveInstances:  make([]*MySQLConfig, 0, len(slaveConfigs)),
	}

	for _, conf := range masterConfigs {
		res.MasterInstances = append(res.MasterInstances, &MySQLConfig{conf})
	}

	for _, conf := range slaveConfigs {
		res.SlaveInstances = append(res.SlaveInstances, &MySQLConfig{conf})
	}
	return res, nil
}

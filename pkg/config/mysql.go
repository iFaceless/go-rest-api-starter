package config

import (
	"fmt"

	"github.com/ifaceless/go-starter/pkg/util/toolkit/env"

	"github.com/ifaceless/go-starter/pkg/util/toolkit/resource"
)

type DBConfig struct {
	Master string
	Slaves []string
}

var (
	MySQLConfig *DBConfig
)

func discoverMySQLResource() {
	res, err := resource.DiscoverMySQL(env.AppName())
	if err != nil {
		panic(err)
	}

	slaveAddrs := make([]string, 0)
	for _, ins := range res.SlaveInstances {
		slaveAddrs = append(slaveAddrs, makeDBAddr(ins.Addr()))
	}
	MySQLConfig = &DBConfig{
		Master: makeDBAddr(res.MasterInstances[0].Addr()),
		Slaves: slaveAddrs,
	}
}

func makeDBAddr(src string) string {
	return fmt.Sprintf("%s?charset=utf8&parseTime=true&loc=Local", src)
}

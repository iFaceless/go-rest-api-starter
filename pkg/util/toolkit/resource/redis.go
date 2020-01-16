package resource

import (
	"fmt"
)

type RedisConfig struct {
	*Config
}

func (rc *RedisConfig) Addr() string {
	var authPart string
	passwd := rc.Password()
	if passwd != "" {
		authPart = fmt.Sprintf(":%s@", passwd)
	}

	return fmt.Sprintf("%s%s:%s", authPart, rc.Host(), rc.Port())
}

type RedisResource struct {
	Instances []*RedisConfig
}

func DiscoverRedis(name string) (*RedisResource, error) {
	configs := findResourceConfigs("redis", name)
	if len(configs) == 0 {
		return nil, errResourceNotFound("redis", name)
	}

	var instances []*RedisConfig
	for _, conf := range configs {
		instances = append(instances, &RedisConfig{conf})
	}
	return &RedisResource{Instances: instances}, nil
}

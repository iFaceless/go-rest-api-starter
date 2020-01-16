package resource

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	up *url.URL
}

func (c Config) Host() string {
	return c.up.Hostname()
}

func (c Config) Port() string {
	return c.up.Port()
}

func (c Config) User() (user string) {
	if c.up.User != nil {
		user = c.up.User.Username()
	}
	return
}

func (c Config) Password() (password string) {
	if c.up.User != nil {
		password, _ = c.up.User.Password()
	}
	return
}

func (c Config) Path() string {
	return c.up.Path
}

func findResourceConfigs(kind, rname string) (configs []*Config) {
	re := regexp.MustCompile("[^0-9a-zA-Z]+")
	rname = strings.ToUpper(re.ReplaceAllString(rname, "_"))
	for index := 0; ; index++ {
		key := fmt.Sprintf("RES_%s_%s_%d_ADDR", strings.ToUpper(kind), rname, index)
		value := os.Getenv(key)
		if value == "" {
			break
		}
		configs = append(configs, parseConfig(kind, value))
	}

	return
}

func parseConfig(kind, addr string) *Config {
	if !strings.Contains(addr, "://") {
		addr = fmt.Sprintf("%s://%s", kind, addr)
	}
	u, err := url.Parse(addr)
	if err != nil {
		panic("failed to parse resource addr")
	}
	return &Config{u}
}

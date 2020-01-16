package config

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func init() {
	envFile := os.Getenv("DOT_ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		logrus.Warnf("dot env file not found: %s", os.Getenv("DOT_ENV_FILE"))
	}

	discoverMySQLResource()
}

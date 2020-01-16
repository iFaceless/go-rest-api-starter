package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	envFile := os.Getenv("DOT_ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println(os.Getenv("DOT_ENV_FILE"))
		panic(err)
	}

	discoverMySQLResource()
}

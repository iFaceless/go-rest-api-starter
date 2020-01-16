package env

import "os"

const (
	deployEnvProduction  = "production"
	deployEnvDevelopment = "development"
)

func IsInProductionEnv() bool {
	return os.Getenv("DEPLOY_ENV") == deployEnvProduction
}

func IsInDevelopmentEnv() bool {
	return os.Getenv("DEPLOY_ENV") == deployEnvDevelopment
}

func AppName() string {
	return os.Getenv("APP_NAME")
}

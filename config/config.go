package config

import (
	"os"
)

type AppConfig struct {
	Port          int
	Driver        string
	Name          string
	Address       string
	DB_Port       int
	Username      string
	Password      string
	TokenMidtrans string
	KeyIDs3       string
	AccessKeyS3   string
	MyRegion      string
}

var appConfig *AppConfig

func InitConfig() *AppConfig {

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "layered_db")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "")
	defaultConfig.TokenMidtrans = getEnv("TOKENMIDTRANS", "")
	defaultConfig.KeyIDs3 = getEnv("KEYIDS3", "")
	defaultConfig.AccessKeyS3 = getEnv("ACCESSKEYS3", "")
	defaultConfig.MyRegion = getEnv("MYREGION", "")
	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

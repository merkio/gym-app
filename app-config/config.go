package config

import (
	loadConfig "gym-app/common/config"
)

const (
	App = "gym"
)

var DataConnectionConfig DataConnectionConf

var AppConfig AppConf

var HealthConfig HealthConf

type DataConnectionConf struct {
	PostgresPort     string `default:"5432" envconfig:"POSTGRES_PORT"`
	PostgresHostname string `default:"127.0.0.1" envconfig:"POSTGRES_HOST"`
	PostgresUser     string `default:"postgres" envconfig:"POSTGRES_USER"`
	PostgresPassword string `default:"postgres" envconfig:"POSTGRES_PASSWORD"`
	PostgresDBName   string `default:"gym" envconfig:"POSTGRES_DB"`
	PostgresSSLMode  string `default:"disable" envconfig:"POSTGRES_SSL_MODE"`
	PostgresSchema   string `default:"gym" envconfig:"POSTGRES_SCHEMA"`
}

type AppConf struct {
	AppVersion string `default:"v1.0.0" envconfig:"APP_VERSION"`
}

type HealthConf struct {
}

func LoadConfig() {

	var cData DataConnectionConf
	var cApp AppConf
	var cHealth HealthConf

	loadConfig.Process(&cData, App)
	loadConfig.Process(&cApp, App)
	loadConfig.Process(&cHealth, App)

	DataConnectionConfig = cData
	AppConfig = cApp
	HealthConfig = cHealth
}

func init() {
	LoadConfig()
}

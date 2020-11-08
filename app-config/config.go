package config

import (
	loadConfig "gym-app/common/config"
)

const (
	App = "gym"
)

var DataConnectionConfig DataConnectionConf

var VkConnectionConfig VkConnectionConf

var AppConfig AppConf

var HealthConfig HealthConf

type VkConnectionConf struct {
	ClientID string `default:"7568447" envconfig:"VK_CLIENT_ID"`
	SecretKey string `default:"ds4oy2WLrUNCT0dduzqy" envconfig:"VK_SECRET_KEY"`
	AccessToken string `default:"b1fb4ef5b1fb4ef5b1fb4ef52fb18832cabb1fbb1fb4ef5eec3764367f4bdbbd150ae40" envconfig:"VK_ACCESS_TOKEN"`
	GroupID string `default:"-62011928" envconfig:"VK_GROUP_ID"`
}

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
	var cVk VkConnectionConf

	loadConfig.Process(&cData, App)
	loadConfig.Process(&cApp, App)
	loadConfig.Process(&cHealth, App)
	loadConfig.Process(&cVk, App)

	DataConnectionConfig = cData
	AppConfig = cApp
	HealthConfig = cHealth
	VkConnectionConfig = cVk
}

func init() {
	LoadConfig()
}

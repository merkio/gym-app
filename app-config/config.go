package config

import (
	loadConfig "gym-app/common/config"
	"strings"
)

const (
	App = "gym"
)

var DataConnectionConfig DataConnectionConf

var VkConnectionConfig VkConnectionConf

var AppConfig AppConf

var HealthConfig HealthConf

type VkConnectionConf struct {
	ClientID    string `default:"" envconfig:"VK_CLIENT_ID"`
	SecretKey   string `default:"" envconfig:"VK_SECRET_KEY"`
	AccessToken string `default:"" envconfig:"VK_ACCESS_TOKEN"`
	Groups      string `default:"" envconfig:"VK_GROUPS"`
	Hour        int    `default:"7" envconfig:"VK_TASK_HOUR"`
	Minute      int    `default:"0" envconfig:"VK_TASK_MINUTE"`
}

func (c VkConnectionConf) GetGroups() map[string]string {
	var groups = make(map[string]string)

	for _, group := range strings.Split(c.Groups, ",") {
		arr := strings.Split(group, ":")
		groups[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
	}
	return groups
}

type DataConnectionConf struct {
	PostgresPort     string `default:"5432" envconfig:"POSTGRES_PORT"`
	PostgresHostname string `default:"127.0.0.1" envconfig:"POSTGRES_HOST"`
	PostgresUser     string `default:"postgres" envconfig:"POSTGRES_USER"`
	PostgresPassword string `default:"postgres" envconfig:"POSTGRES_PASSWORD"`
	PostgresDBName   string `default:"postgres" envconfig:"POSTGRES_DB"`
	PostgresSSLMode  string `default:"disable" envconfig:"POSTGRES_SSL_MODE"`
	PostgresSchema   string `default:"public" envconfig:"POSTGRES_SCHEMA"`
}

type AppConf struct {
	AppVersion string `default:"v0.0.3" envconfig:"APP_VERSION"`
	MetubeUrl  string `default:"http://192.168.1.8:8081" envconfig:"METUBE_URL"`
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

package config

import (
	"fmt"

	"gym-app/common/logger"

	"github.com/kelseyhightower/envconfig"
)

// Process load config and merge with env variables
func Process(conf interface{}, app string) {
	confType := fmt.Sprintf("%T", conf)

	log := logger.NewLogger().WithField("config", confType)

	err := envconfig.Process("", conf)

	if err != nil {
		log.WithError(err).WithField("service", app).Fatal("Configuration could not be processed")
		panic(err)
	}
	log.WithField("service", app).Info("Configuration processed")
}

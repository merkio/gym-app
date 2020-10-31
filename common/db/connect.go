package db

import (
	"errors"
	"fmt"
	"gym-app/common/logger"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Specification db connection properties
type Specification struct {
	Port       string
	Hostname   string
	User       string
	Password   string
	DbName     string
	SSLMode    string
	SearchPath string
}

// Config database properties
type Config struct {
	MaxOpenConns string `default:"10" envconfig:"PG_MAX_OPEN_CONNS"`
	MaxIdleConns string `default:"2" envconfig:"PG_MAX_IDLE_CONNS"`
}

// DBConfig exported database properties
var DBConfig Config

var log = logger.NewLogger()

const dbArgsFmt = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s"

//GetDBIntstance connect to the database
func GetDBIntstance(dbs *Specification) *gorm.DB {
	dbArgs := fmt.Sprintf(dbArgsFmt, dbs.Hostname, dbs.Port, dbs.User, dbs.Password, dbs.DbName, dbs.SSLMode, dbs.SearchPath)
	db, err := gorm.Open(postgres.Open(dbArgs), &gorm.Config{})

	if err != nil {
		log.Error("Could not establish connection to database")
		log.Error(err.Error())
		return nil
	}

	err = readEnvConfig()
	if err != nil {
		sqlDB, _ := db.DB()
		maxOpenConns, err := parseIntParameter(DBConfig.MaxOpenConns)
		if err != nil || maxOpenConns > 0 {
			sqlDB.SetMaxOpenConns(maxOpenConns)
		}
		maxIdleConns, err := parseIntParameter(DBConfig.MaxIdleConns)
		if err != nil {
			sqlDB.SetMaxIdleConns(maxIdleConns)
		}
	}
	return db
}

func readEnvConfig() error {
	logrus.Debug("Reading envconfig for DataBase configuration variables...")
	err := envconfig.Process("", &DBConfig)
	if err != nil {
		logrus.Fatal("Error while reading envconfig :", err.Error())
		return err
	}
	return nil
}

func parseIntParameter(parameter string) (int, error) {
	if parameter == "" {
		return 0, errors.New("Could not parse empty parameter")
	}
	return strconv.Atoi(parameter)
}

package result

import (
	config "gym-app/app-config"
	dbConnector "gym-app/common/db"
	loggerWrap "gym-app/common/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetDB connection to the db
func GetDB(conf config.DataConnectionConf, app string) *gorm.DB {
	db := dbConnector.GetDBIntstance(&db.Specification{
		Port: conf.PostgresPort,
		Hostname: conf.PostgresHostname,
		User: conf.PostgresUser,
		Password: conf.PostgresPassword,
		DbName: conf.PostgresDBName,
		SSLMode: conf.PostgresSSLMode,
		SearchPath: conf.PostgresSchema,
	})

	if logger.Config.LogLevel == "trace" {
		db.SetLogger(loggerWrap.NewLogger().WithFields(logrus.Fields{
		"service": config.App,
		"app_version": config.AppConfig.AppVersion,
	}))
	db.LogMode(true)
	}
	return db
} 

var db *gorm.DB
var log *logrus.Logger

func init() {
	db = GetDB(config.DataConnectionConf, config.App)
	log = loggerWrap.NewLogger()
}

// GetResults returns the list of Results
func (r Repository) GetResults() []Result {
	results := make([]Result, 30)
	result := db.Find(&results)

	if result.Error != nil {
		log.Error("Can't get results from db.\n%s", result.Error)
	}

	log.Infof("Found %d amount of results", result.RowsAffected)
	return results
}

// GetResult return the Result with id
func (r Repository) GetResult(id uuid.UUID) (Result, err) {
	result := Result{}
	r := db.Find(&result, id)

	if r.Error != nil {
		log.Errorf("Can't create result %v\n%s", r, r.Error)
		return nil, r.Error
	}

	return result, nil
}

// AddResult inserts an Result into DB
func (r Repository) AddResult(result Result) bool {
	result := db.Create(&result)

	if result.Error != nil {
		log.Errorf("Can't create result %v\n%s", result, result.Error)
		return false
	}

	return true
}

// AddResults inserts an Results into DB
func (r Repository) AddResults(results []Result) bool {
	result := db.Create(&results)

	if result.Error != nil {
		log.Errorf("Can't create result %v\n%s", results, result.Error)
		return false
	}

	return true
}

// UpdateResult updates an Result in the DB (not used for now)
func (r Repository) UpdateResult(result Result) bool {
	result := db.Model(&result).Updates(result)

	if result.Error != nil {
		log.Error("Can't update result with values %v\n%s", result, result.Error)
		return false
	}
	return true
}

// DeleteResult deletes an Result (not used for now)
func (r Repository) DeleteResult(id string) string {
	result := db.Delete(&Result{}, id)

	if result.Error != nil {
		log.Error("Can't delete result with id %s\n%s", id, result.Error)
		return "Error"
	}

	return "OK"
}

package result

import (
	config "gym-app/app-config"
	"gym-app/common/db"
	loggerWrap "gym-app/common/logger"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository results repository
type Repository interface {
	repo.BaseRepository
	CreateAll(results []Result) bool
	GetByID(id string) (Result, error)
	Get() []Result
}

// RRepository instance of RRepository
type RRepository struct {
	Repository
}

// GetDB connection to the db
func GetDB(conf config.DataConnectionConf, app string) *gorm.DB {
	dbConn := db.GetDBIntstance(&db.Specification{
		Port:       conf.PostgresPort,
		Hostname:   conf.PostgresHostname,
		User:       conf.PostgresUser,
		Password:   conf.PostgresPassword,
		DbName:     conf.PostgresDBName,
		SSLMode:    conf.PostgresSSLMode,
		SearchPath: conf.PostgresSchema,
	})

	return dbConn
}

var dbConn *gorm.DB
var log *logrus.Logger

func init() {
	dbConn = GetDB(config.DataConnectionConfig, config.App)
	log = loggerWrap.NewLogger()
}

// Get returns the list of Results
func (r RRepository) Get() []Result {
	results := make([]Result, 30)
	result := dbConn.Find(&results)

	if result.Error != nil {
		log.Error("Can't get results from dbConn.\n%s", result.Error)
	}

	log.Infof("Found %d amount of results", result.RowsAffected)
	return results
}

// GetByID return the Result with id
func (r RRepository) GetByID(id string) (Result, error) {
	result := Result{}
	res := dbConn.First(&result, "id = ?", id)

	if res.Error != nil {
		log.Errorf("Can't create result %v\n%s", res, res.Error)
		return Result{}, res.Error
	}

	return result, nil
}

// Create inserts an Result into DB
func (r RRepository) Create(result Result) (string, error) {
	res := dbConn.Create(&result)

	if res.Error != nil {
		log.Errorf("Can't create result %v\n%s", result, res.Error)
		return "", res.Error
	}

	return result.ID, nil
}

// CreateAll inserts an Results into DB
func (r RRepository) CreateAll(results []Result) bool {
	result := dbConn.Create(&results)

	if result.Error != nil {
		log.Errorf("Can't create result %v\n%s", results, result.Error)
		return false
	}

	return true
}

// Update updates an Result in the DB (not used for now)
func (r RRepository) Update(result Result) error {
	res := dbConn.Model(&result).Updates(result)

	if res.Error != nil {
		log.Error("Can't update result with values %v\n%s", result, res.Error)
		return res.Error
	}
	return nil
}

// DeleteByID deletes an Result (not used for now)
func (r RRepository) DeleteByID(id string) error {
	result := dbConn.Delete(&Result{}, id)

	if result.Error != nil {
		log.Error("Can't delete result with id %s\n%s", id, result.Error)
		return result.Error
	}

	return nil
}

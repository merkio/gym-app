package exercise

import (
	config "gym-app/app-config"
	"gym-app/app/model"
	"gym-app/common/db"
	loggerWrap "gym-app/common/logger"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository exercise repository
type Repository interface {
	repo.BaseRepository
	CreateAll(exercises []model.Exercise) bool
	GetByID(id string) (model.Exercise, error)
	Get() []model.Exercise
}

// ERepository instance of Repository
type ERepository struct {
	Repository
}

// GetDB get connect to the db
func GetDB(conf config.DataConnectionConf, app string) *gorm.DB {
	dbConn := db.GetDBInstance(&db.Specification{
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

// Get returns the list of Exercises
func (r ERepository) Get() []model.Exercise {
	exercises := make([]model.Exercise, 30)
	result := dbConn.Find(&exercises)

	if result.Error != nil {
		log.Error("Can't get exercises from dbConn.\n", result.Error)
	}

	log.Infof("Found %d amount of exercises", result.RowsAffected)
	return exercises
}

// GetByID return the Exercise with id
func (r ERepository) GetByID(id string) (model.Exercise, error) {
	exercise := model.Exercise{}
	result := dbConn.First(&exercise, "id = ?", id)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%v", exercise, result.Error)
		return model.Exercise{}, result.Error
	}

	return exercise, nil
}

// Create inserts an Exercise into DB
func (r ERepository) Create(exercise model.Exercise) (string, error) {
	result := dbConn.Create(&exercise)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%v", exercise, result.Error)
		return "", result.Error
	}

	return exercise.ID, nil
}

// CreateAll inserts an Exercises into DB
func (r ERepository) CreateAll(exercises []model.Exercise) bool {
	result := dbConn.Create(&exercises)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%v", exercises, result.Error)
		return false
	}

	return true
}

// Update updates an Exercise in the DB (not used for now)
func (r ERepository) Update(exercise model.Exercise) error {
	result := dbConn.Model(&exercise).Updates(exercise)

	if result.Error != nil {
		log.Errorf("Can't update exercise with values %v\n%v", exercise, result.Error)
		return result.Error
	}
	return nil
}

// DeleteByID deletes an Exercise (not used for now)
func (r ERepository) DeleteByID(id string) error {
	result := dbConn.Delete(&model.Exercise{}, id)

	if result.Error != nil {
		log.Errorf("Can't delete exercise with id %s\n%v", id, result.Error)
		return result.Error
	}

	return nil
}

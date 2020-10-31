package exercise

import (
	config "gym-app/app-config"
	dbConnector "gym-app/common/db"
	loggerWrap "gym-app/common/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func init() {
	db = GetDB(config.DataConnectionConf, config.App)
}

// GetExercises returns the list of Exercises
func (r Repository) GetExercises() []Exercise {
	exercises := make([]Exercise, 30)
	result := db.Find(&exercises)

	if result.Error != nil {
		log.Error("Can't get exercises from db.\n%s", result.Error)
	}

	log.Infof("Found %d amount of exercises", result.RowsAffected)
	return exercises
}

// GetExercise return the Exercise with id
func (r Repository) GetExercise(id uuid.UUID) (Exercise, err) {
	exercise := Exercise{}
	result := db.First(&exercise, id)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%s", exercise, result.Error)
		return nil, result.Error
	}

	return exercise, nil
}

// AddExercise inserts an Exercise into DB
func (r Repository) AddExercise(exercise Exercise) bool {
	result := db.Create(&exercise)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%s", exercise, result.Error)
		return false
	}

	return true
}

// AddExercises inserts an Exercises into DB
func (r Repository) AddExercises(exercises []Exercise) bool {
	result := db.Create(&exercises)

	if result.Error != nil {
		log.Errorf("Can't create exercise %v\n%s", exercises, result.Error)
		return false
	}

	return true
}

// UpdateExercise updates an Exercise in the DB (not used for now)
func (r Repository) UpdateExercise(exercise Exercise) bool {
	result := db.Model(&exercise).Updates(exercise)

	if result.Error != nil {
		log.Error("Can't update exercise with values %v\n%s", exercise, result.Error)
		return false
	}
	return true
}

// DeleteExercise deletes an Exercise (not used for now)
func (r Repository) DeleteExercise(id uuid.UUID) string {
	result := db.Delete(&Exercise{}, id)

	if result.Error != nil {
		log.Error("Can't delete exercise with id %s\n%s", id, result.Error)
		return "Error"
	}

	return "OK"
}

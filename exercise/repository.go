package exercise

import (
	dbConnector "gym-app/common/db"
	loggerWrap "gym-app/common/logger"

	"gorm.io/gorm"
)

//Repository ...
type Repository struct{}

var db *gorm.DB
var appLog *loggerWrap.Logger

func init() {
	appLog = loggerWrap.Log
	dbConnector.Connect()
	db = dbConnector.GetDB()
	db.AutoMigrate(&Exercise{})
}

// GetExercises returns the list of Exercises
func (r Repository) GetExercises() []Exercise {
	exercises := make([]Exercise, 30)
	result := db.Find(&exercises)

	if result.Error != nil {
		appLog.Error("Can't get exercises from db.\n%s", result.Error)
	}

	appLog.Infof("Found %d amount of exercises", result.RowsAffected)
	return exercises
}

// AddExercise inserts an Exercise into DB
func (r Repository) AddExercise(exercise Exercise) bool {
	result := db.Create(&exercise)

	if result.Error != nil {
		appLog.Errorf("Can't create exercise %v\n%s", exercise, result.Error)
		return false
	}

	return true
}

// AddExercises inserts an Exercises into DB
func (r Repository) AddExercises(exercises []Exercise) bool {
	result := db.Create(&exercises)

	if result.Error != nil {
		appLog.Errorf("Can't create exercise %v\n%s", exercises, result.Error)
		return false
	}

	return true
}

// UpdateExercise updates an Exercise in the DB (not used for now)
func (r Repository) UpdateExercise(exercise Exercise) bool {
	result := db.Model(&exercise).Updates(exercise)

	if result.Error != nil {
		appLog.Error("Can't update exercise with values %v\n%s", exercise, result.Error)
		return false
	}
	return true
}

// DeleteExercise deletes an Exercise (not used for now)
func (r Repository) DeleteExercise(id string) string {
	result := db.Delete(&Exercise{}, id)

	if result.Error != nil {
		appLog.Error("Can't delete exercise with id %s\n%s", id, result.Error)
		return "Error"
	}

	return "OK"
}

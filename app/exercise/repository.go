package exercise

import (
	"gym-app/app/model"
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

// ERepository instance of PRepository
type ERepository struct {
	Repository
	db  *gorm.DB
	log *logrus.Logger
}

func NewERepository(conn *gorm.DB, logger *logrus.Logger) ERepository {
	return ERepository{
		db:  conn,
		log: logger,
	}
}

// Get returns the list of Exercises
func (r ERepository) Get() []model.Exercise {
	exercises := make([]model.Exercise, 30)
	result := r.db.Find(&exercises)

	if result.Error != nil {
		r.log.Error("Can't get exercises from dbConn.\n", result.Error)
	}

	r.log.Infof("Found %d amount of exercises", result.RowsAffected)
	return exercises
}

// GetByID return the Exercise with id
func (r ERepository) GetByID(id string) (model.Exercise, error) {
	exercise := model.Exercise{}
	result := r.db.First(&exercise, "id = ?", id)

	if result.Error != nil {
		r.log.Errorf("Can't create exercise %v\n%v", exercise, result.Error)
		return model.Exercise{}, result.Error
	}

	return exercise, nil
}

// Create inserts an Exercise into DB
func (r ERepository) Create(exercise model.Exercise) (string, error) {
	result := r.db.Create(&exercise)

	if result.Error != nil {
		r.log.Errorf("Can't create exercise %v\n%v", exercise, result.Error)
		return "", result.Error
	}

	return exercise.ID, nil
}

// CreateAll inserts an Exercises into DB
func (r ERepository) CreateAll(exercises []model.Exercise) bool {
	result := r.db.Create(&exercises)

	if result.Error != nil {
		r.log.Errorf("Can't create exercise %v\n%v", exercises, result.Error)
		return false
	}

	return true
}

// Update updates an Exercise in the DB (not used for now)
func (r ERepository) Update(exercise model.Exercise) error {
	result := r.db.Model(&exercise).Updates(exercise)

	if result.Error != nil {
		r.log.Errorf("Can't update exercise with values %v\n%v", exercise, result.Error)
		return result.Error
	}
	return nil
}

// DeleteByID deletes an Exercise (not used for now)
func (r ERepository) DeleteByID(id string) error {
	result := r.db.Delete(&model.Exercise{}, id)

	if result.Error != nil {
		r.log.Errorf("Can't delete exercise with id %s\n%v", id, result.Error)
		return result.Error
	}

	return nil
}

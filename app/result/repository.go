package result

import (
	"gym-app/app/model"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository results repository
type Repository interface {
	repo.BaseRepository
	CreateAll(results []model.Result) bool
	GetByID(id string) (model.Result, error)
	Get() []model.Result
}

// RRepository instance of PRepository
type RRepository struct {
	Repository
	db  *gorm.DB
	log *logrus.Logger
}

func NewRRepository(conn *gorm.DB, logger *logrus.Logger) RRepository {
	return RRepository{
		db:  conn,
		log: logger,
	}
}

// Get returns the list of Results
func (r RRepository) Get() []model.Result {
	results := make([]model.Result, 30)
	result := r.db.Find(&results)

	if result.Error != nil {
		r.log.Error("Can't get results from DB\n", result.Error)
	}

	r.log.Infof("Found %d amount of results", result.RowsAffected)
	return results
}

// GetByID return the Result with id
func (r RRepository) GetByID(id string) (model.Result, error) {
	result := model.Result{}
	res := r.db.First(&result, "id = ?", id)

	if res.Error != nil {
		r.log.Errorf("Can't create result %v\n%v", res, res.Error)
		return model.Result{}, res.Error
	}

	return result, nil
}

// Create inserts an Result into DB
func (r RRepository) Create(result model.Result) (string, error) {
	res := r.db.Create(&result)

	if res.Error != nil {
		r.log.Errorf("Can't create result %v\n%v", result, res.Error)
		return "", res.Error
	}

	return result.ID, nil
}

// CreateAll inserts an Results into DB
func (r RRepository) CreateAll(results []model.Result) bool {
	result := r.db.Create(&results)

	if result.Error != nil {
		r.log.Errorf("Can't create result %v\n%v", results, result.Error)
		return false
	}

	return true
}

// Update updates an Result in the DB (not used for now)
func (r RRepository) Update(result model.Result) error {
	res := r.db.Model(&result).Updates(result)

	if res.Error != nil {
		r.log.WithField("err", res.Error).
			WithField("values", result).
			Error("Can't update result with values")
		return res.Error
	}
	return nil
}

// DeleteByID deletes an Result (not used for now)
func (r RRepository) DeleteByID(id string) error {
	result := r.db.Delete(&model.Result{}, id)

	if result.Error != nil {
		r.log.WithField("err", result.Error).
			WithField("id", id).
			Error("Can't delete result with id")
		return result.Error
	}

	return nil
}

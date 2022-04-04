package program

import (
	"gym-app/app/model"
	"gym-app/common/utils"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository programs repository
type Repository interface {
	repo.BaseRepository
	CreateAll(programs []model.Program) bool
	GetByID(id string) (model.Program, error)
	Get() []model.Program
	GetByText(text string) bool
	Search(request model.SearchRequest)
}

// PRepository instance of PRepository
type PRepository struct {
	Repository
	db  *gorm.DB
	log *logrus.Logger
}

func NewPRepository(conn *gorm.DB, logger *logrus.Logger) PRepository {
	return PRepository{
		db:  conn,
		log: logger,
	}
}

// Get returns the list of Programs
func (r PRepository) Get() []model.Program {
	programs := make([]model.Program, 30)
	result := r.db.Limit(30).Find(&programs)

	if result.Error != nil {
		r.log.Error("Can't get programs from db.\n", result.Error)
	}

	r.log.Infof("Found %d amount of programs", result.RowsAffected)
	return programs
}

// GetByText get program by text
func (r PRepository) GetByText(text string) bool {
	program := model.Program{}
	result := r.db.Where("text = ?", text).First(&program)

	if result.Error != nil {
		r.log.Errorf("Can't find the program with text %.30s\n%v", text, result.Error)
		return false
	}

	return true
}

// CountByGroupID count messages by group ID
func (r PRepository) CountByGroupID(groupID string) int64 {
	var result int64
	r.db.Table(model.Program{}).Where("group_id = ?", groupID).Count(&result)
	return result
}

// GetByID return the Program with id
func (r PRepository) GetByID(id string) (model.Program, error) {
	program := model.Program{}
	result := r.db.First(&program, "id = ?", id)

	if result.Error != nil {
		r.log.Errorf("Can't create program %v\n%v", program, result.Error)
		return model.Program{}, result.Error
	}

	return program, nil
}

// Create inserts a Program into DB
func (r PRepository) Create(program model.Program) (string, error) {
	result := r.db.Create(&program)

	if result.Error != nil {
		r.log.Errorf("Can't create program %v\n%v", program, result.Error)
		return "", result.Error
	}

	return program.ID, nil
}

// CreateAll inserts an Programs into DB
func (r PRepository) CreateAll(programs []model.Program) bool {
	result := r.db.Create(&programs)

	if result.Error != nil {
		r.log.Errorf("Can't create program %v\n%v", programs, result.Error)
		return false
	}

	return true
}

// Update updates an Program in the DB (not used for now)
func (r PRepository) Update(program model.Program) error {
	result := r.db.Model(&program).Updates(program)

	if result.Error != nil {
		r.log.Errorf("Can't update program with values %v\n%v", program, result.Error)
		return result.Error
	}
	return nil
}

// Search with params
func (r PRepository) Search(params model.SearchRequest) ([]model.Program, error) {
	programs := make([]model.Program, params.Limit)

	if params.Limit == 0 {
		params.Limit = 20
	}
	query := r.db.Limit(params.Limit)
	utils.CreateSearchProgramQuery(&params, query)
	r.log.WithField("params", params).Info("Search programs between dates")

	result := query.Find(&programs)

	if result.Error != nil {
		r.log.Errorf("Can't find program with params %v\n%v", params, result.Error)
		return programs, result.Error
	}
	return programs, nil
}

// DeleteByID deletes an Program (not used for now)
func (r PRepository) DeleteByID(id string) error {
	result := r.db.Delete(&model.Program{}, id)

	if result.Error != nil {
		r.log.Errorf("Can't delete program with id %s\n%v", id, result.Error)
		return result.Error
	}

	return nil
}

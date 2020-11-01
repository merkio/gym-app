package program

import (
	config "gym-app/app-config"
	"gym-app/common/db"
	loggerWrap "gym-app/common/logger"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository programs repository
type Repository interface {
	repo.BaseRepository
	CreateAll(programs []Program) bool
	GetByID(id string) (Program, error)
	Get() []Program
}

// PRepository instance of PRepository
type PRepository struct {
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

// Get returns the list of Programs
func (r PRepository) Get() []Program {
	programs := make([]Program, 30)
	result := dbConn.Find(&programs)

	if result.Error != nil {
		log.Error("Can't get programs from dbConn.\n%s", result.Error)
	}

	log.Infof("Found %d amount of programs", result.RowsAffected)
	return programs
}

// GetByID return the Program with id
func (r PRepository) GetByID(id string) (Program, error) {
	program := Program{}
	result := dbConn.First(&program, "id = ?", id)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", program, result.Error)
		return Program{}, result.Error
	}

	return program, nil
}

// Create inserts an Program into DB
func (r PRepository) Create(program Program) (string, error) {
	result := dbConn.Create(&program)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", program, result.Error)
		return "", result.Error
	}

	return program.ID, nil
}

// CreateAll inserts an Programs into DB
func (r PRepository) CreateAll(programs []Program) bool {
	result := dbConn.Create(&programs)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", programs, result.Error)
		return false
	}

	return true
}

// Update updates an Program in the DB (not used for now)
func (r PRepository) Update(program Program) error {
	result := dbConn.Model(&program).Updates(program)

	if result.Error != nil {
		log.Error("Can't update program with values %v\n%s", program, result.Error)
		return result.Error
	}
	return nil
}

// DeleteByID deletes an Program (not used for now)
func (r PRepository) DeleteByID(id string) error {
	result := dbConn.Delete(&Program{}, id)

	if result.Error != nil {
		log.Error("Can't delete program with id %s\n%s", id, result.Error)
		return result.Error
	}

	return nil
}

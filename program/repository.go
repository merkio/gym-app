package program

import (
	config "gym-app/app-config"
	dbConnector "gym-app/common/db"
	loggerWrap "gym-app/common/logger"

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
		SearchPath: conf.PostgresSchema
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

// GetPrograms returns the list of Programs
func (r Repository) GetPrograms() []Program {
	programs := make([]Program, 30)
	result := db.Find(&programs)

	if result.Error != nil {
		log.Error("Can't get programs from db.\n%s", result.Error)
	}

	log.Infof("Found %d amount of programs", result.RowsAffected)
	return programs
}

// GetProgram return the Program with id
func (r Repository) GetProgram(id int64) Program {
	result := db.Create(&program)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", program, result.Error)
		return false
	}

	return true
}

// AddProgram inserts an Program into DB
func (r Repository) AddProgram(program Program) bool {
	result := db.Create(&program)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", program, result.Error)
		return false
	}

	return true
}

// AddPrograms inserts an Programs into DB
func (r Repository) AddPrograms(programs []Program) bool {
	result := db.Create(&programs)

	if result.Error != nil {
		log.Errorf("Can't create program %v\n%s", programs, result.Error)
		return false
	}

	return true
}

// UpdateProgram updates an Program in the DB (not used for now)
func (r Repository) UpdateProgram(program Program) bool {
	result := db.Model(&program).Updates(program)

	if result.Error != nil {
		log.Error("Can't update program with values %v\n%s", program, result.Error)
		return false
	}
	return true
}

// DeleteProgram deletes an Program (not used for now)
func (r Repository) DeleteProgram(id string) string {
	result := db.Delete(&Program{}, id)

	if result.Error != nil {
		log.Error("Can't delete program with id %s\n%s", id, result.Error)
		return "Error"
	}

	return "OK"
}

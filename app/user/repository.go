package user

import (
	config "gym-app/app-config"
	"gym-app/common/db"
	loggerWrap "gym-app/common/logger"
	repo "gym-app/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository user repository
type Repository interface {
	repo.BaseRepository
	CreateAll(users []User) bool
	GetByID(id string) (User, error)
	Get() []User
}

// URepository instance of Repository
type URepository struct {
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

// Get returns the list of Users
func (r URepository) Get() []User {
	users := make([]User, 30)
	result := dbConn.Find(&users)

	if result.Error != nil {
		log.Error("Can't get users from dbConn.\n", result.Error)
	}

	log.Infof("Found %d amount of users", result.RowsAffected)
	return users
}

// GetByID return the User with id
func (r URepository) GetByID(id string) (User, error) {
	user := User{}
	result := dbConn.First(&user, "id = ?", id)

	if result.Error != nil {
		log.Errorf("Can't create user %v\n%v", user, result.Error)
		return User{}, result.Error
	}

	return user, nil
}

// Create inserts an User into DB
func (r URepository) Create(user User) (string, error) {
	result := dbConn.Create(&user)

	if result.Error != nil {
		log.Errorf("Can't create user %v\n%v", user, result.Error)
		return "", result.Error
	}

	return user.ID, nil
}

// CreateAll inserts an Users into DB
func (r URepository) CreateAll(users []User) bool {
	result := dbConn.Create(&users)

	if result.Error != nil {
		log.Errorf("Can't create user %v\n%v", users, result.Error)
		return false
	}

	return true
}

// Update updates an User in the DB (not used for now)
func (r URepository) Update(user User) error {
	result := dbConn.Model(&user).Updates(user)

	if result.Error != nil {
		log.Errorf("Can't update user with values %v\n%v", user, result.Error)
		return result.Error
	}
	return nil
}

// DeleteByID deletes an User (not used for now)
func (r URepository) DeleteByID(id string) error {
	result := dbConn.Delete(&User{}, id)

	if result.Error != nil {
		log.Errorf("Can't delete user with id %s\n%v", id, result.Error)
		return result.Error
	}

	return nil
}

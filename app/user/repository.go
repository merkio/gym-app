package user

import (
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
	db *gorm.DB
	log *logrus.Logger
}

func NewURepository(conn *gorm.DB, logger *logrus.Logger) URepository {
	return URepository{
		db: conn,
		log: logger,
	}
}

// Get returns the list of Users
func (r URepository) Get() []User {
	users := make([]User, 30)
	result := r.db.Find(&users)

	if result.Error != nil {
		r.log.Error("Can't get users from dbConn.\n", result.Error)
	}

	r.log.Infof("Found %d amount of users", result.RowsAffected)
	return users
}

// GetByID return the User with id
func (r URepository) GetByID(id string) (User, error) {
	user := User{}
	result := r.db.First(&user, "id = ?", id)

	if result.Error != nil {
		r.log.Errorf("Can't create user %v\n%v", user, result.Error)
		return User{}, result.Error
	}

	return user, nil
}

// Create inserts an User into DB
func (r URepository) Create(user User) (string, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		r.log.Errorf("Can't create user %v\n%v", user, result.Error)
		return "", result.Error
	}

	return user.ID, nil
}

// CreateAll inserts an Users into DB
func (r URepository) CreateAll(users []User) bool {
	result := r.db.Create(&users)

	if result.Error != nil {
		r.log.Errorf("Can't create user %v\n%v", users, result.Error)
		return false
	}

	return true
}

// Update updates an User in the DB (not used for now)
func (r URepository) Update(user User) error {
	result := r.db.Model(&user).Updates(user)

	if result.Error != nil {
		r.log.Errorf("Can't update user with values %v\n%v", user, result.Error)
		return result.Error
	}
	return nil
}

// DeleteByID deletes an User (not used for now)
func (r URepository) DeleteByID(id string) error {
	result := r.db.Delete(&User{}, id)

	if result.Error != nil {
		r.log.Errorf("Can't delete user with id %s\n%v", id, result.Error)
		return result.Error
	}

	return nil
}

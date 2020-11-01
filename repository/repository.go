package repository

// BaseRepository - Base interface to create, update and delete operations
type BaseRepository interface {
	Create(entity interface{}) (string, error)
	Update(entity interface{}) error
	DeleteByID(id string) error
}

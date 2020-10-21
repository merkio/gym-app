package exercise

import (
	"time"

	"gorm.io/gorm"
)

// Tabler configuration for gorm to specify table name
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Exercise to `exercises`
func (Exercise) TableName() string {
	return "exercises"
}

//Exercise represents an exercise
type Exercise struct {
	gorm.Model
	ID          int64     `gorm:"primary_key;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	VideoLink   string    `json:"videoLink"`
	Tags        string    `json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

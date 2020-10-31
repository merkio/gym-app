package result

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tabler configuration for gorm to specify table name
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Result to `results`
func (Result) TableName() string {
	return "results"
}

//Result represents an exercise
type Result struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4(); json:"id"`
	Text      string    `gorm:"index"; json:"text"`
	Tags      []string  `gorm:"index"; json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

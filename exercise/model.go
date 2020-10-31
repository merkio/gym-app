package exercise

import (
	"gym-app/result"
	"time"

	"github.com/google/uuid"
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
	ID          uuid.UUID       `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4(); json:"id"`
	Title       string          `gorm:"unique"; json:"title"`
	Description string          `json:"description"`
	VideoLink   string          `json:"videoLink"`
	Tags        string          `gorm:"index"; json:"tags"`
	Results     []result.Result `gorm:"many2many:exercise_results;foreignKey:ID;joinForeignKey:ResultID;References:ExerciseID;JoinReferences:ID"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

package program

import (
	"gym-app/exercise"
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
	return "programs"
}

//Program represents an exercise
type Program struct {
	gorm.Model
	ID        uuid.UUID           `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4(); json:"id"`
	Text      string              `gorm:"index"; json:"text"`
	Tags      []string            `gorm:"index"; json:"tags"`
	Date      time.Time           `json:"date"`
	DateInt   int64               `json:dateInt`
	Exercises []exercise.Exercise `gorm:"many2many:program_exercises;foreignKey:ID;joinForeignKey:ExerciseID;References:ProgramID;JoinReferences:ID"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
}

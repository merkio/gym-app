package program

import (
	"gym-app/exercise"
	"time"

	"gorm.io/gorm"
)

//Program represents an exercise
type Program struct {
	gorm.Model
	ID        string              `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4();" json:"id"`
	Text      string              `json:"text"`
	Tags      string              `gorm:"index" json:"tags"`
	Date      time.Time           `json:"date"`
	DateInt   int64               `json:"dateInt"`
	Exercises []exercise.Exercise `gorm:"many2many:program_exercises;" json:"exercises"`
}

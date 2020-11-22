package exercise

import (
	"gorm.io/gorm"
	"gym-app/result"
)

//Exercise represents an exercise
type Exercise struct {
	gorm.Model
	ID          string          `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	Title       string          `gorm:"unique" json:"title"`
	Description string          `json:"description"`
	VideoLink   string          `json:"videoLink"`
	Image		string			`json:"image"`
	Group		string			`json:"group"`
	Tags        string          `gorm:"index; type:string;" json:"tags"`
	Results     []result.Result `gorm:"many2many:exercise_results" json:"results"`
}

package result

import (
	"gorm.io/gorm"
)

//Result represents an result of the exercise
type Result struct {
	gorm.Model
	ID   string `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4(); json:"id"`
	Text string `json:"text"`
	Tags string `json:"tags"`
}

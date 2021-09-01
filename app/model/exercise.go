package model

import (
	"gorm.io/gorm"
)

//Exercise represents an exercise
type Exercise struct {
	gorm.Model
	ID          string   `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	Title       string   `gorm:"unique" json:"title"`
	Description string   `json:"description"`
	VideoLink   string   `json:"videoLink"`
	Image       string   `json:"image"`
	Labels      []string `gorm:"index; type:string;" json:"group"`
	Tags        string   `gorm:"index; type:string;" json:"tags"`
}

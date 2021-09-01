package model

import "gorm.io/gorm"

//Group represents an exercise
type Group struct {
	gorm.Model
	ID          string   `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Labels      []string `gorm:"index; type:string" json:"group"`
	Tags        string   `gorm:"index; type:string;" json:"tags"`
	Result      Result   `json:"result"`
}

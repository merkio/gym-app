package user

import (
	"gorm.io/gorm"
)

//User represents an exercise
type User struct {
	gorm.Model
	ID         string `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	Name       string `gorm:"unique" json:"name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Image      string `json:"image"`
}

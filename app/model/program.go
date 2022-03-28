package model

import (
	"time"

	"gorm.io/gorm"
)

//Program represents an exercise
type Program struct {
	gorm.Model
	ID        string     `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4();" json:"id"`
	Text      string     `json:"text"`
	Tags      string     `gorm:"index" json:"tags"`
	Date      time.Time  `json:"date"`
	DateInt   int64      `json:"dateInt"`
	GroupName string     `gorm:"default:UDARNIK" json:"group_name"`
	GroupID   string     `gorm:"default:-62011928" json:"group_id"`
	Exercises []Exercise `gorm:"many2many:program_exercises;" json:"exercises"`
}

// SearchRequest request with search params
type SearchRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sort_by"`
	Order     string `json:"order"`
	Text      string `json:"text"`
}

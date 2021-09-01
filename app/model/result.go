package model

import (
	"gorm.io/gorm"
)

//Result represents an result of the exercise
type Result struct {
	gorm.Model
	ID         string     `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	UserID     string     `json:"user_id"`
	MetricType MetricType `json:"metric_type"`
	Type       ResultType `json:"type"`
	EntityID   string     `json:"entity_id"`
}

type MetricType string

const (
	Time        MetricType = "time"
	Repetitions MetricType = "repetitions"
	Weight      MetricType = "weight (kg)"
)

type ResultType string

const (
	ExerciseType ResultType = "exercise"
	GroupType    ResultType = "group"
)

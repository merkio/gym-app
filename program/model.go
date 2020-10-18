package program

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Program represents an exercise
type Program struct {
	ID         bson.ObjectId `bson:"_id"`
	Text       string        `json:"text"`
	Tags       []string      `json:"tags"`
	Date       time.Time     `json:"date"`
	CreatedOn  time.Time     `json:"createdOn"`
	ModifiedOn time.Time     `json:"modifiedOn"`
}

//Programs is an array of Programs
type Programs []Program

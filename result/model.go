package result

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Result represents an exercise
type Result struct {
	ID         bson.ObjectId `bson:"_id"`
	Text       string        `json:"text"`
	Tags       []string      `json:"tags"`
	Date       time.Time     `json:"date"`
	CreatedOn  time.Time     `json:"createdOn"`
	ModifiedOn time.Time     `json:"modifiedOn"`
}

//Results is an array of Results
type Results []Result

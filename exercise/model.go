package exercise

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Exercise represents an exercise
type Exercise struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	VideoLink   string        `json:"videoLink"`
	Tags        []string      `json:"tags"`
	CreatedOn   time.Time     `json:"createdOn"`
	ModifiedOn  time.Time     `json:"modifiedOn"`
}

//Exercises is an array of Exercises
type Exercises []Exercise

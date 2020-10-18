package exercise

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DB the name of the DB instance
const DB = "programs"

// DOCNAME the name of the document
const DOCNAME = "exercises"

// GetExercises returns the list of Exercises
func (r Repository) GetExercises() Exercises {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DB).C(DOCNAME)
	results := Exercises{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddExercise inserts an Exercise in the DB
func (r Repository) AddExercise(exercise Exercise) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	exercise.ID = bson.NewObjectId()
	exercise.CreatedOn = time.Now()
	exercise.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).Insert(exercise)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateExercise updates an Exercise in the DB (not used for now)
func (r Repository) UpdateExercise(exercise Exercise) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	exercise.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).UpdateId(exercise.ID, exercise)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteExercise deletes an Exercise (not used for now)
func (r Repository) DeleteExercise(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DB).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}

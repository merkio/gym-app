package program

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
const DOCNAME = "programs"

// GetPrograms returns the list of Programs
func (r Repository) GetPrograms() Programs {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DB).C(DOCNAME)
	results := Programs{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddProgram inserts an Program in the DB
func (r Repository) AddProgram(program Program) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	program.ID = bson.NewObjectId()
	program.CreatedOn = time.Now()
	program.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).Insert(program)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateProgram updates an Program in the DB (not used for now)
func (r Repository) UpdateProgram(program Program) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	program.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).UpdateId(program.ID, program)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteProgram deletes an Program (not used for now)
func (r Repository) DeleteProgram(id string) string {
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

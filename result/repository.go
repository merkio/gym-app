package result

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
const DOCNAME = "results"

// GetResults returns the list of Results
func (r Repository) GetResults() Results {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DB).C(DOCNAME)
	results := Results{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddResult inserts an Result in the DB
func (r Repository) AddResult(result Result) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	result.ID = bson.NewObjectId()
	result.CreatedOn = time.Now()
	result.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).Insert(result)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateResult updates an Result in the DB (not used for now)
func (r Repository) UpdateResult(result Result) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	result.ModifiedOn = time.Now()
	session.DB(DB).C(DOCNAME).UpdateId(result.ID, result)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteResult deletes an Result (not used for now)
func (r Repository) DeleteResult(id string) string {
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

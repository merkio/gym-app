package exercise

import (
	"encoding/json"
	"gym-app/common/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Controller ...
type Controller struct {
	Repository Repository
}

var log *logrus.Logger

func init() {
	log = logger.NewLogger()
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	exercises := c.Repository.GetExercises() // list of all exercises
	log.Info(exercises)
	data, _ := json.Marshal(exercises)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// GetExercise GET /{id}
func (c *Controller) GetExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                                              // param id
	if exercise, err := c.Repository.GetExercise(id); err != "" { // delete a exercise by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return

	data, _ := json.Marshal(exercise)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddExercise POST /
func (c *Controller) AddExercise(w http.ResponseWriter, r *http.Request) {
	var exercise Exercise
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error AddExercise", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error AddExercise", err)
	}
	if err := json.Unmarshal(body, &exercise); err != nil { // unmarshal body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error AddExercise unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.AddExercise(exercise) // adds the exercise to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateExercise PUT /
func (c *Controller) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise Exercise
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error UpdateExercise", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error UpdateExercise", err)
	}
	if err := json.Unmarshal(body, &exercise); err != nil { // unmarshal body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error UpdateExercise unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.UpdateExercise(exercise) // updates the exercise in the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteExercise DELETE /
func (c *Controller) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                                       // param id
	if err := c.Repository.DeleteExercise(id); err != "" { // delete a exercise by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

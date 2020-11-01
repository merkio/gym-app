package exercise

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository ERepository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	exercises := c.Repository.Get() // list of all exercises
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
	var exercise Exercise
	var err error

	id := vars["id"]
	if exercise, err = c.Repository.GetByID(id); err != nil { // get an exercise by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	id, err := c.Repository.Create(exercise) // adds the exercise to the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
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
	err = c.Repository.Update(exercise) // updates the exercise in the DB
	if err != nil {
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

	id := vars["id"]
	if err := c.Repository.DeleteByID(id); err != nil { // delete a exercise by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
	return
}

package exercise

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	config "gym-app/app-config"
	"gym-app/app/model"
	"gym-app/common/db"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	repository ERepository
	log        *logrus.Logger
}

func NewController(logger *logrus.Logger) Controller {
	return Controller{
		log:        logger,
		repository: NewERepository(db.GetDB(config.DataConnectionConfig), logger),
	}
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	exercises := c.repository.Get() // list of all exercises
	c.log.Info(exercises)
	data, _ := json.Marshal(exercises)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// GetExercise GET /{id}
func (c *Controller) GetExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var exercise model.Exercise
	var err error

	id := vars["id"]
	if exercise, err = c.repository.GetByID(id); err != nil { // get an exercise by id
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
	var exercise model.Exercise
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.Error("Error AddExercise", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		c.log.Error("Error AddExercise", err)
	}
	if err := json.Unmarshal(body, &exercise); err != nil { // unmarshal body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			c.log.Error("Error AddExercise unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	id, err := c.repository.Create(exercise) // adds the exercise to the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(id))
}

// UpdateExercise PUT /
func (c *Controller) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise model.Exercise
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.Error("Error UpdateExercise", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		c.log.Error("Error UpdateExercise", err)
	}
	if err := json.Unmarshal(body, &exercise); err != nil { // unmarshal body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			c.log.Error("Error UpdateExercise unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = c.repository.Update(exercise) // updates the exercise in the DB
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
	if err := c.repository.DeleteByID(id); err != nil { // delete a exercise by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

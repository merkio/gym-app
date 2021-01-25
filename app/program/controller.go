package program

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	repository PRepository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	programs := c.repository.Get() // list of all programs
	log.Info("Found programs: ", len(programs))
	data, _ := json.Marshal(programs)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// GetProgram GET /id
func (c *Controller) GetProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var program Program
	var err error

	if program, err = c.repository.GetByID(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info("Found program: ", program)
	data, _ := json.Marshal(program)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// AddProgram POST /
func (c *Controller) AddProgram(w http.ResponseWriter, r *http.Request) {
	var program Program
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error AddProgram", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error AddProgram", err)
	}
	if err := json.Unmarshal(body, &program); err != nil { // unmarshal body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error AddProgram unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	id, err := c.repository.Create(program) // adds the program to the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(id))
}

// UpdateProgram PUT /
func (c *Controller) UpdateProgram(w http.ResponseWriter, r *http.Request) {
	var program Program
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error UpdateProgram", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error UpdateProgram", err)
	}
	if err := json.Unmarshal(body, &program); err != nil { // unmarshal body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error UpdateProgram unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = c.repository.Update(program) // updates the program in the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteProgram DELETE /
func (c *Controller) DeleteProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if err := c.repository.DeleteByID(id); err != nil { // delete a program by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}
package user

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	repository URepository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	users := c.repository.Get() // list of all users
	log.Info(users)
	data, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// GetUser GET /{id}
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	var err error

	id := vars["id"]
	if user, err = c.repository.GetByID(id); err != nil { // get an user by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// AddUser POST /
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error AddUser", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error AddUser", err)
	}
	if err := json.Unmarshal(body, &user); err != nil { // unmarshal body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error AddUser unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	id, err := c.repository.Create(user) // adds the user to the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(id))
}

// UpdateUser PUT /
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Error("Error UpdateUser", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error("Error UpdateUser", err)
	}
	if err := json.Unmarshal(body, &user); err != nil { // unmarshal body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Error("Error UpdateUser unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = c.repository.Update(user) // updates the user in the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

// DeleteUser DELETE /
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if err := c.repository.DeleteByID(id); err != nil { // delete a user by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

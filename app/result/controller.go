package result

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
	repository RRepository
	log        *logrus.Logger
}

func NewController(logger *logrus.Logger) Controller {
	return Controller{
		log:        logger,
		repository: NewRRepository(db.GetDB(config.DataConnectionConfig), logger),
	}
}
// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	results := c.repository.Get() // list of all results
	c.log.Info("Found results: ", len(results))
	data, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// GetResult GET /id
func (c *Controller) GetResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var result model.Result
	var err error

	if result, err = c.repository.GetByID(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.log.Info("Found result: ", result)
	data, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// AddResult POST /
func (c *Controller) AddResult(w http.ResponseWriter, r *http.Request) {
	var result model.Result
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.Error("Error AddResult", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		c.log.Error("Error AddResult", err)
	}
	if err := json.Unmarshal(body, &result); err != nil { // unmarshal body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			c.log.Error("Error AddResult unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	id, err := c.repository.Create(result) // adds the result to the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(id))
}

// UpdateResult PUT /
func (c *Controller) UpdateResult(w http.ResponseWriter, r *http.Request) {
	var result model.Result
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.log.Error("Error UpdateResult", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		c.log.Error("Error UpdateResult", err)
	}
	if err := json.Unmarshal(body, &result); err != nil { // unmarshal body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			c.log.Error("Error UpdateResult unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = c.repository.Update(result) // updates the result in the DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

// DeleteResult DELETE /
func (c *Controller) DeleteResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if err := c.repository.DeleteByID(id); err != nil { // delete a result by id
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

package main

import (
	"gym-app/result"
	"log"
	"net/http"

	"gym-app/program"

	"gym-app/exercise"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true) // create routes
	exercise.NewSubRouter(router)
	program.NewSubRouter(router)
	result.NewSubRouter(router)

	// these two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	// launch server with CORS validations
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

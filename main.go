package main

import (
	"gym-app/tasks"
	"log"
	"net/http"

	config "gym-app/app-config"
	"gym-app/exercise"
	"gym-app/program"
	"gym-app/result"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true) // create routes
	exercise.NewSubRouter(router)
	program.NewSubRouter(router)
	result.NewSubRouter(router)

	db := exercise.GetDB(config.DataConnectionConfig, config.App)
	err := db.AutoMigrate(&result.Result{}, &exercise.Exercise{}, &program.Program{})
	if err != nil {
		log.Fatal(err)
	}

	// these two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Run vk tasks
	go tasks.CollectVkMessages()

	// launch server with CORS validations
	log.Print("Starting server on port :9000")
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

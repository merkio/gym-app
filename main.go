package main

import (
	"gym-app/app/tasks"
	"log"
	"net/http"

	config "gym-app/app-config"
	"gym-app/app/exercise"
	"gym-app/app/program"
	"gym-app/app/result"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true) // create routes
	exercise.NewSubRouter(router)
	program.NewSubRouter(router)
	result.NewSubRouter(router)
	router.PathPrefix("/.well-known/acme-challenge/").Handler(http.FileServer(http.Dir("./certbot/")))

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
	log.Print("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

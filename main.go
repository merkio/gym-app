package main

import (
	config "gym-app/app-config"
	"gym-app/app/exercise"
	"gym-app/app/model"
	"gym-app/app/program"
	"gym-app/app/result"
	"gym-app/app/tasks"
	"gym-app/common/db"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.LoadConfig()

	router := mux.NewRouter().StrictSlash(true) // create routes
	exercise.NewSubRouter(router)
	program.NewSubRouter(router)
	result.NewSubRouter(router)

	// router.PathPrefix("/.well-known/acme-challenge/").Handler(http.FileServer(http.Dir("./certbot/")))

	conn := db.GetDB(config.DataConnectionConfig)

	err := conn.AutoMigrate(model.Result{}, model.Exercise{}, model.Program{})
	if err != nil {
		log.Fatal(err)
	}

	// these two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
	})

	allowedHeaders := handlers.AllowedHeaders([]string{
		"*",
	})

	// Run vk tasks
	go tasks.CollectVkMessages()

	// launch server with CORS validations
	log.Print("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

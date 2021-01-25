package user

import (
	"net/http"

	"gym-app/common/logger"

	"github.com/gorilla/mux"
)

var controller = &Controller{repository: URepository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"",
		controller.Index,
	},
	Route{
		"Index",
		"GET",
		"/{id}",
		controller.GetUser,
	},
	Route{
		"AddUser",
		"POST",
		"",
		controller.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"",
		controller.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/{id}",
		controller.DeleteUser,
	},
}

//NewSubRouter configures a new router to the API
func NewSubRouter(r *mux.Router) *mux.Router {
	router := r.PathPrefix("/api/user").Subrouter()

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.RequestLogger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

package result

import (
	"net/http"

	"gym-app/common/logger"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

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
		"/",
		controller.Index,
	},
	Route{
		"AddResult",
		"POST",
		"/",
		controller.AddResult,
	},
	Route{
		"UpdateResult",
		"PUT",
		"/",
		controller.UpdateResult,
	},
	Route{
		"DeleteResult",
		"DELETE",
		"/{id}",
		controller.DeleteResult,
	},
}

//NewSubRouter configures a new router to the API
func NewSubRouter(r *mux.Router) *mux.Router {
	router := r.PathPrefix("/result").Subrouter()

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

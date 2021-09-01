package exercise

import (
	"net/http"

	"gym-app/common/logger"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

func routes(c *Controller) []Route {
	return Routes{
		Route{
			"Index",
			"GET",
			"",
			c.Index,
		},
		Route{
			"Index",
			"GET",
			"/{id}",
			c.GetExercise,
		},
		Route{
			"AddExercise",
			"POST",
			"",
			c.AddExercise,
		},
		Route{
			"UpdateExercise",
			"PUT",
			"",
			c.UpdateExercise,
		},
		Route{
			"DeleteExercise",
			"DELETE",
			"/{id}",
			c.DeleteExercise,
		},
	}
}

//NewSubRouter configures a new router to the API
func NewSubRouter(r *mux.Router) *mux.Router {
	router := r.PathPrefix("/api/exercises").Subrouter()
	var controller = NewController(logger.NewLogger())

	for _, route := range routes(&controller) {
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

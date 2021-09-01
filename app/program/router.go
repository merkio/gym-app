package program

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
			c.GetProgram,
		},
		Route{
			"AddProgram",
			"POST",
			"",
			c.AddProgram,
		},
		Route{
			"SearchPrograms",
			"POST",
			"/search",
			c.Search,
		},
		Route{
			"UpdateProgram",
			"PUT",
			"",
			c.UpdateProgram,
		},
		Route{
			"DeleteProgram",
			"DELETE",
			"/{id}",
			c.DeleteProgram,
		},
	}
}

//NewSubRouter configures a new router to the API
func NewSubRouter(r *mux.Router) *mux.Router {
	router := r.PathPrefix("/api/program").Subrouter()
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

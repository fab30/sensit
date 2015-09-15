package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joelvim/sensit/timeseries"
)

// Route is a builder type for routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an alias type for an array of routes
type Routes []Route

// NewRouter builds the application http router. It takes a timeseries.DB connection
func NewRouter(db timeseries.DB, login, password string) *mux.Router {

	var routes = Routes{
		Route{
			"Ping",
			"GET",
			"/ping",
			Ping,
		},
		Route{
			"Temperature",
			"POST",
			"/api/v1/{deviceID}/temperature",
			Temperature(db, login, password),
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joelvim/sensit/callbacks"
	"github.com/joelvim/sensit/timeseries"
)

type requestHandler func(w http.ResponseWriter, r *http.Request)

// Temperature stores the data sent by the sensit server requests.
// It takes a connnection to a timeseries DB and returns a handler
func Temperature(database timeseries.DB, login, password string) func(w http.ResponseWriter, r *http.Request) {
	return basicAuthHandler(login, password, func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		deviceID := vars["deviceID"]

		// Parse the callback to extract the measures
		if measures, err := callbacks.ParseMeasures(deviceID, r.Body); err == nil {

			// Store the measures
			err = database.StoreMeasures(measures)

			if err != nil {
				log.Printf("Error storing measures : %s", err)
				http.Error(w, errors.New("Cannot store measure").Error(), 500)
			}
		} else {
			log.Printf("Error parsing measures : %s", err)
			http.Error(w, err.Error(), 500)
		}
	})
}

// Ping handle requests by returning Pong as text/plain content
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

// basicAuthHandler is a function that decorates an hanlder to add Basic Authentication support
func basicAuthHandler(login string, password string, handler requestHandler) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if l, p, ok := r.BasicAuth(); ok && l == login && p == password {
			handler(w, r)
		} else {
			// Unauthorized, return a response with a header WWW-Authenticate: Basic realm="sensit receiver"
			w.Header().Add("WWW-Authenticate", "Basic realm=\"sensit receiver\"")
			http.Error(w, "Authentication required", http.StatusUnauthorized)
		}
	}
}

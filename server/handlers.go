package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/joelvim/sensit/callbacks"
	"github.com/joelvim/sensit/timeseries"
)

type requestHandler func(w http.ResponseWriter, r *http.Request)

// Temperature stores the data sent by the sensit server requests.
// It takes a connnection to a timeseries DB and returns a handler
func Temperature(database timeseries.DB, authHandler AuthHandler) func(w http.ResponseWriter, r *http.Request) {
	return authHandler.Handle(func(w http.ResponseWriter, r *http.Request) {

		// Parse the callback to extract the measures
		measures, err := callbacks.ParseMeasures(r.Body)
		if err != nil {
			log.Printf("Error parsing measures : %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Log the converted measures
		log.Printf("Measures : len=%d cap=%d %v\n", len(measures), cap(measures), measures)

		// Store the measures
		if err = database.StoreMeasures(measures); err != nil {
			log.Printf("Error storing measures : %s", err)
			http.Error(w, errors.New("Cannot store measure").Error(), http.StatusInternalServerError)
			return
		}
	})
}

// Ping handle requests by returning Pong as text/plain content
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

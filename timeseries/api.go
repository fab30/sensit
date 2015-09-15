// Package timeseries is a package that provides an API to store measures in a
// timeseries DB.
// This API is backend agnostic.
package timeseries

import "github.com/joelvim/sensit/measure"

// DBConfig is a configuration containing the data to connect the timeseries db
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// DB is an interface providing API to store measures
type DB interface {

	// StoreMeasures store the measure as a batch on the backend. In case of
	// failure, an error is returned and is not nil
	StoreMeasures(measure []measure.Measure) error
}

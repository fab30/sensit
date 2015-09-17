// Package influxdb is the implementation of the timeseries API for InfluxDB protocol
package influxdb

import (
	"fmt"
	"net/url"
	"log"

	"github.com/influxdb/influxdb/client"
	"github.com/joelvim/sensit/measure"
	"github.com/joelvim/sensit/timeseries"
)

const (
	temperatureMeasurement = "temperature"
	deviceTag              = "device"
	valueField             = "value"
	minutePrecision       = "m"
	defaultRetentionPolicy = "default"
)

// Connect to an influxdb store
func Connect(config timeseries.DBConfig) (timeseries.DB, error) {
	// Build the server url
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", config.Host, config.Port))
	if err != nil {
		return nil, err
	}

	conf := client.Config{
		URL:      *u,
		Username: config.Login,
		Password: config.Password,
	}

	con, err := client.NewClient(conf)
	return InfluxDB{
		client: con,
		db:     config.Database,
	}, err
}

// InfluxDB is a client for an InfluxDB timeseries database
type InfluxDB struct {
	client *client.Client
	db     string
}

// StoreMeasures stores a batch of measures in an InfluxDB timeseries DB
func (db InfluxDB) StoreMeasures(measures []measure.Measure) error {
	points := make([]client.Point, len(measures))

	// convert measures to points
	for index, measure := range measures {
		points[index] = measureToPoint(measure)
	}

	batch := client.BatchPoints{
		Points:          points,
		Database:        db.db,
		RetentionPolicy: defaultRetentionPolicy,
	}

	log.Printf("Going to write %s", batch)

	// Write the batch
	response, err := db.client.Write(batch)
	log.Printf("InfluxDB response %v", response)

	return err
}

func measureToPoint(measure measure.Measure) client.Point {
	return client.Point{
		Measurement: temperatureMeasurement,
		Tags: map[string]string{
			deviceTag: measure.DeviceID,
		},
		Fields: map[string]interface{}{
			valueField: measure.Value,
		},
		Time:      measure.Time,
		Precision: minutePrecision,
	}
}

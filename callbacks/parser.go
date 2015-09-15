// Package callbacks contains the business logic to convert the input of a
// sensit callback into Measure objects
package callbacks

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/joelvim/sensit/measure"
)

// ParseMeasures extracts the measures from the callback
func ParseMeasures(deviceID string, body io.Reader) ([]measure.Measure, error) {
	var f interface{}
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&f)
	if err != nil {
		return nil, err
	}

	m := f.(map[string]interface{})

	if m["sensors"] == nil {
		return nil, errors.New("Invalid JSON content")
	}

	//sensors is an array of objects
	sensors := m["sensors"].([]interface{})

	var measures []measure.Measure

	// Iterate over the sensors array...
	for sensorIndex := range sensors {
		sensor := sensors[sensorIndex].(map[string]interface{})
		// ... and filter it to keep only temperature
		if sensor["sensor_type"] == "temperature" {
			history := sensor["history"].([]interface{})
			for measureIndex := range history {
				entry := history[measureIndex].(map[string]interface{})
				// Exclude the entries with a date period, we are only interested by periodic entries
				if entry["date_period"] == nil {
					jsonDate := entry["date"]
					data := entry["data"]

					if jsonDate != nil && data != nil {

						time, errTime := time.Parse("2006-01-02T15:04Z", jsonDate.(string))
						value, errValue := strconv.ParseFloat(data.(string), 32)

						// Error handling, still shitty
						if errTime != nil {
							log.Printf("Cannot parse measure time: %s", errTime)
						}

						if errValue != nil {
							log.Printf("Cannot parse value: %s", errValue)
						}

						if errTime == nil && errValue == nil {
							measures = append(measures, measure.Measure{deviceID, float32(value), time})
						}
					} else {
						if jsonDate == nil {
							log.Print("Date field not found")
						}
						if data == nil {
							log.Print("Value field not found")
						}
					}
				}
			}
		}
	}

	return measures, nil
}

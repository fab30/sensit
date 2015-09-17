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
	"fmt"

	"github.com/joelvim/sensit/measure"
)

// ParseMeasures extracts the measures from the callback
func ParseMeasures(body io.Reader) ([]measure.Measure, error) {
	callback := new(callbackMessage)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(callback)
	if err != nil {
		return nil, err
	}

	log.Printf("Parsed callback : %s", callback)

	if callback.Mode != "1" && callback.Mode != "4" {
		return nil, fmt.Errorf("Invalid mode %s", callback.Mode)
	}

	if callback.Sensors == nil {
		return nil, errors.New("Invalid JSON content")
	}

	var measures []measure.Measure

	// Iterate over the sensors array...
	for _, sensor := range callback.Sensors {
		// ... and filter it to keep only temperature
		if sensor.SensorType == temperatureSensorType {
			for _, entry := range sensor.History {
				jsonDate := entry.Date
				data := entry.Data

				if jsonDate != "" && data != "" {

					time, errTime := time.Parse(jsonCallbackTimeFormat, jsonDate)
					value, errValue := strconv.ParseFloat(data, 32)

					// Error handling, still shitty
					if errTime != nil {
						log.Printf("Cannot parse measure time: %s", errTime)
					}

					if errValue != nil {
						log.Printf("Cannot parse value: %s", errValue)
					}

					if errTime == nil && errValue == nil {
						measures = append(measures, measure.Measure{callback.SerialNumber, float32(value), time})
					}
				} else {
					if jsonDate == "" {
						log.Print("Date field not found")
					}
					if data == "" {
						log.Print("Value field not found")
					}
				}
			}
		}
	}

	return measures, nil
}

const (
	temperatureSensorType  = "temperature"
	jsonCallbackTimeFormat = "2006-01-02T15:04Z"
)

type callbackMessage struct {
	ID             string `json:"id"`
	SerialNumber   string `json:"serial_number"`
	ActivationDate string `json:"activation_date"`
	LastCommDate   string `json:"last_comm_date"`
	Mode           string `json:"mode"`

	Sensors []sensor `json:"sensors"`
}

type sensor struct {
	ID         string    `json:"id"`
	SensorType string    `json:"sensor_type"`
	History    []data    `json:"history"`
}

type data struct {
	Data string `json:"data"`
	Date string `json:"date"`
}

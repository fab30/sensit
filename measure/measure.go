// Package measure contains the data model.
package measure

import (
	"time"
	"fmt"
)

// Measure represents a temperature measure at a given time and for a given device
type Measure struct {
	DeviceID string // The id of the device

	Value float32 // The temperature

	Time time.Time // The time of the measure
}

func (m Measure) String() string {
  return fmt.Sprintf("[DeviceID %s Time %s Value %f]", m.DeviceID, m.Time, m.Value)
}

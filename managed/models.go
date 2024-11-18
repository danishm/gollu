package managed

import (
	"time"
)

type BloodGlucoseReading struct {
	Value     int64
	Timestamp time.Time
}

type BloodGlucoseGraph struct {
	PatientID string
	Values    []BloodGlucoseReading
}

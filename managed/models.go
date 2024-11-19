package managed

import (
	"time"
)

type BloodGlucoseReading struct {
	Value     int64     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type BloodGlucoseGraph struct {
	PatientID string                `json:"patient_id"`
	Values    []BloodGlucoseReading `json:"values"`
}

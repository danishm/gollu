package managed

import (
	"github.com/danishm/gollu"
	"time"
)

const (
	ValueUnknown = "unknown"

	ColorGreen   = "green"
	ColorYellow  = "yellow"
	ColorOrange  = "orange"
	ColorRed     = "red"
	ColorUnknown = ValueUnknown

	TrendLow        = "low"
	TrendDecreasing = "decreasing"
	TrendFlat       = "flat"
	TrendIncreasing = "increasing"
	TrendHigh       = "high"
	TrendUnknown    = ValueUnknown
)

type BloodGlucoseGraph struct {
	PatientID string                `json:"patient_id"`
	Values    []BloodGlucoseReading `json:"values"`
}

type BloodGlucoseReading struct {
	Value     int64     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	Color     string    `json:"color,omitempty"`
}

type BloodGlucoseLatestReading struct {
	BloodGlucoseReading
	Trend string `json:"trend,omitempty"`
}

func toBloodGlucoseLatestReading(measurement gollu.LLUGlucoseMeasurement) BloodGlucoseLatestReading {
	bgr := BloodGlucoseLatestReading{
		BloodGlucoseReading: BloodGlucoseReading{
			Value:     measurement.Value,
			Timestamp: time.Time(measurement.Timestamp),
			Color:     getColor(measurement.MeasurementColor),
		},
		Trend: getTrend(measurement.TrendArrow),
	}

	return bgr
}

// getTrend converts the numeric trendArrow returned by the LLU API
// into a string value
func getTrend(trendArrow int64) string {
	switch trendArrow {
	case 1:
		return TrendLow
	case 2:
		return TrendDecreasing
	case 3:
		return TrendFlat
	case 4:
		return TrendIncreasing
	case 5:
		return TrendHigh
	default:
		return TrendUnknown
	}
}

// getColor converts the numeric measurementColor returned by the LLU
// API into a string value
func getColor(measurementColor int64) string {
	switch measurementColor {
	case 1:
		return ColorGreen
	case 2:
		return ColorYellow
	case 3:
		return ColorOrange
	case 4:
		return ColorRed
	default:
		return ColorUnknown
	}
}

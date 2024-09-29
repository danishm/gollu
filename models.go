package gollu

import (
	"strings"
	"time"
)

const (
	LLUTimestampFormat = "1/2/2006 3:04:05 PM"
)

type LLULoginResponse struct {
	Status int64
	Data   LLULoginResponseData
}

type LLULoginResponseData struct {
	AuthTicket LLLULoginResponseAuthTicket
}

type LLUConnectionsResponse struct {
	Status int64
	Data   []LLUConnectionsResponseData `json:"data"`
	Ticket LLLULoginResponseAuthTicket
}

type LLUConnectionsResponseData struct {
	ID                 string
	PatientID          string
	Country            string
	Status             int64
	FirstName          string
	LastName           string
	GlucoseMeasurement LLUGlucoseMeasurement
}

type LLUGlucoseMeasurement struct {
	Timestamp        LLUTimestamp `json:"Timestamp"`
	Type             int64
	TrendArrow       int64 `json:"TrendArrow"`
	MeasurementColor int64 `json:"MeasurementColor"`
	GlucoseUnits     int64 `json:"GlucoseUnits"`
	Value            int64 `json:"Value"`
	IsHigh           bool
	isLow            bool
}

type LLLULoginResponseAuthTicket struct {
	Token    string
	Expires  int64
	Duration int64
}

type LLUTimestamp time.Time

// UnmarshalJSON implements a custom date unmarshaler for the time format used by LibreLinkup
func (lts *LLUTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(LLUTimestampFormat, s)
	if err != nil {
		return err
	}
	*lts = LLUTimestamp(t)
	return nil
}

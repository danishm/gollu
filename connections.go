package gollu

type LLUConnectionsResponse struct {
	Status int64
	Data   []LLUConnection `json:"data"`
	Ticket LLLULoginResponseAuthTicket
}

type LLUConnection struct {
	ID                 string
	PatientID          string
	Country            string
	Status             int64
	FirstName          string
	LastName           string
	GlucoseMeasurement LLUGlucoseMeasurement
	GlucoseItem        LLUGlucoseMeasurement
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

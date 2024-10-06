package gollu

type LLUConnectionsGraphResponse struct {
	Status int64
	Data   LLUConnectionsGraphResponseData `json:"data"`
	Ticket LLLULoginResponseAuthTicket
}

type LLUConnectionsGraphResponseData struct {
	Connection LLUConnection
	GraphData  []LLUGlucoseDataPoint
}

type LLUGlucoseDataPoint struct {
	FactoryTimestamp LLUTimestamp `json:"FactoryTimestamp"`
	Timestamp        LLUTimestamp `json:"Timestamp"`
	Type             int64
	ValueInMgPerDl   int64 `json:"ValueInMgPerDl"`
	MeasurementColor int64 `json:"MeasurementColor"`
	GlucoseUnits     int64 `json:"GlucoseUnits"`
	Value            int64 `json:"Value"`
	IsHigh           bool
	IsLow            bool
}

package managed

import (
	"github.com/danishm/gollu"
	"github.com/pkg/errors"
	"log/slog"
	"time"
)

type LLUClient struct {
	client gollu.LibreLinkUpClient
	ticket *gollu.LLLULoginResponseAuthTicket
}

func NewLLUClient(email, password string) LLUClient {
	client := gollu.NewLibreLinkUpClient(email, password)
	return LLUClient{
		client: client,
	}
}

func (llc *LLUClient) GetLastValue() (*BloodGlucoseReading, error) {

	connections, err := llc.getConnections()
	if err != nil {
		return nil, err
	}

	bgr := BloodGlucoseReading{
		Value:     connections.Data[0].GlucoseMeasurement.Value,
		Timestamp: time.Time(connections.Data[0].GlucoseMeasurement.Timestamp),
	}

	return &bgr, nil
}

func (llc *LLUClient) GetGraphValues() (*BloodGlucoseGraph, error) {

	values := []BloodGlucoseReading{}

	connections, err := llc.getConnections()
	if err != nil {
		return nil, err
	}

	if len(connections.Data) < 1 {
		return nil, errors.New("could not find any patient connections")
	}

	// getting graph data for the first connection
	patientID := connections.Data[0].PatientID
	graphData, err := llc.client.Graph(*llc.ticket, patientID)
	if err != nil {
		return nil, errors.Wrap(err, "Error making graph call")
	}
	for _, item := range graphData.Data.GraphData {
		bgr := BloodGlucoseReading{
			Value:     item.Value,
			Timestamp: time.Time(item.Timestamp),
		}
		values = append(values, bgr)
	}

	bgg := BloodGlucoseGraph{
		PatientID: patientID,
		Values:    values,
	}

	return &bgg, nil
}

func (llc *LLUClient) getConnections() (*gollu.LLUConnectionsResponse, error) {
	ticket, err := llc.getAuthTicket()
	if err != nil {
		return nil, err
	}
	// getting latest value
	connections, err := llc.client.Connections(*ticket)
	if err != nil {
		slog.Error("getConnections()", "error", err.Error())
		return nil, err
	}
	return connections, nil
}

// login attempts to log in to LibreLinkUp API and cache the auth ticket on success
func (llc *LLUClient) login() error {
	slog.Info("login() logging in")

	lr, err := llc.client.Login()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	llc.ticket = &lr.Data.AuthTicket
	return nil
}

// getAuthTicket returns the cached auth ticket if it's still valid. Otherwise,
// it logs in again to get a new ticket
func (llc *LLUClient) getAuthTicket() (*gollu.LLLULoginResponseAuthTicket, error) {

	var expired bool = false

	// check if we have a valid, un-expired auth ticket
	if llc.ticket != nil {
		if llc.ticket.DaysToExpiry() <= 1 {
			expired = true
		}
	} else {
		expired = true
	}

	if !expired {
		slog.Info("getAuthTicket() returning cached auth ticket", "daysToExpiry", llc.ticket.DaysToExpiry())
		return llc.ticket, nil
	}

	slog.Info("getAuthTicket() auth ticket expired")
	err := llc.login()
	if err == nil {
		slog.Info("getAuthTicket() new auth ticket", "daysToExpiry", llc.ticket.DaysToExpiry())
	}
	return llc.ticket, err
}

package main

import (
	"fmt"
	"github.com/danishm/gollu"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide email and password as command line parameters")
		os.Exit(1)
	}

	email := os.Args[1]
	password := os.Args[2]

	// logging in
	client := gollu.NewLibreLinkUpClient(email, password)
	resp, err := client.Login()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}

	// getting the list of connections
	connections, err := client.Connections(resp.Data.AuthTicket)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}

	if len(connections.Data) < 1 {
		fmt.Println("Error: could not find any connections")
		os.Exit(-1)
	}

	// getting graph data for the first connection
	patientID := connections.Data[0].PatientID
	graphData, err := client.Graph(resp.Data.AuthTicket, patientID)
	if err != nil {
		fmt.Printf("Error making graph call: %s\n", err)
		os.Exit(-1)
	}
	for _, item := range graphData.Data.GraphData {
		fmt.Printf("%s\t\t%d\n", item.Timestamp.String(), item.Value)
	}

	os.Exit(0)
}

package main

import (
	"fmt"
	"github.com/danishm/gollu"
	"os"
	"time"
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

	// getting latest value
	connections, err := client.Connections(resp.Data.AuthTicket)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(connections.Data[0].GlucoseMeasurement.Value, time.Time(connections.Data[0].GlucoseMeasurement.Timestamp))

	os.Exit(0)
}

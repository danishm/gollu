package main

import (
	"fmt"
	"github.com/danishm/gollu/managed"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide email and password as command line parameters")
		os.Exit(1)
	}

	email := os.Args[1]
	password := os.Args[2]

	// logging in
	client := managed.NewLLUClient(email, password)
	bgr, err := client.GetLastValue()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}
	valueStr := strconv.FormatInt(bgr.Value, 10)
	timeStr := bgr.Timestamp.Format("2006-01-02 03:04:05 PM")
	fmt.Println("Value Timestamp")
	fmt.Println("----- ----------------------")
	fmt.Printf("%5.5s %s\n", valueStr, timeStr)

	os.Exit(0)
}

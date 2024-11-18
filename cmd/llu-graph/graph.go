package main

import (
	"fmt"
	"github.com/danishm/gollu/managed"
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
	client := managed.NewLLUClient(email, password)

	gv, err := client.GetGraphValues()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for _, item := range gv.Values {
		fmt.Printf("%s\t\t%d\n", item.Timestamp.String(), item.Value)
	}

	os.Exit(0)
}

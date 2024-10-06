GoLLU | Go Client for LibreLinkUp API
=====================================
![build status](https://github.com/danishm/gollu/actions/workflows/go.yml/badge.svg)

Description
-----------
I wanted to be able to collect readings from my own GCM monitor into my own time series
database so that I can do my own time series analysis and use my own visulaization tools
and dashboards.

Credits
-------
 * This library was implemented based on work from [khskekec](https://gist.github.com/khskekec) which,
was generously made available in this [gist](https://gist.github.com/khskekec/6c13ba01b10d3018d816706a32ae8ab2).
 * I also use this `javascript` based API implementation at https://github.com/DiaKEM/libre-link-up-api-client

Usage
-----
Given below are full examples of command line utility implementations of getting the latest value and a range
of values for graph data.

### Get Last Blood Glucose Value

#### Code - `latest.go`

```go
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
	fmt.Println(connections.Data[0].GlucoseMeasurement.Value)

	os.Exit(0)
}
```

#### Output

```shell
$ go run latest.go foo.bar@test.com fakepassword
136
```

### Get Series of Glucose Values

#### Code - `graph.go`

```go
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
```

#### Output

```shell
$go run graph.go foo.bar@test.com fakepassword
10/6/2024 11:29:28 AM           134
10/6/2024 11:34:29 AM           132
10/6/2024 11:39:29 AM           129
10/6/2024 11:44:30 AM           126
10/6/2024 11:49:28 AM           126
10/6/2024 11:54:28 AM           128
10/6/2024 11:59:30 AM           124
10/6/2024 12:04:28 PM           123
10/6/2024 12:09:29 PM           121
10/6/2024 12:14:29 PM           118
10/6/2024 12:19:29 PM           119
10/6/2024 12:24:28 PM           118
10/6/2024 12:29:28 PM           115
10/6/2024 12:34:29 PM           114
10/6/2024 12:39:28 PM           114
10/6/2024 12:44:30 PM           113
10/6/2024 12:49:28 PM           112
... <output trucated for clarity>
```
License
-------
[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)
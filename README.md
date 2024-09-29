GoLLU | Go Client for LibreLinkUp API
=====================================

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
This is a work in progress. The current implementation only supports logging and getting the latest value.
Here is an example.

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
		fmt.Printf("Error: %s", err)
	}

	// getting latest value
	connections, err := client.Connections(resp.Data.AuthTicket)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(connections.Data[0].GlucoseMeasurement.Value,)

	os.Exit(0)
}
```

License
-------
Distributed under the MIT License.
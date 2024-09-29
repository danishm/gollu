package gollu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	LLUUrl                 = "https://api.libreview.io"
	LLULoginEndpoint       = "llu/auth/login"
	LLUConnectionsEndpoint = "llu/connections"
)

// LibreLinkUpClient represents a LibreLinkUp API client, which needs an email address
// and password to make calls. It depends on the primary URL and endpoint URLs identified
// in the const section. It also depends on the addCommonHeaders() function to add the
// common headers needed to make the API calls
type LibreLinkUpClient struct {
	email    string
	password string
}

// NewLibreLinkUpClient creates a new instance of LibreLinkUpClient
func NewLibreLinkUpClient(email, password string) LibreLinkUpClient {
	return LibreLinkUpClient{
		email:    email,
		password: password,
	}
}

// Login performs the login API call and returns a response struct containing the bearer
// token that needs to be referenced to in the subsequent calls
func (llu *LibreLinkUpClient) Login() (*LLULoginResponse, error) {

	// creating the url
	url := fmt.Sprintf("%s/%s", LLUUrl, LLULoginEndpoint)

	// preparing POST data
	data := map[string]string{
		"email":    llu.email,
		"password": llu.password,
	}
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Creating POST body
	postBuffer := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, url, postBuffer)
	if err != nil {
		return nil, err
	}
	addCommonHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	llr := LLULoginResponse{}
	err = json.Unmarshal(body, &llr)
	if err != nil {
		return nil, err
	}
	return &llr, nil
}

// Connections makes the connections API calls that is used to discover all the device
// connections available to the user. It can also be used to get the latest value for all
// the sensors exposed to the user.
func (llu *LibreLinkUpClient) Connections(ticket LLLULoginResponseAuthTicket) (*LLUConnectionsResponse, error) {

	// creating the url
	url := fmt.Sprintf("%s/%s", LLUUrl, LLUConnectionsEndpoint)

	// creating request and setting headers
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	addCommonHeaders(req)
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", ticket.Token))

	// making the call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lcr := LLUConnectionsResponse{}
	err = json.Unmarshal(body, &lcr)
	if err != nil {
		return nil, err
	}
	return &lcr, nil
}

// addCommonHeaders adds the common headers required for working with the LibreLinkUp API
func addCommonHeaders(req *http.Request) {
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("product", "llu.android")
	req.Header.Add("version", "4.7.0")
}

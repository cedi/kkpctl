package bearerclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/savaki/jq"
)

// BearerClient is a simple http.Client that always injects a Bearer Auth token to an Request
type BearerClient interface {
	Do(req *http.Request, out any) (*http.Response, error)
}

// Client holds all config and the http.Client needed to talk to the Kubermatic API
type Client struct {
	httpClient *http.Client
	bearer     string
}

// NewClient creates a new Client for the Kubermatic API
func NewClient(baseURL string, bearer string) *Client {
	return &Client{
		httpClient: &http.Client{},
		bearer:     bearer,
	}
}

// Do injects the stored token into the provided request then sends it
// req is the http.Request to execute
// out can be a pointer to any object to which the result should be de-serialised.
func (c *Client) Do(req *http.Request, out any) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.bearer)

	resp, err := c.httpClient.Do(req)
	if err != nil || out == nil {
		return resp, err
	}

	if resp.Header.Get("content-type") != "application/json" {

		if resp.StatusCode >= 299 {
			return resp, fmt.Errorf("%v", resp.Status)
		}

		return resp, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 299 {
		// Try parsing the body regardless of the error and see if it contains a error message
		op, err := jq.Parse(".error.message")
		if err != nil {
			// If it fails just return the status code then
			return resp, fmt.Errorf("%v", resp.Status)
		}

		value, err := op.Apply(body)
		if err != nil {
			// If it fails to find a error message, just return the status code then
			return resp, fmt.Errorf("%v", resp.Status)
		}

		return resp, fmt.Errorf("%v: %s", resp.Status, string(value))
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return resp, errors.Wrap(err, "Unable parsing response")
	}

	return resp, err
}

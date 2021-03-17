package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	apiV1Path       string = "/api/v1"
	contentTypeJSON string = "application/json"
)

// URLParams is a map of strings which can be added to a Get request
type URLParams map[string]string

// Client holds all config and the http.Client needed to talk to the Kubermatic API
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	bearer     string
}

// NewClient creates a new Client for the Kubermatic API
func NewClient(baseURL string, bearer string) (*Client, error) {
	parsedURL, err := url.Parse(baseURL + apiV1Path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed parsing API URL "+baseURL+apiV1Path)
	}

	httpClient := &http.Client{}

	client := &Client{}
	client.BaseURL = parsedURL
	client.httpClient = httpClient
	client.bearer = bearer

	return client, nil
}

// Do injects the stored token into the provided request then sends it
// req is the http.Request to execute
// out can be a pointer to any object to which the result should be de-serialised.
func (c *Client) Do(req *http.Request, out interface{}) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.bearer) // "/api/v1/projects/6tmbnhdl7h/dc/seed-ix2/clusters/h9dzr6x7wk/node...+11 more"

	resp, err := c.httpClient.Do(req)
	if err != nil || out == nil {
		return resp, err
	}

	if resp.StatusCode >= 299 {
		return resp, errors.New(resp.Status)
	}

	if resp.Header.Get("content-type") != "application/json" {
		return resp, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return resp, errors.Wrap(err, "Unable parsing response")
	}

	return resp, err
}

// Get functions the same as http.Client.Get but injects a stored token into ty header
func (c *Client) Get(requestURL string, out interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, out)
}

// GetWithQueryParams functions the same as http.Client.Get but injects a stored token into ty header
func (c *Client) GetWithQueryParams(requestURL string, queryParams URLParams, out interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	for key, value := range queryParams {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()

	return c.Do(req, out)
}

// Post functions the same as http.Client.Post but injects a stored token into the header
func (c *Client) Post(requestURL string, contentType string, body interface{}, out interface{}) (*http.Response, error) {
	var bodyBuf io.ReadWriter

	if body != nil {
		bodyBuf = new(bytes.Buffer)
		err := json.NewEncoder(bodyBuf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("POST", c.formatURL(requestURL), bodyBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req, out)
}

// Patch functions the same as http.Client.Post but injects a stored token into the header
func (c *Client) Patch(requestURL string, contentType string, body interface{}, out interface{}) (*http.Response, error) {
	var bodyBuf io.ReadWriter

	if body != nil {
		bodyBuf = new(bytes.Buffer)
		err := json.NewEncoder(bodyBuf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("PATCH", c.formatURL(requestURL), bodyBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req, out)
}

// Put functions the same as http.Client.Put but injects a stored token into the header
func (c *Client) Put(requestURL string, contentType string, body interface{}, out interface{}) (*http.Response, error) {
	var bodyBuf io.ReadWriter

	if body != nil {
		bodyBuf = new(bytes.Buffer)
		err := json.NewEncoder(bodyBuf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("PUT", c.formatURL(requestURL), bodyBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req, out)
}

// Delete functions the same as http.Client.Delete but injects a stored token into ty header
func (c *Client) Delete(requestURL string) (*http.Response, error) {
	return c.DeleteWithHeader(requestURL, nil)
}

// DeleteWithHeader functions the same as http.Client.Delete but injects a stored token into ty header
func (c *Client) DeleteWithHeader(requestURL string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	return c.Do(req, nil)
}

func (c *Client) formatURL(requestURL string) string {
	return c.BaseURL.String() + requestURL
}

package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/cedi/kkpctl/pkg/bearerclient"
	"github.com/pkg/errors"
)

const (
	contentTypeJSON string = "application/json"
)

// APIVersion is the version type of the KKP API to use
type APIVersion string

const (
	// V1API is the v1 API
	V1API APIVersion = "v1"

	// V2API is the v2 API
	V2API APIVersion = "v2"
)

// URLParams is a map of strings which can be added to a Get request
type URLParams map[string]string

// Client holds all config and the http.Client needed to talk to the Kubermatic API
type Client struct {
	BaseURL *url.URL

	bc *bearerclient.Client
}

// NewClient creates a new Client for the Kubermatic API
func NewClient(baseURL string, bearer string) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed parsing API URL %s ", baseURL)
	}

	client := &Client{
		BaseURL: parsedURL,
		bc:      bearerclient.NewClient(baseURL, bearer),
	}

	return client, nil
}

// Do injects the correct API Version Path into the request and executes then bearerclient.Do
func (c *Client) Do(req *http.Request, out interface{}, apiVersion APIVersion) (*http.Response, error) {
	req.URL.Path = "/api/" + string(apiVersion) + req.URL.Path
	return c.bc.Do(req, out)
}

// Get functions the same as http.Client.Get but injects a stored token into ty header
func (c *Client) Get(requestURL string, out interface{}, apiVersion APIVersion) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, out, apiVersion)
}

// GetWithQueryParams functions the same as http.Client.Get but injects a stored token into ty header
func (c *Client) GetWithQueryParams(requestURL string, queryParams URLParams, out interface{}, apiVersion APIVersion) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	for key, value := range queryParams {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()

	return c.Do(req, out, apiVersion)
}

// Post functions the same as http.Client.Post but injects a stored token into the header
func (c *Client) Post(requestURL string, contentType string, body interface{}, out interface{}, apiVersion APIVersion) (*http.Response, error) {
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
	return c.Do(req, out, apiVersion)
}

// Patch functions the same as http.Client.Post but injects a stored token into the header
func (c *Client) Patch(requestURL string, contentType string, body interface{}, out interface{}, apiVersion APIVersion) (*http.Response, error) {
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
	return c.Do(req, out, apiVersion)
}

// Put functions the same as http.Client.Put but injects a stored token into the header
func (c *Client) Put(requestURL string, contentType string, body interface{}, out interface{}, apiVersion APIVersion) (*http.Response, error) {
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
	return c.Do(req, out, apiVersion)
}

// Delete functions the same as http.Client.Delete but injects a stored token into ty header
func (c *Client) Delete(requestURL string, apiVersion APIVersion) (*http.Response, error) {
	return c.DeleteWithHeader(requestURL, nil, apiVersion)
}

// DeleteWithHeader functions the same as http.Client.Delete but injects a stored token into ty header
func (c *Client) DeleteWithHeader(requestURL string, headers map[string]string, apiVersion APIVersion) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.formatURL(requestURL), nil)
	if err != nil {
		return nil, err
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	return c.Do(req, nil, apiVersion)
}

func (c *Client) formatURL(requestURL string) string {
	return c.BaseURL.String() + requestURL
}

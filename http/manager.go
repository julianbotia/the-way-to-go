package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Client is wrapper for a HTTP client...
type Client struct {
	HTTPClient *http.Client
}

// DefaultClient you can use the native http implementation
var DefaultClient = http.DefaultClient

// PerformRequest using default http.DefaultClient
func PerformRequest(request *Request) (*Response, error) {
	c := &Client{HTTPClient: DefaultClient}
	return c.PerformRequest(request)
}

// PerformRequest based on the request model
func (c *Client) PerformRequest(request *Request) (*Response, error) {
	// build request data
	req, err := buildNativeRequest(request)
	if err != nil {
		return nil, err
	}
	// do request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	// build response
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	response := Response{
		Status: res.StatusCode,
		Body:   string(body),
	}
	return &response, err
}

func buildNativeRequest(request *Request) (*http.Request, error) {
	method := string(request.Method)
	// build URL
	url := buildURL(request)
	body := bytes.NewBuffer(request.Body)

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return req, err
	}
	//add headers
	req.Header.Set("Content-type", "application/json")
	return req, err
}

func buildURL(request *Request) string {
	// TODO add params
	return request.BaseURL
}

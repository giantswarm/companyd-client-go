package client

import (
	"io"
	"net/http"

	"bytes"
	"encoding/json"
)

// Dial returns a client for the given server.
func Dial(server string) (client *Client, err error) {
	client = new(Client)

	client.endpoint, err = parseEndpoint(server)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Client struct {
	endpoint *Endpoint

	LogGetRequest    func(url string, resp *http.Response, err error)
	LogPostRequest   func(url, contentType string, resp *http.Response, err error)
	LogDeleteRequest func(url string, resp *http.Response, err error)
}

func (c *Client) endpointUrl(url string) string {
	return c.endpoint.String() + url
}

// get requests the given url from the configured endpoint.
func (c *Client) get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if c.LogGetRequest != nil {
		c.LogGetRequest(url, resp, err)
	}

	return resp, err
}

// postJson transforms the body into a JSON stream and sends it to the given URL as a HTTP POST request.
func (c *Client) postJson(url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, Mask(err)
	}

	return c.post(url, "application/json", bytes.NewReader(data))
}

func (c *Client) post(url, contentType string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	if c.LogPostRequest != nil {
		c.LogPostRequest(url, contentType, resp, err)
	}

	return resp, err
}

func (c *Client) delete(url string) (*http.Response, error) {
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if c.LogDeleteRequest != nil {
		c.LogDeleteRequest(url, resp, err)
	}

	return resp, err
}

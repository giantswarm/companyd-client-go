package client

import (
	"fmt"
	"net/url"
	"strings"
)

// Endpoint describes the API endpoint the client uses.
type Endpoint struct {
	url *url.URL
}

func parseEndpoint(connectionString string) (endpoint *Endpoint, err error) {
	endpoint = new(Endpoint)
	endpoint.url, err = url.Parse(connectionString)

	if err != nil {
		return nil, err
	}

	if !(endpoint.url.Scheme == "http" || endpoint.url.Scheme == "https") {
		return nil, fmt.Errorf("Invalid scheme for server URI: %s", endpoint.url.Scheme)
	}

	if endpoint.url.Host == "" {
		return nil, fmt.Errorf("No hostname given: %s", connectionString)
	}

	if strings.HasSuffix(endpoint.url.Path, "/") {
		endpoint.url.Path += "/"
	}

	return endpoint, nil
}

// String returns the http base url. The url always contains a slash (/) at the end.
func (endpoint *Endpoint) String() string {
	return endpoint.url.String()
}

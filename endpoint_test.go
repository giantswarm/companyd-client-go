package client

import "testing"

func TestParseEndpoint(t *testing.T) {
	test := func(url, expected string) {
		endpoint, err := parseEndpoint(url)
		if err != nil {
			t.Fatalf("Got error for url '%s': %v", url, err)
		}

		if endpoint.String() != expected {
			t.Fatalf("Expected '%s'\n     got '%s'", expected, endpoint.String())
		}
	}

	test("http://localhost:8080/v1/", "http://localhost:8080/v1")
	test("http://localhost:8080/v1", "http://localhost:8080/v1")
	test("http://localhost:8080/", "http://localhost:8080")
	test("http://localhost:8080", "http://localhost:8080")
}

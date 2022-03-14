package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization

	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	client.builder.headers = commonHeaders

	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation

	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

}

func TestRequestBodyNilBody(t *testing.T) {

	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {

		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {

		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("xpto", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})
}

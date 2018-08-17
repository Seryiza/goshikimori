package goshikimori_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/seryiza/goshikimori"
)

const (
	jsonContentType = "application/json"
)

type fakeTransport struct {
	OutputStatusCode int
	OutputBody       []byte
	RequestChecker   func(*http.Request)
}

func (fake *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if fake.RequestChecker != nil {
		fake.RequestChecker(req)
	}

	outputReader := ioutil.NopCloser(bytes.NewReader(fake.OutputBody))
	resp := &http.Response{
		StatusCode: fake.OutputStatusCode,
		Body:       outputReader,
	}
	return resp, nil
}

func TestTransport(t *testing.T) {
	appName := "Test App"
	outputBody := []byte("Hello")

	// Check headers: User-Agent and Content-Type
	headersChecker := func(req *http.Request) {
		headerAppName := req.Header.Get("User-Agent")
		if headerAppName != appName {
			t.Error("Header app name is invalid")
		}

		headerContentType := req.Header.Get("Content-Type")
		if headerContentType != jsonContentType {
			t.Error("Request content type is not json")
		}
	}

	client := &fakeTransport{
		OutputStatusCode: http.StatusOK,
		OutputBody:       outputBody,
		RequestChecker:   headersChecker,
	}

	shikiTransport := goshikimori.Transport{
		ApplicationName: appName,
		Target:          client,
	}

	req, err := http.NewRequest("GET", "http://somehost", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := shikiTransport.RoundTrip(req)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(body, outputBody) != 0 {
		t.Error("HTTP body is invalid")
	}
}

func TestInvalidToken(t *testing.T) {
	// Token is invalid when request returns http.StatusUnauthorized
	appName := "Test App"
	outputBody := []byte("Hello")

	client := &fakeTransport{
		OutputStatusCode: http.StatusUnauthorized,
		OutputBody:       outputBody,
	}

	shikiTransport := goshikimori.Transport{
		ApplicationName: appName,
		Target:          client,
	}

	req, err := http.NewRequest("GET", "http://somehost", nil)
	if err != nil {
		t.Error(err)
	}

	_, err = shikiTransport.RoundTrip(req)
	if err == nil {
		t.Error("Token is invalid, but error is nil")
	}
}

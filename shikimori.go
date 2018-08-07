package goshikimori

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	shikimoriAPI_v1 = "https://shikimori.org/api/%s"
	shikimoriAPI_v2 = "https://shikimori.org/api/v2/%s"
)

// Shikimori to send requests to Shikimori API
type Shikimori struct {
	Client *http.Client

	// URLFormat of format template for Shikimori API
	URLFormat string
	Version   string
}

func NewShikimori(client *http.Client, version string) *Shikimori {
	urlFormat := getAPIFormat(version)
	shiki := &Shikimori{
		Client:    client,
		URLFormat: urlFormat,
		Version:   version,
	}
	return shiki
}

func getAPIFormat(version string) string {
	switch version {
	case "1.0":
		return shikimoriAPI_v1
	case "2.0":
		return shikimoriAPI_v2
	default:
		return shikimoriAPI_v1
	}
}

// NewRequest returns request for Shikimori (add to method full address).
// Ex., ("GET", "whoami", nil) => ("GET", "https://shikimori.org/api/whoami", nil)
func (shiki *Shikimori) NewRequest(httpMethod, shikiMethod string, body io.Reader) (*http.Request, error) {
	fullURL := fmt.Sprintf(shiki.URLFormat, shikiMethod)
	return http.NewRequest(httpMethod, fullURL, body)
}

func (shiki *Shikimori) Do(req *http.Request) (*http.Response, error) {
	// todo: добавить обработку не-API ошибок (напр., HTTP 404)
	return shiki.Client.Do(req)
}

func (shiki *Shikimori) Get(method string) (*http.Response, error) {
	req, err := shiki.NewRequest("GET", method, nil)
	if err != nil {
		return nil, err
	}

	return shiki.Do(req)
}

func (shiki *Shikimori) Post(method string, body io.Reader) (*http.Response, error) {
	req, err := shiki.NewRequest("POST", method, body)
	if err != nil {
		return nil, err
	}

	return shiki.Do(req)
}

func (shiki *Shikimori) Delete(method string) (*http.Response, error) {
	req, err := shiki.NewRequest("DELETE", method, nil)
	if err != nil {
		return nil, err
	}

	return shiki.Do(req)
}

func (shiki *Shikimori) Put(method string, body io.Reader) (*http.Response, error) {
	req, err := shiki.NewRequest("PUT", method, body)
	if err != nil {
		return nil, err
	}

	return shiki.Do(req)
}

func (shiki *Shikimori) Patch(method string, body io.Reader) (*http.Response, error) {
	req, err := shiki.NewRequest("PATCH", method, body)
	if err != nil {
		return nil, err
	}

	return shiki.Do(req)
}

func (shiki *Shikimori) JSONGet(method string, outStruct interface{}) (*http.Response, error) {
	resp, err := shiki.Get(method)
	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(outStruct)
	return resp, err
}

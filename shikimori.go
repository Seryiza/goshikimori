package goshikimori

import (
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

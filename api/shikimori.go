package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/seryiza/go-shikimori/auth"
	"golang.org/x/oauth2"
)

// Shikimori to send requests to Shikimori
type Shikimori struct {
	Client *http.Client
}

func DefaultClientByCode(conf *oauth2.Config, appName, authCode string) (*Shikimori, error) {
	ctx := context.Background()
	ctx = auth.AddShikimoriTransport(ctx, appName)

	tok, err := conf.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx, tok)
	shiki := &Shikimori{
		Client: client,
	}

	return shiki, nil
}

func DefaultClientByToken(conf *oauth2.Config, appName string, tok *oauth2.Token) (*Shikimori, error) {
	ctx := context.Background()
	ctx = auth.AddShikimoriTransport(ctx, appName)

	client := conf.Client(ctx, tok)
	shiki := &Shikimori{
		Client: client,
	}

	return shiki, nil
}

// ApiURLWithQuery returns shikimori api url for get-queries
func (shiki *Shikimori) ApiURLWithValues(path string, query url.Values) string {
	url := url.URL{
		Scheme:   "https",
		Host:     "shikimori.org",
		Path:     path,
		RawQuery: query.Encode(),
	}
	return url.String()
}

func (shiki *Shikimori) ApiURLWithString(path string, strQuery string) string {
	url := url.URL{
		Scheme:   "https",
		Host:     "shikimori.org",
		Path:     path,
		RawQuery: strQuery,
	}
	return url.String()
}

// ApiURL returns shikimori api url
func (shiki *Shikimori) ApiURL(path string) string {
	url := url.URL{
		Scheme: "https",
		Host:   "shikimori.org",
		Path:   path,
	}
	return url.String()
}

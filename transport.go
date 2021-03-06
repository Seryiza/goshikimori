package goshikimori

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

// Transport for adding headers
type Transport struct {
	// As User-Agent for Shikimori
	ApplicationName string
	Target          http.RoundTripper
}

// AddTransport to context.
// If ctx.Value(oauth2.HTTPClient) == nil, then using
// DefaultTransport + Shikimori Transport
func AddTransport(ctx context.Context, appName string) context.Context {
	ctxClient := ctx.Value(oauth2.HTTPClient)

	var client *http.Client
	if ctxClient == nil {
		client = &http.Client{}
	} else {
		client = ctxClient.(*http.Client)
	}

	client.Transport = Transport{
		ApplicationName: appName,
		Target:          client.Transport,
	}
	return context.WithValue(ctx, oauth2.HTTPClient, client)
}

// RoundTrip implements RoundTripper. Set User-Agent and call Transport.Target
func (tr Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", tr.ApplicationName)
	req.Header.Set("Content-Type", "application/json")

	resp, err := tr.target().RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if !isTokenCorrect(resp) {
		return resp, errors.New("The access token is invalid")
	}

	return resp, err
}

func isTokenCorrect(resp *http.Response) bool {
	return resp.StatusCode != http.StatusUnauthorized
}

func (tr Transport) target() http.RoundTripper {
	if tr.Target != nil {
		return tr.Target
	}
	return http.DefaultTransport
}

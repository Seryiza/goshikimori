package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// ShikimoriTransport for adding headers
type ShikimoriTransport struct {
	// As User-Agent for Shikimori
	ApplicationName string
	Target          http.RoundTripper
}

// AddShikimoriTransport to context.
// If ctx.Value(oauth2.HTTPClient) == nil, then using
// DefaultTransport + ShikimoriTransport
func AddShikimoriTransport(ctx context.Context, appName string) context.Context {
	ctxClient := ctx.Value(oauth2.HTTPClient)

	var client *http.Client
	if ctxClient == nil {
		client = &http.Client{}
	} else {
		client = ctxClient.(*http.Client)
	}

	client.Transport = ShikimoriTransport{
		ApplicationName: appName,
		Target:          client.Transport,
	}
	return context.WithValue(ctx, oauth2.HTTPClient, client)
}

// RoundTrip implements RoundTripper
func (tr ShikimoriTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", tr.ApplicationName)
	return tr.target().RoundTrip(req)
}

func (tr ShikimoriTransport) target() http.RoundTripper {
	if tr.Target != nil {
		return tr.Target
	}
	return http.DefaultTransport
}

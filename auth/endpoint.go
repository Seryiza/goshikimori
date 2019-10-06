package auth

import "golang.org/x/oauth2"

const (
	// StandaloneRedirectURL for Standalone applications
	StandaloneRedirectURL = "urn:ietf:wg:oauth:2.0:oob"
)

// ShikimoriEndpoint for Shikimori OAuth2
var ShikimoriEndpoint = oauth2.Endpoint{
	AuthURL:  "https://shikimori.one/oauth/authorize",
	TokenURL: "https://shikimori.one/oauth/token",
}

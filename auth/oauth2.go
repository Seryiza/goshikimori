package auth

import (
	"golang.org/x/oauth2"
)

// GetAuthCodeURL returns AuthCodeURL
func GetAuthCodeURL(conf *oauth2.Config) string {
	// empty string because Shikimori API without states
	return conf.AuthCodeURL("")
}

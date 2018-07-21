package auth_test

import (
	"os"
	"testing"

	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/auth"
)

const (
	envLogin    = "SHIKI_LOGIN"
	envPassword = "SHIKI_PASS"
	envClientID = "SHIKI_CLIENTID"
)

func TestLogin(t *testing.T) {
	login, password := os.Getenv(envLogin), os.Getenv(envPassword)
	conf := &oauth2.Config{
		ClientID:    os.Getenv(envClientID),
		RedirectURL: auth.StandaloneRedirectURL,
		Endpoint:    auth.ShikimoriEndpoint,
	}
	url := auth.GetAuthCodeURL(conf)

	code, err := auth.GetCodeByLogin(url, login, password)
	if err != nil {
		t.Error(err.Error())
	}

	if code == "" {
		t.Error("Code is empty")
	}
}

func TestWrongLogin(t *testing.T) {
	login, password := "abc", "xyz"
	conf := &oauth2.Config{
		ClientID:    os.Getenv(envClientID),
		RedirectURL: auth.StandaloneRedirectURL,
		Endpoint:    auth.ShikimoriEndpoint,
	}
	url := auth.GetAuthCodeURL(conf)

	_, err := auth.GetCodeByLogin(url, login, password)
	if err == nil {
		t.Error("Error is nil (want non-nil)")
	}
}

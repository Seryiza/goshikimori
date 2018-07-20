package auth_test

import (
	"testing"

	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/auth"
)

func TestLogin(t *testing.T) {
	login, password := "somebody", "whatsthis"
	conf := &oauth2.Config{
		ClientID:    "your clientid here",
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

package helpers

import (
	"context"
	"errors"
	"os"

	"github.com/seryiza/loadOAuth/conf"
	"github.com/seryiza/loadOAuth/token"
	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/api"
	"github.com/seryiza/go-shikimori/auth"
)

const (
	// envirement variables
	shikiPrefix  = "SHIKI"
	shikiAppName = "SHIKI_APP_NAME"

	// to create config and new token if error
	shikiLogin        = "SHIKI_LOGIN"
	shikiPassword     = "SHIKI_PASS"
	shikiClientID     = "SHIKI_CLIENTID"
	shikiClientSecret = "SHIKI_CLIENTSECRET"
	shikiRedirectURL  = "SHIKI_REDIRECT_URL"
)

// GetShikimori returns api.Shikimori by config and token from files (using loadOAuth).
// If loadOAuth returns error, try create config/token from env-vars
func GetShikimori() (*api.Shikimori, error) {
	conf, err := getConf()
	if err != nil {
		return nil, err
	}

	tok, err := getToken(conf)
	if err != nil {
		return nil, err
	}

	appName := getAppName()
	return api.DefaultClientByToken(conf, appName, tok)
}

func getConf() (*oauth2.Config, error) {
	conf, err := conf.FromFile(shikiPrefix)
	if err != nil {
		// Try create config from env-vars
		return getConfFromEnv()
	}
	return conf, nil
}

func getConfFromEnv() (*oauth2.Config, error) {
	clientID := os.Getenv(shikiClientID)
	clientSecret := os.Getenv(shikiClientSecret)
	redirect := os.Getenv(shikiRedirectURL)

	if clientID == "" || clientSecret == "" || redirect == "" {
		return nil, errors.New("Cannot get clientID, clientSecret and/or redirect url from env")
	}

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     auth.ShikimoriEndpoint,
		RedirectURL:  redirect,
	}
	return conf, nil
}

func getToken(conf *oauth2.Config) (*oauth2.Token, error) {
	tok, err := token.FromFile(shikiPrefix)
	if err != nil {
		// Try get token by login/password
		return getTokenByLogin(conf)
	}

	return tok, nil
}

func getTokenByLogin(conf *oauth2.Config) (*oauth2.Token, error) {
	url := auth.GetAuthCodeURL(conf)
	appName := os.Getenv(shikiAppName)
	login, password := os.Getenv(shikiLogin), os.Getenv(shikiPassword)

	code, err := auth.GetCodeByLogin(url, appName, login, password)
	if err != nil {
		return nil, err
	}

	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func getAppName() string {
	return os.Getenv(shikiAppName)
}

func SaveToken(shiki *api.Shikimori) {
	token.ToFile(shikiPrefix, shiki.Client)
}

package helpers

import (
	"context"
	"errors"
	"os"

	"github.com/seryiza/loadOAuth/conf"
	"github.com/seryiza/loadOAuth/token"
	"golang.org/x/oauth2"

	"github.com/seryiza/goshikimori"
	"github.com/seryiza/goshikimori/auth"
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

// GetShikimori returns goshikimori.Shikimori by config and token from files (using loadOAuth).
// If loadOAuth returns error, try create config/token from env-vars
func GetShikimori(version string) (*goshikimori.Shikimori, error) {
	conf, err := GetConfig()
	if err != nil {
		return nil, err
	}

	tok, err := GetToken(conf)
	if err != nil {
		return nil, err
	}

	appName := GetAppName()
	shiki := CreateShikimoriByToken(conf, tok, appName, version)
	return shiki, nil
}

// GetConfig returns oauth2 config from file. If cannot, tries create from env-vars
func GetConfig() (*oauth2.Config, error) {
	conf, err := conf.FromFile(shikiPrefix)
	if err != nil {
		// Try create config from env-vars
		return getConfigFromEnv()
	}
	return conf, nil
}

// getConfigFromEnv creates oauth2 config from env-variables
func getConfigFromEnv() (*oauth2.Config, error) {
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

// GetToken returns oauth2 token from file. If cannot, try get from env-variables
func GetToken(conf *oauth2.Config) (*oauth2.Token, error) {
	tok, err := token.FromFile(shikiPrefix)
	if err != nil {
		// Try get token by login/password
		return getTokenByLogin(conf)
	}

	return tok, nil
}

// getTokenByLogin gets token by login + password (which get from env-vars)
func getTokenByLogin(conf *oauth2.Config) (*oauth2.Token, error) {
	url := auth.GetAuthCodeURL(conf)
	appName := os.Getenv(shikiAppName)

	login, password, err := getLoginAndPassword()
	if err != nil {
		return nil, err
	}

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

// getLoginAndPassword returns login and password of user from env-vars.
// Return error, if env-vars are empty.
func getLoginAndPassword() (string, string, error) {
	login, password := os.Getenv(shikiLogin), os.Getenv(shikiPassword)
	if login == "" || password == "" {
		return "", "", errors.New("Empty login and/or password from env-vars")
	}

	return login, password, nil
}

// GetAppName returns Shikimori application name
func GetAppName() string {
	return os.Getenv(shikiAppName)
}

// SaveToken to file (if token changed)
func SaveToken(shiki *goshikimori.Shikimori) {
	token.ToFile(shikiPrefix, shiki.Client)
}

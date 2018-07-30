package helpers

import (
	"context"
	"os"

	"github.com/seryiza/loadOAuth/conf"
	"github.com/seryiza/loadOAuth/token"
	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/api"
	"github.com/seryiza/go-shikimori/auth"
)

const (
	shikiPrefix   = "SHIKI"
	shikiLogin    = "SHIKI_LOGIN"
	shikiPassword = "SHIKI_PASS"
	shikiAppName  = "SHIKI_APP_NAME"
)

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
		return nil, err
	}
	return conf, nil
}

func getToken(conf *oauth2.Config) (*oauth2.Token, error) {
	tok, err := token.FromFile(shikiPrefix)
	if err != nil {
		// Try get token by login/password
		tok, err = getTokenByLogin(conf)
		if err != nil {
			return nil, err
		}
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

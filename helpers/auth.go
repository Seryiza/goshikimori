package helpers

import (
	"os"

	"github.com/seryiza/loadOAuth/conf"
	"github.com/seryiza/loadOAuth/token"
	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/api"
)

const (
	shikiPrefix  = "SHIKI"
	shikiAppName = "SHIKI_APP_NAME"
)

func GetShikimori() (*api.Shikimori, error) {
	conf, tok, appName, err := GetAuth()
	if err != nil {
		return nil, err
	}

	return api.DefaultClientByToken(conf, appName, tok)
}

func GetAuth() (*oauth2.Config, *oauth2.Token, string, error) {
	conf, err := conf.FromFile(shikiPrefix)
	if err != nil {
		return nil, nil, "", err
	}

	tok, err := token.FromFile(shikiPrefix)
	if err != nil {
		return nil, nil, "", err
	}

	appName := os.Getenv(shikiAppName)
	return conf, tok, appName, nil
}

func SaveToken(shiki *api.Shikimori) {
	token.ToFile(shikiPrefix, shiki.Client)
}

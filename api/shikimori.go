package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/seryiza/go-shikimori/auth"
	"golang.org/x/oauth2"
)

// Shikimori to send requests to Shikimori
type Shikimori struct {
	Client *http.Client
}

func DefaultClientByCode(conf *oauth2.Config, appName, authCode string) (*Shikimori, error) {
	ctx := context.Background()
	ctx = auth.AddShikimoriTransport(ctx, `Seryiza's Test API`)

	tok, err := conf.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx, tok)
	shiki := &Shikimori{
		Client: client,
	}

	return shiki, nil
}

func DefaultClientByToken(conf *oauth2.Config, appName string, tok *oauth2.Token) (*Shikimori, error) {
	ctx := context.Background()
	ctx = auth.AddShikimoriTransport(ctx, `Seryiza's Test API`)

	client := conf.Client(ctx, tok)
	shiki := &Shikimori{
		Client: client,
	}

	return shiki, nil
}

// GetRequestURLWithQuery returns shikimori url
// todo: сделать две функции ниже более красивыми. Напр., вынести URL в отдельную
// 			 сущность, а добавление/отсутствие query в другие сущности (функции)
func GetRequestURLWithQuery(path string, query url.Values) string {
	url := url.URL{
		Scheme:   "https",
		Host:     "shikimori.org",
		Path:     path,
		RawQuery: query.Encode(),
	}
	return url.String()
}

// GetRequestURL returns shikimori url
func GetRequestURL(path string) string {
	url := url.URL{
		Scheme: "https",
		Host:   "shikimori.org",
		Path:   path,
	}
	return url.String()
}

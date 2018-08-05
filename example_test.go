package goshikimori_test

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/api"
	"github.com/seryiza/go-shikimori/auth"
	"github.com/seryiza/go-shikimori/helpers"
)

func ExampleShikimori() {
	conf := &oauth2.Config{
		ClientID:     "your shikimori client id",
		ClientSecret: "your shikimori client secret",
		RedirectURL:  auth.StandaloneRedirectURL,
		Endpoint:     auth.ShikimoriEndpoint,
	}

	url := auth.GetAuthCodeURL(conf)
	fmt.Println("Enter code from here: ", url)

	var code string
	if _, err := fmt.Scanln(&code); err != nil {
		panic(err)
	}

	shiki, err := api.DefaultClientByCode(conf, "your shikimori app name", code)
	if err != nil {
		panic(err)
	}

	user, err := shiki.Whoami()
	fmt.Println(user, err)
}

func ExampleShikimori_customHTTP() {
	conf := &oauth2.Config{
		ClientID:     "your shikimori client id",
		ClientSecret: "your shikimori client secret",
		RedirectURL:  auth.StandaloneRedirectURL,
		Endpoint:     auth.ShikimoriEndpoint,
	}

	ctx := context.Background()
	// You can add custom http.Client into context as oauth2.HTTPClient.
	// If ctx.Value(oauth2.HTTPClient) == nil, then using DefaultTransport.
	ctx = auth.AddShikimoriTransport(ctx, "your shikimori app name")

	url := auth.GetAuthCodeURL(conf)
	fmt.Println("Enter code from here: ", url)

	var code string
	if _, err := fmt.Scanln(&code); err != nil {
		panic(err)
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}

	shiki := &api.Shikimori{
		Client: conf.Client(ctx, tok),
	}

	user, err := shiki.Whoami()
	fmt.Println(user, err)
}

func ExampleShikimori_withHelpers() {
	// Get Shikimori config and token from files (by env-vars). See `helpers/auth.go`
	shiki, err := helpers.GetShikimori()
	if err != nil {
		panic(err)
	}

	user, err := shiki.Whoami()
	fmt.Println(user, err)
}

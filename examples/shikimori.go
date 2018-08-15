package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"

	"github.com/seryiza/goshikimori"
	"github.com/seryiza/goshikimori/auth"
	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/models"
)

// ExampleShikimori creates and using Shikimori object
func ExampleShikimori() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("SHIKI_CLIENTID"),
		ClientSecret: os.Getenv("SHIKI_CLIENTSECRET"),
		RedirectURL:  auth.StandaloneRedirectURL,
		Endpoint:     auth.ShikimoriEndpoint,
	}

	url := auth.GetAuthCodeURL(conf)
	fmt.Println("Enter code from here: ", url)

	var code string
	if _, err := fmt.Scanln(&code); err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx = goshikimori.AddShikimoriTransport(ctx, os.Getenv("SHIKI_APP_NAME"))

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, tok)
	shiki := goshikimori.NewShikimori(client, "1.0")

	resp, err := shiki.Get("users/whoami")
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(user); err != nil {
		panic(err)
	}

	fmt.Printf("I'm %s", user.Nickname)
}

// ExampleShikimoriWithHelpers creates Shikimori using `helpers` package
func ExampleShikimoriWithHelpers() {
	// GetShikimori gets oauth2 config & token from files or env-variables
	// and create Shikimori from default http client.
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		panic(err)
	}
	// Save token into file, if changed.
	defer helpers.SaveToken(shiki)

	user := &models.User{}
	_, err = shiki.JSONGet("users/whoami", user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("I'm %s", user.Nickname)
}

// PrintShikiGet print GET response from Shikimori
func PrintShikiGet(method string) {
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		panic(err)
	}
	defer helpers.SaveToken(shiki)

	resp, err := shiki.Get(method)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}

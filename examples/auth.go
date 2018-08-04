package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/seryiza/go-shikimori/auth"
)

func getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("SHIKI_CLIENTID"),
		ClientSecret: os.Getenv("SHIKI_CLIENTSECRET"),
		Endpoint:     auth.ShikimoriEndpoint,
		RedirectURL:  os.Getenv("SHIKI_REDIRECT_URL"),
	}
}

func ExampleGetConfig() {
	conf := getConfig()
	json, _ := json.Marshal(conf)
	fmt.Println(string(json))
}

func ExampleGetToken() {
	conf := getConfig()
	url := auth.GetAuthCodeURL(conf)

	var code string
	fmt.Println("Введите код из этой ссылки: ", url)
	if _, err := fmt.Scanln(&code); err != nil {
		panic(err)
	}

	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		panic(err)
	}

	json, _ := json.Marshal(tok)
	fmt.Println(string(json))
}

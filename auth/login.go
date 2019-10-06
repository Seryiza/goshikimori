package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"html"

	"github.com/headzoo/surf"
)

const (
	createSessionURL = `https://shikimori.one/api/sessions`
	jsonErrorKey     = "error"

	authAgreeForm = "form.authorize"
	authCodeField = "#authorization_code"
)

type userSession struct {
	User userCredentials `json:"user"`
}

type userCredentials struct {
	Login    string `json:"nickname"`
	Password string `json:"password"`
}

// GetCodeByLogin gets auth code by login + password.
// Application redirect url must be StandaloneRedirectURL (see endpoind.go).
func GetCodeByLogin(url, appName, login, password string) (string, error) {
	// todo: Разбить эту длинную функцию
	bow := surf.NewBrowser()
	bow.SetUserAgent(appName)

	if login == "" || password == "" {
		return "", errors.New("Empty login and/or password")
	}

	user := &userSession{
		User: userCredentials{
			Login:    login,
			Password: password,
		},
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	err = bow.Post(
		createSessionURL,
		`application/json`,
		bytes.NewReader(userBytes),
	)
	if err != nil {
		return "", err
	}

	loginResponse := make(map[string]interface{})
	respJSON := html.UnescapeString(bow.Body())
	if err = json.Unmarshal([]byte(respJSON), &loginResponse); err != nil {
		return "", err
	}

	if errMsg, ok := loginResponse[jsonErrorKey]; ok {
		return "", errors.New(errMsg.(string))
	}

	err = bow.Open(url)
	if err != nil {
		return "", err
	}

	// if need click to "Agree"
	agreeForm, _ := bow.Form(authAgreeForm)
	if agreeForm != nil {
		agreeForm.Submit()
	}

	code := bow.Find(authCodeField).Text()
	return code, nil
}

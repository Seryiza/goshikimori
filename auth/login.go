package auth

import (
	"errors"

	"github.com/headzoo/surf"
)

const (
	signinURL     = "/users/sign_in"
	signinForm    = "form.new_user"
	loginField    = "user[nickname]"
	passwordField = "user[password]"

	authAgreeForm = "form.authorize"
	authCodeField = "#authorization_code"
)

// GetCodeByLogin gets auth code by login + password
func GetCodeByLogin(url, login, password string) (string, error) {
	bow := surf.NewBrowser()
	err := bow.Open(url)
	if err != nil {
		return "", err
	}

	// if need to login
	if bow.Url().RequestURI() == signinURL {
		form, err := bow.Form(signinForm)
		if err != nil {
			return "", err
		}

		form.Input(loginField, login)
		form.Input(passwordField, password)
		if err = form.Submit(); err != nil {
			return "", err
		}

		err = bow.Open(url)
		if err != nil {
			return "", err
		}

		if bow.Url().RequestURI() == signinURL {
			return "", errors.New("Wrong login and/or password")
		}
	}

	// if need click to "Agree"
	agreeForm, _ := bow.Form(authAgreeForm)
	if agreeForm != nil {
		agreeForm.Submit()
	}

	code := bow.Find(authCodeField).Text()
	return code, nil
}

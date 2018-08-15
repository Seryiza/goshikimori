package goshikimori_test

import (
	"fmt"
	"testing"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/models"
)

// Test Shikimori struct and api call (not models)
func TestGetUser(t *testing.T) {
	userID := 206253
	wantNickname := "Seryiza"
	method := fmt.Sprintf("users/%d", userID)

	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	user := &models.User{}
	if _, err = shiki.JSONGet(method, user); err != nil {
		t.Error(err)
	}

	if user.Nickname != wantNickname {
		t.Errorf("Given nickname '%v' is not '%v'", user.Nickname, wantNickname)
	}
}

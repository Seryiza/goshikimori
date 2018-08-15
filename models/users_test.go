package models_test

import (
	"testing"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/models"
)

func TestGetUsers(t *testing.T) {
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	users := make(models.Users, 0)
	if _, err = shiki.JSONGet("users", &users); err != nil {
		t.Error(err)
	}

	if len(users) == 0 {
		t.Error("Users slice is empty")
	}

	if users[0].ID == 0 {
		t.Error("User ID is default int value")
	}
}

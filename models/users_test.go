package models_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/methods"
	"github.com/seryiza/goshikimori/models"
)

// Test GET /api/users
func TestGetUsers(t *testing.T) {
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	users := make(models.Users, 0)
	if _, err = shiki.JSONGet(methods.Users, &users); err != nil {
		t.Error(err)
	}

	if len(users) == 0 {
		t.Error("Users slice is empty")
	}

	if users[0].ID == 0 {
		t.Error("User ID is default int value")
	}
}

// Test GET /api/users/:id/info
func TestGetUserInfo(t *testing.T) {
	haveID := 386084
	haveMethod := fmt.Sprintf(methods.UsersInfo, haveID)

	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	userBrief := &models.UserBrief{}
	if _, err = shiki.JSONGet(haveMethod, userBrief); err != nil {
		t.Error(err)
	}

	if userBrief.LastOnlineAt.IsZero() {
		t.Error("Last online is zero")
	}

	// Время онлайн меняется, поэтому нельзя сравнивать его с константой
	someUserBrief.LastOnlineAt = userBrief.LastOnlineAt

	if !reflect.DeepEqual(userBrief, someUserBrief) {
		t.Errorf("UserBrief is invalid.\nWant: %v\nGot: %v\n", someUserBrief, userBrief)
	}
}

var someUserBrief = &models.UserBrief{
	Locale: "ru",
	User: models.User{
		ID:       386084,
		Nickname: "SeryizasBot",
		Avatar:   "https://dere.shikimori.org/system/users/x48/386084.png?1532088739",
		Image: models.UserImage{
			X160: "https://dere.shikimori.org/system/users/x160/386084.png?1532088739",
			X148: "https://dere.shikimori.org/system/users/x148/386084.png?1532088739",
			X80:  "https://dere.shikimori.org/system/users/x80/386084.png?1532088739",
			X64:  "https://dere.shikimori.org/system/users/x64/386084.png?1532088739",
			X48:  "https://dere.shikimori.org/system/users/x48/386084.png?1532088739",
			X32:  "https://dere.shikimori.org/system/users/x32/386084.png?1532088739",
			X16:  "https://dere.shikimori.org/system/users/x16/386084.png?1532088739",
		},

		LastOnlineAt: time.Time{},
	},
}

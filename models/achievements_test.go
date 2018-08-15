package models_test

import (
	"net/url"
	"testing"

	"github.com/seryiza/goshikimori"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/methods"
	"github.com/seryiza/goshikimori/models"
)

func TestGetAchievement(t *testing.T) {
	// todo: Пока указан User ID основного аккаунта. Как система достижений
	//       официально заработает, сменить на аккаунт бота.
	//       И сделать более подробную проверку через reflect.DeepEqual.

	haveVals := url.Values{"user_id": {"206253"}}
	haveMethod := goshikimori.FormatQuery(methods.Achievements, haveVals)

	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	achievements := make(models.Achievements, 0)
	if _, err := shiki.JSONGet(haveMethod, &achievements); err != nil {
		t.Error(err)
	}

	if len(achievements) == 0 {
		t.Error("Achievement list is empty")
	}

	if achievements[0].UserID == 0 {
		t.Error("Achievement user ID is 0 (default int value)")
	}
}

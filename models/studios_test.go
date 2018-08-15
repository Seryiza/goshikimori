package models_test

import (
	"reflect"
	"testing"

	"github.com/seryiza/goshikimori/methods"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/models"
)

func TestGetStudios(t *testing.T) {
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	studios := make(models.Studios, 0)
	if _, err = shiki.JSONGet(methods.Studios, &studios); err != nil {
		t.Error(err)
	}

	if len(studios) == 0 {
		t.Error("List of studios is empty")
	}

	emptyStudio := models.Studio{}
	if reflect.DeepEqual(studios[0], emptyStudio) {
		t.Error("First studio is empty struct")
	}
}

package api_test

import (
	"testing"

	"github.com/seryiza/go-shikimori/helpers"
)

func TestGetStudios(t *testing.T) {
	shiki, err := helpers.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	studios, err := shiki.GetStudiosList()
	if err != nil {
		t.Error(err)
	}

	if len(studios) == 0 {
		t.Error("List of studios is empty")
	}

	if studios[0].Name == "" {
		t.Error("Studio name is empty")
	}
}

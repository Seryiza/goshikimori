package api_test

import (
	"testing"

	"github.com/seryiza/go-shikimori/helpers"
)

func TestGetGenres(t *testing.T) {
	shiki, err := helpers.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	genres, err := shiki.GetGenresList()
	if err != nil {
		t.Error(err)
	}

	if len(genres) == 0 {
		t.Error("List of genres is empty")
	}

	if genres[0].Name == "" {
		t.Error("Genre name is empty")
	}
}

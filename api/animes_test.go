package api_test

import (
	"testing"

	"github.com/seryiza/go-shikimori/api"
	"github.com/seryiza/go-shikimori/helpers"
)

func TestSimpleGetAnimes(t *testing.T) {
	shiki, err := helpers.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	animes, err := shiki.GetAnimes(&api.GetAnimeOpts{
		IDFilter: "1",
	})

	if err != nil {
		t.Error(err)
	}

	if animes[0].Name != "Cowboy Bebop" {
		t.Error("Anime name isn't Cowboy Bebop")
	}
}

func TestSimpleGetOneAnime(t *testing.T) {
	shiki, err := helpers.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	anime, err := shiki.GetAnime(1)

	if err != nil {
		t.Error(err)
	}

	if anime.Name != "Cowboy Bebop" {
		t.Error("Anime name isn't Cowboy Bebop")
	}
}

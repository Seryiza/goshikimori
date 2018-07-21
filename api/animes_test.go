package api_test

import (
	"testing"

	"github.com/seryiza/go-shikimori/api"
	"github.com/seryiza/go-shikimori/api/shikiTesting"
)

func TestSimpleGetAnimes(t *testing.T) {
	shiki, err := shikiTesting.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer shikiTesting.RefreshToken(shiki)

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

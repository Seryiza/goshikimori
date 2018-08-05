package api_test

import (
	"testing"

	"github.com/seryiza/go-shikimori/helpers"
)

func TestGetVideos(t *testing.T) {
	shiki, err := helpers.GetShikimori()
	if err != nil {
		t.Error(err)
	}
	defer helpers.SaveToken(shiki)

	videos, err := shiki.GetVideos(1)
	if err != nil {
		t.Error(err)
	}

	if len(videos) == 0 {
		t.Error("List of videos is empty")
	}
}

package goshikimori_test

import (
	"net/url"
	"testing"

	"github.com/seryiza/goshikimori"
)

func TestFormatQuery(t *testing.T) {
	haveMethod := "users"
	haveValues := url.Values{
		"limit": {"100"},
	}
	wantURL := "users?limit=100"

	gotURL := goshikimori.FormatQuery(haveMethod, haveValues)
	if gotURL != wantURL {
		t.Errorf("FormatQuery is invalid; Want %v, got %v", wantURL, gotURL)
	}
}

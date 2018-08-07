package helpers

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/seryiza/goshikimori"
)

// CreateShikimoriByToken returns Shikimori based on http.DefaultClient
func CreateShikimoriByToken(conf *oauth2.Config, tok *oauth2.Token, appName, version string) *goshikimori.Shikimori {
	ctx := context.Background()
	ctx = goshikimori.AddShikimoriTransport(ctx, appName)

	client := conf.Client(ctx, tok)
	shiki := goshikimori.NewShikimori(client, version)
	return shiki
}

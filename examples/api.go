package examples

import (
	"fmt"

	"github.com/seryiza/goshikimori/helpers"
	"github.com/seryiza/goshikimori/methods"
	"github.com/seryiza/goshikimori/models"
)

// SendMessage to the Shikimori user
func SendMessage(from, to, body string) {
	shiki, err := helpers.GetShikimori("1.0")
	if err != nil {
		panic(err)
	}
	defer helpers.SaveToken(shiki)

	message := models.Message{
		Inner: models.InnerMessage{
			FromID: from,
			ToID:   to,
			Body:   body,
			Kind:   models.MessageKindPrivate,
		},
	}
	result := models.MessageResult{}

	if _, err = shiki.JSONPost(methods.PostMessage, message, &result); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}

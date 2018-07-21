package api

import (
	"encoding/json"

	"github.com/pasztorpisti/qs"

	"github.com/seryiza/go-shikimori/api/structs"
)

type GetAnimeOpts struct {
	// Search phrase to filter animes by name.
	SearchText string `qs:"search,omitempty"`

	// Number of search page.
	// Must be a number between 1 and 100000.
	Page int32 `qs:"page,omitempty"`

	// Anime maximum count.
	// Must be a number between 1 and 50.
	Limit int8 `qs:"limit,omitempty"`

	// Minimal anime score.
	// todo: проверить, может ли это быть float
	MinScore int `qs:"score,omitempty"`

	// Duration of episode.
	// See structs/animes.go for consts.
	Duration string `qs:"duration,omitempty"`

	// Age rating of anime.
	// See structs/animes.go for consts.
	Rating string `qs:"rating,omitempty"`

	// Current anime status (ongoing/released/anons).
	// See structs/animes.go for consts.
	Status string `qs:"status,omitempty"`

	// Status of anime in user list.
	// See structs/animes.go for consts.
	MyListStatus string `qs:"mylist,omitempty"`

	// Sort order.
	// See structs/animes.go for consts.
	Order string `qs:"order,omitempty"`

	// Type/kind of anime.
	// See structs/animes.go for consts.
	Kind string `qs:"kind,omitempty"`

	// Expression for search by seasons.
	// Ex.: "summer_2017", "2016", "2014_2016", "199x"
	SeasonFilter string `qs:"season,omitempty"`

	// List of genre ids separated by comma.
	GenreFilter string `qs:"genre,omitempty"`

	// List of studio ids separated by comma.
	StudioFilter string `qs:"studio,omitempty"`

	// List of franchises separated by comma.
	FranchiseFilter string `qs:"franchise,omitempty"`

	// Lists of anime ids separated by comma.
	IDFilter        string `qs:"ids,omitempty"`
	IDExcludeFilter string `qs:"exclude_ids,omitempty"`

	// Set to false to allow hentai, yaoi and yuri.
	// todo: Можно ли здесь как-нибудь красиво-удобно добавить omitempty?
	Censored bool `qs:"censored"`
}

// GetAnimes implements GET /api/anime
// https://shikimori.org/api/doc/1.0/animes/index
func (shiki *Shikimori) GetAnimes(opts *GetAnimeOpts) (structs.Animes, error) {
	urlVals, err := qs.Marshal(opts)
	if err != nil {
		return nil, err
	}

	url := shiki.ApiURLWithString(getAnimePath, urlVals)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	jr := json.NewDecoder(resp.Body)
	animes := make(structs.Animes, 0)
	if err = jr.Decode(&animes); err != nil {
		return nil, err
	}

	return animes, nil
}

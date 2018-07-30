package api

import (
	"encoding/json"

	"github.com/seryiza/go-shikimori/api/structs"
)

// GetGenresList implements /api/genres
// https://shikimori.org/api/doc/1.0/genres/index
func (shiki *Shikimori) GetGenresList() (structs.Genres, error) {
	url := shiki.ApiURL(getGenresList)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	genres := make(structs.Genres, 0)
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(&genres); err != nil {
		return nil, err
	}

	return genres, nil
}

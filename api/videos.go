package api

import (
	"encoding/json"
	"fmt"

	"github.com/seryiza/go-shikimori/api/structs"
)

// GetVideos implemets GET /api/animes/:anime_id/videos
// https://shikimori.org/api/doc/1.0/videos/index
func (shiki *Shikimori) GetVideos(animeID int32) (structs.Videos, error) {
	path := fmt.Sprintf(getVideoFormat, animeID)
	url := shiki.ApiURL(path)

	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	videos := make(structs.Videos, 0)
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(&videos); err != nil {
		return nil, err
	}

	return videos, nil
}

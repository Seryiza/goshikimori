package api

import (
	"encoding/json"

	"github.com/seryiza/go-shikimori/api/structs"
)

// GetStudiosList implements /api/studios
// https://shikimori.org/api/doc/1.0/studios/index
func (shiki *Shikimori) GetStudiosList() (structs.Studios, error) {
	url := shiki.ApiURL(getStudiosList)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	studios := make(structs.Studios, 0)
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(&studios); err != nil {
		return nil, err
	}

	return studios, nil
}

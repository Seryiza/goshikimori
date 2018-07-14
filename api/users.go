package api

import (
	"encoding/json"

	"github.com/seryiza/go-shikimori/api/structs"
)

const (
	whoamiPath = "/api/users/whoami"
)

// Whoami implements GET /api/users/whoami
// https://shikimori.org/api/doc/1.0/users/whoami
func (shiki *Shikimori) Whoami() (*structs.User, error) {
	url := shiki.ApiURL(whoamiPath)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	user := &structs.User{}
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

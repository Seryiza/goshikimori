package api

import (
	"encoding/json"
	"time"
)

const (
	whoamiPath = "/api/users/whoami"
)

// User of Shikimori
type User struct {
	ID         int32     `json:"id"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Image      UserImage `json:"image"`
	LastOnline time.Time `json:"last_online_at"`

	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Website   string `json:"website"`
	BirthDate string `json:"birth_on"`
	Locale    string `json:"locale"`
}

// UserImage of Shikimori user
type UserImage struct {
	X160 string `json:"x160"`
	X148 string `json:"x148"`
	X80  string `json:"x80"`
	X64  string `json:"x64"`
	X48  string `json:"x48"`
	X32  string `json:"x32"`
	X16  string `json:"x16"`
}

// Whoami implements GET /api/users/whoami
// https://shikimori.org/api/doc/1.0/users/whoami
func (shiki *Shikimori) Whoami() (*User, error) {
	url := GetRequestURL(whoamiPath)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	user := &User{}
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

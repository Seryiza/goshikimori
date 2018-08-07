package structs

import "time"

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

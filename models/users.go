package models

import "time"

type Users []User

// User of Shikimori.
// Ex., GET /api/users
type User struct {
	ID       int32  `json:"id"`
	Nickname string `json:"nickname"`

	Avatar string    `json:"avatar"`
	Image  UserImage `json:"image"`

	LastOnlineAt time.Time `json:"last_online_at"`
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

// UserBrief of Shikimori.
// Ex., GET /api/users/:id/info
type UserBrief struct {
	User

	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Website   string `json:"website"`
	BirthDate string `json:"birth_on"`
	Locale    string `json:"locale"`
}

// UserDetailed of Shikimori.
// Ex., GET /api/users/:id
type UserDetailed struct {
	// todo
	UserBrief
}

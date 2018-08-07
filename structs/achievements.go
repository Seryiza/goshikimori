package structs

import "time"

// Achievements of some user
type Achievements []Achievement

// Achievement of Shikimori
type Achievement struct {
	ID     int32  `json:"id"`
	NekoID string `json:"neko_id"`
	Level  int32  `json:"level"`

	// Progress has int8 because it changes 0..100 (?)
	Progress int8  `json:"progress"`
	UserID   int32 `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

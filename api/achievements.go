package api

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

const (
	getAchievementsPath = "/api/achievements"
)

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

// GetAchievements implemets GET /api/achievements
// https://shikimori.org/api/doc/1.0/achievements/index
func (shiki *Shikimori) GetAchievements(userID int) (Achievements, error) {
	strUserID := strconv.Itoa(userID)
	query := url.Values{
		"user_id": {strUserID},
	}

	url := GetRequestURLWithQuery(getAchievementsPath, query)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	jr := json.NewDecoder(resp.Body)
	achs := make(Achievements, 0)
	if err = jr.Decode(&achs); err != nil {
		return nil, err
	}

	return achs, nil
}

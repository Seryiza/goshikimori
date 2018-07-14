package api

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/seryiza/go-shikimori/api/structs"
)

const (
	getAchievementsPath = "/api/achievements"
)

// GetAchievements implemets GET /api/achievements
// https://shikimori.org/api/doc/1.0/achievements/index
func (shiki *Shikimori) GetAchievements(userID int) (structs.Achievements, error) {
	strUserID := strconv.Itoa(userID)
	query := url.Values{
		"user_id": {strUserID},
	}

	url := shiki.ApiURLWithQuery(getAchievementsPath, query)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	jr := json.NewDecoder(resp.Body)
	achs := make(structs.Achievements, 0)
	if err = jr.Decode(&achs); err != nil {
		return nil, err
	}

	return achs, nil
}

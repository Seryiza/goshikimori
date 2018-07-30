package api

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/seryiza/go-shikimori/api/structs"
)

// GetAchievements implemets GET /api/achievements
// https://shikimori.org/api/doc/1.0/achievements/index
func (shiki *Shikimori) GetAchievements(userID int) (structs.Achievements, error) {
	urlVals := url.Values{
		"user_id": {strconv.Itoa(userID)},
	}

	url := shiki.ApiURLWithValues(getAchievementsPath, urlVals)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	achievements := make(structs.Achievements, 0)
	jd := json.NewDecoder(resp.Body)
	if err = jd.Decode(&achievements); err != nil {
		return nil, err
	}

	return achievements, nil
}

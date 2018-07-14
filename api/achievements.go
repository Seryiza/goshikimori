package api

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

const (
	getAchievementsPath = "/api/achievements"
)

// GetAchievements implemets GET /api/achievements
// https://shikimori.org/api/doc/1.0/achievements/index
func (shiki *Shikimori) GetAchievements(userID int) (json.RawMessage, error) {
	// todo: как появится больше информации об этом, сделать более удобное
	// 		   получение информации (без json)

	strUserID := strconv.Itoa(userID)

	query := url.Values{
		"user_id": {strUserID},
	}

	url := GetRequestURLWithQuery(getAchievementsPath, query)
	resp, err := shiki.Client.Get(url)
	if err != nil {
		return nil, err
	}

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return json, nil
}

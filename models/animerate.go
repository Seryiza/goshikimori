package models

type AnimeRates []AnimeRate

type AnimeRate struct {
	ID     int32  `json:"id"`
	Score  int8   `json:"score"`
	Status string `json:"status"`

	Text     string `json:"text"`
	HTMLText string `json:"text_html"`

	Episodes  int16 `json:"episodes"`
	Rewatches int8  `json:"rewatches"`
}

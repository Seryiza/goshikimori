package models

type Videos []Video

type Video struct {
	ID        int32  `json:"id"`
	URL       string `json:"url"`
	Image     string `json:"image_url"`
	PlayerURL string `json:"player_url"`
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Hosting   string `json:"hosting"`
}

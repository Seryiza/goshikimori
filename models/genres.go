package models

type Genres []Genre

type Genre struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}

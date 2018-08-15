package models

// Studios on Shikimori
type Studios []Studio

// Studio on Shikimori
type Studio struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	FilteredName string `json:"filtered_name"`
	Real         bool   `json:"real"`
	Image        string `json:"image"`
}

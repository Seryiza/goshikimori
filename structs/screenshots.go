package structs

type Screenshots []Screenshot

type Screenshot struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
}

package structs

type Animes []Anime

type Anime struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	RussianName string `json:"russian"`

	Image  AnimeImage `json:"image"`
	URL    string     `json:"url"`
	Kind   string     `json:"kind"`
	Status string     `json:"status"`

	Episodes      int16 `json:"episodes"`
	EpisodesAired int16 `json:"episodes_aired"`

	// todo: посмотреть, можно ли их как-нибудь легко/удобно перевести в time.Time
	AiredOn    string `json:"aired_on"`
	ReleasedOn string `json:"released_on"`
}

type AnimeImage struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}

const (
	// kinds of anime
	AnimeKindTV      = "tv"
	AnimeKindMovie   = "movie"
	AnimeKindOVA     = "ova"
	AnimeKindONA     = "ona"
	AnimeKindSpecial = "special"
	AnimeKindMusic   = "music"
	AnimeKindTV13    = "tv_13"
	AnimeKindTV24    = "tv_24"
	AnimeKindTV48    = "tv_48"

	// status of anime
	AnimeStatusAnons    = "anons"
	AnimeStatusOngoing  = "ongoing"
	AnimeStatusReleased = "released"
)

package structs

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

	// durations of episode
	AnimeDurationLess10Min = "S"
	AnimeDurationLess30Min = "D"
	AnimeDurationMore30Min = "F"

	// age rating of episode
	AnimeRatingNone  = "none"   // none – No rating
	AnimeRatingG     = "g"      // G - All ages
	AnimeRatingPG    = "pg"     // PG - Children
	AnimeRatingPG13  = "pg_13"  // PG-13 - Teens 13 or older
	AnimeRatingR     = "r"      // R - 17+ recommended (violence & profanity)
	AnimeRatingRPlus = "r_plus" // R+ - Mild Nudity (may also contain violence & profanity)
	AnimeRatingRx    = "rx"     // Rx - Hentai (extreme sexual content/nudity)

	// user list statuses
	AnimeListPlanned    = "planned"
	AnimeListWatching   = "watching"
	AnimeListRewatching = "rewatching"
	AnimeListCompleted  = "completed"
	AnimeListOnHold     = "on_hold"
	AnimeListDropped    = "dropped"

	// sort order
	AnimeOrderByID         = "id"
	AnimeOrderByRanked     = "ranked"
	AnimeOrderByKind       = "kind"
	AnimeOrderByPopularity = "popularity"
	AnimeOrderByName       = "name"
	AnimeOrderByAiredOn    = "aired_on"
	AnimeOrderByEpisodes   = "episodes"
	AnimeOrderByStatus     = "status"
	AnimeOrderByRandom     = "random"
)

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

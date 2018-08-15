package models

const (
	// kind of anime video
	AnimeVideoKindRaw       = "raw"
	AnimeVideoKindSubtitles = "subtitles"
	AnimeVideoKindFandub    = "fandub"
	AnimeVideoKindUnknown   = "unknown"

	// languages of anime video
	AnimeVideoLanguageRussian  = "russian"
	AnimeVideoLanguageEnglish  = "english"
	AnimeVideoLanguageOriginal = "original"
	AnimeVideoLanguageUnknown  = "unknown"

	// quality of video
	AnimeVideoQualityBD      = "bd"
	AnimeVideoQualityWeb     = "web"
	AnimeVideoQualityTV      = "tv"
	AnimeVideoQualityDVD     = "dvd"
	AnimeVideoQualityUnknown = "unknown"
)

type AnimeVideo struct {
	AnimeID  int32  `json:"anime_id"`
	Episode  int16  `json:"episode"`
	Kind     string `json:"kind"`
	Language string `json:"language"`
	Quality  string `json:"quality"`
	Source   string `json:"source"`
	URL      string `json:"url"`

	// Формат записи: Название_проекта/студии (Ник_даббера1 & Ник_даббера2).
	// Поле необязательное.
	AuthorName string `json:"author_name,omitempty"`
}

package jsonfile

type Repository struct {
	Provider Slug `json:"provider"`
	Url      Url  `json:"url"`
}

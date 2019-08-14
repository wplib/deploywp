package jsonfile

import (
	"net/url"
)

type Repository struct {
	Provider Slug    `json:"provider"`
	Url      Url     `json:"url"`
	url      url.URL `json:"-"`
}

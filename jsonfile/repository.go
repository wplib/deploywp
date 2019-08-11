package jsonfile

import (
	"github.com/wplib/deploywp/providers"
	"net/url"
)

type Repository struct {
	ProviderSlug Slug               `json:"provider"`
	provider     providers.Provider `json:"-"`
	Url          Url                `json:"url"`
	url          url.URL            `json:"-"`
}

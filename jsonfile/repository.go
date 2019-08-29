package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
	"net/url"
)

var _ deploywp.RepositoryGetter = (*Repository)(nil)

type Repository struct {
	Provider Identifier `json:"provider"`
	Url      Url        `json:"url"`
}

func (me Repository) GetProvider() *providers.Provider {
	return nil
}

func (me Repository) GetUrl() *url.URL {
	return nil
}

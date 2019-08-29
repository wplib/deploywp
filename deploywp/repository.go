package deploywp

import (
	"github.com/wplib/deploywp/providers"
	"net/url"
)

type Repository struct {
	Provider *providers.Provider
	Url      *url.URL
}

type RepositoryGetter interface {
	GetProvider() *providers.Provider
	GetUrl() *url.URL
}

func NewRepositoryFromGetter(rg RepositoryGetter) (r *Repository) {
	return &Repository{
		Provider: rg.GetProvider(),
		Url:      rg.GetUrl(),
	}
}

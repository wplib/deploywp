package deploywp

import (
	"github.com/wplib/deploywp/providers"
)

type Repository struct {
	Provider providers.Provider
	Url      Url
}

type RepositoryGetter interface {
	GetProvider() providers.Provider
	GetUrl() Url
}

func NewRepositoryFromGetter(rg RepositoryGetter) (r *Repository) {
	return &Repository{
		Provider: rg.GetProvider(),
		Url:      rg.GetUrl(),
	}
}

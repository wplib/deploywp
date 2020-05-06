package deploywp

import (
	"github.com/wplib/deploywp/providers"
)

type Repository = providers.Repository
type RepositoryGetter = providers.RepositoryGetter

func NewRepositoryFromGetter(rg RepositoryGetter) (r *Repository) {
	return &Repository{
		Provider: rg.GetProvider(),
		Url:      rg.GetUrl(),
	}
}

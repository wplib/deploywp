package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
)

var _ deploywp.RepositoryGetter = (*Repository)(nil)

type Repository struct {
	ProviderId    Identifier `json:"provider"`
	Url           Url        `json:"url"`
	provider      providers.Provider
	urlNormalized bool
}

func (me Repository) GetProviderId() (pid providers.ProviderId) {
	for range Once {
		if me.provider != nil {
			pid = me.provider.GetId()
			break
		}
		if me.ProviderId != "" {
			pid = me.ProviderId
			break
		}
		p := providers.DetectByUrl(me.GetUrl())
		pid = p.GetId()
	}
	return pid
}

func (me Repository) GetProvider() providers.Provider {
	return providers.Dispense(me.GetProviderId())
}

func (me Repository) GetUrl() providers.Url {
	for range Once {
		if me.Url == "" {
			break
		}
		if me.urlNormalized {
			break
		}
		me.urlNormalized = true
		p := providers.Dispense(me.GetProviderId())
		me.Url = p.NormalizeUrl(me.Url)
	}
	return me.Url
}

func NewRepository(pid providers.ProviderId, u Url) *Repository {
	return &Repository{
		ProviderId: pid,
		Url:        u,
	}
}

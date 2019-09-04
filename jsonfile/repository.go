package jsonfile

import (
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
	"net/url"
)

var _ deploywp.RepositoryGetter = (*Repository)(nil)

type Repository struct {
	Provider Identifier         `json:"provider"`
	Url      Url                `json:"url"`
	provider providers.Provider `json:"-"`
}

func (me Repository) GetProvider() providers.Provider {
	for range Once {
		if me.Provider != "" {
			me.provider = providers.Dispense(me.Provider)
			break
		}
		me.provider, me.Url = providers.DetectByUrl(me.Url)
		me.Provider = me.provider.GetId()
	}
	if me.provider == nil {
		app.Fail("provider not specified or detectable for '%s'", me.Url)
	}
	return me.provider
}

func (me Repository) GetUrl() *url.URL {
	if me.Provider == "" {
		_, me.Url = providers.DetectByUrl(me.Url)
	}
	return providers.ParseUrl(me.Url, "repository")
}

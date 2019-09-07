package jsonfile

import (
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
)

var _ deploywp.RepositoryGetter = (*Repository)(nil)

type Repository struct {
	Provider Identifier         `json:"provider"`
	Url      Url                `json:"url"`
	provider providers.Provider `json:"-"`
}

func (me Repository) GetProvider() providers.Provider {
	for range Once {
		if me.provider != nil {
			break
		}
		if me.Provider != "" {
			me.provider = providers.Dispense(me.Provider)
		}
		if me.provider != nil {
			break
		}
		me.SetProvider(me.detectProvider())
	}
	return me.provider
}

func (me Repository) GetUrl() Url {
	me.normalizeUrl()
	return me.Url
}

func (me Repository) detectProvider() (p providers.Provider) {
	return providers.DetectByUrl(me.Url)
}

func (me *Repository) SetProvider(p providers.Provider) {
	me.provider = p
	me.Provider = p.GetId()
}

func (me *Repository) SetUrl(u Url) {
	me.Url = u
}

func (me *Repository) normalizeUrl() {
	me.ensureProvider()
	me.ensureURL()
	me.SetUrl(me.provider.NormalizeUrl(me.Url))
}

func (me *Repository) ensureURL() {
	if me.Url == "" {
		app.Fail("URL not specified for repository")
	}
}

func (me *Repository) ensureProvider() {
	for range Once {
		p := me.provider
		if p == nil {
			p = me.detectProvider()
		}
		me.SetProvider(p)
	}
}

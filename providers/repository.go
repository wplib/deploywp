package providers

import (
	"github.com/wplib/deploywp/app"
)

var _ RepositoryGetter = (*Repository)(nil)

type Repository struct {
	Provider Provider
	Url      Url
}

func (me Repository) GetUrl() Url {
	return me.Url
}

func (me Repository) GetProvider() Provider {
	return me.Provider
}

func (me Repository) detectProvider() (p Provider) {
	for range Once {
		if me.Provider != nil {
			break
		}
		p = DetectByUrl(me.Url)
	}
	return p
}

func (me Repository) SetProvider(p Provider) {
	me.Provider = p
}

func (me Repository) SetUrl(u Url) {
	me.Url = u
}

func (me Repository) normalizeUrl() {
	me.ensureProvider()
	me.ensureURL()
	me.SetUrl(me.Provider.NormalizeUrl(me.Url))
}

func (me Repository) ensureURL() {
	if me.Url == "" {
		app.Fail("URL not specified for repository")
	}
}

func (me Repository) ensureProvider() {
	for range Once {
		p := me.Provider
		if p != nil {
			break
		}
		me.SetProvider(me.detectProvider())
	}
}

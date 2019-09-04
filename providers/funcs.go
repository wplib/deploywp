package providers

import (
	"github.com/asaskevich/govalidator"
	"github.com/wplib/deploywp/app"
	"net/url"
)

var providersMap Map = Map{
	GitHubId:   NewGitHubProvider(),
	PantheonId: NewPantheonProvider(),
}

func Register(pid ProviderId, p Provider) {
	providersMap[pid] = p
}

func Dispense(pid ProviderId) (p Provider) {
	var ok bool
	p, ok = providersMap[pid]
	if !ok {
		app.Fail("Invalid provider ID '%s'", pid)
	}
	return p
}

func DetectByUrl(u Url) (p Provider, nu Url) {
	for range Once {
		for _, _p := range providersMap {
			var d bool
			d, nu = _p.DetectByUrl(u)
			if !d {
				continue
			}
			p = _p
			break
		}
		if govalidator.IsURL(nu) {
			break
		}
		app.Fail("Invalid URL '%s'", nu)
	}
	if p == nil {
		app.Fail("Cannot detect provider by URL '%s'", u)
	}
	return p, nu
}

func ParseUrl(u Url, kind ReadableName) *url.URL {
	uu, err := url.Parse(u)
	if err != nil {
		app.Fail("Invalid URL '%s' for %s; parse error: %s", u, kind, err.Error())
	}
	return uu
}

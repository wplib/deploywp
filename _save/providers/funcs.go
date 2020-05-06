package providers

import (
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

func Dispense(pid ProviderId) Provider {
	p, ok := providersMap[pid]
	if !ok {
		app.Fail("Invalid provider ID '%s'", pid)
	}
	return p
}

func DetectByUrl(u Url) (p Provider) {
	for _, _p := range providersMap {
		if !_p.DetectByUrl(u) {
			continue
		}
		p = _p
		break
	}
	if p == nil {
		app.Fail("Cannot detect provider by URL '%s'", u)
	}
	return p
}

func ParseUrl(u Url, kind ReadableName) *url.URL {
	uu, err := url.Parse(u)
	if err != nil {
		app.Fail("invalid URL '%s' for '%s'; parse error: %s", u, kind, err.Error())
	}
	return uu
}

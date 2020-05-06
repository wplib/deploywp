package providers

import (
	"fmt"
	"github.com/wplib/deploywp/app"
	"regexp"
	"strings"
)

var _ Provider = (*Pantheon)(nil)

type Pantheon struct {
	BaseProvider
}

func NewPantheonProvider() *Pantheon {
	return &Pantheon{
		BaseProvider: *NewBaseProvider(PantheonId, BaseProvider{
			Type:    WebHostingProvider,
			Name:    PantheonName,
			Website: ParseUrl(PantheonWebsite, ""),
		}),
	}
}

//
// URL Format: ssh://codeserver.{env}.{guid}@codeserver.{env}.{guid}.drush.in:2222/~/repository.git
//
var panRe = regexp.MustCompile("^(ssh://)?(codeserver\\.)([a-z-_]+)(\\.)(.+)(@codeserver\\.)([a-z-_]+)(\\.)(.+)(\\.drush\\.in:2222/~/repository)(\\.git)?$")

func (me *Pantheon) DetectByUrl(u Url) bool {
	detected := true
	for range Once {
		if panRe.MatchString(u) {
			break
		}
		detected = false
	}
	return detected
}

func (me *Pantheon) NormalizeUrl(u Url) Url {
	m := panRe.FindStringSubmatch(u)
	for range Once {
		if m == nil {
			break
		}
		m[0] = ""       // Clear out the full match so it is not included in the join
		m[1] = "ssh://" // Replace http:// or https:// with ssh://
		m[11] = ".git"  // Add .git extension since it is missing
		u = strings.Join(m, "")
	}
	return u
}

func (me *Pantheon) ValidateHostDefaults(hg HostGetterSetter) {
	for range Once {
		if hg.GetSiteGuid() == "" {
			app.Fail("targets.defaults.site_guid not set")
		}
		if hg.GetSiteGuid() == "" {
			app.Fail("targets.defaults.site_guid not set")
		}
		if hg.GetSiteGuid() == "" {
			app.Fail("targets.defaults.site_guid not set")
		}

	}
}

func (me *Pantheon) ValidateHost(hg HostGetterSetter) {
	for range Once {
		if hg.GetId() == "" {
			app.Fail("targets.host[n].id not set")
		}
		if hg.GetName() == "" {
			app.Fail("targets.host[n].name not set")
		}
	}
}

func (me *Pantheon) initializeRepository(r *Repository) *Repository {
	if r == nil {
		r = &Repository{}
	}
	if r.Provider == nil {
		r.Provider = Dispense("pantheon")
	}
	if r.Url == "" {
		r.Url = "codeserver.dev.{..site_guid}@codeserver.dev.{..site_guid}.drush.in:2222/~/repository.git"
	}
	r.normalizeUrl()
	return r
}

func (me *Pantheon) InitializeHost(hg HostGetterSetter) {
	for range Once {
		r := me.initializeRepository(hg.GetRepository())
		hg.SetRepository(r)

		if hg.GetDomainSuffix() == "" {
			hg.SetDomainSuffix("{.site.id}.pantheonsite.io")
		}

		if hg.GetWebRoot() == "" {
			hg.SetWebRoot("/code")
		}

		hostid := hg.GetId()
		if hostid == "" {
			break
		}

		if hg.GetDomain() == "" {
			hg.SetDomain(fmt.Sprintf("%s-{domain_suffix}", hostid))
		}

		if hg.GetBranch() == "" {
			hg.SetBranch(hostid)
		}
	}
}

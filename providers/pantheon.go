package providers

import (
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

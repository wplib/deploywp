package providers

import "regexp"

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
var p1re = regexp.MustCompile("^ssh://codeserver\\.(?P<env>.+)\\.(?P<guid>.+)@codeserver\\.(.+)\\.(.+)\\.drush\\.in:2222/~/repository\\.git$")

func (me *Pantheon) DetectByUrl(u Url) (bool, Url) {
	detected := true
	for range Once {
		if p1re.MatchString(u) {
			break
		}
		detected = false
	}
	return detected, u
}

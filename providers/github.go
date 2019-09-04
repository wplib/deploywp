package providers

import (
	"fmt"
	"regexp"
)

type Github struct {
	BaseProvider
}

var _ Provider = (*Github)(nil)

func NewGitHubProvider() *Github {
	return &Github{
		BaseProvider: *NewBaseProvider(GitHubId, BaseProvider{
			Type:    VersionControlProvider,
			Name:    GitHubName,
			Website: ParseUrl(GitHubWebsite, ""),
		}),
	}
}

var gh1re = regexp.MustCompile("^https://github.com")
var gh2re = regexp.MustCompile("^git@github.com")
var gh3re = regexp.MustCompile("^github.com")

//
// See https://help.github.com/en/articles/which-remote-url-should-i-use
//
func (me *Github) DetectByUrl(u Url) (bool, Url) {
	detected := true
	for range Once {
		if gh3re.MatchString(u) {
			u = fmt.Sprintf("https://%s", u)
			break
		}
		if gh1re.MatchString(u) {
			break
		}
		if gh2re.MatchString(u) {
			break
		}
		detected = false
	}
	return detected, u
}

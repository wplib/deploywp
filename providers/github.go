package providers

import (
	"regexp"
	"strings"
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

//var forTestingNormalizeUrl1 = []string{
//"github.com:mikeschinkel/tview.git",
//"github.com:mikeschinkel/tview",
//"git@github.com:mikeschinkel/tview.git",
//"git@github.com:mikeschinkel/tview",
//"https://github.com/mikeschinkel/tview.git",
//"http://github.com/mikeschinkel/tview.git",
//"https://github.com/mikeschinkel/tview",
//"http://github.com/mikeschinkel/tview",
//}

var ghRe = regexp.MustCompile("^(git@|https?://)?(github.com)([:/])([^/]+?/.+?)(\\.git)?$")

//
// See https://help.github.com/en/articles/which-remote-url-should-i-use
//
func (me *Github) DetectByUrl(u Url) bool {
	detected := true
	for range Once {
		if ghRe.MatchString(u) {
			break
		}
		detected = false
	}
	return detected
}

//
// Currently forces everything to SSH
// @todo Add config option that allows SSH or HTTPS
//
func (me *Github) NormalizeUrl(u Url) Url {
	m := ghRe.FindStringSubmatch(u)
	for range Once {
		if m == nil {
			break
		}
		m[0] = ""     // Clear out the full match so it is not included in the join
		m[1] = "git@" // Replace http:// or https:// with git@
		m[3] = ":"    // Replace / with : for git@ syntax
		m[5] = ".git" // Add .git extension since it is missing
		u = strings.Join(m, "")
	}
	return u
}

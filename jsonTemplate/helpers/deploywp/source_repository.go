package deploywp

import (
	"github.com/wplib/deploywp/only"
	"strings"
)


type Repository struct {
	Provider string `json:"provider"`
	URL      URL `json:"url"`

	Valid bool
	Error error
}


type URL string
func (me *URL) ToString() string {
	return string(*me)
}

type String string
func (me *String) ToString() string {
	return string(*me)
}


func (me *Repository) New() Repository {

	if me == nil {
		me = &Repository{
			Provider: "",
			URL:      "",
		}
	}

	return *me
}


func (me *Repository) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			ok = true
		}
		// @TODO - perform other validity checks here.

		ok = false
	}

	return ok
}


func (me *Repository) GetProvider() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Provider
	}

	return ret
}


func (me *Repository) GetUrl() URL {
	var ret URL

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.URL
	}

	return ret
}


func (me *Repository) IsGitProvider() bool {
	var ok bool

	for range only.Once {
		if me.IsNil() {
			break
		}

		switch strings.ToLower(me.Provider) {
			case "git":
				fallthrough
			case "github":
				fallthrough
			case "gitlab":
				fallthrough
			case "gitlabs":
				ok = true
		}
	}

	return ok
}

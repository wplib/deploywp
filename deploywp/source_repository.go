package deploywp

import (
	"github.com/wplib/deploywp/ux"
	"strings"
)


type Repository struct {
	Provider string `json:"provider"`
	URL      URL `json:"url"`

	Valid bool
	State *ux.State
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
	me = &Repository{
		Provider: "",
		URL:      "",
		State: ux.NewState(false),
	}
	return *me
}


func (e *Repository) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (me *Repository) GetProvider() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.Provider
}


func (me *Repository) GetUrl() URL {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.URL
}


func (me *Repository) IsGitProvider() bool {
	var ok bool
	if state := me.IsNil(); state.IsError() {
		return false
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

	return ok
}

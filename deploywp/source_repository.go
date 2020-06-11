package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


type Repository struct {
	Provider string `json:"provider"`
	URL      URL `json:"url"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (r *Repository) New(runtime *toolRuntime.TypeRuntime) *Repository {
	runtime = runtime.EnsureNotNil()
	return &Repository{
		Provider: "",
		URL:      "",

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (r *Repository) IsNil() *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}
	r.state = r.state.EnsureNotNil()
	return r.state
}


type URL string
func (u *URL) String() string {
	return string(*u)
}

type String string
func (me *String) ToString() string {
	return string(*me)
}


func (r *Repository) GetProvider() string {
	if state := r.IsNil(); state.IsError() {
		return ""
	}
	return r.Provider
}


func (r *Repository) GetUrl() URL {
	if state := r.IsNil(); state.IsError() {
		return ""
	}
	return r.URL
}


func (r *Repository) IsGitProvider() bool {
	var ok bool
	if state := r.IsNil(); state.IsError() {
		return false
	}

	switch strings.ToLower(r.Provider) {
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

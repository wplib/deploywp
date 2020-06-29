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

func (r *Repository) IsValid() bool {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return false
	}
	for range onlyOnce {
		if r.Provider == "" {
			r.state.SetError("Empty repository.%s", GetStructTag(r, "Provider"))
			r.Valid = false
			break
		}
		if r.URL == "" {
			r.state.SetError("Empty repository.%s", GetStructTag(r, "URL"))
			r.Valid = false
			break
		}
		r.Valid = true
	}
	return r.Valid
}
func (r *Repository) IsNotValid() bool {
	return !r.IsValid()
}

func (r *Repository) GetUrlAsDir() string {
	return r.URL.GetAsDir()
}


type URL string
func (u *URL) String() string {
	return string(*u)
}

func (u *URL) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(u); state.IsError() {
		return ok
	}
	for range onlyOnce {
		if u == nil {
			ok = false
			break
		}
		if *u == "" {
			ok = false
			break
		}
		ok = true
	}
	return ok
}
func (u *URL) IsNotValid() bool {
	return !u.IsValid()
}

func (u *URL) GetAsDir() string {
	ts := strings.TrimPrefix(u.String(), "https://")
	ts = strings.TrimSuffix(ts, ".git")
	ts = strings.ReplaceAll(ts, "/", "_")
	return ts
}


type String string
func (s *String) ToString() string {
	return string(*s)
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

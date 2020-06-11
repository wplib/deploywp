package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type Source struct {
	Build      Build      `json:"build"`
	Paths      Paths      `json:"paths"`
	Repository Repository `json:"repository"`
	Revision   Revision   `json:"revision"`

	AbsPaths   Paths
	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (s *Source) New(runtime *toolRuntime.TypeRuntime) *Source {
	runtime = runtime.EnsureNotNil()
	return &Source{
		Build:      *((*Build).New(&Build{}, runtime)),
		Paths:      *((*Paths).New(&Paths{}, runtime)),
		Repository: *((*Repository).New(&Repository{}, runtime)),
		Revision:   *((*Revision).New(&Revision{}, runtime)),
		AbsPaths:   *((*Paths).New(&Paths{}, runtime)),

		Valid:   false,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (s *Source) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.state = s.state.EnsureNotNil()
	return s.state
}
func (s *Source) Process() *ux.State {
	if state := s.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		s.AbsPaths = s.Paths
		s.state = s.AbsPaths.ExpandPaths()
		if s.state.IsError() {
			break
		}
		s.Valid = true
	}

	return s.state
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (s *Source) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths
	if state := s.IsNil(); state.IsError() {
		return &Paths{}
	}

	for range onlyOnce {
		if toolTypes.ReflectBoolArg(abs) {
			ret = &s.AbsPaths
			break
		}
		ret = &s.Paths
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Repository
func (s *Source) GetRepository() *Repository {
	return &s.Repository
}
func (s *Source) GetRepositoryProvider() string {
	var ret string
	if state := s.IsNil(); state.IsError() {
		return ret
	}
	ret = s.Repository.GetProvider()
	return ret
}
func (s *Source) GetRepositoryUrl() URL {
	var ret URL
	if state := s.IsNil(); state.IsError() {
		return ret
	}
	ret = s.Repository.GetUrl()
	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revision
func (s *Source) GetRevision() *Revision {
	return &s.Revision
}
func (s *Source) GetRevisionType() string {
	var ret string
	if state := s.IsNil(); state.IsError() {
		return ret
	}
	ret = s.Revision.GetType()
	return ret
}
func (s *Source) GetRevisionName() string {
	var ret string
	if state := s.IsNil(); state.IsError() {
		return ret
	}
	ret = s.Revision.GetName()
	return ret
}
func IsValidVersionType(t string) bool {
	var ok bool
	for range onlyOnce {
		if t == "branch" {
			ok = true
			break
		}
		if t == "tag" {
			ok = true
			break
		}
	}
	return ok
}


// ////////////////////////////////////////////////////////////////////////////////
// Build
func (s *Source) GetBuild() bool {
	var ret bool
	if state := s.IsNil(); state.IsError() {
		return false
	}

	for range onlyOnce {
		ret = s.Build.GetBuild()
	}

	return ret
}

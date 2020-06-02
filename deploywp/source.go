package deploywp

import (
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type Source struct {
	Build      Build      `json:"build"`
	Paths      Paths      `json:"paths"`
	Repository Repository `json:"repository"`
	Revision   Revision   `json:"revision"`

	AbsPaths   Paths
	Valid bool
	State *ux.State
}


func (me *Source) New() Source {
	me.Build.New()
	me.Paths.New()
	me.Repository.New()
	me.Revision.New()

	me.AbsPaths.New()
	me.State = ux.NewState(false)

	return *me
}

func (me *Source) Process() *ux.State {
	if state := me.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		me.AbsPaths = me.Paths
		me.State = me.AbsPaths.ExpandPaths()
		if me.State.IsError() {
			break
		}

		me.Valid = true
	}

	return me.State
}

func (e *Source) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Source) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths
	if state := me.IsNil(); state.IsError() {
		return &Paths{}
	}

	for range OnlyOnce {
		if helperTypes.ReflectBoolArg(abs) {
			ret = &me.AbsPaths
			break
		}

		ret = &me.Paths
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Repository
func (me *Source) GetRepository() *Repository {
	return &me.Repository
}
func (me *Source) GetRepositoryProvider() string {
	var ret string
	if state := me.IsNil(); state.IsError() {
		return ""
	}

	for range OnlyOnce {
		ret = me.Repository.GetProvider()
	}

	return ret
}
func (me *Source) GetRepositoryUrl() URL {
	var ret URL
	if state := me.IsNil(); state.IsError() {
		return ""
	}

	for range OnlyOnce {
		ret = me.Repository.GetUrl()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revision
func (me *Source) GetRevision() *Revision {
	return &me.Revision
}
func (me *Source) GetRevisionType() string {
	var ret string
	if state := me.IsNil(); state.IsError() {
		return ""
	}

	for range OnlyOnce {
		ret = me.Revision.GetType()
	}

	return ret
}
func (me *Source) GetRevisionName() string {
	var ret string
	if state := me.IsNil(); state.IsError() {
		return ""
	}

	for range OnlyOnce {
		ret = me.Revision.GetName()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Build
func (me *Source) GetBuild() bool {
	var ret bool
	if state := me.IsNil(); state.IsError() {
		return false
	}

	for range OnlyOnce {
		ret = me.Build.GetBuild()
	}

	return ret
}

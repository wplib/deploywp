package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
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
	me.State = ux.NewState()

	return *me
}

func (me *Source) Process() *ux.State {
	for range only.Once {
		if me.IsNil() {
			break
		}

		me.AbsPaths = me.Paths
		me.State = me.AbsPaths.ExpandPaths()
		if me.State.IsError() {
			break
		}

		me.Valid = true
	}

	return me.State
}

func (me *Source) IsNil() bool {
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


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Source) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

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

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Repository.GetProvider()
	}

	return ret
}
func (me *Source) GetRepositoryUrl() URL {
	var ret URL

	for range only.Once {
		if me.IsNil() {
			break
		}

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

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revision.GetType()
	}

	return ret
}
func (me *Source) GetRevisionName() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revision.GetName()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Build
func (me *Source) GetBuild() bool {
	var ret bool

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Build.GetBuild()
	}

	return ret
}

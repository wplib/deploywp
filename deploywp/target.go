package deploywp

import (
	"github.com/jinzhu/copier"
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type Target struct {
	Files     Files           `json:"files"`
	Paths     Paths           `json:"paths"`
	Providers Providers       `json:"providers"`
	Revisions TargetRevisions `json:"revisions"`

	AbsPaths  Paths
	AbsFiles  Files

	Valid bool
	State *ux.State
}

func (me *Target) New() Target {
	me.Files.New()
	me.Paths.New()
	me.Providers.New()
	me.Revisions.New()

	me.AbsPaths.New()
	me.AbsFiles.New()

	me.State = ux.NewState(false)

	return *me
}

func (me *Target) Process() *ux.State {
	if state := me.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		err := copier.Copy(&me.AbsPaths, &me.Paths)
		me.State.SetError(err)
		if me.State.IsError() {
			break
		}

		me.State = me.AbsPaths.ExpandPaths()
		me.State.SetError(err)
		if me.State.IsError() {
			break
		}

		me.AbsFiles.Copy = append(me.AbsFiles.Copy, me.Files.Copy...)
		me.AbsFiles.Delete = append(me.AbsFiles.Delete, me.Files.Delete...)
		me.AbsFiles.Exclude = append(me.AbsFiles.Exclude, me.Files.Exclude...)
		me.AbsFiles.Keep = append(me.AbsFiles.Keep, me.Files.Keep...)

		me.State = me.AbsFiles.Process(me.AbsPaths)
		if me.State.IsError() {
			break
		}

		me.State = me.Files.Process(me.Paths)
		if me.State.IsError() {
			break
		}

		me.Valid = true
	}

	return me.State
}

func (e *Target) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


// ////////////////////////////////////////////////////////////////////////////////
// Files
func (me *Target) GetFiles(ftype interface{}) *FilesArray {
	var ret *FilesArray
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}

	for range OnlyOnce {
		ret = me.Files.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Target) GetPaths(abs ...interface{}) *Paths {
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
// Providers
func (me *Target) GetProvider(provider interface{}) *Provider {
	var ret *Provider
	if state := me.IsNil(); state.IsError() {
		return &Provider{}
	}

	for range OnlyOnce {
		ret = me.Providers.GetProvider(provider)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revisions
func (me *Target) GetRevision(host interface{}) *TargetRevision {
	var ret *TargetRevision
	if state := me.IsNil(); state.IsError() {
		return &TargetRevision{}
	}

	for range OnlyOnce {
		ret = me.Revisions.GetRevision(host)
	}

	return ret
}

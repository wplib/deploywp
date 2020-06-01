package deploywp

import (
	"github.com/newclarity/JsonToConfig/ux"
)


type Build struct {
	Empty bool

	Valid bool
	State *ux.State
}

//var _ deploywp.BuildGetter = (*Build)(nil)

func (me *Build) New() Build {
	me = &Build {
		Empty: false,
		State: ux.NewState(false),
	}
	return *me
}


func (e *Build) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (me *Build) GetBuild() bool {
	if state := me.IsNil(); state.IsError() {
		return false
	}
	return me.Empty
}

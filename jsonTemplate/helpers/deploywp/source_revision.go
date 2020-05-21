package deploywp

import (
	"github.com/wplib/deploywp/ux"
)


type Revision struct {
	RefName string `json:"ref_name" mapstructure:"ref_name"`
	RefType string `json:"ref_type" mapstructure:"ref_type"`

	Valid bool
	State *ux.State
}


func (me *Revision) New() Revision {
	me = &Revision {
		RefName: "",
		RefType: "",
		State: ux.NewState(false),
	}
	return *me
}

func (e *Revision) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (me *Revision) GetType() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.RefType
}

func (me *Revision) GetName() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.RefName
}

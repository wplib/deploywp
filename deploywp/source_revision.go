package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Revision struct {
	RefName string `json:"ref_name" mapstructure:"ref_name"`
	RefType string `json:"ref_type" mapstructure:"ref_type"`

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (r *Revision) New(runtime *toolRuntime.TypeRuntime) *Revision {
	runtime = runtime.EnsureNotNil()
	return &Revision {
		RefName: "",
		RefType: "",

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (r *Revision) IsNil() *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}
	r.state = r.state.EnsureNotNil()
	return r.state
}

func (r *Revision) IsValid() bool {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return false
	}
	for range onlyOnce {
		if r.RefName == "" {
			r.state.SetError("Empty revision.%s", GetStructTag(r, "RefName"))
			r.Valid = false
			break
		}
		if r.RefType == "" {
			r.state.SetError("Empty revision.%s", GetStructTag(r, "RefType"))
			r.Valid = false
			break
		}
		r.Valid = true
	}
	return r.Valid
}
func (r *Revision) IsNotValid() bool {
	return !r.IsValid()
}


func (r *Revision) GetType() string {
	if state := r.IsNil(); state.IsError() {
		return ""
	}
	return r.RefType
}

func (r *Revision) GetName() string {
	if state := r.IsNil(); state.IsError() {
		return ""
	}
	return r.RefName
}

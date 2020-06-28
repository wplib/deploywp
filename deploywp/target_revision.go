package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type TargetRevision struct {
	AutoDeploy bool   `json:"auto_deploy" mapstructure:"auto_deploy"`
	HostName   string `json:"host_name" mapstructure:"host_name"`
	RefName    string `json:"ref_name" mapstructure:"ref_name"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (tr *TargetRevision) New(runtime *toolRuntime.TypeRuntime) *TargetRevision {
	runtime = runtime.EnsureNotNil()
	return &TargetRevision {
		HostName:     "",
		RefName:     "",

		Valid: false,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (tr *TargetRevision) IsNil() *ux.State {
	if state := ux.IfNilReturnError(tr); state.IsError() {
		return state
	}
	tr.state = tr.state.EnsureNotNil()
	return tr.state
}

func (tr *TargetRevision) IsValid() bool {
	if state := ux.IfNilReturnError(tr); state.IsError() {
		return false
	}
	for range onlyOnce {
		if tr.HostName == "" {
			tr.state.SetError("Empty target.revision.%s", GetStructTag(tr, "HostName"))
			tr.Valid = false
			break
		}
		if tr.RefName == "" {
			tr.state.SetError("Empty target.revision.%s", GetStructTag(tr, "RefName"))
			tr.Valid = false
			break
		}
		tr.Valid = true
	}
	return tr.Valid
}
func (tr *TargetRevision) IsNotValid() bool {
	return !tr.IsValid()
}

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


type TargetRevisions []TargetRevision
func (tr *TargetRevisions) New() *TargetRevisions {
	if tr == nil {
		return &TargetRevisions{}
	}
	return tr
}
func (tr *TargetRevisions) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		for i, _ := range *tr {
			//	(*tr)[i] = *((*tr)[i].New(runtime))
			(*tr)[i].Valid = true
			(*tr)[i].runtime = runtime
			(*tr)[i].state = ux.NewState(runtime.CmdName, runtime.Debug)
		}
	}
	return state
}


func (tr *TargetRevisions) GetByHost(host string) *TargetRevision {
	ret := (*TargetRevision).New(&TargetRevision{}, nil)

	for range onlyOnce {
		if host == "" {
			ret.state.SetError("GetRevision hostname not a string")
			break
		}

		var ok bool
		for _, v := range *tr {
			if v.HostName == host {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetRevision hostname not found")
			break
		}
	}

	return ret
}


func (tr *TargetRevisions) GetByRefName(ref string) *TargetRevision {
	ret := (*TargetRevision).New(&TargetRevision{}, nil)

	for range onlyOnce {
		if ref == "" {
			ret.state.SetError("GetRevision hostname not a string")
			break
		}

		var ok bool
		for _, v := range *tr {
			if v.HostName == ref {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetRevision hostname not found")
			break
		}
	}

	return ret
}

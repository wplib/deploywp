package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


//type TargetRevisions struct{
//	revisionsArray
//
//	Valid   bool
//	runtime *toolRuntime.TypeRuntime
//	state   *ux.State
//}
//type revisionsArray []TargetRevision
type TargetRevisions []TargetRevision


func (tr *TargetRevisions) New(runtime *toolRuntime.TypeRuntime) *TargetRevisions {
	runtime = runtime.EnsureNotNil()
	return &TargetRevisions {
		//revisionsArray: revisionsArray{},
		//
		//Valid: false,
		//runtime: runtime,
		//state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (tr *TargetRevisions) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(tr); state.IsError() {
		return ok
	}
	for range onlyOnce {
		ok = true
		for _, t := range *tr {
			if t.IsNotValid() {
				ok = false
				break
			}
		}
	}
	return ok
}
func (tr *TargetRevisions) IsNotValid() bool {
	return !tr.IsValid()
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

package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


//type Targets struct{
//	revisionsArray
//
//	Valid   bool
//	runtime *toolRuntime.TypeRuntime
//	state   *ux.State
//}
//type revisionsArray []Target
type Targets []Target


func (tr *Targets) New(runtime *toolRuntime.TypeRuntime) *Targets {
	runtime = runtime.EnsureNotNil()
	return &Targets{
		//revisionsArray: revisionsArray{},
		//
		//Valid: false,
		//runtime: runtime,
		//state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (tr *Targets) IsValid() bool {
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
func (tr *Targets) IsNotValid() bool {
	return !tr.IsValid()
}

func (tr *Targets) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		for i := range *tr {
			//	(*tr)[i] = *((*tr)[i].New(runtime))
			(*tr)[i].Valid = true
			(*tr)[i].runtime = runtime
			(*tr)[i].state = ux.NewState(runtime.CmdName, runtime.Debug)
		}
	}
	return state
}


func (tr *Targets) GetByHost(host string) *Target {
	ret := (*Target).New(&Target{}, nil)

	for range onlyOnce {
		if host == "" {
			ret.state.SetError("GetRevision hostname not valid")
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


func (tr *Targets) GetByRefName(ref string) *Target {
	ret := (*Target).New(&Target{}, nil)

	for range onlyOnce {
		if ref == "" {
			ret.state.SetError("GetRevision RefName not valid")
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
			ret.state.SetError("GetRevision RefName not found")
			break
		}
	}

	return ret
}

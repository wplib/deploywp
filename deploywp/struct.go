package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


const onlyOnce = "1"

type DeployWpGetter interface {
}


type TypeDeployWp struct {
	Hosts       Hosts       `json:"hosts"`	// mapstructure:",squash"`
	Source      Source      `json:"source"`
	Destination Destination `json:"destination"`

	Print   UxPrint
	Valid   bool
	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
}

func (dwp *TypeDeployWp) New(runtime *toolRuntime.TypeRuntime) *TypeDeployWp {
	runtime = runtime.EnsureNotNil()
	return &TypeDeployWp{
		Hosts:       *((*Hosts).New(&Hosts{}, runtime)),
		Source:      *((*Source).New(&Source{}, runtime)),
		Destination: *((*Destination).New(&Destination{}, runtime)),

		Valid:   false,
		Runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (dwp *TypeDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(dwp); state.IsError() {
		return state
	}
	dwp.State = dwp.State.EnsureNotNil()
	return dwp.State
}

func (dwp *TypeDeployWp) IsValid() bool {
	if state := ux.IfNilReturnError(dwp); state.IsError() {
		return false
	}
	for range onlyOnce {
		//if dwp.Hosts.IsNotValid() {
		//	dwp.State = dwp.Hosts.state
		//	dwp.Valid = false
		//	break
		//}
		if dwp.Source.IsNotValid() {
			dwp.State = dwp.Source.state
			dwp.Valid = false
			break
		}
		if dwp.Destination.IsNotValid() {
			dwp.State = dwp.Destination.state
			dwp.Valid = false
			break
		}
		dwp.Valid = true
	}
	return dwp.Valid
}
func (dwp *TypeDeployWp) IsNotValid() bool {
	return !dwp.IsValid()
}


func ReflectDeployWp(ref interface{}) *TypeDeployWp {
	return ref.(*TypeDeployWp)
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}

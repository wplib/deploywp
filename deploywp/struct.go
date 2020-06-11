package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


const onlyOnce = "1"

type DeployWpGetter interface {
}


type TypeDeployWp struct {
	Hosts  Hosts  `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`

	Valid   bool
	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (dwp *TypeDeployWp) New(runtime *toolRuntime.TypeRuntime) *TypeDeployWp {
	runtime = runtime.EnsureNotNil()
	return &TypeDeployWp{
		Hosts:  *((*Hosts).New(&Hosts{})),
		Source: *((*Source).New(&Source{}, runtime)),
		Target: *((*Target).New(&Target{}, runtime)),

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


func ReflectDeployWp(ref interface{}) *TypeDeployWp {
	return ref.(*TypeDeployWp)
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}

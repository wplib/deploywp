package deploywp

import (
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/ux"
)

const OnlyOnce = "1"

type DeployWpGetter interface {
}


type TypeDeployWp struct {
	Hosts  Hosts  `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`

	Runtime

	Valid  bool
	State  *ux.State
}

type Runtime struct {
	Exec *runtime.Exec
}


func ReflectDeployWp(ref interface{}) *TypeDeployWp {
	return ref.(*TypeDeployWp)
}


func (e *TypeDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func NewJsonFile() *TypeDeployWp {
	var jf TypeDeployWp

	jf.State = ux.NewState(false)
	jf.Runtime = Runtime{}

	jf.Hosts.New()
	jf.Source.New()
	jf.Target.New()

	return &jf
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}

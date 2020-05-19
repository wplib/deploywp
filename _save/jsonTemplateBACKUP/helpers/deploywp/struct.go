package deploywp

import "github.com/wplib/deploywp/ux"


type DeployWpGetter interface {
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}

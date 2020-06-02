package deploywp

import "github.com/newclarity/scribeHelpers/ux"


type DeployWpGetter interface {
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}

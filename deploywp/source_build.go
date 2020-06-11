package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Build struct {
	Empty bool

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (b *Build) New(runtime *toolRuntime.TypeRuntime) *Build {
	runtime = runtime.EnsureNotNil()
	return &Build {
		Empty: false,

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (b *Build) IsNil() *ux.State {
	if state := ux.IfNilReturnError(b); state.IsError() {
		return state
	}
	b.state = b.state.EnsureNotNil()
	return b.state
}


func (b *Build) GetBuild() bool {
	if state := b.IsNil(); state.IsError() {
		return false
	}
	return b.Empty
}

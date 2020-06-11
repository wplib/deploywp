package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Host struct {
	HostName string `json:"host_name" mapstructure:"host_name"`
	Label    string `json:"label"`
	Provider string `json:"provider"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (h *Host) New(runtime *toolRuntime.TypeRuntime) *Host {
	runtime = runtime.EnsureNotNil()
	return &Host{
		HostName: "",
		Label:    "",
		Provider: "",

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (h *Host) IsNil() *ux.State {
	if state := ux.IfNilReturnError(h); state.IsError() {
		return state
	}
	h.state = h.state.EnsureNotNil()
	return h.state
}

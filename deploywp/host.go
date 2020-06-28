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

func (h *Host) IsValid() bool {
	if state := ux.IfNilReturnError(h); state.IsError() {
		return false
	}
	for range onlyOnce {
		if h.HostName == "" {
			h.state.SetError("Empty .host.%s", GetStructTag(h, "HostName"))
			h.Valid = false
			break
		}
		if h.Label == "" {
			h.state.SetError("Empty .host.%s", GetStructTag(h, "Label"))
			h.Valid = false
			break
		}
		if h.Provider == "" {
			h.state.SetError("Empty .host.%s", GetStructTag(h, "Provider"))
			h.Valid = false
			break
		}
		h.Valid = true
	}
	return h.Valid
}
func (h *Host) IsNotValid() bool {
	return !h.IsValid()
}

//func (h *Host) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
//	state := ux.NewState(runtime.CmdName, runtime.Debug)
//	for range onlyOnce {
//		h.runtime = runtime
//		state =
//	}
//	return state
//}

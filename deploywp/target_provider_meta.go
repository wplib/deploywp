package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Meta struct {
	SiteID   string `json:"site_id" mapstructure:"site_id"`
	SiteName string `json:"site_name" mapstructure:"site_name"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (m *Meta) New(runtime *toolRuntime.TypeRuntime) *Meta {
	runtime = runtime.EnsureNotNil()
	return &Meta{
		SiteID:   "",
		SiteName: "",

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (m *Meta) IsNil() *ux.State {
	if state := ux.IfNilReturnError(m); state.IsError() {
		return state
	}
	m.state = m.state.EnsureNotNil()
	return m.state
}

func (m *Meta) IsValid() bool {
	if state := ux.IfNilReturnError(m); state.IsError() {
		return false
	}
	for range onlyOnce {
		if m.SiteID == "" {
			m.state.SetError("Empty target.provider.meta.%s", GetStructTag(m, "SiteID"))
			m.Valid = false
			break
		}
		if m.SiteName == "" {
			m.state.SetError("Empty target.provider.meta.%s", GetStructTag(m, "SiteName"))
			m.Valid = false
			break
		}
		m.Valid = true
	}
	return m.Valid
}
func (m *Meta) IsNotValid() bool {
	return !m.IsValid()
}

func (m *Meta) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	runtime = runtime.EnsureNotNil()
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		m.Valid = true
		m.runtime = runtime
		m.state = ux.NewState(runtime.CmdName, runtime.Debug)
	}
	return state
}

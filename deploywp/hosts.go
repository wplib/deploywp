package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Hosts []Host
func (h *Hosts) New() *Hosts {
	if h == nil {
		return &Hosts{}
	}
	return h
}
func (h *Hosts) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	runtime = runtime.EnsureNotNil()
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		for i, _ := range *h {
			(*h)[i].Valid = true
			(*h)[i].runtime = runtime
			(*h)[i].state = ux.NewState(runtime.CmdName, runtime.Debug)
		}
	}
	return state
}


func (h *Hosts) Count() int {
	return len(*h)
}


func (h *Hosts) GetByName(host string) *Host {
	ret := (*Host).New(&Host{}, nil)

	for range onlyOnce {
		if host == "" {
			ret.state.SetError("GetHostByName - hostname empty")
			break
		}

		var ok bool
		for _, v := range *h {
			if v.HostName == host {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetHostByName - hostname not found")
			break
		}
	}

	return ret
}


func (h *Hosts) GetByProvider(provider string) *Host {
	ret := (*Host).New(&Host{}, nil)

	for range onlyOnce {
		if provider == "" {
			ret.state.SetError("GetHostByProvider - provider empty")
			break
		}

		var ok bool
		for _, v := range *h {
			if v.Provider == provider {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetHostByProvider - provider not found")
			break
		}
	}

	return ret
}

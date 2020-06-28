package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


//type Hosts struct {
//	hostsArray
//
//	Valid   bool
//	runtime *toolRuntime.TypeRuntime
//	state   *ux.State
//}
//type hostsArray []Host

type Hosts []Host

func (h *Hosts) New(runtime *toolRuntime.TypeRuntime) *Hosts {
	runtime = runtime.EnsureNotNil()
	return &Hosts{
		//hostsArray: []Host{},

		//Valid:   false,
		//runtime: runtime,
		//state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (h *Hosts) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(h); state.IsError() {
		return ok
	}
	for range onlyOnce {
		ok = true
		//for _, f := range h.hostsArray {
		//	if f.IsNotValid() {
		//		ok = false
		//		break
		//	}
		//}
		for _, f := range *h {
			if f.IsNotValid() {
				ok = false
				break
			}
		}
	}
	return ok
}
func (h *Hosts) IsNotValid() bool {
	return !h.IsValid()
}

func (h *Hosts) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	runtime = runtime.EnsureNotNil()
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		//for i, _ := range h.hostsArray {
		//	h.hostsArray[i].Valid = true
		//	h.hostsArray[i].runtime = runtime
		//	h.hostsArray[i].state = ux.NewState(runtime.CmdName, runtime.Debug)
		//}
		for i, _ := range *h {
			(*h)[i].Valid = true
			(*h)[i].runtime = runtime
			(*h)[i].state = ux.NewState(runtime.CmdName, runtime.Debug)
		}
	}
	return state
}


func (h *Hosts) Count() int {
	//return len(h.hostsArray)
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
		//for _, v := range h.hostsArray {
		//	if v.HostName == host {
		//		ret = &v
		//		ok = true
		//		break
		//	}
		//}
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
		//for _, v := range h.hostsArray {
		//	if v.Provider == provider {
		//		ret = &v
		//		ok = true
		//		break
		//	}
		//}
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


//func (h *Hosts) GetByName(host string) *Host {
//	ret := (*Host).New(&Host{}, nil)
//	return ret
//}
//func (h *Hosts) GetByProvider(provider string) *Host {
//	ret := (*Host).New(&Host{}, nil)
//	return ret
//}

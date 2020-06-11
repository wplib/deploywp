package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Providers []Provider
func (ps *Providers) New() *Providers {
	if ps == nil {
		return &Providers{}
	}
	return ps
}
func (ps *Providers) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	runtime = runtime.EnsureNotNil()
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		for i, _ := range *ps {
			(*ps)[i].Process(runtime)
		}
	}
	return state
}


func (ps *Providers) GetByName(provider string) *Provider {
	ret := (*Provider).New(&Provider{}, nil)

	for range onlyOnce {
		if provider == "" {
			ret.state.SetError("GetProvider name empty")
			break
		}

		var ok bool
		for _, v := range *ps {
			if v.Name == provider {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetProvider name not found")
			break
		}
	}

	return ret
}


func (ps *Providers) GetBySiteId(siteId string) *Provider {
	ret := (*Provider).New(&Provider{}, nil)

	for range onlyOnce {
		if siteId == "" {
			ret.state.SetError("GetProvider siteId empty")
			break
		}

		var ok bool
		for _, v := range *ps {
			if v.Meta.SiteID == siteId {
				ret = &v
				ok = true
				break
			}
		}

		if !ok {
			ret.state.SetError("GetProvider siteId not found")
			break
		}
	}

	return ret
}

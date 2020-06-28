package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


//type Providers struct{
//	providersArray
//
//	Valid   bool
//	runtime *toolRuntime.TypeRuntime
//	state   *ux.State
//}
//type providersArray []Provider
type Providers []Provider


func (ps *Providers) New(runtime *toolRuntime.TypeRuntime) *Providers {
	runtime = runtime.EnsureNotNil()
	return &Providers {
		//providersArray: providersArray{},
		//
		//Valid: false,
		//runtime: runtime,
		//state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (ps *Providers) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(ps); state.IsError() {
		return ok
	}
	for range onlyOnce {
		ok = true
		for _, f := range *ps {
			if f.IsNotValid() {
				ok = false
				break
			}
		}
	}
	return ok
}
func (ps *Providers) IsNotValid() bool {
	return !ps.IsValid()
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

package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
)


type Provider struct {
	Name     string   `json:"name"`
	Meta     Meta     `json:"meta"`
	Defaults Defaults `json:"defaults"`

	Valid bool
	State *ux.State
}
type Providers []Provider

func (me *Provider) New() Provider {
	if me == nil {
		me = &Provider {
			Name:     "",
			Meta:     me.Meta.New(),
			Defaults: me.Defaults.New(),
			State: ux.NewState(false),
		}
	}

	return *me
}

func (me *Providers) New() Providers {
	if me == nil {
		me = &Providers {}
	}

	return *me
}


//func (e *Providers) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(e); state.IsError() {
//		return state
//	}
//	e.State = e.State.EnsureNotNil()
//	return e.State
//}


type Meta struct {
	SiteID   string `json:"site_id"`
	SiteName string `json:"site_name"`

	State    *ux.State
}
func (me *Meta) New() Meta {
	if me == nil {
		me = &Meta{
			SiteID:   "",
			SiteName: "",
		}
	}

	return *me
}


func (e *Meta) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


type Defaults struct {
	Paths DefaultsPaths `json:"paths"`
	Repository DefaultsRepository `json:"repository"`
}
type DefaultsPaths struct {
	WebrootDir string `json:"webroot_dir"`
}
type DefaultsRepository struct {
	URL string `json:"url"`
}
func (me *Defaults) New() Defaults {
	if me == nil {
		me = &Defaults{
			Paths:      DefaultsPaths{},
			Repository: DefaultsRepository{},
		}
	}

	return *me
}


func (me *Providers) GetProvider(provider interface{}) *Provider {
	var ret Provider

	for range OnlyOnce {
		value := helperTypes.ReflectString(provider)
		if value == nil {
			ret.State.SetError("GetProvider arg not a string")
			break
		}

		for _, v := range *me {
			if v.Name == *value {
				ret = v
				ret.Valid = true
				break
			}
		}

		if !ret.Valid {
			ret.State.SetError("GetProvider hostname not found")
			break
		}
	}

	return &ret
}

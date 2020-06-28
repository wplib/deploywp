package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type Defaults struct {
	Paths      DefaultsPaths `json:"paths"`
	Repository DefaultsRepository `json:"repository"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (d *Defaults) New(runtime *toolRuntime.TypeRuntime) *Defaults {
	return &Defaults{
		Paths:      DefaultsPaths{},
		Repository: DefaultsRepository{},

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (d *Defaults) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(d); state.IsError() {
		return ok
	}
	for range onlyOnce {
		if d.Paths.IsNotValid() {
			d.state.SetError("Empty target.provider.defaults.%s", GetStructTag(d, "Paths"))
			ok = false
			break
		}
		if d.Repository.IsNotValid() {
			d.state.SetError("Empty target.provider.defaults.%s", GetStructTag(d, "Repository"))
			ok = false
			break
		}
		ok = true
	}
	return ok
}
func (d *Defaults) IsNotValid() bool {
	return !d.IsValid()
}


type DefaultsPaths struct {
	WebrootDir string `json:"webroot_dir" mapstructure:"webroot_dir"`
}

func (d *DefaultsPaths) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(d); state.IsError() {
		return ok
	}
	for range onlyOnce {
		if d.WebrootDir == "" {
			ok = false
			break
		}
		ok = true
	}
	return ok
}
func (d *DefaultsPaths) IsNotValid() bool {
	return !d.IsValid()
}


type DefaultsRepository struct {
	URL string `json:"url"`
}

func (d *DefaultsRepository) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(d); state.IsError() {
		return ok
	}
	for range onlyOnce {
		if d.URL == "" {
			ok = false
			break
		}
		ok = true
	}
	return ok
}
func (d *DefaultsRepository) IsNotValid() bool {
	return !d.IsValid()
}

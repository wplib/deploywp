package deploywp

import (
	"github.com/jinzhu/copier"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


const DefaultDestinationBasePath = "/tmp/deploywp"

type Destination struct {
	Files     Files     `json:"files"`
	Paths     Paths     `json:"paths"`
	Providers Providers `json:"providers"` // mapstructure:",squash"`
	Targets   Targets   `json:"targets"`   // mapstructure:",squash"`

	AbsPaths  Paths
	AbsFiles  Files

	selectedHost  *Host			// Should be set to the selected host_name from arg[0]
	selectedHostName  string	// Should be set to the selected host_name from arg[0]

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (d *Destination) New(runtime *toolRuntime.TypeRuntime) *Destination {
	runtime = runtime.EnsureNotNil()
	return &Destination{
		Files:     *((*Files).New(&Files{}, runtime)),
		Paths:     *((*Paths).New(&Paths{}, runtime)),
		Providers: *((*Providers).New(&Providers{}, runtime)),
		Targets:   *((*Targets).New(&Targets{}, runtime)),
		AbsPaths:  *((*Paths).New(&Paths{}, runtime)),
		AbsFiles:  *((*Files).New(&Files{}, runtime)),

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (d *Destination) IsNil() *ux.State {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return state
	}
	d.state = d.state.EnsureNotNil()
	return d.state
}

func (d *Destination) IsValid() bool {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return false
	}
	for range onlyOnce {
		if d.Files.IsNotValid() {
			d.state = d.Files.state
			d.Valid = false
			break
		}
		if d.Paths.IsNotValid() {
			d.state = d.Paths.state
			d.Valid = false
			break
		}
		if d.Providers.IsNotValid() {
			//d.state = d.Providers.state
			d.Valid = false
			break
		}
		if d.Targets.IsNotValid() {
			//d.state = d.Revisions.state
			d.Valid = false
			break
		}

		if d.AbsPaths.IsNotValid() {
			d.state = d.AbsPaths.state
			d.Valid = false
			break
		}
		if d.AbsFiles.IsNotValid() {
			d.state = d.AbsFiles.state
			d.Valid = false
			break
		}

		d.Valid = true
	}
	return d.Valid
}
func (d *Destination) IsNotValid() bool {
	return !d.IsValid()
}

func (d *Destination) Process() *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		err := copier.Copy(&d.AbsPaths, &d.Paths)
		d.state.SetError(err)
		if d.state.IsError() {
			break
		}

		d.state = d.Providers.Process(d.runtime)
		if d.state.IsError() {
			break
		}

		d.state = d.Targets.Process(d.runtime)
		if d.state.IsError() {
			break
		}

		if d.Paths.BasePath == "" {
			d.Paths.BasePath = DefaultDestinationBasePath
		}
		//d.state = d.Paths.GetBasePath(d.Repository.GetUrlAsDir())
		selectedProvider := d.Providers.GetByName(d.selectedHost.Provider)
		d.state = d.Paths.AppendBasePath(d.selectedHost.HostName, selectedProvider.Meta.SiteName)
		if d.state.IsError() {
			break
		}

		d.AbsPaths.BasePath = d.Paths.BasePath
		d.state = d.AbsPaths.ExpandPaths()
		d.state.SetError(err)
		if d.state.IsError() {
			break
		}

		d.AbsFiles.Copy.Append(&d.Files.Copy)
		d.AbsFiles.Delete.Append(&d.Files.Delete)
		d.AbsFiles.Exclude.Append(&d.Files.Exclude)
		d.AbsFiles.Keep.Append(&d.Files.Keep)

		d.state = d.AbsFiles.Process(d.AbsPaths)
		if d.state.IsError() {
			break
		}

		d.state = d.Files.Process(d.Paths)
		if d.state.IsError() {
			break
		}

		d.Valid = true
	}

	return d.state
}


// ////////////////////////////////////////////////////////////////////////////////
// Files
func (d *Destination) GetFiles(ftype string) *FilesArray {
	var ret *FilesArray
	if state := d.IsNil(); state.IsError() {
		return &FilesArray{}
	}

	for range onlyOnce {
		ret = d.Files.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (d *Destination) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths
	if state := d.IsNil(); state.IsError() {
		return &Paths{}
	}

	for range onlyOnce {
		if toolTypes.ReflectBoolArg(abs) {
			ret = &d.AbsPaths
			break
		}

		ret = &d.Paths
	}

	return ret
}


func (d *Destination) GetBasePath() string {
	if state := d.IsNil(); state.IsError() {
		return ""
	}
	return d.Paths.BasePath
}


// ////////////////////////////////////////////////////////////////////////////////
// Providers
func (d *Destination) GetProviderByName(provider string) *Provider {
	var ret *Provider
	if state := d.IsNil(); state.IsError() {
		return ret
	}
	return d.Providers.GetByName(provider)
}
func (d *Destination) GetProviderBySiteId(siteId string) *Provider {
	var ret *Provider
	if state := d.IsNil(); state.IsError() {
		return ret
	}
	ret = d.Providers.GetBySiteId(siteId)
	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revisions
func (d *Destination) GetTargetByHost(host string) *Target {
	var ret *Target
	if state := d.IsNil(); state.IsError() {
		return &Target{state: state}
	}
	ret = d.Targets.GetByHost(host)
	return ret
}

func (d *Destination) GetTargetByName(ref string) *Target {
	var ret *Target
	if state := d.IsNil(); state.IsError() {
		return &Target{state: state}
	}
	ret = d.Targets.GetByRefName(ref)
	return ret
}

//func (d *Destination) SelectHost(host string) *ux.State {
//	if state := d.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		d.selectedHost = d.GetTargetByHost(host)
//		d.selectedHostName = host
//	}
//
//	return d.state
//}


func (d *Destination) GetSelectedHost() *Host {
	if state := d.IsNil(); state.IsError() {
		return &Host{state: state}
	}
	return d.selectedHost
}


func (d *Destination) SetDestinationHost(host *Host) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if host == nil {
			d.state.SetError("Selected host is nil.")
			break
		}

		d.selectedHost = host
		d.selectedHostName = host.HostName

		d.state.SetOk()
	}

	return d.state
}

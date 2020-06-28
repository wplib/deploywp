package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"path/filepath"
	"strings"
)


type Provider struct {
	Name     string   `json:"name"`
	Meta     Meta     `json:"meta"`
	Defaults Defaults `json:"defaults"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (p *Provider) New(runtime *toolRuntime.TypeRuntime) *Provider {
	runtime = runtime.EnsureNotNil()
	return &Provider {
		// *((*Defaults).New(&Defaults{}, runtime))
		Name:     "",
		Meta:     *((*Meta).New(&Meta{}, runtime)),
		Defaults: *((*Defaults).New(&Defaults{}, runtime)),

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (p *Provider) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.state = p.state.EnsureNotNil()
	return p.state
}

func (p *Provider) IsValid() bool {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return false
	}
	for range onlyOnce {
		if p.Name == "" {
			p.state.SetError("Empty target.provider.%s", GetStructTag(p, "Name"))
			p.Valid = false
			break
		}
		if p.Meta.IsNotValid() {
			p.state = p.Meta.state
			p.Valid = false
			break
		}
		if p.Defaults.IsNotValid() {
			p.state = p.Defaults.state
			p.Valid = false
			break
		}
		p.Valid = true
	}
	return p.Valid
}
func (p *Provider) IsNotValid() bool {
	return !p.IsValid()
}

func (p *Provider) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	runtime = runtime.EnsureNotNil()
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		p.Valid = true
		p.runtime = runtime
		p.state = ux.NewState(runtime.CmdName, runtime.Debug)
		p.Meta.Process(runtime)
	}
	return state
}


func (p *Provider) GetRepository() string {
	var url string
	if state := p.IsNil(); state.IsError() {
		return url
	}

	for range onlyOnce {
		p.state = p.IsNil()
		if p.state.IsError() {
			break
		}

		url = strings.ReplaceAll(p.Defaults.Repository.URL, "{site_id}", p.Meta.SiteID)
		if url == "" {
			p.state.SetError(".target.providers.defaults.repository is nil")
			break
		}
		p.state.SetOk()
	}

	return url
}


func (p *Provider) GetWebroot(addPrefix ...string) string {
	var path string
	if state := p.IsNil(); state.IsError() {
		return path
	}

	for range onlyOnce {
		p.state = p.IsNil()
		if p.state.IsError() {
			break
		}

		if p.Defaults.Paths.WebrootDir == "" {
			p.state.SetError(".target.providers.defaults.paths.webroot_dir is nil")
			break
		}

		path = filepath.Join(addPrefix...)
		path = filepath.Join(path, p.Defaults.Paths.WebrootDir)

		p.state.SetOk()
	}

	return path
}


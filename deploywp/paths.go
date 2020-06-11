package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"path/filepath"
)


type Paths struct {
	BasePath    string `json:"base_path" mapstructure:"base_path"`
	WebrootPath string `json:"webroot_path" mapstructure:"webroot_path"`
	Wordpress   Wordpress `json:"wordpress"`

	//BaseAbsPath string

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (p *Paths) New(runtime *toolRuntime.TypeRuntime) *Paths {
	runtime = runtime.EnsureNotNil()
	return &Paths {
		WebrootPath: "",
		Wordpress:   *((*Wordpress).New(&Wordpress{}, runtime)),

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (p *Paths) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.state = p.state.EnsureNotNil()
	return p.state
}


type Wordpress struct {
	ContentPath string `json:"content_path" mapstructure:"content_path"`
	CorePath    string `json:"core_path" mapstructure:"core_path"`
	RootPath    string `json:"root_path" mapstructure:"root_path"`
	VendorPath  string `json:"vendor_path" mapstructure:"vendor_path"`

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (wp *Wordpress) New(runtime *toolRuntime.TypeRuntime) *Wordpress {
	runtime = runtime.EnsureNotNil()
	return &Wordpress {
		ContentPath: "",
		CorePath:    "",
		RootPath:    "",
		VendorPath:  "",

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (wp *Wordpress) IsNil() *ux.State {
	if state := ux.IfNilReturnError(wp); state.IsError() {
		return state
	}
	wp.state = wp.state.EnsureNotNil()
	return wp.state
}


func _FileToAbs(f ...string) string {
	var ret string

	for range onlyOnce {
		ret = filepath.Join(f...)

		if filepath.IsAbs(ret) {
			break
		}

		var err error
		ret, err = filepath.Abs(ret)
		if err != nil {
			ret = ""
			break
		}
	}

	return ret
}


func (p *Paths) ExpandPaths() *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}
	p.BasePath = _FileToAbs(p.BasePath)
	return p.state
}


func (p *Paths) GetBasePath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return p.BasePath
}

func (p *Paths) GetWebRootPath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(p.BasePath, p.WebrootPath)
}

func (p *Paths) GetContentPath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(p.BasePath, p.WebrootPath, p.Wordpress.ContentPath)
}

func (p *Paths) GetCorePath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(p.BasePath, p.WebrootPath, p.Wordpress.CorePath)
}

func (p *Paths) GetRootPath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(p.BasePath, p.WebrootPath, p.Wordpress.RootPath)
}

func (p *Paths) GetVendorPath() string {
	if state := p.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(p.BasePath, p.WebrootPath, p.Wordpress.VendorPath)
}

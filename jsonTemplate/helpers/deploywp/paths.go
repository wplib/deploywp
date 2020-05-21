package deploywp

import (
	"github.com/wplib/deploywp/ux"
	"path/filepath"
)


type Paths struct {
	BasePath    string `json:"base_path" mapstructure:"base_path"`
	WebrootPath string `json:"webroot_path" mapstructure:"webroot_path"`
	Wordpress   Wordpress `json:"wordpress"`

	//BaseAbsPath string

	Valid bool
	State *ux.State
}

type Wordpress struct {
	ContentPath string `json:"content_path" mapstructure:"content_path"`
	CorePath    string `json:"core_path" mapstructure:"core_path"`
	RootPath    string `json:"root_path" mapstructure:"root_path"`
	VendorPath  string `json:"vendor_path" mapstructure:"vendor_path"`

	Valid bool
	State *ux.State
}


func (me *Paths) New() Paths {
	me = &Paths {
		WebrootPath: "",
		Wordpress:   me.Wordpress.New(),
		State: ux.NewState(false),
	}
	return *me
}

func (me *Wordpress) New() Wordpress {
	me = &Wordpress {
		ContentPath: "",
		CorePath:    "",
		RootPath:    "",
		VendorPath:  "",
		State: ux.NewState(false),
	}
	return *me
}

func (e *Paths) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}

func (e *Wordpress) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}

func _FileToAbs(f ...string) string {
	var ret string

	for range OnlyOnce {
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


func (me *Paths) ExpandPaths() *ux.State {
	if state := me.IsNil(); state.IsError() {
		return state
	}
	me.BasePath = _FileToAbs(me.BasePath)
	return me.State
}


func (me *Paths) GetBasePath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.BasePath
}

func (me *Paths) GetWebRootPath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(me.BasePath, me.WebrootPath)
}

func (me *Paths) GetContentPath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.ContentPath)
}

func (me *Paths) GetCorePath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.CorePath)
}

func (me *Paths) GetRootPath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.RootPath)
}

func (me *Paths) GetVendorPath() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.VendorPath)
}

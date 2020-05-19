package deploywp

import (
	"github.com/wplib/deploywp/only"
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
		State: ux.NewState(),
	}
	return *me
}

func (me *Wordpress) New() Wordpress {
	me = &Wordpress {
		ContentPath: "",
		CorePath:    "",
		RootPath:    "",
		VendorPath:  "",
		State: ux.NewState(),
	}
	return *me
}

func (me *Paths) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			ok = true
		}
		// @TODO - perform other validity checks here.

		ok = false
	}

	return ok
}

func _FileToAbs(f ...string) string {
	var ret string

	for range only.Once {
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
	for range only.Once {
		if me.IsNil() {
			break
		}

		me.BasePath = _FileToAbs(me.BasePath)
		me.State = ux.NewState()
	}

	return me.State
}


func (me *Paths) GetBasePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = me.BasePath
	}

	return ret
}

func (me *Paths) GetWebRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = filepath.Join(me.BasePath, me.WebrootPath)
	}

	return ret
}

func (me *Paths) GetContentPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.ContentPath)
	}

	return ret
}

func (me *Paths) GetCorePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.CorePath)
	}

	return ret
}

func (me *Paths) GetRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.RootPath)
	}

	return ret
}

func (me *Paths) GetVendorPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}
		ret = filepath.Join(me.BasePath, me.WebrootPath, me.Wordpress.VendorPath)
	}

	return ret
}

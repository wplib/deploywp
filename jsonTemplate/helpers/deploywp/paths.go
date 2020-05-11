package deploywp

import (
	"github.com/wplib/deploywp/only"
	"path/filepath"
)


type Paths struct {
	BasePath    string `json:"base_path" mapstructure:"base_path"`
	WebrootPath string `json:"webroot_path" mapstructure:"webroot_path"`
	Wordpress   Wordpress `json:"wordpress"`

	//BaseAbsPath string

	Valid bool
	Error error
}

type Wordpress struct {
	ContentPath string `json:"content_path" mapstructure:"content_path"`
	CorePath    string `json:"core_path" mapstructure:"core_path"`
	RootPath    string `json:"root_path" mapstructure:"root_path"`
	VendorPath  string `json:"vendor_path" mapstructure:"vendor_path"`

	Valid bool
	Error error
}


//var _ deploywp.PathsGetter = (*Paths)(nil)

func (me *Paths) New() Paths {
	if me == nil {
		me = &Paths {
			WebrootPath: "",
			Wordpress:   me.Wordpress.New(),
		}
	}

	return *me
}

func (me *Wordpress) New() Wordpress {
	if me == nil {
		me = &Wordpress {
			ContentPath: "",
			CorePath:    "",
			RootPath:    "",
			VendorPath:  "",
		}
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

func (me *Paths) ExpandPaths() error {
	var err error

	for range only.Once {
		if me.IsNil() {
			break
		}

		me.BasePath = _FileToAbs(me.BasePath)
	}

	return err
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

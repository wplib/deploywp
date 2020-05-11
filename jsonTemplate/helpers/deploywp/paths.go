package deploywp

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"path/filepath"
)


type Paths struct {
	Prefix string `json:"prefix"`
	WebrootPath string `json:"webroot_path" mapstructure:"webroot_path"`
	Wordpress Wordpress `json:"wordpress"`

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


func (me *Paths) GetWebRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		var err error
		ret, err = filepath.Abs(fmt.Sprintf("%s/%s", me.Prefix, me.WebrootPath))
		if err != nil {
			ret = ""
		}
	}

	return ret
}

func (me *Paths) GetContentPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		var err error
		ret, err = filepath.Abs(fmt.Sprintf("%s/%s/%s", me.Prefix, me.WebrootPath, me.Wordpress.ContentPath))
		if err != nil {
			ret = ""
		}
	}

	return ret
}

func (me *Paths) GetCorePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		var err error
		ret, err = filepath.Abs(fmt.Sprintf("%s/%s/%s", me.Prefix, me.WebrootPath, me.Wordpress.CorePath))
		if err != nil {
			ret = ""
		}
	}

	return ret
}

func (me *Paths) GetRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		var err error
		ret, err = filepath.Abs(fmt.Sprintf("%s/%s/%s", me.Prefix, me.WebrootPath, me.Wordpress.RootPath))
		if err != nil {
			ret = ""
		}
	}

	return ret
}

func (me *Paths) GetVendorPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		var err error
		ret, err = filepath.Abs(fmt.Sprintf("%s/%s/%s", me.Prefix, me.WebrootPath, me.Wordpress.VendorPath))
		if err != nil {
			ret = ""
		}
	}

	return ret
}

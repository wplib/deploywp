package deploywp

import (
	"github.com/wplib/deploywp/only"
)


type Paths struct {
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
			ok = false
		}
		// @TODO - perform other validity checks here.

		ok = true
	}

	return ok
}


func (me *Paths) GetWebRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.WebrootPath
	}

	return ret
}

func (me *Paths) GetContentPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Wordpress.ContentPath
	}

	return ret
}

func (me *Paths) GetCorePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Wordpress.CorePath
	}

	return ret
}

func (me *Paths) GetRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Wordpress.RootPath
	}

	return ret
}

func (me *Paths) GetVendorPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Wordpress.VendorPath
	}

	return ret
}

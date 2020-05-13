package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


type Source struct {
	Build      Build      `json:"build"`
	Paths      Paths      `json:"paths"`
	Repository Repository `json:"repository"`
	Revision   Revision   `json:"revision"`

	AbsPaths   Paths
	Valid bool
	Error error
}


func (me *Source) New() Source {
	me.Build.New()
	me.Paths.New()
	me.Repository.New()
	me.Revision.New()

	return *me
}

func (me *Source) Process() error {
	for range only.Once {
		if me.IsNil() {
			break
		}

		me.AbsPaths = me.Paths
		me.Error = me.AbsPaths.ExpandPaths()
	}

	return me.Error
}

func (me *Source) IsNil() bool {
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


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Source) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		if helperTypes.ReflectBoolArg(abs) {
			ret = &me.AbsPaths
			break
		}

		ret = &me.Paths
	}

	return ret
}
//func (me *Source) GetBasePath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetBasePath()
//			break
//		}
//
//		ret = me.Paths.GetBasePath()
//	}
//
//	return ret
//}
//func (me *Source) GetWebRootPath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetWebRootPath()
//			break
//		}
//
//		ret = me.Paths.GetWebRootPath()
//	}
//
//	return ret
//}
//func (me *Source) GetContentPath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetContentPath()
//			break
//		}
//		ret = me.Paths.GetContentPath()
//	}
//
//	return ret
//}
//func (me *Source) GetCorePath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetCorePath()
//			break
//		}
//
//		ret = me.Paths.GetCorePath()
//	}
//
//	return ret
//}
//func (me *Source) GetRootPath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetRootPath()
//			break
//		}
//
//		ret = me.Paths.GetRootPath()
//	}
//
//	return ret
//}
//func (me *Source) GetVendorPath(abs ...interface{}) string {
//	var ret string
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if helperTypes.ReflectBoolArg(abs) {
//			ret = me.AbsPaths.GetVendorPath()
//			break
//		}
//
//		ret = me.Paths.GetVendorPath()
//	}
//
//	return ret
//}


// ////////////////////////////////////////////////////////////////////////////////
// Repository
func (me *Source) GetRepositoryProvider() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Repository.GetProvider()
	}

	return ret
}
func (me *Source) GetRepositoryUrl() URL {
	var ret URL

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Repository.GetUrl()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revision
func (me *Source) GetRevisionType() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revision.GetType()
	}

	return ret
}
func (me *Source) GetRevisionName() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revision.GetName()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Build
func (me *Source) GetBuild() bool {
	var ret bool

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Build.GetBuild()
	}

	return ret
}

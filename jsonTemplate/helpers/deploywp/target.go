package deploywp

import (
	"github.com/jinzhu/copier"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


type Target struct {
	Files     Files           `json:"files"`
	Paths     Paths           `json:"paths"`
	Providers Providers       `json:"providers"`
	Revisions TargetRevisions `json:"revisions"`

	AbsPaths  Paths
	AbsFiles  Files
	Valid bool
	Error error
}

func (me *Target) New() Target {
	if me == nil {
		me = &Target {
			Files:     me.Files.New(),
			Paths:     me.Paths.New(),
			Providers: me.Providers.New(),
			Revisions: me.Revisions.New(),
		}
	}

	return *me
}

func (me *Target) Process() error {
	for range only.Once {
		if me.IsNil() {
			break
		}

		me.Error = copier.Copy(&me.AbsPaths, &me.Paths)
		if me.Error != nil {
			break
		}

		me.Error = me.AbsPaths.ExpandPaths()
		if me.Error != nil {
			break
		}

		me.AbsFiles.Copy = append(me.AbsFiles.Copy, me.Files.Copy...)
		me.AbsFiles.Delete = append(me.AbsFiles.Delete, me.Files.Delete...)
		me.AbsFiles.Exclude = append(me.AbsFiles.Exclude, me.Files.Exclude...)
		me.AbsFiles.Keep = append(me.AbsFiles.Keep, me.Files.Keep...)

		me.Error = me.AbsFiles.Process(me.AbsPaths)
		if me.Error != nil {
			break
		}

		me.Error = me.Files.Process(me.Paths)
	}

	return me.Error
}

func (me *Target) IsNil() bool {
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
// Files
func (me *Target) GetFiles(ftype interface{}) *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Files.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Target) GetPaths(abs ...interface{}) *Paths {
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
//func (me *Target) GetBasePath(abs ...interface{}) string {
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
//func (me *Target) GetWebRootPath(abs ...interface{}) string {
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
//func (me *Target) GetContentPath(abs ...interface{}) string {
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
//
//		ret = me.Paths.GetContentPath()
//	}
//
//	return ret
//}
//func (me *Target) GetCorePath(abs ...interface{}) string {
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
//func (me *Target) GetRootPath(abs ...interface{}) string {
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
//func (me *Target) GetVendorPath(abs ...interface{}) string {
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
// Providers
func (me *Target) GetProvider(provider interface{}) *Provider {
	var ret *Provider

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Providers.GetProvider(provider)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revisions
func (me *Target) GetRevision(host interface{}) *TargetRevision {
	var ret *TargetRevision

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revisions.GetRevision(host)
	}

	return ret
}

package helperSystem

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"os"
)

var _ helperTypes.TypeOsPathGetter = (*TypeOsPath)(nil)
type TypeOsPath helperTypes.TypeOsPath

var _ helperTypes.TypeErrorGetter = (*TypeError)(nil)
type TypeError helperTypes.TypeError


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func HelperChdir(dir ...interface{}) *TypeOsPath {
	var cd TypeOsPath

	for range only.Once {
		f := helperTypes.ReflectPath(dir...)
		if f == nil {
			cd.Error = errors.New("directory empty")
			break
		}

		cd = *ResolveAbsPath(*f)
		if cd.IsError() {
			cd.Error = cd.Error
			break
		}
		if !cd.Exists {
			cd.Error = errors.New("directory not found")
			break
		}
		if cd.IsFile {
			cd.Error = errors.New("directory is a file")
			break
		}

		cd.Error = os.Chdir(cd.Dirname)
		if cd.Error != nil {
			break
		}

		var cwd string
		cwd, cd.Error = os.Getwd()
		if cd.Error != nil {
			break
		}
		if cwd != cd.Dirname {
			break
		}
	}

	return &cd
}


// Usage:
//		{{ $ret := Getwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperGetwd() *TypeOsPath {
	var ret TypeOsPath

	for range only.Once {
		ret.Path, ret.Error = os.Getwd()
		if ret.Error != nil {
			break
		}

		ret = *ResolveAbsPath(ret.Path)
	}

	return &ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperChmod(mode interface{}, name ...interface{}) *TypeOsPath {
	var cd TypeOsPath

	for range only.Once {
		f := helperTypes.ReflectPath(name...)
		if f == nil {
			cd.Error = errors.New("directory empty")
			break
		}

		m := helperTypes.ReflectFileMode(mode)
		if m == nil {
			break
		}

		cd = *ResolveAbsPath(*f)
		if cd.IsError() {
			cd.Error = cd.Error
			break
		}
		if !cd.Exists {
			cd.Error = errors.New("directory not found")
			break
		}
		if cd.IsFile {
			cd.Error = errors.New("directory is a file")
			break
		}

		cd.Error = os.Chmod(cd.Path, *m)
		if cd.Error != nil {
			break
		}
	}

	return &cd
}

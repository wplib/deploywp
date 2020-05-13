package helperSystem

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"os"
)


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func HelperChdir(dir ...interface{}) *helperTypes.TypeOsPath {
	var cd helperTypes.TypeOsPath

	for range only.Once {
		f := helperTypes.ReflectPath(dir...)
		if f == nil {
			cd.SetError("directory empty")
			break
		}

		cd = *ResolveAbsPath(*f)
		if cd.IsError() {
			break
		}
		if !cd.Exists {
			cd.SetError("directory not found")
			break
		}
		if cd.IsFile {
			cd.SetError("directory is a file")
			break
		}

		cd.ErrorValue = os.Chdir(cd.Dirname)
		if cd.IsError() {
			break
		}

		var cwd string
		cwd, cd.ErrorValue = os.Getwd()
		if cd.IsError() {
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
func HelperGetwd() *helperTypes.TypeOsPath {
	var ret helperTypes.TypeOsPath

	for range only.Once {
		ret.Path, ret.ErrorValue = os.Getwd()
		if ret.IsError() {
			break
		}

		ret = *ResolveAbsPath(ret.Path)
	}

	return &ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperChmod(mode interface{}, name ...interface{}) *helperTypes.TypeOsPath {
	var cd helperTypes.TypeOsPath

	for range only.Once {
		f := helperTypes.ReflectPath(name...)
		if f == nil {
			cd.SetError("directory empty")
			break
		}

		m := helperTypes.ReflectFileMode(mode)
		if m == nil {
			break
		}

		cd = *ResolveAbsPath(*f)
		if cd.IsError() {
			break
		}
		if !cd.Exists {
			cd.SetError("directory not found")
			break
		}
		if cd.IsFile {
			cd.SetError("directory is a file")
			break
		}

		cd.ErrorValue = os.Chmod(cd.Path, *m)
		if cd.IsError() {
			break
		}
	}

	return &cd
}

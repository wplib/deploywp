// High level helper functions available within templates - general file related.
package helperPath

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


type HelperOsPath TypeOsPath


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func HelperReadFile(file ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetError("filename empty")
			break
		}

		ret.SetPath(*f)
		ret.State = (*ux.State)(ret.ReadFile())
		if ret.State.IsError() {
			break
		}
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $return := WriteFile .Data.Source 0644 "dir1" "dir2/dir3" "filename.txt" }}
func HelperWriteFile(contents interface{}, perms interface{}, file ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetError("filename is nil")
			break
		}
		ret.SetPath(*f)

		c := helperTypes.ReflectByteArray(contents)
		if c == nil {
			ret.State.SetError("content string is nil")
			break
		}
		ret.LoadContents(*c)

		p := ReflectFileMode(perms)
		if p == nil {
			ret.SetMode(0)
		} else {
			ret.SetMode(*p)
		}

		ret.State = (*ux.State)(ret.WriteFile())
		if ret.State.IsError() {
			break
		}
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func HelperChdir(dir ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		f := ReflectPath(dir...)
		if f == nil {
			ret.State.SetError("directory is empty")
			break
		}
		ret.SetPath(*f)

		ret.State = (*ux.State)(ret.Chdir())
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperGetCwd() *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		cwd, state := ret.GetCwd()
		if (*ux.State)(state).IsError() {
			break
		}
		ret.SetPath(cwd)
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperIsCwd() *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		if ret.IsCwd() {
			break
		}
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperChmod(mode interface{}, name ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		f := ReflectPath(name...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		m := ReflectFileMode(mode)
		if m == nil {
			break
		}

		ret.State = (*ux.State)(ret.Chmod(*m))
		if ret.State.IsError() {
			break
		}
	}

	return (*HelperOsPath)(ret)
}

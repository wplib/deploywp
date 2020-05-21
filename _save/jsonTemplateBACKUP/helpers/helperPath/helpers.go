// High level helper functions available within templates - general file related.
package helperPath

import (
	"github.com/wplib/deploywp/only"
)


type HelperOsPath TypeOsPath


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func HelperNewPath(file ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}

		if !ret.SetPath(*f) {
			ret.State.SetError("path error")
			break
		}

		ret.State.SetState(ret.StatPath())
		if ret.State.IsError() {
			break
		}
	}

	return ReflectHelperOsPath(ret)
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperChmod(mode interface{}, path ...interface{}) *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		m := ReflectFileMode(mode)
		if m == nil {
			break
		}

		ret.State.SetState(ret.Chmod(*m))
		if ret.State.IsError() {
			break
		}
	}

	return ret
}

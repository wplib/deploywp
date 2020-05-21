package helperPath

import "github.com/wplib/deploywp/only"


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func HelperChdir(dir ...interface{}) *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(dir...)
		if f == nil {
			ret.State.SetError("directory is empty")
			break
		}
		ret.SetPath(*f)

		ret.State.SetState(ret.Chdir())
	}

	return ret
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperGetCwd() *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		state := ret.GetCwd()
		ret.State.SetState(state)
		if ret.State.IsError() {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperIsCwd() *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		if ret.IsCwd() {
			ret.State.Response = true
			break
		}
		ret.State.Response = false
	}

	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperCreateDir(path ...interface{}) *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		ret.State.SetState(ret.Mkdir())
		if ret.State.IsError() {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperRemoveDir(force bool, path ...interface{}) *TypeOsPath {
	ret := NewOsPath()

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		if force {
			ret.SetRemoveable()
		}
		ret.State.SetState(ret.RemoveDir())
		if ret.State.IsError() {
			break
		}
	}

	return ret
}

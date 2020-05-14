package helperFile

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
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


//// Usage:
////		{{  }}
//func (me *HelperOsPath) Path(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(path...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		me.SetPath(*f)
//	}
//
//	return me
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) SetDir(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) SetFile(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) IsDir(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) IsFile(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) Exists(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) ModTime(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) Mode(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *HelperOsPath) Size(path ...interface{}) *HelperOsPath {
//	for range only.Once {
//		f := ReflectPath(file...)
//		if f == nil {
//			me.State.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.State.SetError("content string is nil")
//			break
//		}
//
//		p := ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		me._Path = ResolveAbsPath(*f)
//		//if ret._File.IsError() {
//		//	break
//		//}
//		//if !ret._File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me._IsDir {
//			me.State.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me._Path, *c, *p)
//		if err != nil {
//			me.State.SetError(err)
//			break
//		}
//	}
//
//	return (*HelperOsPath)(me)
//}


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func HelperChdir(dir ...interface{}) *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		f := ReflectPath(dir...)
		if f == nil {
			ret.State.SetError("directory is nil")
			break
		}
		ret.SetPath(*f)

		ret.State = (*ux.State)(ret.Chdir())
	}

	return (*HelperOsPath)(ret)
}


// Usage:
//		{{ $ret := Getwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func HelperGetwd() *HelperOsPath {
	ret := NewOsPath()

	for range only.Once {
		ret._Path, ret.ErrorValue = os.Getwd()
		if ret.IsError() {
			break
		}

		ret = *ResolveAbsPath(ret._Path)
	}

	return &ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func HelperChmod(mode interface{}, name ...interface{}) *HelperOsPath {
	var cd HelperOsPath

	for range only.Once {
		f := ReflectPath(name...)
		if f == nil {
			cd.SetError("directory empty")
			break
		}

		m := ReflectFileMode(mode)
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


func (me *HelperOsPath) _IsValid() bool {
	var ok bool

	for range only.Once {
		if !me._Valid {
			me.State.SetError("path not valid")
			break
		}

		if me._Path == "" {
			me.State.SetError("path not set")
			break
		}
	}

	return ok
}
func (me *HelperOsPath) _IsNotValid() bool {
	return !me._IsValid()
}

package helperSystem

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"io/ioutil"
	"strings"
)

var _ helperTypes.TypeOsPathGetter = (*TypeOsPath)(nil)
type TypeOsPath helperTypes.TypeOsPath
var _ helperTypes.TypeOsPathGetter = (*TypeReadFile)(nil)
type TypeReadFile helperTypes.TypeReadFile
var _ helperTypes.TypeOsPathGetter = (*TypeWriteFile)(nil)
type TypeWriteFile helperTypes.TypeWriteFile


//type TypeReadFile struct {
//	TypeError
//	File *TypeOsPath
//	String string
//	Array  []string
//}
//
//type TypeWriteFile struct {
//	TypeError
//	File *TypeOsPath
//}


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func HelperReadFile(file ...interface{}) *helperTypes.TypeReadFile {
	var rf helperTypes.TypeReadFile

	for range only.Once {
		f := helperTypes.ReflectPath(file...)
		if f == nil {
			rf.SetError("filename empty")
			break
		}

		rf.File = ResolveAbsPath(*f)
		if rf.File.IsError() {
			rf.SetError(rf.File.ErrorValue)
			break
		}
		if !rf.File.Exists {
			rf.SetError("filename not found")
			break
		}
		if rf.File.IsDir {
			rf.SetError("filename is a directory")
			break
		}

		var d []byte
		var err error
		d, err = ioutil.ReadFile(rf.File.Path)
		if err != nil {
			rf.SetError(err)
			break
		}

		rf.String = string(d)
		rf.Array = strings.Split(string(d), "\n")
	}

	return &rf
}


// Usage:
//		{{ $return := WriteFile .Data.Source 0644 "dir1" "dir2/dir3" "filename.txt" }}
func HelperWriteFile(contents interface{}, perms interface{}, file ...interface{}) *helperTypes.TypeWriteFile {
	var ret helperTypes.TypeWriteFile

	for range only.Once {
		f := helperTypes.ReflectPath(file...)
		if f == nil {
			ret.SetError("filename is nil")
			break
		}

		c := helperTypes.ReflectByteArray(contents)
		if c == nil {
			ret.SetError("content string is nil")
			break
		}

		p := helperTypes.ReflectFileMode(perms)
		if p == nil {
			break
		}
		if *p == 0 {
			*p = 0644
		}


		ret.File = ResolveAbsPath(*f)
		//if ret.File.IsError() {
		//	break
		//}
		//if !ret.File.Exists {
		//	ret.Error = errors.New("filename not found")
		//	break
		//}
		if ret.File.IsDir {
			ret.SetError("filename is a directory")
			break
		}


		err := ioutil.WriteFile(ret.File.Path, c, *p)
		if err != nil {
			ret.SetError(err)
			break
		}
	}

	return &ret
}

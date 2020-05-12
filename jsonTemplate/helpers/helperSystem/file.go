package helperSystem

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"io/ioutil"
	"strings"
)

type TypeReadFile struct {
	TypeError
	File *TypeOsPath
	String string
	Array  []string
}

type TypeWriteFile struct {
	TypeError
	File *TypeOsPath
}


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func HelperReadFile(file ...interface{}) *TypeReadFile {
	var rf TypeReadFile

	for range only.Once {
		f := helperTypes.ReflectPath(file...)
		if f == nil {
			rf.Error = errors.New("filename empty")
			break
		}

		rf.File = ResolveAbsPath(*f)
		if rf.File.IsError() {
			rf.Error = rf.File.Error
			break
		}
		if !rf.File.Exists {
			rf.Error = errors.New("filename not found")
			break
		}
		if rf.File.IsDir {
			rf.Error = errors.New("filename is a directory")
			break
		}

		var d []byte
		d, rf.Error = ioutil.ReadFile(rf.File.Path)
		if rf.Error != nil {
			break
		}

		rf.String = string(d)
		rf.Array = strings.Split(string(d), "\n")
	}

	return &rf
}


// Usage:
//		{{ $return := WriteFile .Data.Source 0644 "dir1" "dir2/dir3" "filename.txt" }}
func HelperWriteFile(contents interface{}, perms interface{}, file ...interface{}) *TypeWriteFile {
	var ret TypeWriteFile

	for range only.Once {
		f := helperTypes.ReflectPath(file...)
		if f == nil {
			ret.Error = errors.New("filename is nil")
			break
		}

		c := helperTypes.ReflectByteArray(contents)
		if c == nil {
			ret.Error = errors.New("content string is nil")
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
		if ret.File.IsError() {
			ret.Error = ret.File.Error
			break
		}
		//if !ret.File.Exists {
		//	ret.Error = errors.New("filename not found")
		//	break
		//}
		if ret.File.IsDir {
			ret.Error = errors.New("filename is a directory")
			break
		}


		ret.Error = ioutil.WriteFile(ret.File.Path, c, *p)
		if ret.Error != nil {
			break
		}
	}

	return &ret
}

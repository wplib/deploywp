package helperSystem

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type TypeFile struct {
	Error error
	String string
	Array []string
}


func _FileToAbs(f ...string) string {
	var ret string

	for range only.Once {
		ret = filepath.Join(f...)

		if filepath.IsAbs(ret) {
			break
		}

		var err error
		ret, err = filepath.Abs(ret)
		if err != nil {
			ret = ""
			break
		}
	}
	//ret = strings.ReplaceAll(ret, "//", "/")

	return ret
}


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func ReadFile(file interface{}) TypeFile {
	var ret TypeFile

	for range only.Once {
		f := helperTypes.ReflectString(file)
		if f == nil {
			ret.Error = errors.New("file name empty")
			break
		}

		rf := _FileToAbs(*f)
		if rf == "" {
			ret.Error = errors.New("file name not defined")
			break
		}

		var d []byte
		d, ret.Error = ioutil.ReadFile(rf)
		if ret.Error != nil {
			break
		}

		ret.String = string(d)
		ret.Array = strings.Split(string(d), "\n")
	}

	return ret
}


// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func WriteFile(file interface{}, contents interface{}, perms interface{}) TypeError {
	var ret TypeError

	for range only.Once {
		f := helperTypes.ReflectString(file)
		if f == nil {
			ret.Error = errors.New("file name empty")
			break
		}

		c := helperTypes.ReflectByteArray(contents)
		if c == nil {
			ret.Error = errors.New("contents empty")
			break
		}

		p := helperTypes.ReflectFileMode(perms)
		if p == nil {
			break
		}
		if *p == 0 {
			*p = 0644
		}

		wf := _FileToAbs(*f)
		if wf == "" {
			ret.Error = errors.New("file name not defined")
			break
		}

		ret.Error = ioutil.WriteFile(wf, c, *p)
		if ret.Error != nil {
			break
		}
	}

	return ret
}

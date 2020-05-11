package helperSystem

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"os"
)

type TypeDir struct {
	Error error
	Dir string
}


// Usage:
//		{{ $response := Chdir "/root" }}
func Chdir(dir interface{}) TypeDir {
	var ret TypeDir

	for range only.Once {
		d := helperTypes.ReflectString(dir)
		if d == nil {
			break
		}

		wd := _FileToAbs(*d)
		if wd == "" {
			ret.Error = errors.New("file name not defined")
			break
		}

		err := os.Chdir(wd)
		if err != nil {
			break
		}

		var cwd string
		cwd, err = os.Getwd()
		if err != nil {
			break
		}
		if cwd != wd {
			break
		}

		ret.Dir = cwd
	}

	return ret
}


// Usage:
//		{{ $response := Getwd }}
func Getwd() TypeDir {
	var ret TypeDir

	for range only.Once {
		ret.Dir, ret.Error = os.Getwd()
		if ret.Error != nil {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $response := Getwd }}
func Chmod(name interface{}, mode interface{}) TypeDir {
	var ret TypeDir

	for range only.Once {
		n := helperTypes.ReflectString(name)
		if n == nil {
			break
		}

		m := helperTypes.ReflectFileMode(mode)
		if m == nil {
			break
		}

		ret.Error = os.Chmod(*n, *m)
		if ret.Error != nil {
			break
		}
	}

	return ret
}

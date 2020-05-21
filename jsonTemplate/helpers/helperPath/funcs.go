package helperPath

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"os"
	"path/filepath"
	"reflect"
)


func ReflectFileMode(ref interface{}) *os.FileMode {
	var fm os.FileMode

	for range OnlyOnce {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.Uint32 {
			break
		}

		fm = os.FileMode(value.Uint())
	}

	return &fm
}


func ReflectPath(ref ...interface{}) *string {
	var fp string

	for range OnlyOnce {
		var path []string
		for _, r := range ref {
			// Sometimes we can have dirs within each string slice.
			// EG: [0] = "dir1/dir2" OR [0] = "dir1\dir2"
			// This handles paths across O/S sanely.
			p := filepath.SplitList(*helperTypes.ReflectString(r))

			path = append(path, p...)
		}
		fp = filepath.Join(path...)
	}

	return &fp
}


func ReflectAbsPath(ref ...interface{}) *string {
	var fp string

	for range OnlyOnce {
		path := ReflectPath(ref...)

		var err error
		fp, err = filepath.Abs(*path)
		if err != nil {
			fp = *path
		}
	}

	return &fp
}


func _GetAbsPath(p ...string) string {
	var ret string

	for range OnlyOnce {
		ret = filepath.Join(p...)

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

	return ret
}

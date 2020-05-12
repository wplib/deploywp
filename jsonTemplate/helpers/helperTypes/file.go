package helperTypes

import (
	"github.com/wplib/deploywp/only"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

type TypeOsPathGetter interface {
}

type TypeOsPath struct {
	TypeError
	Path     string
	Filename string
	Dirname  string
	IsDir    bool
	IsFile   bool
	Exists   bool
	ModTime  time.Time
	Mode     os.FileMode
	Size     int64
}


func ReflectFileMode(ref interface{}) *os.FileMode {
	var fm os.FileMode

	for range only.Once {
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

	for range only.Once {
		var path []string
		for _, r := range ref {
			// Sometimes we can have dirs within each string slice.
			// EG: [0] = "dir1/dir2" OR [0] = "dir1\dir2"
			// This handles paths across O/S sanely.
			p := filepath.SplitList(*ReflectString(r))
			path = append(path, p...)
		}
		fp = filepath.Join(path...)
	}

	return &fp
}

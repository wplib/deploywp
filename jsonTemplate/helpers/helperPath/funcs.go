package helperFile

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"path/filepath"
	"reflect"
	"time"
)


const DefaultSeparator = "\n"

type OsPathGetter interface {
}


type TypeOsPath struct {
	State     *ux.State

	_Path     string
	_Filename string
	_Dirname  string
	_IsDir    bool
	_IsFile   bool
	_Exists   bool
	_ModTime  time.Time
	_Name     string
	_Mode     os.FileMode
	_Size     int64

	_String    string
	_Array     []string
	_Separator string
	_Valid     bool
	_Overwrite bool
}

type State ux.State


func NewOsPath() *TypeOsPath {
	return &TypeOsPath{
		State:     ux.New(),
		_Path:     "",
		_Filename: "",
		_Dirname:  "",
		_IsDir:    false,
		_IsFile:   false,
		_Exists:   false,
		_ModTime:  time.Time{},
		_Mode:     0,
		_Size:     0,
		_String:   "",
		_Array:    nil,
		_Separator: DefaultSeparator,
		_Valid:     false,
		_Overwrite: false,
	}
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
			p := filepath.SplitList(*helperTypes.ReflectString(r))
			path = append(path, p...)
		}
		fp = filepath.Join(path...)
	}

	return &fp
}


func _GetAbsPath(p ...string) string {
	var ret string

	for range only.Once {
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

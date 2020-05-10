package deploywp

import (
	"github.com/wplib/deploywp/only"
)


type Files struct {
	Copy    FilesArray `json:"copy"`
	Delete  FilesArray `json:"delete"`
	Exclude FilesArray `json:"exclude"`
	Keep    FilesArray `json:"keep"`

	Valid bool
	Error error
}
type FilesArray []string


func (me *Files) New() Files {
	if me == nil {
		me = &Files{
			Copy:    FilesArray{},
			Delete:  FilesArray{},
			Exclude: FilesArray{},
			Keep:    FilesArray{},
		}
	}

	return *me
}

func (me *Files) Process() error {
	for range only.Once {
		if me.IsNil() {
			break
		}
	}

	return me.Error
}

func (me *Files) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			ok = false
		}
		// @TODO - perform other validity checks here.

		ok = true
	}

	return ok
}


const (
	TargetActionCopy = "copy"
	TargetActionDelete = "delete"
	TargetActionExclude = "exclude"
	TargetActionKeep = "keep"
)

func (me *Files) GetFiles(action interface{}) *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		value := ReflectString(action)
		if value == nil {
			//ret.Error = errors.New("GetTargetFiles arg not a string")
			break
		}

		switch *value {
			case TargetActionCopy:
				ret = &me.Copy
			case TargetActionDelete:
				ret = &me.Delete
			case TargetActionExclude:
				ret = &me.Exclude
			case TargetActionKeep:
				ret = &me.Keep

			default:
				//ret.Error = errors.New("GetTargetFiles file type not defined")
		}
	}

	return ret
}

func (me *Files) GetCopyFiles() *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Copy
	}

	return ret
}

func (me *Files) GetDeleteFiles() *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Delete
	}

	return ret
}

func (me *Files) GetExcludeFiles() *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Exclude
	}

	return ret
}

func (me *Files) GetKeepFiles() *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Keep
	}

	return ret
}

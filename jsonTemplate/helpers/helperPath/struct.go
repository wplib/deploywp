package helperPath

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"path/filepath"
	"strings"
	"time"
)


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
	_Remote    bool
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}


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


func (p *TypeOsPath) SetPath(path ...string) bool {
	var ok bool

	for range only.Once {
		if p._IsRemotePath(path...) {
			ok = p._SetRemotePath(path...)
			break
		}

		ok = p._SetLocalPath(path...)
	}

	return ok
}
func (p *TypeOsPath) GetPath() string {
	return p._Path
}
func (p *TypeOsPath) _SetLocalPath(path ...string) bool {
	for range only.Once {
		p._Valid = false
		p._Path = _GetAbsPath(path...)
		if p._Path == "" {
			break
		}
		p._Valid = true
		p._Remote = false
	}

	return p._Valid
}
func (p *TypeOsPath) _SetRemotePath(path ...string) bool {
	for range only.Once {
		p._Valid = false
		// @TODO - May have to change this logic to:
		// @TODO - p._Path = strings.Join(path, "")
		p._Path = filepath.Join(path...)
		if p._Path == "" {
			break
		}
		p._Valid = true
		p._Remote = true
	}

	return p._Valid
}
func (p *TypeOsPath) _IsRemotePath(path ...string) bool {
	return strings.ContainsAny(strings.Join(path, ""), ":@")
}


//func (p *TypeOsPath) SetRemote() {
//	// @TODO - Add in extra logic to convert filename to path.
//	p._Remote = true
//}
func (p *TypeOsPath) IsRemote() bool {
	return p._Remote
}


//func (p *TypeOsPath) SetFilename(filename string) {
//	// @TODO - Add in extra logic to convert filename to path.
//	p._Filename = filename
//}
func (p *TypeOsPath) GetFilename() string {
	return p._Filename
}


//func (p *TypeOsPath) SetDirname(dirname string) {
//	// @TODO - Add in extra logic to convert dirname to path.
//	p._Dirname = dirname
//}
func (p *TypeOsPath) GetDirname() string {
	return p._Dirname
}


func (p *TypeOsPath) SetModTime(time time.Time) {
	p._ModTime = time
}
func (p *TypeOsPath) GetModTime() time.Time {
	return p._ModTime
}


func (p *TypeOsPath) SetMode(mode os.FileMode) {
	p._Mode = mode
}
func (p *TypeOsPath) GetMode() os.FileMode {
	return p._Mode
}


//func (p *TypeOsPath) SetSize(size int64) {
//	p._Size = size
//}
func (p *TypeOsPath) GetSize() int64 {
	return p._Size
}


//func (p *TypeOsPath) SetExists() {
//	p._Exists = true
//}
func (p *TypeOsPath) Exists() bool {
	var ok bool

	for range only.Once {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("path does not exist")
			break
		}
		p.State.SetOk("path exists")
		ok = p._Exists
	}

	return ok
}
func (p *TypeOsPath) FileExists() bool {
	var ok bool

	for range only.Once {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("file does not exist")
			break
		}
		if !p._IsFile {
			p.State.SetError("file is a dir")
			break
		}
		p.State.SetOk("file exists")
		ok = p._Exists
	}

	return ok
}
func (p *TypeOsPath) DirExists() bool {
	var ok bool

	for range only.Once {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("dir does not exist")
			break
		}
		if !p._IsDir {
			p.State.SetError("dir is a file")
			break
		}
		p.State.SetOk("dir exists")
		ok = p._Exists
	}

	return ok
}


func (p *TypeOsPath) ThisIsAFile() {
	p._IsFile = true
	p._IsDir = false
	p.State.Clear()
}
func (p *TypeOsPath) IsAFile() bool {
	return p._IsFile
}


func (p *TypeOsPath) ThisIsADir() {
	p._IsFile = false
	p._IsDir = true
	p.State.Clear()
}
func (p *TypeOsPath) IsADir() bool {
	return p._IsDir
}


func (p *TypeOsPath) _SetValid() {
	p._Valid = true
}
func (p *TypeOsPath) _SetInvalid() {
	p._Valid = false
}
func (p *TypeOsPath) IsValid() bool {
	var ok bool

	for range only.Once {
		if !p._Valid {
			p.State.SetError("path not valid")
			break
		}

		if p._Path == "" {
			p.State.SetError("path not set")
			break
		}
	}

	return ok
}
func (p *TypeOsPath) IsInvalid() bool {
	return !p.IsValid()
}

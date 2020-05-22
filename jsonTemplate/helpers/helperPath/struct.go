package helperPath

import (
	"github.com/wplib/deploywp/ux"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const OnlyOnce = "1"


type OsPathGetter interface {
}


type TypeOsPath struct {
	State         *ux.State

	_Path         string
	_Filename     string
	_Dirname      string
	_IsDir        bool
	_IsFile       bool
	_Exists       bool
	_ModTime      time.Time
	_Name         string
	_Mode         os.FileMode
	_Size         int64

	_String       string
	_Array        []string
	_Separator    string
	fileHandle    *os.File

	_Valid        bool
	_CanOverwrite bool
	_CanRemove    bool
	_Remote       bool
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectHelperOsPath(p *TypeOsPath) *HelperOsPath {
	return (*HelperOsPath)(p)
}

func (c *TypeOsPath) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func NewOsPath(debugMode bool) *TypeOsPath {
	p := &TypeOsPath{
		State:         ux.NewState(debugMode),
		_Path:         "",
		_Filename:     "",
		_Dirname:      "",
		_IsDir:        false,
		_IsFile:       false,
		_Exists:       false,
		_ModTime:      time.Time{},
		_Mode:         0,
		_Size:         0,
		_String:       "",
		_Array:        nil,
		_Separator:    DefaultSeparator,
		_Valid:        false,
		_CanRemove:    false,
		_CanOverwrite: false,
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return p
}


func (p *TypeOsPath) GetPath() string {
	return p._Path
}
func (p *TypeOsPath) SetPath(path ...string) bool {
	p._Path = ""
	return p.AppendPath(path...)
}
func (p *TypeOsPath) AppendPath(path ...string) bool {
	var ok bool

	for range OnlyOnce {
		if p._IsRemotePath(p._Path) {
			ok = p._AppendRemotePath(path...)
			break
		}
		if p._IsRemotePath(path...) {
			ok = p._AppendRemotePath(path...)
			break
		}

		ok = p._AppendLocalPath(path...)
	}

	return ok
}
func (p *TypeOsPath) _AppendLocalPath(path ...string) bool {
	for range OnlyOnce {
		p._Valid = false
		p._Path = _GetAbsPath(path...)
		if p._Path == "" {
			p.State.SetError("src path empty")
			break
		}
		//p._Valid = true
		p._Remote = false

		// Reset these until a later StatPath()
		p._Dirname = ""
		p._Filename = ""
		p._IsDir = false
		p._IsFile = false
		p._Exists = false
	}

	return p._Valid
}
func (p *TypeOsPath) _AppendRemotePath(path ...string) bool {
	for range OnlyOnce {
		p._Valid = false
		// @TODO - May have to change this logic to:
		// @TODO - p._Path = strings.Join(path, "")
		p._Path = filepath.Join(path...)
		if p._Path == "" {
			p.State.SetError("src path empty")
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
func (p *TypeOsPath) GetModTimeString() string {
	return p._ModTime.Format("2006-01-02T15:04:05-0700")
}
func (p *TypeOsPath) GetModTimeEpoch() int64 {
	return p._ModTime.Unix()
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

	for range OnlyOnce {
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
func (p *TypeOsPath) NotExists() bool {
	return !p.Exists()
}
func (p *TypeOsPath) FileExists() bool {
	var ok bool

	for range OnlyOnce {
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

	for range OnlyOnce {
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
	for range OnlyOnce {
		//if !p._Valid {
		//	p.State.SetError("path not valid")
		//	break
		//}

		if p._Path == "" {
			p.State.SetError("path not set")
			break
		}

		p._Valid = true
	}

	return p._Valid
}
func (p *TypeOsPath) IsInvalid() bool {
	return !p.IsValid()
}
func (p *TypeOsPath) IsNotValid() bool {
	return !p.IsValid()
}


func (p *TypeOsPath) SetOverwriteable() {
	p._CanOverwrite = true
}
func (p *TypeOsPath) CanOverwrite() bool {
	return p._CanOverwrite
}
func (p *TypeOsPath) IsOverwriteable() bool {
	return p._CanOverwrite
}


func (p *TypeOsPath) SetRemoveable() {
	p._CanRemove = true
}
func (p *TypeOsPath) CanRemove() bool {
	return p._CanRemove
}
func (p *TypeOsPath) IsRemoveable() bool {
	return p._CanRemove
}

package helperFile

import (
	"github.com/wplib/deploywp/only"
	"os"
	"time"
)


func (p *TypeOsPath) SetPath(path ...string) {
	for range only.Once {
		p._Path = _GetAbsPath(path...)
		if p._Path == "" {
			break
		}
	}
}
func (p *TypeOsPath) GetPath() string {
	return p._Path
}


func (p *TypeOsPath) SetFilename(filename string) {
	p._Filename = filename
}
func (p *TypeOsPath) GetFilename() string {
	return p._Filename
}


func (p *TypeOsPath) SetDirname(dir string) {
	p._Dirname = dir
}
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
		if !p._IsValid() {
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
		if !p._IsValid() {
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
		if !p._IsValid() {
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


func (p *TypeOsPath) _IsValid() bool {
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

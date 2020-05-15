package helperPath

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"path/filepath"
)


func (p *TypeOsPath) StatPath() *State {
	for range only.Once {
		p.State.Clear()

		if p._Path == "" {
			p.State.SetError("path is empty")
			break
		}

		if p._Remote {
			// @TODO - Maybe add in some remote checks?
			p._Valid = true
			p._Exists = true
			p.State.SetOk("path is remote")
			break
		}

		var stat os.FileInfo
		stat, p.State.Error = os.Stat(p._Path)
		if os.IsNotExist(p.State.Error) {
			p.State.SetError("path does not exist - %s", p.State.Error)
			p._Exists = false
			break
		}
		if p.State.Error != nil {
			break
		}

		p._Valid = true
		p._Exists = true
		p._ModTime = stat.ModTime()
		p._Name = stat.Name()
		p._Mode = stat.Mode()
		p._Size = stat.Size()

		if stat.IsDir() {
			p._IsDir = true
			p._IsFile = false
			p._Dirname = p._Path
			p._Filename = ""

		} else {
			p._IsDir = false
			p._IsFile = true
			p._Dirname = filepath.Dir(p._Path)
			p._Filename = filepath.Base(p._Path)
		}

		p.State.SetOk("stat OK")
	}

	return (*State)(p.State)
}


func (p *TypeOsPath) Chmod(m os.FileMode) *State {
	for range only.Once {
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		p.State = (*ux.State)(p.StatPath())
		if p.State.IsError() {
			break
		}

		p.State.Error = os.Chmod(p._Path, m)
		if p.State.IsError() {
			break
		}

		p.State = (*ux.State)(p.StatPath())
		if p.State.IsError() {
			break
		}

		p.State.SetOk("chmod OK")
	}

	return (*State)(p.State)
}


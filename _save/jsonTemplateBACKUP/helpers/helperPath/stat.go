package helperPath

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
)


func (p *TypeOsPath) StatPath() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
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
		var err error
		stat, err = os.Stat(p._Path)
		if os.IsNotExist(err) {
			p.State.SetError("path does not exist - %s", err)
			p._Exists = false
			break
		}
		p.State.SetError(err)
		if p.State.IsError() {
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
			p._Dirname = fmt.Sprintf("%s%c", p._Path, filepath.Separator)
			p._Path = p._Dirname
			p._Filename = ""

		} else {
			p._IsDir = false
			p._IsFile = true
			p._Dirname = fmt.Sprintf("%s%c", filepath.Dir(p._Path), filepath.Separator)
			p._Filename = filepath.Base(p._Path)
		}

		p.State.SetOk("stat OK")
	}

	return p.State
}


func (p *TypeOsPath) Chmod(m os.FileMode) *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		p.State.SetState(p.StatPath())
		if p.State.IsError() {
			break
		}

		var err error
		err = os.Chmod(p._Path, m)
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State.SetState(p.StatPath())
		if p.State.IsError() {
			break
		}

		p.State.SetOk("chmod OK")
	}

	return p.State
}

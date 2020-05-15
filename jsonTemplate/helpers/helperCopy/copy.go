package helperCopy

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
)


func (p *TypeOsPath) Copy() *State {
	for range only.Once {
		if !p._IsValid() {
			break
		}

		p.State = (*ux.State)(p.StatPath())
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("directory not found")
			break
		}
		if p._IsFile {
			p.State.SetError("directory is a file")
			break
		}


		p.State.Error = os.Chdir(p._Path)
		if p.State.IsError() {
			break
		}


		var cwd string
		cwd, p.State.Error = os.Getwd()
		if p.State.IsError() {
			break
		}
		if cwd != p._Path {
			p.State.SetError("destination directory doesn't match")
			break
		}

		p.State.SetOk("chdir OK")
	}

	return (*State)(p.State)
}

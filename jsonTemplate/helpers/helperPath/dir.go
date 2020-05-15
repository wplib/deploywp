package helperPath

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
)


func (p *TypeOsPath) Chdir() *State {
	for range only.Once {
		p.State.Clear()

		if !p.IsValid() {
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


func (p *TypeOsPath) GetCwd() (string, *State) {
	var cwd string

	for range only.Once {
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		cwd, p.State.Error = os.Getwd()
		if p.State.IsError() {
			break
		}

		//p.State.Output = cwd
		p.State.Clear()
	}

	return cwd, (*State)(p.State)
}


func (p *TypeOsPath) IsCwd() bool {
	var ok bool

	for range only.Once {
		cwd, state := p.GetCwd()
		if (*ux.State)(state).IsError() {
			break
		}

		if cwd != p._Path {
			break
		}
	}

	return ok
}

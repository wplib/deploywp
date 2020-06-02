package helperPath

import (
	"github.com/wplib/deploywp/only"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)


func (p *TypeOsPath) Chdir() *ux.State {
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
		if !p._Exists {
			p.State.SetError("directory not found")
			break
		}
		if p._IsFile {
			p.State.SetError("directory is a file")
			break
		}


		var err error
		err = os.Chdir(p._Path)
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		//var cwd string
		//cwd, err = os.Getwd()
		//p.State.SetError(err)
		//if p.State.IsError() {
		//	break
		//}
		//if cwd != p._Path {
		//	p.State.SetError("destination directory doesn't match")
		//	break
		//}

		p.State.SetOk("chdir OK")
	}

	return p.State
}


func (p *TypeOsPath) GetCwd() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}


		p.State.Response = ""
		var cwd string
		var err error
		cwd, err = os.Getwd()
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State.Response = cwd
		p.State.Clear()
	}

	return p.State
}


func (p *TypeOsPath) IsCwd() bool {
	var ok bool

	for range OnlyOnce {
		p.State.SetFunction("")

		state := p.GetCwd()
		if state.IsError() {
			break
		}

		if state.Response != p._Path {
			break
		}
	}

	return ok
}


func (p *TypeOsPath) Mkdir() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}


		if p._Mode == 0 {
			p._Mode = 0644
		}

		p.State.Response = false
		var err error
		err = os.Mkdir(p._Path, p._Mode)
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State.Response = true
		p.State.Clear()
		p.State.SetOk("mkdir OK")
	}

	return p.State
}

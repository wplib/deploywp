package helperPath

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/ux"
	"os"
)


func (p *TypeOsPath) Remove() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range OnlyOnce {
			p.StatPath()
			if !p._Exists {
				p.State.SetWarning("path '%s' doesn't exist", p._Path)
				break
			}
			if p._CanRemove {
				break
			}

			if !helperSystem.HelperUserPromptBool("Remove path '%s'? (Y|N) ", p._Path) {
				p.State.SetWarning("not removing path '%s'", p._Path)
				break
			}
			p.State.Clear()
		}
		if p.State.IsNotOk() {
			break
		}

		err := os.Remove(p._Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("path '%s' removed OK", p._Path)
	}

	return p.State
}


func (p *TypeOsPath) RemoveFile() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range OnlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path is a directory")
				break
			}
			if !p._Exists {
				p.State.SetWarning("file '%s' doesn't exist", p._Path)
				break
			}
			if p._CanRemove {
				break
			}

			p.State.Clear()
			if !helperSystem.HelperUserPromptBool("Remove file '%s'? (Y|N) ", p._Path) {
				p.State.SetWarning("not removing file '%s'", p._Path)
				break
			}
		}
		if p.State.IsNotOk() {
			break
		}

		err := os.Remove(p._Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("file '%s' removed OK", p._Path)
	}

	return p.State
}


func (p *TypeOsPath) RemoveDir() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range OnlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path '%s' is a directory", p._Path)
				break
			}
			if !p._Exists {
				p.State.SetWarning("directory '%s' doesn't exist", p._Path)
				break
			}
			if p._CanRemove {
				break
			}

			if !helperSystem.HelperUserPromptBool("Remove directory '%s'? (Y|N) ", p._Path) {
				p.State.SetWarning("not removing file '%s'", p._Path)
				break
			}
			p.State.Clear()
		}
		if p.State.IsError() {
			break
		}

		err := os.Remove(p._Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("file '%s' removed OK", p._Path)
	}

	return p.State
}

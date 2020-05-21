package helperCopy

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


func (p *TypeOsCopy) Copy() *ux.State {
	for range OnlyOnce {
		if !p.Source.Exists() {
			p.State.SetError("src path not found")
			break
		}

		for range OnlyOnce {
			if p.Destination.NotExists() {
				break
			}
			if p.Destination.CanOverwrite() {
				break
			}
			p.State.SetError("cannot overwrite destination")
		}

		if p.State.IsError() {
			break
		}

		// @TODO - do copying of files here

		p.State.SetOk("chdir OK")
	}

	return p.State
}

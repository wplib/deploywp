package deploywp

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


type Build struct {
	Empty bool

	Valid bool
	State *ux.State
}

//var _ deploywp.BuildGetter = (*Build)(nil)

func (me *Build) New() Build {
	me = &Build {
		Empty: false,
		State: ux.NewState(),
	}
	return *me
}

func (me *Build) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			ok = true
		}
		// @TODO - perform other validity checks here.

		ok = false
	}

	return ok
}


func (me *Build) GetBuild() bool {
	var ret bool

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Empty
	}

	return ret
}

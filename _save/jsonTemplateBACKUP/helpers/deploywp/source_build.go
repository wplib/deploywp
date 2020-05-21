package deploywp

import (
	"github.com/wplib/deploywp/only"
)


type Build struct {
	Empty bool

	Valid bool
	Error error
}

//var _ deploywp.BuildGetter = (*Build)(nil)

func (me *Build) New() Build {
	if me == nil {
		me = &Build {
			Empty: false,
		}
	}

	return *me
}

func (me *Build) IsNil() bool {
	var ok bool

	for range OnlyOnce {
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

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.Empty
	}

	return ret
}

package deploywp

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


type Revision struct {
	RefName string `json:"ref_name" mapstructure:"ref_name"`
	RefType string `json:"ref_type" mapstructure:"ref_type"`

	Valid bool
	State *ux.State
}


func (me *Revision) New() Revision {
	me = &Revision {
		RefName: "",
		RefType: "",
		State: ux.NewState(),
	}
	return *me
}

func (me *Revision) IsNil() bool {
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


func (me *Revision) GetType() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.RefType
	}

	return ret
}

func (me *Revision) GetName() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.RefName
	}

	return ret
}

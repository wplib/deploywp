package deploywp

import (
	"github.com/wplib/deploywp/only"
)


type Revision struct {
	RefName string `json:"ref_name" mapstructure:"ref_name"`
	RefType string `json:"ref_type" mapstructure:"ref_type"`

	Valid bool
	Error error
}


func (me *Revision) New() Revision {
	if me == nil {
		me = &Revision {
			RefName: "",
			RefType: "",
		}
	}

	return *me
}

func (me *Revision) IsNil() bool {
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


func (me *Revision) GetType() string {
	var ret string

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.RefType
	}

	return ret
}

func (me *Revision) GetName() string {
	var ret string

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.RefName
	}

	return ret
}

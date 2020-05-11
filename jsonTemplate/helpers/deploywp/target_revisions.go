package deploywp

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


type TargetRevision struct {
	HostName string `json:"host_name" mapstructure:"host_name"`
	RefName  string `json:"ref_name" mapstructure:"ref_name"`

	Valid bool
	Error error
}
type TargetRevisions []TargetRevision

func (me *TargetRevision) New() TargetRevision {
	if me == nil {
		me = &TargetRevision {
			HostName:     "",
			RefName:     "",
		}
	}

	return *me
}

func (me *TargetRevisions) New() TargetRevisions {
	if me == nil {
		me = &TargetRevisions {}
	}

	return *me
}

func (me *TargetRevisions) IsNil() bool {
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


func (me *TargetRevisions) GetRevision(host interface{}) *TargetRevision {
	var ret TargetRevision

	for range only.Once {
		if me.IsNil() {
			break
		}

		value := helperTypes.ReflectString(host)
		if value == nil {
			ret.Error = errors.New("GetRevision arg not a string")
			break
		}

		for _, v := range *me {
			if v.HostName == *value {
				ret = v
				ret.Valid = true
				break
			}
		}

		if !ret.Valid {
			ret.Error = errors.New("GetRevision host_name not found")
			break
		}
	}

	return &ret
}

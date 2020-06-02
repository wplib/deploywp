package deploywp

import (
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type TargetRevision struct {
	AutoDeploy bool   `json:"auto_deploy" mapstructure:"auto_deploy"`
	HostName   string `json:"host_name" mapstructure:"host_name"`
	RefName    string `json:"ref_name" mapstructure:"ref_name"`

	Valid bool
	State *ux.State
}
type TargetRevisions []TargetRevision

func (me *TargetRevision) New() TargetRevision {
	if me == nil {
		me = &TargetRevision {
			HostName:     "",
			RefName:     "",
			State: ux.NewState(false),
		}
	}

	return *me
}

func (me *TargetRevisions) New() TargetRevisions {
	if me == nil {
		me = &TargetRevisions{}
	}

	return *me
}

//func (e *TargetRevisions) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(e); state.IsError() {
//		return state
//	}
//	e.State = e.State.EnsureNotNil()
//	return e.State
//}


func (me *TargetRevisions) GetRevision(host interface{}) *TargetRevision {
	var ret TargetRevision

	for range OnlyOnce {
		value := helperTypes.ReflectString(host)
		if value == nil {
			ret.State.SetError("GetRevision arg not a string")
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
			ret.State.SetError("GetRevision host_name not found")
			break
		}
	}

	return &ret
}

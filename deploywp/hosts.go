package deploywp

import (
	"github.com/wplib/deploywp/jtc/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
)


type Host struct {
	HostName string `json:"host_name" mapstructure:"host_name"`
	Label    string `json:"label"`
	Provider string `json:"provider"`

	Valid bool
	State *ux.State
}
type Hosts []Host


func NewHost() *Host {
	me := &Host{
		HostName: "",
		Label:    "",
		Provider: "",
		Valid: false,
		State: ux.NewState(false),
	}
	return me
}

func (me *Host) New() *Host {
	me = &Host{
		HostName: "",
		Label:    "",
		Provider: "",
		Valid: true,
		State: ux.NewState(false),
	}
	return me
}

func (me *Hosts) New() Hosts {
	if me == nil {
		me = &Hosts{ }
	}

	return *me
}

func (me *Hosts) Count() int {
	return len(*me)
}

func (me *Hosts) Process() *ux.State {
	state := ux.NewState(false)

	for range OnlyOnce {
		for h, _ := range *me {
			(*me)[h].State = ux.NewState(false)
			(*me)[h].Valid = true
		}
	}

	return state
}


//func (e *Hosts) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(e); state.IsError() {
//		return state
//	}
//	e.State = e.State.EnsureNotNil()
//	return e.State
//}


func (me *Hosts) GetHost(host interface{}) *Host {
	ret := NewHost()

	for range OnlyOnce {
		value := helperTypes.ReflectString(host)
		if value == nil {
			ret.State.SetError("GetHost arg not a string")
			break
		}

		for _, v := range *me {
			if v.HostName == *value {
				ret = &v
				break
			}
		}

		if !ret.Valid {
			ret.State.SetError("GetHost hostname not found")
			break
		}
	}

	return ret
}

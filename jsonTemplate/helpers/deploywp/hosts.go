package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
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
		State: ux.NewState(),
	}
	return me
}

func (me *Host) New() *Host {
	me = &Host{
		HostName: "",
		Label:    "",
		Provider: "",
		Valid: true,
		State: ux.NewState(),
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
	state := ux.NewState()

	for range only.Once {
		if me.IsNil() {
			break
		}

		for h, _ := range *me {
			(*me)[h].State = ux.NewState()
			(*me)[h].Valid = true
		}
	}

	return state
}

func (me *Hosts) IsNil() bool {
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


func (me *Hosts) GetHost(host interface{}) *Host {
	ret := NewHost()

	for range only.Once {
		if me.IsNil() {
			break
		}

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

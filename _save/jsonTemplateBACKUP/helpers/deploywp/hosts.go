package deploywp

import (
	"errors"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


type Host struct {
	HostName string `json:"host_name" mapstructure:"host_name"`
	Label    string `json:"label"`
	Provider string `json:"provider"`

	Valid bool
	Error error
}
type Hosts []Host


func (me *Host) New() Host {
	if me == nil {
		me = &Host{
			HostName: "",
			Label:    "",
			Provider: "",
		}
	}
	return *me
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

func (me *Hosts) Process() error {
	var err error

	for range OnlyOnce {
		if me.IsNil() {
			break
		}
	}

	return err
}

func (me *Hosts) IsNil() bool {
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


func (me *Hosts) GetHost(host interface{}) *Host {
	var ret Host

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		value := helperTypes.ReflectString(host)
		if value == nil {
			ret.Error = errors.New("GetHost arg not a string")
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
			ret.Error = errors.New("GetHost hostname not found")
			break
		}
	}

	return &ret
}

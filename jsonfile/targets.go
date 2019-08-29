package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
)

var _ deploywp.TargetsGetter = (*Targets)(nil)

type Targets struct {
	Defaults *Host `json:"defaults"`
	Hosts    Hosts `json:"hosts"`
}

func (me Targets) GetDefaults() *deploywp.Host {
	return deploywp.NewHostFromGetter(me.Defaults)
}

func (me Targets) GetHosts() deploywp.Hosts {
	return deploywp.NewHostsFromGetter(me.Hosts)
}

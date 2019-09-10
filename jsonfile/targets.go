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
	d := me.Defaults
	if d.Id == "" {
		d.Id = "master"
	}
	if d.Name == "" {
		d.Name = "defaults"
	}
	if d.Label == "" {
		d.Label = "Defaults"
	}
	if d.WebRoot == "" {
		d.WebRoot = "/www"
	}
	if d.Paths != nil {
		d.Paths.ApplyDefaults(NewWordPressPaths())
	} else {
		d.Paths = NewWordPressPaths()
	}
	me.Defaults = d
	return deploywp.NewHostFromGetter(d)
}

func (me Targets) GetHosts() deploywp.Hosts {
	return deploywp.NewHostsFromGetter(me.Hosts)
}

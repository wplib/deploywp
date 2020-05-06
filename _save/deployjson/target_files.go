package deployjson

type Files struct {
	Copy    []string `json:"copy"`
	Delete  []string `json:"delete"`
	Exclude []string `json:"exclude"`
	Keep    []string `json:"keep"`
}


func (me *Files) New() Files {
	if me == nil {
		me = &Files{
			Copy:    []string{},
			Delete:  []string{},
			Exclude: []string{},
			Keep:    []string{},
		}
	}

	return *me
}



//var _ deploywp.FilesGetter = (*Files)(nil)
//
//func (me Files) GetDefaults() *deploywp.Host {
//	d := me.Defaults
//	if d.Id == "" {
//		d.Id = "master"
//	}
//	if d.Name == "" {
//		d.Name = "defaults"
//	}
//	if d.Label == "" {
//		d.Label = "Defaults"
//	}
//	if d.WebRoot == "" {
//		d.WebRoot = "/www"
//	}
//	if d.Paths != nil {
//		d.Paths.ApplyDefaults(NewWordPressPaths())
//	} else {
//		d.Paths = NewWordPressPaths()
//	}
//	me.Defaults = d
//
//	return deploywp.NewHostFromGetter(d)
//}
//
//func (me Files) GetHosts() deploywp.Hosts {
//	return deploywp.NewHostsFromGetter(me.Hosts)
//}

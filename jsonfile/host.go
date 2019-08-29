package jsonfile

import (
	"fmt"
	"github.com/wplib/deploywp/deploywp"
	"reflect"
)

var _ deploywp.HostsGetter = (*Hosts)(nil)
var _ deploywp.HostGetter = (*Host)(nil)

type Hosts []*Host

func (me Hosts) Count() int {
	return len(me)
}
func (me Hosts) Hosts() (hs deploywp.Hosts) {
	hs = make(deploywp.Hosts, me.Count())
	for i, h := range me {
		hs[i] = deploywp.NewHostFromGetter(h)
	}
	return hs
}

type Host struct {
	Id           Identifier       `json:"id"`
	SiteGuid     Guid             `json:"site_guid"`
	Domain       Domain           `json:"domain"`
	DomainSuffix Domain           `json:"domain_suffix"`
	Provider     Identifier       `json:"provider"`
	Name         ReadableName     `json:"name"`
	Label        Label            `json:"label"`
	Branch       Identifier       `json:"branch"`
	WebRoot      Path             `json:"web_root"`
	Repository   Repository       `json:"repository"`
	Paths        WordPressPaths   `json:"wp_paths"`
	Files        FileDispositions `json:"files"`
	After        string           `json:"after"`
}

func (me *Host) ApplyDefaults(defaults *Host) (err error) {

	hv := reflect.ValueOf(me).Elem()
	dv := reflect.ValueOf(defaults).Elem()

	for i := 0; i < hv.NumField(); i++ {
		fh := hv.Field(i)
		//fmt.Println(hv.Type().Field(i).Tag.Get("json"))
		if !fh.CanSet() {
			err = fmt.Errorf("unable to set host field '%s'", fh.Type().Name())
			break
		}
		fd := dv.Field(i)

		if fh.Kind() == reflect.Ptr {
			fh = fh.Elem()
			if fh.IsNil() {
				continue
			}
		}

		switch fh.Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice:
			fh.Set(fd)

		case reflect.String:
			if fh.String() == "" && fd.String() != "" {
				fh.SetString(fd.String())
			}
		}
	}
	return err
}

func (me *Host) GetId() deploywp.Identifier {
	return me.Id
}
func (me *Host) GetSiteGuid() deploywp.Guid {
	return me.SiteGuid
}
func (me *Host) GetDomain() deploywp.Domain {
	return me.Domain
}
func (me *Host) GetDomainSuffix() deploywp.Domain {
	return me.DomainSuffix
}
func (me *Host) GetProvider() deploywp.Identifier {
	return me.Provider
}
func (me *Host) GetName() deploywp.ReadableName {
	return me.Name
}
func (me *Host) GetLabel() deploywp.Label {
	return me.Label
}
func (me *Host) GetBranch() deploywp.Identifier {
	return me.Branch
}
func (me *Host) GetWebRoot() deploywp.Path {
	return me.WebRoot
}
func (me *Host) GetRepository() *deploywp.Repository {
	return deploywp.NewRepositoryFromGetter(me.Repository)
}
func (me *Host) GetPaths() *deploywp.WordPressPaths {
	return deploywp.NewWordPressPathsFromGetter(me.Paths)
}
func (me *Host) GetAfter() string {
	return me.After
}
func (me *Host) GetFiles() *deploywp.FileDispositions {
	return deploywp.NewFileDispositionsFromGetter(me.Files)
}

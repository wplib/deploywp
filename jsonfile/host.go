package jsonfile

import (
	"fmt"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
	"reflect"
)

var _ deploywp.HostsGetter = (*Hosts)(nil)
var _ deploywp.HostGetter = (*Host)(nil)
var _ providers.HostGetterSetter = (*Host)(nil)

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
	Id           Identifier        `json:"id"`
	SiteGuid     Guid              `json:"site_guid"`
	Domain       Domain            `json:"domain"`
	DomainSuffix Domain            `json:"domain_suffix"`
	ProviderId   Identifier        `json:"provider"`
	Name         ReadableName      `json:"name"`
	Label        Label             `json:"label"`
	Branch       Identifier        `json:"branch"`
	WebRoot      Path              `json:"web_root"`
	Repository   *Repository       `json:"repository"`
	Paths        *WordPressPaths   `json:"wp_paths"`
	Files        *FileDispositions `json:"files"`
	After        string            `json:"after"`
}

func (me *Host) GetProviderType() ProviderType {
	return providers.Dispense(me.ProviderId).GetType()
}

func (me *Host) SetRepository(r *providers.Repository) {
	me.Repository = NewRepository(r.Provider.GetId(), r.Url)
}

func (me *Host) SetDomain(d Domain) {
	me.Domain = d
}

func (me *Host) SetDomainSuffix(ds Domain) {
	me.DomainSuffix = ds
}

func (me *Host) SetWebRoot(wr Path) {
	me.WebRoot = wr
}

func (me *Host) SetBranch(b Identifier) {
	me.Branch = b
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

		if fd.Kind() == reflect.Ptr {
			fd = fd.Elem()
		}

		if fh.Kind() == reflect.Ptr {
			if fh.IsNil() {
				continue
			}
			fh = fh.Elem()
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

func (me *Host) GetId() Identifier {
	return me.Id
}
func (me *Host) GetSiteGuid() Guid {
	return me.SiteGuid
}
func (me *Host) GetDomain() Domain {
	return me.Domain
}
func (me *Host) GetDomainSuffix() Domain {
	return me.DomainSuffix
}
func (me *Host) GetProviderId() Identifier {
	return me.ProviderId
}
func (me *Host) GetName() ReadableName {
	return me.Name
}
func (me *Host) GetLabel() Label {
	return me.Label
}
func (me *Host) GetBranch() Identifier {
	return me.Branch
}
func (me *Host) GetWebRoot() Path {
	return me.WebRoot
}
func (me *Host) GetRepository() *providers.Repository {
	if me.Repository == nil {
		me.Repository = NewRepository(me.ProviderId, "")
	}
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

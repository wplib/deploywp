package jsonfile

import (
	"fmt"
	"reflect"
)

type Hosts []*Host

type Host struct {
	Id           Slug             `json:"id"`
	SiteGuid     Guid             `json:"site_guid"`
	Domain       Domain           `json:"domain"`
	DomainSuffix Domain           `json:"domain_suffix"`
	Provider     Slug             `json:"provider"`
	Name         ReadableName     `json:"name"`
	Label        Label            `json:"label"`
	Branch       Slug             `json:"branch"`
	WebRoot      Path             `json:"web_root"`
	Repository   Repository       `json:"repository"`
	Paths        WordPressPaths   `json:"wp_paths"`
	After        string           `json:"after"`
	Files        FileDispositions `json:"files"`
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

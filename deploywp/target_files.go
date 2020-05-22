package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
	"strings"
)


type Files struct {
	Copy    FilesArray `json:"copy"`
	Delete  FilesArray `json:"delete"`
	Exclude FilesArray `json:"exclude"`
	Keep    FilesArray `json:"keep"`

	Valid bool
	State *ux.State
}
type FilesArray []string


func (me *Files) New() Files {
	me.Copy =    FilesArray{}
	me.Delete =  FilesArray{}
	me.Exclude = FilesArray{}
	me.Keep =    FilesArray{}
	me.State = ux.NewState(false)
	return *me
}

func (me *Files) Process(paths Paths) *ux.State {
	if state := me.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		for i, p := range me.Copy {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			me.Copy[i] = p
		}

		for i, p := range me.Delete {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			me.Delete[i] = p
		}

		for i, p := range me.Exclude {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			me.Exclude[i] = p
		}

		for i, p := range me.Keep {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			me.Keep[i] = p
		}

		//bp := paths.GetWebRootPath()
		//fmt.Printf("EXPANDPATHS Files.Process()\n", bp)
	}

	return me.State
}


func (e *Files) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


const (
	TargetActionCopy = "copy"
	TargetActionDelete = "delete"
	TargetActionExclude = "exclude"
	TargetActionKeep = "keep"
)

func (me *Files) GetFiles(action interface{}) *FilesArray {
	var ret *FilesArray
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}

	for range OnlyOnce {
		value := helperTypes.ReflectString(action)
		if value == nil {
			//ret.Error = errors.New("GetTargetFiles arg not a string")
			break
		}

		switch *value {
			case TargetActionCopy:
				ret = &me.Copy
			case TargetActionDelete:
				ret = &me.Delete
			case TargetActionExclude:
				ret = &me.Exclude
			case TargetActionKeep:
				ret = &me.Keep

			default:
				//ret.Error = errors.New("GetTargetFiles file type not defined")
		}
	}

	return ret
}

func (me *Files) GetCopyFiles() *FilesArray {
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &me.Copy
}

func (me *Files) GetDeleteFiles() *FilesArray {
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &me.Delete
}

func (me *Files) GetExcludeFiles() *FilesArray {
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &me.Exclude
}

func (me *Files) GetKeepFiles() *FilesArray {
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &me.Keep
}

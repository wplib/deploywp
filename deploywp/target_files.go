package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


type Files struct {
	Copy    FilesArray `json:"copy"`
	Delete  FilesArray `json:"delete"`
	Exclude FilesArray `json:"exclude"`
	Keep    FilesArray `json:"keep"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (f *Files) New(runtime *toolRuntime.TypeRuntime) *Files {
	runtime = runtime.EnsureNotNil()
	return &Files{
		Copy:    *(*FilesArray).New(&FilesArray{}),
		Delete:  *(*FilesArray).New(&FilesArray{}),
		Exclude: *(*FilesArray).New(&FilesArray{}),
		Keep:    *(*FilesArray).New(&FilesArray{}),

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (f *Files) IsNil() *ux.State {
	if state := ux.IfNilReturnError(f); state.IsError() {
		return state
	}
	f.state = f.state.EnsureNotNil()
	return f.state
}
func (f *Files) Process(paths Paths) *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		for i, p := range f.Copy {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			f.Copy[i] = p
		}

		for i, p := range f.Delete {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			f.Delete[i] = p
		}

		for i, p := range f.Exclude {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			f.Exclude[i] = p
		}

		for i, p := range f.Keep {
			p = strings.ReplaceAll(p, "{webroot_path}", paths.GetWebRootPath())
			p = strings.ReplaceAll(p, "{wordpress.content_path}", paths.GetContentPath())
			p = strings.ReplaceAll(p, "{wordpress.vendor_path}", paths.GetVendorPath())
			p = strings.ReplaceAll(p, "{wordpress.core_path}", paths.GetCorePath())
			p = strings.ReplaceAll(p, "{wordpress.root_path}", paths.GetRootPath())
			f.Keep[i] = p
		}

		//bp := paths.GetWebRootPath()
		//fmt.Printf("EXPANDPATHS Files.Process()\n", bp)
	}

	return f.state
}


type FilesArray []string
func (fa *FilesArray) New() *FilesArray {
	if fa == nil {
		return &FilesArray{}
	}
	return fa
}
func (fa *FilesArray) Process(runtime *toolRuntime.TypeRuntime) *ux.State {
	state := ux.NewState(runtime.CmdName, runtime.Debug)
	for range onlyOnce {
		//for i, _ := range *fa {
		//	(*fa)[i] = *((*fa)[i].New(runtime))
		//}
	}
	return state
}
func (fa *FilesArray) FindBySomething() *FilesArray {
	ret := (*FilesArray).New(&FilesArray{})
	for range onlyOnce {
		//for i, _ := range *fa {
		//	(*fa)[i] = *((*fa)[i].New(runtime))
		//}
	}
	return ret
}


const (
	TargetActionCopy = "copy"
	TargetActionDelete = "delete"
	TargetActionExclude = "exclude"
	TargetActionKeep = "keep"
)

func (f *Files) GetFiles(action string) *FilesArray {
	ret := (*FilesArray).New(&FilesArray{})
	if state := f.IsNil(); state.IsError() {
		return ret
	}

	for range onlyOnce {
		if action == "" {
			//ret.Error = errors.New("GetTargetFiles arg not a string")
			break
		}

		switch action {
			case TargetActionCopy:
				ret = &f.Copy
			case TargetActionDelete:
				ret = &f.Delete
			case TargetActionExclude:
				ret = &f.Exclude
			case TargetActionKeep:
				ret = &f.Keep

			default:
				//ret.Error = errors.New("GetTargetFiles file type not defined")
		}
	}

	return ret
}

func (f *Files) GetCopyFiles() *FilesArray {
	if state := f.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &f.Copy
}

func (f *Files) GetDeleteFiles() *FilesArray {
	if state := f.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &f.Delete
}

func (f *Files) GetExcludeFiles() *FilesArray {
	if state := f.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &f.Exclude
}

func (f *Files) GetKeepFiles() *FilesArray {
	if state := f.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return &f.Keep
}

package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


type Files struct {
	Copy    FilesArray `json:"copy"`	// mapstructure:",squash"`
	Delete  FilesArray `json:"delete"`	// mapstructure:",squash"`
	Exclude FilesArray `json:"exclude"`	// mapstructure:",squash"`
	Keep    FilesArray `json:"keep"`	// mapstructure:",squash"`

	Valid   bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}

func (f *Files) New(runtime *toolRuntime.TypeRuntime) *Files {
	runtime = runtime.EnsureNotNil()
	return &Files{
		Copy:    *(*FilesArray).New(&FilesArray{}, runtime),
		Delete:  *(*FilesArray).New(&FilesArray{}, runtime),
		Exclude: *(*FilesArray).New(&FilesArray{}, runtime),
		Keep:    *(*FilesArray).New(&FilesArray{}, runtime),

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

func (f *Files) IsValid() bool {
	if state := ux.IfNilReturnError(f); state.IsError() {
		return false
	}
	for range onlyOnce {
		if f.Copy.IsNotValid() {
			//f.state = f.Copy.state
			f.Valid = false
			break
		}
		if f.Delete.IsNotValid() {
			//f.state = f.Delete.state
			f.Valid = false
			break
		}
		if f.Exclude.IsNotValid() {
			//f.state = f.Exclude.state
			f.Valid = false
			break
		}
		if f.Keep.IsNotValid() {
			//f.state = f.Keep.state
			f.Valid = false
			break
		}
		f.Valid = true
	}
	return f.Valid
}
func (f *Files) IsNotValid() bool {
	return !f.IsValid()
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


//type FilesArray struct{
//	Array []string `json:"array"`
//
//	Valid   bool
//	runtime *toolRuntime.TypeRuntime
//	state   *ux.State
//}
type FilesArray []string

func (fa *FilesArray) New(runtime *toolRuntime.TypeRuntime) *FilesArray {
	runtime = runtime.EnsureNotNil()
	return &FilesArray {
		//Array: []string{},
		//
		//Valid: false,
		//runtime: runtime,
		//state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}

func (fa *FilesArray) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(fa); state.IsError() {
		return ok
	}
	for range onlyOnce {
		ok = true
		//for _, f := range fa.Array {
		for _, f := range *fa {
			if f == "" {
				ok = false
				break
			}
		}
	}
	return ok
}
func (fa *FilesArray) IsNotValid() bool {
	return !fa.IsValid()
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

func (fa *FilesArray) Append(add *FilesArray) {
	//fa.Array = append(fa.Array, add.Array...)
	*fa = append(*fa, *add...)
}

//func (fa *FilesArray) FindBySomething() *FilesArray {
//	ret := (*FilesArray).New(&FilesArray{})
//	for range onlyOnce {
//		//for i, _ := range *fa {
//		//	(*fa)[i] = *((*fa)[i].New(runtime))
//		//}
//	}
//	return ret
//}


const (
	TargetActionCopy = "copy"
	TargetActionDelete = "delete"
	TargetActionExclude = "exclude"
	TargetActionKeep = "keep"
)

func (f *Files) GetFiles(action string) *FilesArray {
	ret := (*FilesArray).New(&FilesArray{}, f.runtime)
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

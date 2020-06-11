package deploywp

import (
	"github.com/jinzhu/copier"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type Target struct {
	Files     Files           `json:"files"`
	Paths     Paths           `json:"paths"`
	Providers Providers       `json:"providers"`
	Revisions TargetRevisions `json:"revisions"`

	AbsPaths  Paths
	AbsFiles  Files

	Valid bool
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
func (t *Target) New(runtime *toolRuntime.TypeRuntime) *Target {
	runtime = runtime.EnsureNotNil()
	return &Target{
		Files:     *((*Files).New(&Files{}, runtime)),
		Paths:     *((*Paths).New(&Paths{}, runtime)),
		Providers: *((*Providers).New(&Providers{})),
		Revisions: *((*TargetRevisions).New(&TargetRevisions{})),
		AbsPaths:  *((*Paths).New(&Paths{}, runtime)),
		AbsFiles:  *((*Files).New(&Files{}, runtime)),

		Valid:   true,
		runtime: runtime,
		state:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
}
func (t *Target) IsNil() *ux.State {
	if state := ux.IfNilReturnError(t); state.IsError() {
		return state
	}
	t.state = t.state.EnsureNotNil()
	return t.state
}
func (t *Target) Process() *ux.State {
	if state := t.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		err := copier.Copy(&t.AbsPaths, &t.Paths)
		t.state.SetError(err)
		if t.state.IsError() {
			break
		}

		t.state = t.AbsPaths.ExpandPaths()
		t.state.SetError(err)
		if t.state.IsError() {
			break
		}

		t.AbsFiles.Copy = append(t.AbsFiles.Copy, t.Files.Copy...)
		t.AbsFiles.Delete = append(t.AbsFiles.Delete, t.Files.Delete...)
		t.AbsFiles.Exclude = append(t.AbsFiles.Exclude, t.Files.Exclude...)
		t.AbsFiles.Keep = append(t.AbsFiles.Keep, t.Files.Keep...)

		t.state = t.AbsFiles.Process(t.AbsPaths)
		if t.state.IsError() {
			break
		}

		t.state = t.Files.Process(t.Paths)
		if t.state.IsError() {
			break
		}

		t.state = t.Providers.Process(t.runtime)
		if t.state.IsError() {
			break
		}

		t.state = t.Revisions.Process(t.runtime)
		if t.state.IsError() {
			break
		}

		t.Valid = true
	}

	return t.state
}


// ////////////////////////////////////////////////////////////////////////////////
// Files
func (t *Target) GetFiles(ftype string) *FilesArray {
	var ret *FilesArray
	if state := t.IsNil(); state.IsError() {
		return &FilesArray{}
	}

	for range onlyOnce {
		ret = t.Files.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (t *Target) GetPaths(abs ...interface{}) *Paths {
	var ret *Paths
	if state := t.IsNil(); state.IsError() {
		return &Paths{}
	}

	for range onlyOnce {
		if toolTypes.ReflectBoolArg(abs) {
			ret = &t.AbsPaths
			break
		}

		ret = &t.Paths
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Providers
func (t *Target) GetProviderByName(provider string) *Provider {
	var ret *Provider
	if state := t.IsNil(); state.IsError() {
		return ret
	}
	return t.Providers.GetByName(provider)
}
func (t *Target) GetProviderBySiteId(siteId string) *Provider {
	var ret *Provider
	if state := t.IsNil(); state.IsError() {
		return ret
	}
	ret = t.Providers.GetBySiteId(siteId)
	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revisions
func (t *Target) GetRevisionByHost(host string) *TargetRevision {
	var ret *TargetRevision
	if state := t.IsNil(); state.IsError() {
		return &TargetRevision{}
	}
	ret = t.Revisions.GetByHost(host)
	return ret
}

func (t *Target) GetRevisionByName(ref string) *TargetRevision {
	var ret *TargetRevision
	if state := t.IsNil(); state.IsError() {
		return &TargetRevision{}
	}
	ret = t.Revisions.GetByRefName(ref)
	return ret
}

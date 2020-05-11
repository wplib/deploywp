package deploywp

import (
	"github.com/wplib/deploywp/only"
)


type Target struct {
	Files     Files           `json:"files"`
	Paths     Paths           `json:"paths"`
	Providers Providers       `json:"providers"`
	Revisions TargetRevisions `json:"revisions"`

	Valid bool
	Error error
}

func (me *Target) New() Target {
	if me == nil {
		me = &Target {
			Files:     me.Files.New(),
			Paths:     me.Paths.New(),
			Providers: me.Providers.New(),
			Revisions: me.Revisions.New(),
		}
	}

	return *me
}

func (me *Target) Process() error {
	for range only.Once {
		if me.IsNil() {
			break
		}

		me.Error = me.Paths.ExpandPaths()
		me.Error = me.Files.Process()
	}

	return me.Error
}

func (me *Target) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			ok = true
		}
		// @TODO - perform other validity checks here.

		ok = false
	}

	return ok
}


// ////////////////////////////////////////////////////////////////////////////////
// Files
func (me *Target) GetFiles(ftype interface{}) *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Files.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Paths
func (me *Target) GetPaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Paths
	}

	return ret
}
func (me *Target) GetBasePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetBasePath()
	}

	return ret
}
func (me *Target) GetWebRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetWebRootPath()
	}

	return ret
}
func (me *Target) GetContentPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetContentPath()
	}

	return ret
}
func (me *Target) GetCorePath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetCorePath()
	}

	return ret
}
func (me *Target) GetRootPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetRootPath()
	}

	return ret
}
func (me *Target) GetVendorPath() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Paths.GetVendorPath()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Providers
func (me *Target) GetProvider(provider interface{}) *Provider {
	var ret *Provider

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Providers.GetProvider(provider)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Revisions
func (me *Target) GetRevision(host interface{}) *TargetRevision {
	var ret *TargetRevision

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Revisions.GetRevision(host)
	}

	return ret
}

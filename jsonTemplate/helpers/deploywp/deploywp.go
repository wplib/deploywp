package deploywp

import (
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/ux"
	//"github.com/davecgh/go-spew/spew"
	"github.com/wplib/deploywp/only"
)


type TypeDeployWp struct {
	Hosts  Hosts  `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`

	Runtime

	Valid  bool
	State  *ux.State
}

type Runtime struct {
	Exec runtime.Exec
}


func ReflectDeployWp(ref interface{}) *TypeDeployWp {
	return ref.(*TypeDeployWp)
}


func NewJsonFile() *TypeDeployWp {
	var jf TypeDeployWp

	jf.State = ux.NewState()
	jf.Runtime = Runtime{}

	jf.Hosts.New()
	jf.Source.New()
	jf.Target.New()

	return &jf
}


func (me *TypeDeployWp) IsNil() bool {
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
// Source
func (me *TypeDeployWp) GetSource() *Source {
	return &me.Source
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Paths
func (me *TypeDeployWp) GetSourcePaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Source.Paths
	}

	return ret
}
//func (me *TypeDeployWp) GetSourcePaths(abs ...interface{}) *Paths {
//	var ret *Paths
//
//	for range only.Once {
//		if me.IsNil() {
//			break
//		}
//
//		if len(abs) > 0 {
//			ok := helperTypes.ReflectBoolArg(abs[0])
//			if ok {
//				ret = &me.Source.AbsPaths
//				break
//			}
//		}
//
//		ret = &me.Source.Paths
//	}
//
//	return ret
//}
func (me *TypeDeployWp) GetSourceAbsPaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Source.AbsPaths
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Repository
func (me *TypeDeployWp) GetSourceRepositoryProvider() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRepositoryProvider()
	}

	return ret
}
func (me *TypeDeployWp) GetSourceRepositoryUrl() URL {
	var ret URL

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRepositoryUrl()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Revision
func (me *TypeDeployWp) GetSourceRevisionType() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRevisionType()
	}

	return ret
}
func (me *TypeDeployWp) GetSourceRevisionName() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRevisionName()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Build
func (me *TypeDeployWp) GetSourceBuild() bool {
	var ret bool

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetBuild()
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Target
func (me *TypeDeployWp) GetTarget() *Target {
	return &me.Target
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Files
func (me *TypeDeployWp) GetTargetFiles(ftype interface{}) *FilesArray {
	var ret *FilesArray

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Target.GetFiles(ftype)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Paths
func (me *TypeDeployWp) GetTargetPaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Target.Paths
	}

	return ret
}
func (me *TypeDeployWp) GetTargetAbsPaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Target.AbsPaths
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Revisions
func (me *TypeDeployWp) GetTargetRevision(host interface{}) *TargetRevision {
	var ret *TargetRevision

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Target.GetRevision(host)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Providers
func (me *TypeDeployWp) GetTargetProvider(provider interface{}) *Provider {
	var ret *Provider

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Target.GetProvider(provider)
	}

	return ret
}


// ////////////////////////////////////////////////////////////////////////////////
// Hosts
func (me *TypeDeployWp) GetHosts() *Hosts {
	return &me.Hosts
}

func (me *TypeDeployWp) GetHost(host interface{}) *Host {
	var ret *Host

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Hosts.GetHost(host)
	}

	return ret
}

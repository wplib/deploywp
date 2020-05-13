package deploywp

import (
	"github.com/mitchellh/mapstructure"
	//"github.com/davecgh/go-spew/spew"
	"github.com/wplib/deploywp/only"
)


type DeployWp struct {
	Hosts Hosts `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`

	Meta RuntimeMeta
	Valid bool
	Error error
}

type RuntimeMeta struct {
}

func _NewJsonFile() *DeployWp {
	var jf DeployWp

	jf.Hosts.New()
	jf.Source.New()
	jf.Target.New()

	return &jf
}

func HelperLoadDeployWp(str interface{}) *DeployWp {
	var j DeployWp

	for range only.Once {
		j.Error = mapstructure.Decode(str, &j)
		if j.Error != nil {
			break
		}

		j.Error = j.Source.Process()
		if j.Error != nil {
			break
		}

		j.Error = j.Target.Process()
		if j.Error != nil {
			break
		}

		j.Error = j.Hosts.Process()
		if j.Error != nil {
			break
		}

		j.Valid = true
	}

	return &j
}


func (me *DeployWp) IsNil() bool {
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
func (me *DeployWp) GetSource() *Source {
	return &me.Source
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Paths
func (me *DeployWp) GetSourcePaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Source.Paths
	}

	return ret
}
//func (me *DeployWp) GetSourcePaths(abs ...interface{}) *Paths {
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
func (me *DeployWp) GetSourceAbsPaths() *Paths {
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
func (me *DeployWp) GetSourceRepositoryProvider() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRepositoryProvider()
	}

	return ret
}
func (me *DeployWp) GetSourceRepositoryUrl() URL {
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
func (me *DeployWp) GetSourceRevisionType() string {
	var ret string

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRevisionType()
	}

	return ret
}
func (me *DeployWp) GetSourceRevisionName() string {
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
func (me *DeployWp) GetSourceBuild() bool {
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
func (me *DeployWp) GetTarget() *Target {
	return &me.Target
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Files
func (me *DeployWp) GetTargetFiles(ftype interface{}) *FilesArray {
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
func (me *DeployWp) GetTargetPaths() *Paths {
	var ret *Paths

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = &me.Target.Paths
	}

	return ret
}
func (me *DeployWp) GetTargetAbsPaths() *Paths {
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
func (me *DeployWp) GetTargetRevision(host interface{}) *TargetRevision {
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
func (me *DeployWp) GetTargetProvider(provider interface{}) *Provider {
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
func (me *DeployWp) GetHosts() *Hosts {
	return &me.Hosts
}

func (me *DeployWp) GetHost(host interface{}) *Host {
	var ret *Host

	for range only.Once {
		if me.IsNil() {
			break
		}

		ret = me.Hosts.GetHost(host)
	}

	return ret
}

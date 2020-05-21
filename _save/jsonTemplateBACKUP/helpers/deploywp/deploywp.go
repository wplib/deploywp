package deploywp

import (
	"github.com/mitchellh/mapstructure"
	//"github.com/davecgh/go-spew/spew"
	"github.com/wplib/deploywp/only"
)


type TypeDeployWp struct {
	Hosts Hosts `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`

	Meta RuntimeMeta
	Valid bool
	Error error
}

type RuntimeMeta struct {
}


func NewJsonFile() *TypeDeployWp {
	var jf TypeDeployWp

	jf.Hosts.New()
	jf.Source.New()
	jf.Target.New()

	return &jf
}

func HelperLoadDeployWp(str interface{}) *TypeDeployWp {
	j := NewJsonFile()

	for range OnlyOnce {
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

	return j
}


func (me *TypeDeployWp) IsNil() bool {
	var ok bool

	for range OnlyOnce {
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

	for range OnlyOnce {
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
//	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRepositoryProvider()
	}

	return ret
}
func (me *TypeDeployWp) GetSourceRepositoryUrl() URL {
	var ret URL

	for range OnlyOnce {
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

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.Source.GetRevisionType()
	}

	return ret
}
func (me *TypeDeployWp) GetSourceRevisionName() string {
	var ret string

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = &me.Target.Paths
	}

	return ret
}
func (me *TypeDeployWp) GetTargetAbsPaths() *Paths {
	var ret *Paths

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
		if me.IsNil() {
			break
		}

		ret = me.Hosts.GetHost(host)
	}

	return ret
}

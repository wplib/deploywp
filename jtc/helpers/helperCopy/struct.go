package helperCopy

import (
	"github.com/wplib/deploywp/jtc/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
)

const OnlyOnce = "1"


// @TODO - Look at several other copy options that provide "cloud" based copies.
// @TODO - https://rclone.org/
// @TODO - https://pkg.go.dev/github.com/Redundancy/go-sync?tab=doc
// @TODO - https://github.com/Redundancy/go-sync
// @TODO - https://pkg.go.dev/bitbucket.org/kardianos/rsync?tab=doc
// @TODO - https://github.com/zloylos/grsync


type OsCopyGetter interface {
}

type TypeOsPath helperPath.TypeOsPath

type TypeOsCopy struct {
	State        *ux.State

	Source       *helperPath.TypeOsPath
	Destination  *helperPath.TypeOsPath

	Exclude PathArray
	Include PathArray

	//RsyncOptions []string
	//TarOptions   []string
	//CpioOptions  []string
	//CpOptions    []string
	//SftpOptions  []string

	_Valid       bool
	Method      *TypeCopyMethods
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
//func ReflectState(p *ux.State) *ux.State {
//	return (*State)(p)
//}
func ReflectHelperOsCopy(p *TypeOsCopy) *HelperOsCopy {
	return (*HelperOsCopy)(p)
}

func (c *TypeOsCopy) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func NewOsCopy() *TypeOsCopy {
	c := &TypeOsCopy{
		State:        ux.NewState(false),

		Source:       helperPath.NewOsPath(false),
		Destination:  helperPath.NewOsPath(false),

		Exclude: PathArray{},
		Include: PathArray{},

		_Valid:       false,
		Method:      NewCopyMethod(),
	}
	c.State.SetPackage("")
	c.State.SetFunctionCaller()

	return c
}
func (me *TypeOsCopy) EnsureNotNil() {
	if me == nil {
		me = NewOsCopy()
	}
}


func (p *TypeOsCopy) SetSourcePath(path ...string) bool {
	var ok bool

	for range OnlyOnce {
		ok = p.Source.SetPath(path...)
		if !ok {
			break
		}

		if p.Destination.IsValid() {
			p._Valid = true
		}

		ok = true
	}

	return ok
}
func (p *TypeOsCopy) GetSourcePath() string {
	return p.Source.GetPath()
}


func (p *TypeOsCopy) SetDestinationPath(path ...string) bool {
	var ok bool

	for range OnlyOnce {
		ok = p.Destination.SetPath(path...)
		if !ok {
			break
		}

		if p.Source.IsValid() {
			p._Valid = true
		}

		ok = true
	}

	return ok
}
func (p *TypeOsCopy) GetDestinationPath() string {
	return p.Destination.GetPath()
}


func (p *TypeOsCopy) SetExcludePaths(paths ...string) bool {
	return p.Exclude.SetPaths(paths...)
}
func (p *TypeOsCopy) AddExcludePaths(paths ...string) bool {
	return p.Exclude.AddPaths(paths...)
}
func (p *TypeOsCopy) GetExcludePaths() *PathArray {
	return p.Exclude.GetPaths()
}


func (p *TypeOsCopy) SetIncludePaths(paths ...string) bool {
	return p.Include.SetPaths(paths...)
}
func (p *TypeOsCopy) AddIncludePaths(paths ...string) bool {
	return p.Include.AddPaths(paths...)
}
func (p *TypeOsCopy) GetIncludePaths() *PathArray {
	return p.Include.GetPaths()
}


func (p *TypeOsCopy) SetMethodRsync() bool {
	return p.Method.SelectMethod(ConstMethodRsync)
}
func (p *TypeOsCopy) SetMethodTar() bool {
	return p.Method.SelectMethod(ConstMethodTar)
}
func (p *TypeOsCopy) SetMethodCpio() bool {
	return p.Method.SelectMethod(ConstMethodCpio)
}
func (p *TypeOsCopy) SetMethodSftp() bool {
	return p.Method.SelectMethod(ConstMethodSftp)
}
func (p *TypeOsCopy) SetMethodCp() bool {
	return p.Method.SelectMethod(ConstMethodCp)
}


func (p *TypeOsCopy) GetMethodOptions() interface{} {
	return p.Method.GetOptions()
}
func (p *TypeOsCopy) GetMethodName() string {
	return p.Method.GetName()
}
func (p *TypeOsCopy) GetMethodPath() string {
	return p.Method.GetPath()
}
func (p *TypeOsCopy) GetMethodAllowRemote() bool {
	return p.Method.GetAllowRemote()
}
func (p *TypeOsCopy) GetMethodAvailable() bool {
	return p.Method.GetAvailable()
}

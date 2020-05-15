package helperCopy

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"github.com/zloylos/grsync"
)


// @TODO - Look at several other copy options that provide "cloud" based copies.
// @TODO - https://rclone.org/
// @TODO - https://pkg.go.dev/github.com/Redundancy/go-sync?tab=doc
// @TODO - https://github.com/Redundancy/go-sync
// @TODO - https://pkg.go.dev/bitbucket.org/kardianos/rsync?tab=doc
// @TODO - https://github.com/zloylos/grsync


type OsCopyGetter interface {
}


type TypeOsCopy struct {
	State        *ux.State

	Source       *helperPath.TypeOsPath
	Destination  *helperPath.TypeOsPath

	Exclude PathArray
	Include PathArray

	RsyncOptions []string
	TarOptions   []string
	CpioOptions  []string
	CpOptions    []string
	SftpOptions  []string

	_Valid       bool
	Method      *TypeCopyMethods
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}


func NewOsCopy() *TypeOsCopy {
	return &TypeOsCopy{
		State:        ux.New(),

		Source:       helperPath.NewOsPath(),
		Destination:  helperPath.NewOsPath(),

		Exclude: PathArray{},
		Include: PathArray{},

		//RsyncOptions: []string{"-HvaxP", "-n"},
		//TarOptions:   []string{},
		//CpioOptions:  []string{},
		//SftpOptions:  []string{"-rf"},
		//CpOptions:    []string{},

		_Valid:       false,
		Method:      NewCopyMethod(),
	}
}


func (p *TypeOsCopy) SetSourcePath(path ...string) bool {
	var ok bool

	for range only.Once {
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

	for range only.Once {
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
	var ok bool

	for range only.Once {
		if !p.Method.SelectMethod(ConstMethodRsync) {
			break
		}

		//task := grsync.NewTask(
		//	"username@server.com:/source/folder",
		//	"/home/user/destination",
		//	grsync.RsyncOptions{},
		//)
		task := grsync.NewTask(
			p.Source.GetPath(),
			p.Destination.GetPath(),
			p.Method.Selected.Options.(grsync.RsyncOptions),
		)

		ok = true
	}

	return ok
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

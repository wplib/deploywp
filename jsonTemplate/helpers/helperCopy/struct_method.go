package helperCopy

import (
	"github.com/wplib/deploywp/only"
	"github.com/zloylos/grsync"
	"os/exec"
)

const (
	ConstMethodDefault = ConstMethodRsync
	ConstMethodRsync = "rsync"
	ConstMethodTar = "tar"
	ConstMethodCpio = "cpio"
	ConstMethodSftp = "sftp"
	ConstMethodCp = "cp"
)


// GoLang enums - much better than plain old C type enums!
type TypeCopyMethod struct
{
	Name        string
	Path        string
	AllowRemote bool
	Available   bool
	Options     interface{}
}
type TypeCopyMethods struct {
	Selected *TypeCopyMethod
	All      []*TypeCopyMethod
}

func _CopyMethodDefault() *TypeCopyMethod {
	return _CopyMethodRsync()
}
func _CopyMethodRsync() *TypeCopyMethod {
	path, ok := _ExecExists(ConstMethodRsync)

	opts := grsync.RsyncOptions {
		HardLinks:         true,
		Verbose:           true,
		Archive:           true,
		OneFileSystem:     true,
		Progress:          true,
		//RsyncProgramm:     path,
	}

	return &TypeCopyMethod { Name: ConstMethodRsync, Path: path, Available: ok, Options: opts, AllowRemote: true }
}
func _CopyMethodTar() *TypeCopyMethod {
	path, ok := _ExecExists(ConstMethodTar)

	opts := []string{""}

	return &TypeCopyMethod { Name: ConstMethodTar, Path: path, Available: ok, Options: opts, AllowRemote: true }
}
func _CopyMethodCpio() *TypeCopyMethod {
	path, ok := _ExecExists(ConstMethodCpio)

	opts := []string{""}

	return &TypeCopyMethod { Name: ConstMethodCpio, Path: path, Available: ok, Options: opts, AllowRemote: true }
}
func _CopyMethodSftp() *TypeCopyMethod {
	path, ok := _ExecExists(ConstMethodSftp)

	opts := []string{"-rf"}

	return &TypeCopyMethod { Name: ConstMethodSftp, Path: path, Available: ok, Options: opts, AllowRemote: true }
}
func _CopyMethodCp() *TypeCopyMethod {
	path, ok := _ExecExists(ConstMethodCp)

	opts := []string{"-rip"}

	return &TypeCopyMethod { Name: ConstMethodCp, Path: path, Available: ok, Options: opts, AllowRemote: false }
}

func _ExecExists(e string) (string, bool) {
	var path string
	var ok bool

	for range only.Once {
		var err error
		path, err = exec.LookPath(e)
		if err != nil {
			break
		}
		ok = true
	}

	return path, ok
}


func NewCopyMethod() *TypeCopyMethods {
	var ret TypeCopyMethods

	for range only.Once {
		// Set priority of use.
		ret.All = append(ret.All, _CopyMethodRsync())
		ret.All = append(ret.All, _CopyMethodTar())
		ret.All = append(ret.All, _CopyMethodCpio())
		ret.All = append(ret.All, _CopyMethodSftp())
		ret.All = append(ret.All, _CopyMethodCp())

		for _, m := range ret.All {
			if m.Available {
				ret.Selected = m
				break
			}
		}
	}

	return &ret
}


func (p *TypeCopyMethods) GetOptions() interface{} {
	return p.Selected.Options
}
func (p *TypeCopyMethods) GetName() string {
	return p.Selected.Name
}
func (p *TypeCopyMethods) GetPath() string {
	return p.Selected.Path
}
func (p *TypeCopyMethods) GetAllowRemote() bool {
	return p.Selected.AllowRemote
}
func (p *TypeCopyMethods) GetAvailable() bool {
	return p.Selected.Available
}


func (p *TypeCopyMethods) SelectMethod(method string) bool {
	var ok bool

	for range only.Once {
		for _, m := range p.All {
			if m.Name != method {
				continue
			}

			if !m.Available {
				continue
			}

			ok = true
			break
		}
	}

	return ok
}

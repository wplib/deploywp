package helperSystem

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os/exec"
	"strings"
	"syscall"
)

type TypeRsync struct {
	SourcePath string
	DestinationPath string
	ExcludeFiles []string
	//IncludeFiles []string
	Options []string

	Valid bool
	Error error
}

// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func NewRsync(src interface{}, dest interface{}, options interface{}, exclude ...interface{}) *TypeRsync {
	var ret TypeRsync

	for range only.Once {
		s := helperTypes.ReflectString(src)
		if s == nil {
			ret.Error = errors.New("rsync source empty")
			break
		}
		ret.SourcePath = _FileToAbs(*s) + "/"	// Always add a "/" postfix
		if ret.SourcePath == "" {
			ret.Error = errors.New("rsync source empty")
			break
		}

		d := helperTypes.ReflectString(dest)
		if d == nil {
			ret.Error = errors.New("rsync destination empty")
			break
		}
		ret.DestinationPath = _FileToAbs(*d) + "/"	// Always add a "/" postfix
		if ret.DestinationPath == "" {
			ret.Error = errors.New("rsync destination empty")
			break
		}

		o := helperTypes.ReflectString(options)
		switch {
			case o == nil:
				fallthrough
			case *o == "":
				ret.Options = []string{"-HvaxPn"}
			default:
				ret.Options = []string{*o}
		}

		e := helperTypes.ReflectStrings(exclude...)
		if e == nil {
			break
		}
		ret.ExcludeFiles = *e

		for _, es := range ret.ExcludeFiles {
			ret.Options = append(ret.Options, fmt.Sprintf("--exclude='%s'", es))
		}
	}

	return &ret
}

// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func (me *TypeRsync) Run() *TypeExecCommand {
	var ret TypeExecCommand

	for range only.Once {
		opts := []string{}
		opts = append(opts, me.Options...)
		opts = append(opts, me.SourcePath)
		opts = append(opts, me.DestinationPath)
		fmt.Printf("rsync %s\n", strings.Join(opts, " "))

		c := exec.Command("rsync", opts...)

		var out []byte
		out, ret.Error = c.CombinedOutput()
		ret.Output = string(out)

		if ret.Error != nil {
			if exitError, ok := ret.Error.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				ret.Exit = waitStatus.ExitStatus()
			}

			fmt.Printf("%s\n", ret.PrintError())
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		ret.Exit = waitStatus.ExitStatus()

		if ret.Error != nil {
			ret.PrintError()
			break
		}

		ret.Output = ux.SprintfGreen("%s\n", ret.Output)
	}

	return &ret
}

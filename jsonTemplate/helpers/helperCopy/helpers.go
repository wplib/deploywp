// High level helper functions available within templates - file copy.
package helperCopy

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
	"os/exec"
	"strings"
	"syscall"
)


type HelperOsCopy TypeOsCopy

func (c *HelperOsCopy) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Alias of Rsync || Tar || whatever - basically determine what tool to use based on availability.
// @TODO - To be implemented.
// Usage:
//		{{ $copy := CopyFiles }}
func HelperCopyFiles() *HelperOsCopy {
	ret := NewOsCopy()

	for range OnlyOnce {
		ret.State.Clear()
	}

	return (*HelperOsCopy)(ret)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetSourcePath(src ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")

	for range OnlyOnce {
		p := helperTypes.ReflectStrings(src...)
		if p == nil {
			c.State.SetError("%s source empty", c.Method.GetName())
			break
		}
		if !c.Source.SetPath(*p...) {
			c.State.SetError("%s source empty", c.Method.GetName())
			break
		}
		c.State.Clear()
	}

	return c.State
}
func (c *HelperOsCopy) SetSource(dest ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")
	return c.SetSourcePath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetDestinationPath "filename.txt" }}
func (c *HelperOsCopy) SetDestinationPath(dest ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")

	for range OnlyOnce {
		p := helperTypes.ReflectStrings(dest...)
		if p == nil {
			c.State.SetError("%s destination empty", c.Method.GetName())
			break
		}
		if !c.Destination.SetPath(*p...) {
			c.State.SetError("%s destination empty", c.Method.GetName())
			break
		}
		c.State.Clear()
	}

	return c.State
}
func (c *HelperOsCopy) SetTarget(dest ...interface{}) *ux.State {
	return c.SetDestinationPath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetExcludePaths(exclude ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")

	for range OnlyOnce {
		e := helperTypes.ReflectStrings(exclude...)
		if e == nil {
			break
		}
		if !c.Exclude.SetPaths(*e...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return c.State
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetIncludePaths(include ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")

	for range OnlyOnce {
		i := helperTypes.ReflectStrings(include...)
		if i == nil {
			break
		}
		if !c.Include.SetPaths(*i...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return c.State
}


// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func (c *HelperOsCopy) Run() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction("")

	for range OnlyOnce {
		c.State.SetState(c.Source.StatPath())
		if c.State.IsError() {
			break
		}

		opts := []string{}
		//opts = append(opts, c.RsyncOptions...)
		opts = append(opts, c.Source.GetPath())
		opts = append(opts, c.Destination.GetPath())

		cmd := exec.Command("rsync", opts...)

		out, err := cmd.CombinedOutput()
		c.State.SetOutput(out)
		c.State.SetError(err)

		if c.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				c.State.ExitCode = waitStatus.ExitStatus()
			}

			//fmt.Printf("%s\n", ret.PrintError())
			break
		}

		waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		c.State.ExitCode = waitStatus.ExitStatus()

		fmt.Printf("\nrsync %s\n", strings.Join(opts, " "))
		c.State.SetOk("%s", c.State.Output)
	}

	return c.State
}


//// Usage:
////		{{ $copy := CopyFiles }}
////		{{ $state := SetSourcePath "filename.txt" }}
//func (c *HelperOsCopy) SetOptions(src interface{}) *ux.State {
//	for range OnlyOnce {
//		e := helperTypes.ReflectStrings(exclude...)
//		if e == nil {
//			break
//		}
//		ret.ExcludeFiles = *e
//
//		for _, es := range ret.ExcludeFiles {
//			ret.RsyncOptions = append(ret.RsyncOptions, fmt.Sprintf("--exclude='%s'", es))
//		}
//	}
//
//	return (*State)(c.State)
//}
//

// High level helper functions available within templates - file copy.
package helperCopy

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os/exec"
	"strings"
	"syscall"
)


type HelperOsCopy TypeOsCopy


// Alias of Rsync || Tar || whatever - basically determine what tool to use based on availability.
// @TODO - To be implemented.
// Usage:
//		{{ $copy := CopyFiles }}
func HelperCopyFiles() *HelperOsCopy {
	ret := NewOsCopy()

	for range only.Once {
		ret.State.Clear()
	}

	return (*HelperOsCopy)(ret)
}


// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func (me *HelperOsCopy) Run() *State {
	var ret State

	for range only.Once {
		me.State = (*ux.State)(me.Source.StatPath())
		if me.State.IsError() {
			break
		}

		opts := []string{}
		opts = append(opts, me.RsyncOptions...)
		opts = append(opts, me.SourcePath)
		opts = append(opts, me.DestinationPath)

		c := exec.Command("rsync", opts...)

		out, err := c.CombinedOutput()
		ret.Output = string(out)
		ret.SetError(err)

		if ret.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				ret.Exit = waitStatus.ExitStatus()
			}

			//fmt.Printf("%s\n", ret.PrintError())
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		ret.Exit = waitStatus.ExitStatus()

		if ret.IsError() {
			//ret.PrintError()
			break
		}

		fmt.Printf("\nrsync %s\n", strings.Join(opts, " "))
		ret.Output = ux.SprintfGreen("%s\n", ret.Output)
	}

	return &ret
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetSourcePath(src ...interface{}) *State {
	for range only.Once {
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

	return (*State)(c.State)
}
func (c *HelperOsCopy) SetSource(dest ...interface{}) *State {
	return c.SetSourcePath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetDestinationPath "filename.txt" }}
func (c *HelperOsCopy) SetDestinationPath(dest ...interface{}) *State {
	for range only.Once {
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

	return (*State)(c.State)
}
func (c *HelperOsCopy) SetTarget(dest ...interface{}) *State {
	return c.SetDestinationPath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetExcludePaths(exclude ...interface{}) *State {
	for range only.Once {
		e := helperTypes.ReflectStrings(exclude...)
		if e == nil {
			break
		}
		if !c.Exclude.SetPaths(*e...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return (*State)(c.State)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *HelperOsCopy) SetIncludePaths(include ...interface{}) *State {
	for range only.Once {
		i := helperTypes.ReflectStrings(include...)
		if i == nil {
			break
		}
		if !c.Include.SetPaths(*i...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return (*State)(c.State)
}


//// Usage:
////		{{ $copy := CopyFiles }}
////		{{ $state := SetSourcePath "filename.txt" }}
//func (c *HelperOsCopy) SetOptions(src interface{}) *State {
//	for range only.Once {
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



// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
//func HelperRsync(src interface{}, dest interface{}, options interface{}, exclude ...interface{}) *HelperOsCopy {
//	ret := NewOsCopy()
//
//	for range only.Once {
//		s := helperTypes.ReflectString(src)
//		if s == nil {
//			ret.State.SetError("rsync source empty")
//			break
//		}
//		if ret.Source.SetPath(*s) {
//			ret.State.SetError("rsync source empty")
//		}
//
//
//		d := helperTypes.ReflectString(dest)
//		if d == nil {
//			ret.State.SetError("rsync destination empty")
//			break
//		}
//		if ret.Source.SetPath(*s) {
//			ret.State.SetError("rsync destination empty")
//		}
//
//
//		o := helperTypes.ReflectString(options)
//		switch {
//			case o == nil:
//				fallthrough
//			case *o == "":
//				ret.RsyncOptions = []string{"-HvaxPn"}
//			default:
//				ret.RsyncOptions = []string{*o}
//		}
//
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
//	return (*HelperOsCopy)(ret)
//}

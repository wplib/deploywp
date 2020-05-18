package helperCopy

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"github.com/zloylos/grsync"
	"time"
)


// Usage:
//		{{ $return := WriteFile "filename.txt" .Data.Source 0644 }}
func HelperCopyRsync(src interface{}, dest interface{}, exclude ...interface{}) *ux.State {
	c := NewOsCopy()

	for range only.Once {
		s := helperPath.ReflectAbsPath(src)
		if s == nil {
			break
		}
		if !c.Source.SetPath(*s) {
			break
		}
		c.State.SetState(c.Source.StatPath())
		if !c.Source.Exists() {
			//c.State.SetError("src path not found")
			break
		}


		c.Destination.SetOverwriteable()

		d := helperPath.ReflectAbsPath(dest)
		if d == nil {
			break
		}
		if !c.Destination.SetPath(*d) {
			break
		}
		for range only.Once {
			c.State.SetState(c.Destination.StatPath())
			if c.Destination.NotExists() {
				c.State.Clear()
				break
			}
			if c.Destination.CanOverwrite() {
				break
			}
			c.State.SetError("cannot overwrite destination '%s'", c.Destination.GetPath())
		}
		if c.State.IsError() {
			break
		}


		if !c.Method.SelectMethod(ConstMethodRsync) {
			c.State.SetError("rsync method unavailable")
			break
		}


		task := grsync.NewTask(
			c.Source.GetPath(),
			c.Destination.GetPath(),
			c.Method.Selected.Options.(grsync.RsyncOptions),
		)

		loop := true
		go func() {
			ux.Printf("\n")
			for ;loop; {
				state := task.State()
				ux.PrintfGreen(
					"Copy progress: %.2f / rem. %d / tot. %d / sp. %s \n",
					state.Progress,
					state.Remain,
					state.Total,
					state.Speed,
				)
				time.Sleep(time.Second)
			}
			ux.Printf("\n")
		}()

		err := task.Run()
		loop = false
		c.State.SetOutput(task.Log().Stdout)
		//l := task.Log().Stdout
		//fmt.Print("%s\n", l)
		c.State.SetError(err)
		if c.State.IsError() {
			break
		}


		//opts := []string{}
		////opts = append(opts, c.RsyncOptions...)
		//opts = append(opts, c.Source.GetPath())
		//opts = append(opts, c.Destination.GetPath())
		//cmd := exec.Command("rsync", opts...)
		//out, err := cmd.CombinedOutput()
		//c.State.SetOutput(out)
		//c.State.SetError(err)
		//
		//if c.State.IsError() {
		//	if exitError, ok := err.(*exec.ExitError); ok {
		//		waitStatus := exitError.Sys().(syscall.WaitStatus)
		//		c.State.ExitCode = waitStatus.ExitStatus()
		//	}
		//
		//	//fmt.Printf("%s\n", ret.PrintError())
		//	break
		//}
		//waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		//c.State.ExitCode = waitStatus.ExitStatus()

		c.State.SetOk("%s", c.State.Output)
	}

	return c.State
}


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

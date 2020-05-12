package helperSystem

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"strings"
)

var _ helperTypes.TypeExecCommandGetter = (*TypeExecCommand)(nil)
type TypeExecCommand helperTypes.TypeExecCommand


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func HelperExecCommand(cmd ...interface{}) *TypeExecCommand {
	var ret *TypeExecCommand

	for range only.Once {
		ec := helperTypes.ReflectExecCommand(cmd...)
		if ec == nil {
			break
		}
		ecp := TypeExecCommand(*ec)

		ret = ExecCommand(&ecp)
		ret.PrintError()

		////c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		//c := exec.Command(ec.Exe, ec.Args...)
		//
		//var out []byte
		//out, ret.Error = c.CombinedOutput()
		//ret.Output = string(out)
		//
		//if ret.Error != nil {
		//	if exitError, ok := ret.Error.(*exec.ExitError); ok {
		//		waitStatus := exitError.Sys().(syscall.WaitStatus)
		//		ret.Exit = waitStatus.ExitStatus()
		//	}
		//	break
		//}
		//
		//waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		//ret.Exit = waitStatus.ExitStatus()
	}

	return ret
}
// Alias of ExecCommand
func HelperExec(cmd ...interface{}) *TypeExecCommand {
	return HelperExecCommand(cmd...)
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintError }}
func (me *TypeExecCommand) PrintError() string {
	var ret string

	for range only.Once {
		if me.Exit != 0 {
			ret = ux.SprintfRed("ERROR: %s - %s", me.Error, me.Output)
		}
	}

	return ret
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ if $cmd.ParseOutput "%s" "mick" ... }}found string{{ end }}
func (me *TypeExecCommand) ParseOutput(search interface{}, args ...interface{}) bool {
	var ret bool

	for range only.Once {
		sp := helperTypes.ReflectString(search)
		if sp == nil {
			break
		}
		s := fmt.Sprintf(*sp, args...)

		ret = strings.Contains(me.Output, s)
	}

	return ret
}


//// Usage:
////		{{ $cmd := ExecCommand "ps %s" "-eaf" }}
////		{{ if $cmd.IsError }}found error{{ end }}
//func (me *TypeExecCommand) IsError() bool {
//	return me.Error.IsError()
//}
//
//
//// Usage:
////		{{ $cmd := ExecCommand "ps %s" "-eaf" }}
////		{{ if $cmd.IsOk }}OK{{ end }}
//func (me *TypeExecCommand) IsOk() bool {
//	return me.Error.IsOk()
//}


// Usage:
//		{{ OsExit 1 }}
func HelperOsExit(e ...interface{}) bool {
	var ret bool

	for range only.Once {
		value := helperTypes.ReflectInt(e)
		os.Exit(int(*value))
	}

	return ret
}

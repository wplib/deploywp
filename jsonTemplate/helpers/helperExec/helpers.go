package helperExec

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func HelperExecCommand(cmd ...interface{}) *ux.State {
	ret := NewExecCommand()

	for range only.Once {
		ec := ReflectExecCommand(cmd...)
		if ec == nil {
			break
		}
		ret.Exe = ec.Exe
		ret.Args = ec.Args

		ret = ExecCommand(ret)
	}

	return ret.State
}
// Alias of ExecCommand
func HelperExec(cmd ...interface{}) *ux.State {
	return HelperExecCommand(cmd...)
}


func (e *TypeExecCommand) IsNil() bool {
	if e == nil {
		return true
	}
	return false
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintError }}
func (e *TypeExecCommand) PrintError() string {
	return e.State.SprintError()
}


// Usage:
//		{{ $cmd.ExitOnError }}
func (e *TypeExecCommand) ExitOnError() string {
	e.State.ExitOnError()
	return ""
}


// Usage:
//		{{ $cmd.ExitOnWarning }}
func (e *TypeExecCommand) ExitOnWarning() string {
	e.State.ExitOnWarning()
	return ""
}


// Usage:
//		{{ OsExit 1 }}
func HelperOsExit(e ...interface{}) string {
	for range only.Once {
		value := helperTypes.ReflectInt(e)
		ux.Exit(*value)
	}
	return ""
}

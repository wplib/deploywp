package helperExec

import (
	"github.com/wplib/deploywp/ux"
)


type HelperExecCommand TypeExecCommand
func (g *HelperExecCommand) Reflect() *TypeExecCommand {
	return (*TypeExecCommand)(g)
}
func (g *TypeExecCommand) Reflect() *HelperExecCommand {
	return (*HelperExecCommand)(g)
}

func (c *HelperExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func HelperExecCmd(cmd ...interface{}) *ux.State {
	ret := NewExecCommand(false)

	for range OnlyOnce {
		ec := ReflectExecCommand(cmd...)
		if ec == nil {
			break
		}
		ret.exe = ec.exe
		ret.args = ec.args

		ret = execCommand(ret)
	}

	return ret.State
}
// Alias of ExecCommand
func HelperExec(cmd ...interface{}) *ux.State {
	return HelperExecCmd(cmd...)
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
	for range OnlyOnce {
		value := ux.ReflectInt(e)
		ux.Exit(*value)
	}
	return ""
}

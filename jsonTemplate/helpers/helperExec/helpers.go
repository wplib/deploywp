package helperExec

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"strings"
)


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func HelperExecCommand(cmd ...interface{}) *TypeExecCommand {
	var ret *TypeExecCommand

	for range only.Once {
		ec := ReflectExecCommand(cmd...)
		if ec == nil {
			break
		}
		ecp := *ec

		ret = ExecCommand(&ecp)
	}

	return ret
}
// Alias of ExecCommand
func HelperExec(cmd ...interface{}) *TypeExecCommand {
	return HelperExecCommand(cmd...)
}


func (me *TypeExecCommand) IsNil() bool {
	if me == nil {
		return true
	}
	return false
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintError }}
func (me *TypeExecCommand) PrintError() string {
	return me.State.SprintError()
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintResponse }}
func (me *TypeExecCommand) PrintResponse() string {
	return me.State.Sprint()
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


// Usage:
//		{{ HelperGrep .This.Output "uid=%s" "mick" ... }}
func (me *TypeExecCommand) GrepArray(format interface{}, a ...interface{}) []string {
	var ret []string

	for range only.Once {
		if me.Output == "" {
			break
		}
		ret = helperTypes.HelperGrepArray(me.Output, format, a...)
	}

	return ret
}


// Usage:
//		{{ HelperGrep .This.Output "uid=%s" "mick" ... }}
func (me *TypeExecCommand) Grep(format interface{}, a ...interface{}) string {
	var ret string

	for range only.Once {
		if me.Output == "" {
			break
		}
		ret = helperTypes.HelperGrep(me.Output, format, a...)
	}

	return ret
}


// Usage:
//		{{ $cmd.ExitOnError }}
func (me *TypeExecCommand) ExitOnError() string {
	me.State.ExitOnError()
	return ""
}


// Usage:
//		{{ $cmd.ExitOnWarning }}
func (me *TypeExecCommand) ExitOnWarning() string {
	me.State.ExitOnWarning()
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

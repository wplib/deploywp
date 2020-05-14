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
	var ret string

	for range only.Once {
		ev := ""
		if me.ErrorValue != nil {
			ev = fmt.Sprintf("'%s'", me.ErrorValue)
		}
		switch {
			case me.Exit != 0:
				ret = ux.SprintfError("ERROR: Exit(%d) %v\n%s", me.Exit, ev, me.Output)
			case me.ErrorValue != nil:
				ret = ux.SprintfError("ERROR: %v\n%s", ev, me.Output)
		}
	}

	return ret
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintResponse }}
func (me *TypeExecCommand) PrintResponse() string {
	var ret string

	for range only.Once {
		ev := ""
		if me.ErrorValue != nil {
			ev = fmt.Sprintf("'%s'", me.ErrorValue)
		}
		switch {
			case me.Exit != 0:
				ret = ux.SprintfError("ERROR: Exit(%d) %v\n%s", me.Exit, ev, me.Output)
			case me.ErrorValue != nil:
				ret = ux.SprintfError("ERROR: %v\n%s", ev, me.Output)
			default:
				ret = ux.SprintfOk("%s", me.Output)
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


// Usage:
//		{{ HelperGrep .This.String "uid=%s" "mick" ... }}
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
//		{{ HelperGrep .This.String "uid=%s" "mick" ... }}
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
	var ret string

	switch {
		case me.ErrorValue != nil:
			fallthrough
		case me.Exit > 0:
			_, _ = fmt.Fprintf(os.Stderr,"%s", me.PrintError())
			os.Exit(me.Exit)
	}

	return ret
}


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

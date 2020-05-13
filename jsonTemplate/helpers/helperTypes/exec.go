package helperTypes

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
)

type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	Exe    string
	Args   []string
	Exit   int
	Output string
	Data   interface{}

	TypeError
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	var ec TypeExecCommand

	for range only.Once {
		for i, r := range ref {
			s := *ReflectString(r)

			if i == 0 {
				ec.Exe = s
			} else {
				ec.Args = append(ec.Args, s)
			}
		}
	}

	return &ec
}


func (me *TypeExecCommand) EnsureNotNil() *TypeExecCommand {
	for range only.Once {
		if me != nil {
			break
		}

		me = &TypeExecCommand {
			Exe:    "",
			Args:   nil,
			Exit:   0,
			Output: "",
			Data:   nil,
			TypeError: TypeError{},
		}
	}

	return me
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
//		{{ $cmd.ExitOnError }}
func (me *TypeExecCommand) ExitOnError() string {
	var ret string

	for range only.Once {
		if me.Exit == 0 {
			break
		}

		_, _ = fmt.Fprintf(os.Stderr,"%s", me.PrintError())
		os.Exit(me.Exit)
	}

	return ret
}

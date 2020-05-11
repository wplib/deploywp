package helperSystem

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"os/exec"
	"strings"
	"syscall"
)


type TypeExecCommand struct {
	Exit int
	Error error
	Output string
}


// Usage:
//		{{ OsExit 1 }}
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
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func ExecCommand(cmd ...interface{}) *TypeExecCommand {
	var ret TypeExecCommand

	for range only.Once {
		cmds := helperTypes.ReflectStrings(cmd...)
		if cmds == nil {
			break
		}

		c := exec.Command((*cmds)[0], (*cmds)[1:]...)

		var out []byte
		out, ret.Error = c.CombinedOutput()
		ret.Output = string(out)

		if ret.Error != nil {
			if exitError, ok := ret.Error.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				ret.Exit = waitStatus.ExitStatus()
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		ret.Exit = waitStatus.ExitStatus()
	}

	return &ret
}


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
//		{{ if ExecParseOutput $output "uid=%s" "mick" ... }}YES{{ end }}
func ExecParseOutput(output interface{}, search interface{}, args ...interface{}) bool {
	var ret bool

	for range only.Once {
		sp := helperTypes.ReflectString(search)
		if sp == nil {
			break
		}
		s := fmt.Sprintf(*sp, args...)

		op := helperTypes.ReflectString(output)
		if op == nil {
			break
		}
		p := fmt.Sprintf(*op, args...)

		ret = strings.Contains(p, s)
	}

	return ret
}


// Usage:
//		{{ OsExit 1 }}
func OsExit(e ...interface{}) bool {
	var ret bool

	for range only.Once {
		value := helperTypes.ReflectInt(e)
		os.Exit(int(*value))
	}

	return ret
}

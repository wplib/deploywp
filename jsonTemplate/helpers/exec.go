package helpers

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"os/exec"
	"strings"
)


// Usage: {{ ExecCommand "ps %s" "-eaf" ... }}
type TypeExecCommand struct {
	Error error
	Output string
}
func ExecCommand(cmd ...interface{}) TypeExecCommand {
	var ret TypeExecCommand

	for range only.Once {
		cmds := ReflectStrings(cmd...)
		if cmds == nil {
			break
		}

		var out []byte
		out, ret.Error = exec.Command((*cmds)[0], (*cmds)[1:]...).Output()
		ret.Output = string(out)
	}

	return ret
}


// Usage: {{ if ExecParseOutput $output "uid=%s" "mick" ... }}YES{{ end }}
func ExecParseOutput(output interface{}, search interface{}, args ...interface{}) bool {
	var ret bool

	for range only.Once {
		sp := ReflectString(search)
		if sp == nil {
			break
		}
		s := fmt.Sprintf(*sp, args...)

		op := ReflectString(output)
		if op == nil {
			break
		}
		p := fmt.Sprintf(*op, args...)

		ret = strings.Contains(p, s)
	}

	return ret
}


//func ExecCommand2(cmd interface{}, args ...interface{}) TypeExecCommand {
//	var ret TypeExecCommand
//
//	for range only.Once {
//		cs := ReflectString(cmd)
//		if cs == nil {
//			break
//		}
//
//		ce := fmt.Sprintf(*cs, args...)
//		cea := strings.Fields(ce)
//
//		var out []byte
//		out, ret.Error = exec.Command(cea).Output()
//		ret.Output = string(out)
//	}
//
//	return ret
//}

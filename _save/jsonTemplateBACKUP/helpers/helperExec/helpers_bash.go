package helperExec

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"io/ioutil"
	"log"
	"os"
)


func HelperExecBash(cmd ...interface{}) *ux.State {
	ret := NewExecCommand()

	for range OnlyOnce {
		ret.Exe = "bash"
		ret.Args = []string{"-c"}

		a := helperTypes.ReflectStrings(cmd...)
		ret.Args = append(ret.Args, *a...)

		ret = ExecCommand(ret)
	}

	return ret.State
}


func HelperNewBash(cmd ...interface{}) *HelperExecCommand {
	ret := NewExecCommand()

	for range OnlyOnce {
		ret.Exe = "bash"
		ret.Args = []string{"-c"}

		a := helperTypes.ReflectStrings(cmd...)
		ret.Args = append(ret.Args, *a...)
	}

	return ret.Reflect()
}


// Intent: TBD
//
// Template examples:
//  {{ . }}
//  fmt.Println("Hello")
func (e *HelperExecCommand) AppendCommands(cmd ...interface{}) *ux.State {
	for range OnlyOnce {
		a := helperTypes.ReflectStrings(cmd...)
		e.Args = append(e.Args, *a...)
	}

	return e.State
}
func (e *HelperExecCommand) Append(cmd ...interface{}) *ux.State {
	return e.AppendCommands(cmd...)
}


func (e *HelperExecCommand) Run() *ux.State {
	for range OnlyOnce {
		file, err := ioutil.TempFile("tmp", "deploywp-shell")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())

		e.Exe = "bash"
		e.Args = []string{"-c"}

		e = ExecCommand(e.Reflect()).Reflect()
	}

	return e.State
}

package helperExec

import (
	"github.com/wplib/deploywp/jtc/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
	"io/ioutil"
	"log"
	"os"
)


func HelperExecBash(cmd ...interface{}) *ux.State {
	ret := NewExecCommand(false)

	for range OnlyOnce {
		ret.exe = "bash"
		ret.args = []string{"-c"}

		a := helperTypes.ReflectStrings(cmd...)
		ret.args = append(ret.args, *a...)

		ret = execCommand(ret)
	}

	return ret.State
}


func HelperNewBash(cmd ...interface{}) *HelperExecCommand {
	ret := NewExecCommand(false)

	for range OnlyOnce {
		ret.exe = "bash"
		ret.args = []string{"-c"}

		a := helperTypes.ReflectStrings(cmd...)
		ret.args = append(ret.args, *a...)
	}

	return ret.Reflect()
}


// Intent: TBD
//
// Template examples:
//  {{ . }}
//  fmt.Println("Hello")
func (e *HelperExecCommand) AppendCommands(cmd ...interface{}) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range OnlyOnce {
		a := helperTypes.ReflectStrings(cmd...)
		e.args = append(e.args, *a...)
	}

	return e.State
}
func (e *HelperExecCommand) Append(cmd ...interface{}) *ux.State {
	return e.AppendCommands(cmd...)
}


func (e *HelperExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range OnlyOnce {
		file, err := ioutil.TempFile("tmp", "deploywp-shell")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())

		e.exe = "bash"
		e.args = []string{"-c"}

		e = execCommand(e.Reflect()).Reflect()
	}

	return e.State
}

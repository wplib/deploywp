package helperExec

import (
	"github.com/wplib/deploywp/jtc/helpers/helperPath"
	"github.com/wplib/deploywp/jtc/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
	"os/exec"
	"syscall"
)

const OnlyOnce = "1"

type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	State        *ux.State

	exe    string
	args   []string
	output []byte
}


func NewExecCommand(debugMode bool) *TypeExecCommand {
	ret := &TypeExecCommand {
		exe:  "",
		args: nil,
		//Exit:   0,
		//_Output: "",
		//Data:   nil,
		State: ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	var ec TypeExecCommand

	for range OnlyOnce {
		for i, r := range ref {
			s := *helperTypes.ReflectString(r)

			if i == 0 {
				ec.exe = s
			} else {
				ec.args = append(ec.args, s)
			}
		}
	}

	return &ec
}


func (e *TypeExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (e *TypeExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range OnlyOnce {
		if e.State == nil {
			e.State = ux.NewState(false)
		}

		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(e.exe, e.args...)

		var out []byte
		var err error
		out, err = c.CombinedOutput()
		e.State.SetError(err)
		e.State.SetOutput(out)

		if e.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				e.State.SetExitCode(waitStatus.ExitStatus())
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		e.State.SetExitCode(waitStatus.ExitStatus())
	}

	return e.State
}


func (e *TypeExecCommand) SetPath(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	ep := helperPath.HelperNewPath()
	ep.SetPath(path...)
	e.exe = ep.GetPath()
	return e.State
}


func (e *TypeExecCommand) SetArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}
	e.args = []string{}
	e.AddArgs(args...)
	return e.State
}
func (e *TypeExecCommand) AddArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	e.args = append(e.args, args...)
	return e.State
}
func (e *TypeExecCommand) AppendArgs(args ...string) *ux.State {
	return e.AddArgs(args...)
}


func (e *TypeExecCommand) GetOutput() []byte {
	return e.output
}

func (e *TypeExecCommand) GetArgs() []string {
	return e.args
}

func (e *TypeExecCommand) GetExe() string {
	return e.exe
}


//func (e *TypeExecCommand) EnsureNotNil() *TypeExecCommand {
//	for range OnlyOnce {
//		if e != nil {
//			break
//		}
//
//		e = NewExecCommand()
//	}
//
//	return e
//}

package helperExec

import (
	"github.com/wplib/deploywp/ux"
	"os/exec"
	"syscall"
)


func ExecCommand(ec *TypeExecCommand) *TypeExecCommand {
	if state := ec.IsNil(); state.IsError() {
		return nil
	}

	for range OnlyOnce {
		if ec.State == nil {
			ec.State = ux.NewState(false)
		}

		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(ec.Exe, ec.Args...)

		var out []byte
		var err error
		out, err = c.CombinedOutput()
		ec.State.SetError(err)
		ec.State.SetOutput(out)

		if ec.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				ec.State.SetExitCode(waitStatus.ExitStatus())
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		ec.State.SetExitCode(waitStatus.ExitStatus())
	}

	return ec
}

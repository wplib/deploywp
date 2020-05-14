package helperSystem

import (
	"github.com/wplib/deploywp/only"
	"os/exec"
	"syscall"
)

//var _ helperTypes.TypeOsPathGetter = (*TypeOsPath)(nil)
//var _ helperTypes.TypeExecCommandGetter = (*TypeExecCommand)(nil)


func ExecCommand(ec *TypeExecCommand) *TypeExecCommand {
	for range only.Once {
		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(ec.Exe, ec.Args...)

		var out []byte
		var err error
		out, err = c.CombinedOutput()
		ec.SetError(err)
		ec.Output = string(out)

		if ec.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				ec.Exit = waitStatus.ExitStatus()
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		ec.Exit = waitStatus.ExitStatus()
	}

	return ec
}

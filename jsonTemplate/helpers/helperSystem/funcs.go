package helperSystem

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"os"
	"os/exec"
	"path/filepath"
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


func FileToAbs(f ...string) string {
	var ret string

	for range only.Once {
		ret = filepath.Join(f...)

		if filepath.IsAbs(ret) {
			break
		}

		var err error
		ret, err = filepath.Abs(ret)
		if err != nil {
			ret = ""
			break
		}
	}
	//ret = strings.ReplaceAll(ret, "//", "/")

	return ret
}


func ResolvePath(path ...string) *helperTypes.TypeOsPath {
	var ret helperTypes.TypeOsPath

	for range only.Once {
		ret.Path = FileToAbs(path...)

		var stat os.FileInfo
		stat, ret.ErrorValue = os.Stat(ret.Path)
		//if err != nil {
		//	break
		//}

		if os.IsNotExist(ret.ErrorValue) {
			ret.Exists = false
			break
		}

		ret.Exists = true
		ret.ModTime = stat.ModTime()
		//ret.ModTime = stat.Name()
		ret.Mode = stat.Mode()
		ret.Size = stat.Size()

		if stat.IsDir() {
			ret.IsDir = true
			ret.IsFile = false
			ret.Dirname = ret.Path
			ret.Filename = ""

		} else {
			ret.IsDir = false
			ret.IsFile = true
			ret.Dirname = filepath.Dir(ret.Path)
			ret.Filename = filepath.Base(ret.Path)
		}
	}

	return &ret
}


func ResolveAbsPath(path ...string) *helperTypes.TypeOsPath {
	return ResolvePath(FileToAbs(path...))
}

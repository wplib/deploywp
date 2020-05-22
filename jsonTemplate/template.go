package jsonTemplate

import (
	"fmt"
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/jsonTemplate/helpers"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperExec"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
	"os"
	"text/template"
)


func (at *ArgTemplate) CreateTemplate() (*template.Template, *ux.State) {
	var t *template.Template
	if state := at.IsNil(); state.IsError() {
		return nil, state
	}

	for range OnlyOnce {
		// Define additional template functions.
		tfm := helpers.DiscoverHelpers()
		if tfm.IsNotOk() {
			break
		}

		tfm.Response.(template.FuncMap)["PrintHelpers"] = helpers.PrintHelpers

		t = template.New("JSON").Funcs(tfm.Response.(template.FuncMap))
	}

	return t, at.State
}


func (at *ArgTemplate) ProcessTemplate() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if at.JsonStruct == nil {
			at.JsonStruct = &jsonStruct{}
		}


		//at.JsonStruct.Exec.Cmd, err = os.Executable()
		//at.JsonStruct.Exec.CmdDir = path.Dir(at.JsonStruct.Exec.Cmd)
		//at.JsonStruct.Exec.CmdFile = path.Base(at.JsonStruct.Exec.Cmd)
		//at.JsonStruct.Exec.CmdVersion = at.exec.CmdVersion
		//at.JsonStruct.Exec.FullArgs = at.exec.FullArgs
		//at.JsonStruct.Exec.args = at.exec.args
		e := runtime.NewExec()
		at.JsonStruct.Exec = e
		at.JsonStruct.CreationEpoch = e.TimeStampEpoch()
		at.JsonStruct.CreationDate = e.TimeStampString()
		at.JsonStruct.Env = e.GetEnvMap()

		at.State = at.LoadJsonFile()
		if at.State.IsNotOk() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}

		at.State = at.LoadTemplateFile()
		if at.State.IsNotOk() {
			at.State.SetError("Template file: %s", at.State.GetError())
			break
		}

		at.JsonStruct.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", at.JsonStruct.CreationDate, at.JsonStruct.TemplateFile.Name, at.JsonStruct.JsonFile.Name)
		at.JsonStruct.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"


		if at.OutFile == nil {
			err := at.TemplateRef.Execute(os.Stdout, &at.JsonStruct)
			if err != nil {
				at.State.SetError("Processing error: %s", err)
			}
			break
		}


		at.State = at.CheckOutFile()
		if at.State.IsNotOk() {
			at.State.SetError("Out file: %s", at.State.GetError())
			break
		}


		at.State = at.OutFile.OpenFile()
		if at.State.IsNotOk() {
			at.State.SetError("Out file: %s", at.State.GetError())
			break
		}
		fh := at.OutFile.State.Response.(*os.File)

		err := at.TemplateRef.Execute(fh, &at.JsonStruct)
		if err != nil {
			at.State.SetError("Out file: %s", err)
			break
		}

		at.State = at.OutFile.CloseFile()
		if at.State.IsNotOk() {
			break
		}


		// Are we treating this as a shell script?
		if at.ExecShell {
			outFile := helperPath.HelperNewPath(at.OutFile)
			at.State = outFile.State
			if at.State.IsNotOk() {
				at.State.SetError("Shell script error: %s", err)
				break
			}
			outFile.Chmod(0755)

			exe := helperExec.NewExecCommand(false)
			at.State = exe.State
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", err)
				break
			}

			at.State = exe.SetPath(outFile.GetPath())
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", err)
				break
			}

			at.State = exe.SetArgs()
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", err)
				break
			}

			at.State = exe.Run()
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", err)
				break
			}
			fmt.Printf("# STDOUT from script:\n%s\n", exe.GetExe())
			fmt.Printf("%v\n", exe.GetOutput())

			if at.RemoveFiles {
				outFile.RemoveFile()
				if at.State.IsError() {
					at.State.SetError("Shell script error: %s", err)
					break
				}
			}
		}

		if at.RemoveFiles {
			at.State = at.TemplateFile.RemoveFile()
			if at.State.IsNotOk() {
				break
			}
		}
	}

	return at.State
}


//func (at *ArgTemplate) SetVersion(s string) error {
//	var err error
//
//	for range OnlyOnce {
//		at.Exec.CmdVersion = s
//	}
//
//	return err
//}

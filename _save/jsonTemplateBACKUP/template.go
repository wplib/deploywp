package jsonTemplate

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Template struct {
	//version string
	exec Exec

	jsonFile string
	jsonString string

	createFlag string
	templateFile string
	outFile string

	templateString string
	removeFiles bool
	overWrite bool
	execShell bool

	valid bool
}


func (me *Template) CreateTemplate() (*template.Template, ux.State) {
	var state ux.State
	var t *template.Template

	for range OnlyOnce {
		var tfm template.FuncMap
		var err error

		// Define additional template functions.
		tfm, err = helpers.DiscoverHelpers()
		if err != nil {
			break
		}

		tfm["PrintHelpers"] = helpers.PrintHelpers

		t = template.New("JSON").Funcs(tfm)
	}

	return t, state
}


func (me *Template) ProcessTemplate() ux.State {
	var state ux.State
	var err error

	for range OnlyOnce {
		var jsonStr jsonStruct
		//jsonStr.ExecName, err = os.Executable()
		//jsonStr.DirPath = path.Dir(jsonStr.ExecName)
		//jsonStr.ExecVersion = me.version
		//jsonStr.ExecArgs = me.args
		jsonStr.Exec.Cmd, err = os.Executable()
		jsonStr.Exec.CmdDir = path.Dir(jsonStr.Exec.Cmd)
		jsonStr.Exec.CmdFile = path.Base(jsonStr.Exec.Cmd)
		jsonStr.Exec.CmdVersion = me.exec.CmdVersion
		jsonStr.Exec.Args = me.exec.Args

		now := time.Now()
		jsonStr.CreationEpoch = now.Unix()
		jsonStr.CreationDate = now.Format("2006-01-02T15:04:05-0700")
		jsonStr.Env, _ = helperSystem.GetEnv()


		// Pull in JSON data.
		var js []byte
		if me.jsonFile != "" {
			js, err = fileToString(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}
			jsonStr.JsonString = strings.ReplaceAll(string(js), "\n", "")
			jsonStr.JsonString = strings.ReplaceAll(jsonStr.JsonString, "\t", "")

			jsonStr.Json = make(map[string]interface{})
			err = json.Unmarshal(js, &jsonStr.Json)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			err = jsonStr.JsonFile.getPaths(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

		} else if me.jsonString != "" {
			jsonStr.JsonFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}
			js = []byte(me.jsonString)

		} else {
			state.SetError("no json file or string")
			break
		}


		// Pull in template file.
		if me.templateFile != "" {
			err = jsonStr.TemplateFile.getPaths(me.templateFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			var ts []byte
			ts, err = fileToString(me.templateFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}
			me.templateString = string(ts)

		} else if me.templateString != "" {
			jsonStr.TemplateFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}

		} else {
			state.SetError("no template file or string")
			break
		}
		me.templateString = UnescapeString(me.templateString)

		// Check on output file.
		if me.outFile != "" {
			err = jsonStr.OutFile.getPaths(me.outFile)
			if err != nil {
				// break - IGNORE as it shouldn't be there.
			}

		} else {
			jsonStr.OutFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}
		}


		jsonStr.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", jsonStr.CreationDate, jsonStr.TemplateFile.Name, jsonStr.JsonFile.Name)
		jsonStr.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"


		if me.overWrite {
			_, err = os.Stat(me.outFile)
			if os.IsNotExist(err) {
				state.SetError("Processing error: %s", err)
				break
			}
		}

		var t *template.Template
		t, state = me.CreateTemplate()
		t.Option("missingkey=error")

		var tt *template.Template
		tt, err = t.Parse(me.templateString)
		if err != nil {
			state.SetError("Processing error: %s", err)
			break
		}
		tt.Option("missingkey=error")

		if me.outFile == "" {
			err = tt.Execute(os.Stdout, &jsonStr)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

		} else {
			var f *os.File

			f, err = os.Create(me.outFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			err = tt.Execute(f, &jsonStr)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			_ = f.Sync()
			_ = f.Close()
		}

		// Are we treating this as a shell script?
		if me.execShell {
			fn := _FileToAbs(jsonStr.OutFile.Dir, jsonStr.OutFile.Name)
			//fn := fmt.Sprintf("%s/%s", jsonStr.OutFile.Dir, jsonStr.OutFile.Name)

			err = os.Chmod(fn, 0755)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			var out []byte
			out, err = exec.Command(fn).Output()
			fmt.Printf("# STDOUT from script:\n%s\n", out)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			if me.removeFiles {
				err = os.Remove(fn)
				if err != nil {
					state.SetError("Processing error: %s", err)
					break
				}
			}
		}

		if me.removeFiles {
			fn := _FileToAbs(jsonStr.TemplateFile.Dir, jsonStr.TemplateFile.Name)
			//fn := fmt.Sprintf("%s/%s", jsonStr.TemplateFile.Dir, jsonStr.TemplateFile.Name)

			err = os.Remove(fn)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}
		}

	}

	return state
}


func (me *Template) SetVersion(s string) error {
	var err error

	for range OnlyOnce {
		me.exec.CmdVersion = s
	}

	return err
}


func (me *Template) SetArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		me.exec.Args = a
	}

	return err
}
func (me *Template) AddArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		me.exec.Args = append(me.exec.Args, a...)
	}

	return err
}


func (me *Template) SetValid() error {
	var err error

	for range OnlyOnce {
		me.valid = true
	}

	return err
}

func (me *Template) SetInvalid() error {
	var err error

	for range OnlyOnce {
		me.valid = false
	}

	return err
}


func (me *Template) SetJsonFile(s string) error {
	var err error

	for range OnlyOnce {
		s, err = filepath.Abs(s)
		if err != nil {
			break
		}

		// Check JSON file exists.
		_, err = os.Stat(s)
		if os.IsNotExist(err) {
			me.valid = false
			break
		}

		me.jsonFile = s
	}

	return err
}
func (me *Template) GetJsonFile() string {
	return me.jsonFile
}


func (me *Template) SetJsonString(s string) error {
	var err error

	for range OnlyOnce {
		me.jsonString = s
	}

	return err
}

func (me *Template) SetTemplateFile(s string) error {
	var err error

	for range OnlyOnce {
		s, err = filepath.Abs(s)
		if err != nil {
			break
		}

		// Check template file exists.
		_, err = os.Stat(s)
		if os.IsNotExist(err) {
			me.valid = false
			break
		}

		me.templateFile = s
	}

	return err
}
func (me *Template) GetTemplateFile() string {
	return me.templateFile
}


func (me *Template) SetTemplateString(s string) error {
	var err error

	for range OnlyOnce {
		me.templateString = s
	}

	return err
}


func UnescapeString(s string) string {

	// \a	Alert or bell
	// \b	Backspace
	// \\	Backslash
	// \t	Horizontal tab
	// \n	Line feed or newline
	// \f	Form feed
	// \r	Carriage return
	// \v	Vertical tab
	// \'	Single quote (only in rune literals)
	// \"	Double quote (only in string literals)

	s = strings.ReplaceAll(s, `\a`, "\a")
	s = strings.ReplaceAll(s, `\b`, "\b")
	s = strings.ReplaceAll(s, `\\`, "\\")
	s = strings.ReplaceAll(s, `\t`, "\t")
	s = strings.ReplaceAll(s, `\n`, "\n")
	s = strings.ReplaceAll(s, `\f`, "\f")
	s = strings.ReplaceAll(s, `\r`, "\r")
	s = strings.ReplaceAll(s, `\v`, "\v")
	s = strings.ReplaceAll(s, `\'`, `'`)
	s = strings.ReplaceAll(s, `\"`, `"`)

	return s
}


func (me *Template) PrintHelpers() {
	_, _ = fmt.Fprintf(os.Stderr, helpers.PrintHelpers())
}

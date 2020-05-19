package jsonTemplate

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/cmd/runtime"
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
	exec           runtime.Exec

	jsonFile       string
	jsonString     string

	templateFile   string
	templateString string

	createFlag     string
	outFile        string
	removeFiles    bool
	overWrite      bool
	execShell      bool

	JsonStruct     *jsonStruct
	valid          bool
}


func (me *Template) CreateTemplate() (*template.Template, ux.State) {
	var state ux.State
	var t *template.Template

	for range only.Once {
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


func (me *Template) LoadJson() *ux.State {
	state := ux.NewState()
	var err error

	for range only.Once {
		if me.JsonStruct == nil {
			me.JsonStruct = &jsonStruct{}
		}

		me.JsonStruct.Exec.Cmd, err = os.Executable()
		me.JsonStruct.Exec.CmdDir = path.Dir(me.JsonStruct.Exec.Cmd)
		me.JsonStruct.Exec.CmdFile = path.Base(me.JsonStruct.Exec.Cmd)
		me.JsonStruct.Exec.CmdVersion = me.exec.CmdVersion
		me.JsonStruct.Exec.FullArgs = me.exec.FullArgs
		me.JsonStruct.Exec.Args = me.exec.Args

		now := time.Now()
		me.JsonStruct.CreationEpoch = now.Unix()
		me.JsonStruct.CreationDate = now.Format("2006-01-02T15:04:05-0700")
		me.JsonStruct.Env, _ = helperSystem.GetEnv()


		// Pull in JSON data.
		var js []byte
		if me.jsonFile != "" {
			js, err = fileToString(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}
			me.JsonStruct.JsonString = strings.ReplaceAll(string(js), "\n", "")
			me.JsonStruct.JsonString = strings.ReplaceAll(me.JsonStruct.JsonString, "\t", "")

			me.JsonStruct.Json = make(map[string]interface{})
			err = json.Unmarshal(js, &me.JsonStruct.Json)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			err = me.JsonStruct.JsonFile.getPaths(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

		} else {
			state.SetError("no json file or string")
			break
		}

		me.JsonStruct.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", me.JsonStruct.CreationDate, me.JsonStruct.TemplateFile.Name, me.JsonStruct.JsonFile.Name)
		me.JsonStruct.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"
	}

	return state
}


func (me *Template) ProcessTemplate() ux.State {
	var state ux.State
	var err error

	for range only.Once {
		//var jsonStr jsonStruct
		//me.JsonStruct.ExecName, err = os.Executable()
		//me.JsonStruct.DirPath = path.Dir(me.JsonStruct.ExecName)
		//me.JsonStruct.ExecVersion = me.version
		//me.JsonStruct.ExecArgs = me.args
		me.JsonStruct.Exec.Cmd, err = os.Executable()
		me.JsonStruct.Exec.CmdDir = path.Dir(me.JsonStruct.Exec.Cmd)
		me.JsonStruct.Exec.CmdFile = path.Base(me.JsonStruct.Exec.Cmd)
		me.JsonStruct.Exec.CmdVersion = me.exec.CmdVersion
		me.JsonStruct.Exec.FullArgs = me.exec.FullArgs
		me.JsonStruct.Exec.Args = me.exec.Args

		now := time.Now()
		me.JsonStruct.CreationEpoch = now.Unix()
		me.JsonStruct.CreationDate = now.Format("2006-01-02T15:04:05-0700")
		me.JsonStruct.Env, _ = helperSystem.GetEnv()


		// Pull in JSON data.
		var js []byte
		if me.jsonFile != "" {
			js, err = fileToString(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}
			me.JsonStruct.JsonString = strings.ReplaceAll(string(js), "\n", "")
			me.JsonStruct.JsonString = strings.ReplaceAll(me.JsonStruct.JsonString, "\t", "")

			me.JsonStruct.Json = make(map[string]interface{})
			err = json.Unmarshal(js, &me.JsonStruct.Json)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			err = me.JsonStruct.JsonFile.getPaths(me.jsonFile)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

		} else if me.jsonString != "" {
			me.JsonStruct.JsonFile = FileInfo{
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
			err = me.JsonStruct.TemplateFile.getPaths(me.templateFile)
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
			me.JsonStruct.TemplateFile = FileInfo{
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
			err = me.JsonStruct.OutFile.getPaths(me.outFile)
			if err != nil {
				// break - IGNORE as it shouldn't be there.
			}

		} else {
			me.JsonStruct.OutFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}
		}


		me.JsonStruct.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", me.JsonStruct.CreationDate, me.JsonStruct.TemplateFile.Name, me.JsonStruct.JsonFile.Name)
		me.JsonStruct.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"


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
			err = tt.Execute(os.Stdout, &me.JsonStruct)
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

			err = tt.Execute(f, &me.JsonStruct)
			if err != nil {
				state.SetError("Processing error: %s", err)
				break
			}

			_ = f.Sync()
			_ = f.Close()
		}

		// Are we treating this as a shell script?
		if me.execShell {
			fn := _FileToAbs(me.JsonStruct.OutFile.Dir, me.JsonStruct.OutFile.Name)
			//fn := fmt.Sprintf("%s/%s", me.JsonStruct.OutFile.Dir, me.JsonStruct.OutFile.Name)

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
			fn := _FileToAbs(me.JsonStruct.TemplateFile.Dir, me.JsonStruct.TemplateFile.Name)
			//fn := fmt.Sprintf("%s/%s", me.JsonStruct.TemplateFile.Dir, me.JsonStruct.TemplateFile.Name)

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

	for range only.Once {
		me.exec.CmdVersion = s
	}

	return err
}


func (me *Template) SetArgs(a ...string) error {
	var err error

	for range only.Once {
		me.exec.Args = a
	}

	return err
}
func (me *Template) GetArgs() []string {
	return me.exec.Args
}
func (me *Template) AddArgs(a ...string) error {
	var err error

	for range only.Once {
		me.exec.Args = append(me.exec.Args, a...)
	}

	return err
}


func (me *Template) SetFullArgs(a ...string) error {
	var err error

	for range only.Once {
		me.exec.FullArgs = a
	}

	return err
}
func (me *Template) GetFullArgs() []string {
	return me.exec.FullArgs
}
func (me *Template) AddFullArgs(a ...string) error {
	var err error

	for range only.Once {
		me.exec.FullArgs = append(me.exec.FullArgs, a...)
	}

	return err
}


func (me *Template) SetValid() error {
	var err error

	for range only.Once {
		me.valid = true
	}

	return err
}

func (me *Template) SetInvalid() error {
	var err error

	for range only.Once {
		me.valid = false
	}

	return err
}


func (me *Template) SetJsonFile(s string) error {
	var err error

	for range only.Once {
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

	for range only.Once {
		me.jsonString = s
	}

	return err
}

func (me *Template) SetTemplateFile(s string) error {
	var err error

	for range only.Once {
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

	for range only.Once {
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

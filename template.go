package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig"
	"only"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"
)

type Template struct {
	jsonFile *string
	jsonString *string

	createFlag *string
	templateFile *string
	outFile *string

	templateString *string
	removeFiles *bool
	overWrite *bool
	execShell *bool
}

func (me *Template) ProcessTemplate() error {
	var err error

	for range only.Once {
		var jsonStr jsonStruct
		jsonStr.ExecName, err = os.Executable()
		jsonStr.DirPath = path.Dir(jsonStr.ExecName)
		jsonStr.ExecVersion = Version
		now := time.Now()
		jsonStr.CreationEpoch = now.Unix()
		jsonStr.CreationDate = now.Format("2006-01-02T15:04:05-0700")
		jsonStr.Env, _ = getEnv()

		// Pull in JSON data.
		var js []byte
		if *me.jsonFile != "" {
			js, err = fileToString(*me.jsonFile)
			if err != nil {
				break
			}
			jsonStr.JsonString = strings.ReplaceAll(string(js), "\n", "")
			jsonStr.JsonString = strings.ReplaceAll(jsonStr.JsonString, "\t", "")

			jsonStr.Json = make(map[string]interface{})
			err = json.Unmarshal(js, &jsonStr.Json)
			if err != nil {
				break
			}

			err = jsonStr.JsonFile.getPaths(*me.jsonFile)
			if err != nil {
				break
			}

		} else if *me.jsonString != "" {
			jsonStr.JsonFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}
			js = []byte(*me.jsonString)

		} else {
			err = errors.New("no json file or string")
			break
		}

		// Pull in template file.
		if *me.templateFile != "" {
			err = jsonStr.TemplateFile.getPaths(*me.templateFile)
			if err != nil {
				break
			}

			var ts []byte
			ts, err = fileToString(*me.templateFile)
			if err != nil {
				break
			}
			*me.templateString = string(ts)

		} else if *me.templateString != "" {
			jsonStr.TemplateFile = FileInfo{
				Dir:           "",
				Name:          "",
				CreationEpoch: 0,
				CreationDate:  "",
			}

		} else {
			err = errors.New("no template file or string")
			break
		}
		*me.templateString = UnescapeString(*me.templateString)

		// Check on output file.
		if *me.outFile != "" {
			err = jsonStr.OutFile.getPaths(*me.outFile)
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

		if *me.overWrite {
			_, err = os.Stat(*me.outFile)
			if os.IsNotExist(err) {
				break
			}
		}

		// Define additional template functions.
		tfm := sprig.TxtFuncMap()

		// Additional functions.
		tfm["isInt"] = isInt
		tfm["isString"] = isString
		tfm["isSlice"] = isSlice
		tfm["isArray"] = isArray
		tfm["isMap"] = isMap
		tfm["ToUpper"] = ToUpper
		tfm["ToLower"] = ToLower
		tfm["ToString"] = ToString
		tfm["FindInMap"] = FindInMap
		tfm["ReadFile"] = ReadFile
		tfm["PrintEnv"] = PrintEnv

		t := template.New("JSON").Funcs(tfm)

		var tt *template.Template
		tt, err = t.Parse(*me.templateString)
		if err != nil {
			break
		}

		if *me.outFile == "" {
			err = tt.Execute(os.Stdout, &jsonStr)
			if err != nil {
				break
			}

		} else {
			var f *os.File

			f, err = os.Create(*me.outFile)
			if err != nil {
				break
			}

			err = tt.Execute(f, &jsonStr)
			if err != nil {
				break
			}

			_ = f.Sync()
			_ = f.Close()
		}

		// Are we treating this as a shell script?
		if *me.execShell {
			fn := fmt.Sprintf("%s/%s", jsonStr.OutFile.Dir, jsonStr.OutFile.Name)

			err = os.Chmod(fn, 0755)
			if err != nil {
				break
			}

			var out []byte
			out, err = exec.Command(fn).Output()
			fmt.Printf("# STDOUT from script:\n%s\n", out)
			if err != nil {
				break
			}

			if *me.removeFiles {
				err = os.Remove(fn)
				if err != nil {
					break
				}
			}
		}

		if *me.removeFiles {
			fn := fmt.Sprintf("%s/%s", jsonStr.TemplateFile.Dir, jsonStr.TemplateFile.Name)

			err = os.Remove(fn)
			if err != nil {
				break
			}
		}

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
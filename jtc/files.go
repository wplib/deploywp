package jtc

import (
	"encoding/json"
	"github.com/wplib/deploywp/jtc/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
	"strings"
	"text/template"
)


func (at *ArgTemplate) SetJsonFile(s string) *ux.State {

	for range OnlyOnce {
		if at.JsonFile == nil {
			at.JsonFile = helperPath.HelperNewPath(s)
			if at.JsonFile.State.IsNotOk() {
				break
			}
		}

		at.JsonFile.SetPath(s)

		if at.OverWrite {
			at.JsonFile.SetOverwriteable()
		}

		at.State = at.JsonFile.StatPath()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) LoadJsonFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if at.JsonStruct == nil {
			at.State.SetError("Json structure is nil")
			break
		}


		// Stat JSON file.
		if at.JsonFile.IsNotValid() {
			at.State = at.JsonFile.StatPath()
		}
		if at.JsonFile.NotExists() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		if at.State.IsNotOk() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		at.JsonStruct.JsonFile.SetFileInfo(at.JsonFile)


		// Read JSON data from file.
		at.State = at.JsonFile.ReadFile()
		if at.State.IsNotOk() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		at.JsonString = at.JsonFile.GetContentString()
		at.JsonStruct.JsonString = at.JsonString
		at.JsonStruct.JsonString = strings.ReplaceAll(at.JsonStruct.JsonString, "\n", "")
		at.JsonStruct.JsonString = strings.ReplaceAll(at.JsonStruct.JsonString, "\t", "")


		// Process JSON string.
		js := []byte(at.JsonString)
		at.JsonStruct.Json = make(map[string]interface{})
		err := json.Unmarshal(js, &at.JsonStruct.Json)
		if err != nil {
			at.State.SetError("Processing error: %s", err)
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) SetTemplateFile(s string) *ux.State {

	for range OnlyOnce {
		if at.TemplateFile == nil {
			at.TemplateFile = helperPath.HelperNewPath(s)
			if at.TemplateFile.State.IsNotOk() {
				break
			}
		}

		at.TemplateFile.SetPath(s)
		if at.OverWrite {
			at.TemplateFile.SetOverwriteable()
		}

		at.State = at.TemplateFile.StatPath()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) LoadTemplateFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if at.JsonStruct == nil {
			at.State.SetError("Json structure is nil")
			break
		}


		// Stat JSON file.
		if at.TemplateFile.IsNotValid() {
			at.State = at.TemplateFile.StatPath()
		}
		if at.TemplateFile.NotExists() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		if at.State.IsNotOk() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		at.JsonStruct.TemplateFile.SetFileInfo(at.TemplateFile)


		// Read JSON data from file.
		at.State = at.TemplateFile.ReadFile()
		if at.State.IsNotOk() {
			at.State.SetError("Json file: %s", at.State.GetError())
			break
		}
		at.TemplateString = at.TemplateFile.GetContentString()


		// Create template instance.
		var t *template.Template
		t, at.State = at.CreateTemplate()
		t.Option("missingkey=error")


		// Do it again - may have to perform recursion here.
		var err error
		at.TemplateRef, err = t.Parse(at.TemplateString)
		if err != nil {
			at.State.SetError("Processing error: %s", err)
			break
		}
		at.TemplateRef.Option("missingkey=error")
	}

	return at.State
}


func (at *ArgTemplate) SetOutFile(s string) *ux.State {

	for range OnlyOnce {
		if at.OutFile == nil {
			at.State.Clear()
			// Special case.
			//outFile := helperPath.HelperNewPath(at.OutFile)
			break
		}

		at.OutFile.SetPath(s)
		if at.OutFile.State.IsNotOk() {
			break
		}
		if at.OverWrite {
			at.OutFile.SetOverwriteable()
		}

		at.State = at.OutFile.StatPath()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *ArgTemplate) CheckOutFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if at.OutFile == nil {
			at.State.Clear()
			// Special case.
			//outFile := helperPath.HelperNewPath(at.OutFile)
			break
		}


		// Stat JSON file.
		if at.OutFile.IsNotValid() {
			at.State = at.OutFile.StatPath()
		}
		//if at.OutFile.NotExists() {
		//	at.State.SetError("Json file: %s", at.State.GetError())
		//	break
		//}
		//if at.State.IsNotOk() {
		//	at.State.SetError("Json file: %s", at.State.GetError())
		//	break
		//}
		at.JsonStruct.OutFile.SetFileInfo(at.OutFile)
	}

	return at.State
}

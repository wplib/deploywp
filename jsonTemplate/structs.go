package jsonTemplate

import (
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
)

type jsonStruct struct {
	//ExecName    string
	//ExecVersion string
	Exec         runtime.Exec

	TemplateFile FileInfo
	JsonFile     FileInfo
	OutFile      FileInfo
	Env          helperSystem.Environment

	JsonString      string
	CreationEpoch   int64
	CreationDate    string
	CreationInfo    string
	CreationWarning string

	Json map[string]interface{}
}

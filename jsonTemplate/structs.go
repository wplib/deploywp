package jsonTemplate

import "github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"

type jsonStruct struct {
	DirPath string
	ExecName string
	ExecVersion string
	TemplateFile FileInfo
	JsonFile FileInfo
	OutFile FileInfo
	Env helperSystem.Environment

	JsonString string
	CreationEpoch int64
	CreationDate string
	CreationInfo string
	CreationWarning string

	Json map[string]interface{}
}


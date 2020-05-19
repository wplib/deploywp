package jsonTemplate

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"strings"
)

type jsonStruct struct {
	//ExecName    string
	//ExecVersion string
	Exec    Exec

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

type Exec struct {
	CmdVersion string
	Cmd        string
	CmdDir     string
	CmdFile    string
	Args       ExecArgs
}

type ExecArgs []string

func (me *ExecArgs) ToString() string {
	return strings.Join(*me, " ")
}

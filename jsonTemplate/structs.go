package jsonTemplate

import (
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
	"text/template"
)

const OnlyOnce = "1"


type ArgTemplate struct {
	Exec           *runtime.Exec

	JsonFile       *helperPath.TypeOsPath
	JsonString     string

	TemplateFile   *helperPath.TypeOsPath
	TemplateString string

	OutFile        *helperPath.TypeOsPath

	CreateFlag     string
	RemoveFiles    bool
	OverWrite      bool
	ExecShell      bool

	JsonStruct     *jsonStruct
	TemplateRef    *template.Template

	State          *ux.State
	valid          bool
}


func NewArgTemplate() *ArgTemplate {

	//p := ArgTemplate{
	//	State:        ux.NewState(false),
	//	Url:          "",
	//	Base:         helperPath.NewOsPath(false),
	//	GitConfig:    nil,
	//	GitOptions:   nil,
	//	skipDirCheck: false,
	//	client:       nil,
	//	repository:   nil,
	//	Cmd:          nil,
	//}

	p := ArgTemplate{
		Exec:           runtime.NewExec(),
		JsonFile:       helperPath.NewOsPath(false),
		JsonString:     "",
		TemplateFile:   helperPath.NewOsPath(false),
		TemplateString: "",
		CreateFlag:     "",
		OutFile:        helperPath.NewOsPath(false),
		RemoveFiles:    false,
		OverWrite:      false,
		ExecShell:      false,
		JsonStruct:     nil,
		State:          ux.NewState(false),
		valid:          true,
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return &p
}

func (at *ArgTemplate) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}


func (at *ArgTemplate) SetValid() {
	at.valid = true
}

func (at *ArgTemplate) SetInvalid() {
	at.valid = false
}

package helperExec

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/ux"
)

const OnlyOnce = "1"

type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	State        *ux.State

	Exe          string
	Args         []string
}


func NewExecCommand(debugMode bool) *TypeExecCommand {
	ret := &TypeExecCommand {
		Exe:    "",
		Args:   nil,
		//Exit:   0,
		//_Output: "",
		//Data:   nil,
		State: ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	var ec TypeExecCommand

	for range OnlyOnce {
		for i, r := range ref {
			s := *helperTypes.ReflectString(r)

			if i == 0 {
				ec.Exe = s
			} else {
				ec.Args = append(ec.Args, s)
			}
		}
	}

	return &ec
}


func (e *TypeExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


//func (e *TypeExecCommand) EnsureNotNil() *TypeExecCommand {
//	for range OnlyOnce {
//		if e != nil {
//			break
//		}
//
//		e = NewExecCommand()
//	}
//
//	return e
//}

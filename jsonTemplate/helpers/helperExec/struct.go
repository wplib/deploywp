package helperExec

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)

type TypeExecCommandGetter interface {
}


type TypeExecCommand struct {
	Exe    string
	Args   []string

	Exit   int
	Output string
	Data   interface{}

	State  *ux.State
}


func NewExecCommand() *TypeExecCommand {
	ret := &TypeExecCommand {
		Exe:    "",
		Args:   nil,
		Exit:   0,
		Output: "",
		Data:   nil,
		State: ux.New(),
	}

	return ret
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	var ec TypeExecCommand

	for range only.Once {
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


func (me *TypeExecCommand) EnsureNotNil() *TypeExecCommand {
	for range only.Once {
		if me != nil {
			break
		}

		me = NewExecCommand()
	}

	return me
}

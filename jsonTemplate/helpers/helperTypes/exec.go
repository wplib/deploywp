package helperTypes

import (
	"github.com/wplib/deploywp/only"
)

type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	Exe    string
	Args   []string
	Exit   int
	Output string
	Data   interface{}

	TypeError
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	var ec TypeExecCommand

	for range only.Once {
		for i, r := range ref {
			s := *ReflectString(r)

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

		me = &TypeExecCommand {
			Exe:       "",
			Args:      nil,
			Exit:      0,
			Output:    "",
			Data:      nil,
			TypeError: TypeError{},
		}
	}

	return me
}


func (me *TypeExecCommand) IsNil() bool {
	if me == nil {
		return true
	}
	return false
}

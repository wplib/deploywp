package helperExec

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)

type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	Exe          string
	Args         []string

	Exit         int
	//_Output      string
	//Data         interface{}

	State        *ux.State
}


func NewExecCommand() *TypeExecCommand {
	ret := &TypeExecCommand {
		Exe:    "",
		Args:   nil,
		Exit:   0,
		//_Output: "",
		//Data:   nil,
		State: ux.NewState(),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

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


func (e *TypeExecCommand) EnsureNotNil() *TypeExecCommand {
	for range only.Once {
		if e != nil {
			break
		}

		e = NewExecCommand()
	}

	return e
}


//func (p *TypeExecCommand) LoadContents(data ...interface{}) {
//	for range only.Once {
//		p._Output = ""
//		p._OutputArray = []string{}
//
//		p._AppendContents(data...)
//	}
//}
//
//
//func (p *TypeExecCommand) _AppendContents(data ...interface{}) {
//	for range only.Once {
//		if p._Separator == "" {
//			p._Separator = DefaultSeparator
//		}
//
//		for _, d := range data {
//			//value := reflect.ValueOf(d)
//			//switch value.Kind() {
//			//	case reflect._Output:
//			//		p._Array = append(p._Array, value._Output())
//			//	case reflect.Array:
//			//		p._Array = append(p._Array, d.([]string)...)
//			//	case reflect.Slice:
//			//		p._Array = append(p._Array, d.([]string)...)
//			//}
//
//			var sa []string
//			switch d.(type) {
//				case []string:
//					for _, s := range d.([]string) {
//						sa = append(sa, strings.Split(s, p._Separator)...)
//					}
//				case string:
//					sa = append(sa, strings.Split(d.(string), p._Separator)...)
//			}
//
//			p._OutputArray = append(p._OutputArray, sa...)
//		}
//	}
//}
//
//
//func (p *TypeExecCommand) GetContentString() string {
//	if p._Separator == "" {
//		p._Separator = DefaultSeparator
//	}
//
//	return strings.Join(p._OutputArray, p._Separator)
//}
//
//
//func (p *TypeExecCommand) GetContentArray() []string {
//	return p._OutputArray
//}
//
//
//func (p *TypeExecCommand) SetSeparator(separator string) {
//	for range only.Once {
//		p._Separator = separator
//		p._OutputArray = strings.Split(p._Output, p._Separator)
//	}
//}
//
//
//func (p *TypeExecCommand) GetSeparator() string {
//	return p._Separator
//}

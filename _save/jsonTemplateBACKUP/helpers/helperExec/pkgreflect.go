// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package helperExec

import "reflect"

var Types = map[string]reflect.Type{
	"HelperExecCommand": reflect.TypeOf((*HelperExecCommand)(nil)).Elem(),
	"TypeExecCommand": reflect.TypeOf((*TypeExecCommand)(nil)).Elem(),
	"TypeExecCommandGetter": reflect.TypeOf((*TypeExecCommandGetter)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"ExecCommand": reflect.ValueOf(ExecCommand),
	"HelperExec": reflect.ValueOf(HelperExec),
	"HelperExecBash": reflect.ValueOf(HelperExecBash),
	"HelperNewBash": reflect.ValueOf(HelperNewBash),
	"HelperOsExit": reflect.ValueOf(HelperOsExit),
	"NewExecCommand": reflect.ValueOf(NewExecCommand),
	"ReflectExecCommand": reflect.ValueOf(ReflectExecCommand),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
}


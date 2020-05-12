// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package helperTypes

import "reflect"

var Types = map[string]reflect.Type{
	"TypeError": reflect.TypeOf((*TypeError)(nil)).Elem(),
	"TypeErrorGetter": reflect.TypeOf((*TypeErrorGetter)(nil)).Elem(),
	"TypeExecCommand": reflect.TypeOf((*TypeExecCommand)(nil)).Elem(),
	"TypeExecCommandGetter": reflect.TypeOf((*TypeExecCommandGetter)(nil)).Elem(),
	"TypeGenericString": reflect.TypeOf((*TypeGenericString)(nil)).Elem(),
	"TypeGenericStringArray": reflect.TypeOf((*TypeGenericStringArray)(nil)).Elem(),
	"TypeOsPath": reflect.TypeOf((*TypeOsPath)(nil)).Elem(),
	"TypeOsPathGetter": reflect.TypeOf((*TypeOsPathGetter)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"HelperContains": reflect.ValueOf(HelperContains),
	"HelperFindInMap": reflect.ValueOf(HelperFindInMap),
	"HelperIsArray": reflect.ValueOf(HelperIsArray),
	"HelperIsInt": reflect.ValueOf(HelperIsInt),
	"HelperIsMap": reflect.ValueOf(HelperIsMap),
	"HelperIsSlice": reflect.ValueOf(HelperIsSlice),
	"HelperIsString": reflect.ValueOf(HelperIsString),
	"HelperSprintf": reflect.ValueOf(HelperSprintf),
	"HelperToLower": reflect.ValueOf(HelperToLower),
	"HelperToString": reflect.ValueOf(HelperToString),
	"HelperToUpper": reflect.ValueOf(HelperToUpper),
	"ReflectByteArray": reflect.ValueOf(ReflectByteArray),
	"ReflectExecCommand": reflect.ValueOf(ReflectExecCommand),
	"ReflectFileMode": reflect.ValueOf(ReflectFileMode),
	"ReflectInt": reflect.ValueOf(ReflectInt),
	"ReflectPath": reflect.ValueOf(ReflectPath),
	"ReflectString": reflect.ValueOf(ReflectString),
	"ReflectStrings": reflect.ValueOf(ReflectStrings),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
}


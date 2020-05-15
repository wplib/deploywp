// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package helperPath

import "reflect"

var Types = map[string]reflect.Type{
	"HelperOsPath": reflect.TypeOf((*HelperOsPath)(nil)).Elem(),
	"OsPathGetter": reflect.TypeOf((*OsPathGetter)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"TypeOsPath": reflect.TypeOf((*TypeOsPath)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"HelperChdir": reflect.ValueOf(HelperChdir),
	"HelperChmod": reflect.ValueOf(HelperChmod),
	"HelperGetCwd": reflect.ValueOf(HelperGetCwd),
	"HelperIsCwd": reflect.ValueOf(HelperIsCwd),
	"HelperReadFile": reflect.ValueOf(HelperReadFile),
	"HelperWriteFile": reflect.ValueOf(HelperWriteFile),
	"NewOsPath": reflect.ValueOf(NewOsPath),
	"ReflectAbsPath": reflect.ValueOf(ReflectAbsPath),
	"ReflectFileMode": reflect.ValueOf(ReflectFileMode),
	"ReflectPath": reflect.ValueOf(ReflectPath),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"DefaultSeparator": reflect.ValueOf(DefaultSeparator),
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
}


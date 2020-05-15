// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package helperCopy

import "reflect"

var Types = map[string]reflect.Type{
	"HelperOsCopy": reflect.TypeOf((*HelperOsCopy)(nil)).Elem(),
	"OsCopyGetter": reflect.TypeOf((*OsCopyGetter)(nil)).Elem(),
	"PathArray": reflect.TypeOf((*PathArray)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"TypeCopyMethod": reflect.TypeOf((*TypeCopyMethod)(nil)).Elem(),
	"TypeCopyMethods": reflect.TypeOf((*TypeCopyMethods)(nil)).Elem(),
	"TypeOsCopy": reflect.TypeOf((*TypeOsCopy)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"HelperCopyFiles": reflect.ValueOf(HelperCopyFiles),
	"NewCopyMethod": reflect.ValueOf(NewCopyMethod),
	"NewOsCopy": reflect.ValueOf(NewOsCopy),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"ConstMethodCp": reflect.ValueOf(ConstMethodCp),
	"ConstMethodCpio": reflect.ValueOf(ConstMethodCpio),
	"ConstMethodDefault": reflect.ValueOf(ConstMethodDefault),
	"ConstMethodRsync": reflect.ValueOf(ConstMethodRsync),
	"ConstMethodSftp": reflect.ValueOf(ConstMethodSftp),
	"ConstMethodTar": reflect.ValueOf(ConstMethodTar),
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
}


// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package helperSystem

import "reflect"

var Types = map[string]reflect.Type{
	"Environment": reflect.TypeOf((*Environment)(nil)).Elem(),
	"Prompt": reflect.TypeOf((*Prompt)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"GetEnv": reflect.ValueOf(GetEnv),
	"HelperUserPrompt": reflect.ValueOf(HelperUserPrompt),
	"HelperUserPromptBool": reflect.ValueOf(HelperUserPromptBool),
	"HelperUserPromptHidden": reflect.ValueOf(HelperUserPromptHidden),
	"PrintEnv": reflect.ValueOf(PrintEnv),
	"UserPrompt": reflect.ValueOf(UserPrompt),
	"UserPromptHidden": reflect.ValueOf(UserPromptHidden),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
	"OnlyOnce": reflect.ValueOf(OnlyOnce),
}


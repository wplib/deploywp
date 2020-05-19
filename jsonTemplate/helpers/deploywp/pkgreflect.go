// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package deploywp

import "reflect"

var Types = map[string]reflect.Type{
	"Build": reflect.TypeOf((*Build)(nil)).Elem(),
	"Defaults": reflect.TypeOf((*Defaults)(nil)).Elem(),
	"DefaultsPaths": reflect.TypeOf((*DefaultsPaths)(nil)).Elem(),
	"DefaultsRepository": reflect.TypeOf((*DefaultsRepository)(nil)).Elem(),
	"DeployWpGetter": reflect.TypeOf((*DeployWpGetter)(nil)).Elem(),
	"Files": reflect.TypeOf((*Files)(nil)).Elem(),
	"FilesArray": reflect.TypeOf((*FilesArray)(nil)).Elem(),
	"HelperDeployWp": reflect.TypeOf((*HelperDeployWp)(nil)).Elem(),
	"Host": reflect.TypeOf((*Host)(nil)).Elem(),
	"Hosts": reflect.TypeOf((*Hosts)(nil)).Elem(),
	"Meta": reflect.TypeOf((*Meta)(nil)).Elem(),
	"Paths": reflect.TypeOf((*Paths)(nil)).Elem(),
	"Provider": reflect.TypeOf((*Provider)(nil)).Elem(),
	"Providers": reflect.TypeOf((*Providers)(nil)).Elem(),
	"Repository": reflect.TypeOf((*Repository)(nil)).Elem(),
	"Revision": reflect.TypeOf((*Revision)(nil)).Elem(),
	"Runtime": reflect.TypeOf((*Runtime)(nil)).Elem(),
	"Source": reflect.TypeOf((*Source)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"String": reflect.TypeOf((*String)(nil)).Elem(),
	"Target": reflect.TypeOf((*Target)(nil)).Elem(),
	"TargetRevision": reflect.TypeOf((*TargetRevision)(nil)).Elem(),
	"TargetRevisions": reflect.TypeOf((*TargetRevisions)(nil)).Elem(),
	"TypeDeployWp": reflect.TypeOf((*TypeDeployWp)(nil)).Elem(),
	"URL": reflect.TypeOf((*URL)(nil)).Elem(),
	"Wordpress": reflect.TypeOf((*Wordpress)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"HelperLoadDeployWp": reflect.ValueOf(HelperLoadDeployWp),
	"NewHost": reflect.ValueOf(NewHost),
	"NewJsonFile": reflect.ValueOf(NewJsonFile),
	"ReflectDeployWp": reflect.ValueOf(ReflectDeployWp),
}

var Variables = map[string]reflect.Value{
	"GetHelpers": reflect.ValueOf(&GetHelpers),
}

var Consts = map[string]reflect.Value{
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
	"TargetActionCopy": reflect.ValueOf(TargetActionCopy),
	"TargetActionDelete": reflect.ValueOf(TargetActionDelete),
	"TargetActionExclude": reflect.ValueOf(TargetActionExclude),
	"TargetActionKeep": reflect.ValueOf(TargetActionKeep),
}


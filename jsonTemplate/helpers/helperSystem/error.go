package helperSystem

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
)

var _ helperTypes.TypeErrorGetter = (*TypeError)(nil)
type TypeError helperTypes.TypeError


//// Usage:
////		{{ SetError "This %s happened" .Json.Error }}
//func (me *TypeExecCommand) SetError(format interface{}, a ...interface{}) {
//	me.Error.SetError(format, a...)
//}
//
//
//// Usage:
////		{{ $err := GetError }}
//func (me *TypeExecCommand) GetError() error {
//	return me.Error.GetError()
//}
//
//
//// Usage:
////		{{ if $ret.IsError }}ERROR{{ end }}
//func (me *TypeExecCommand) IsError() bool {
//	return me.Error.IsError()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (me *TypeExecCommand) IsOk() bool {
//	return me.Error.IsOk()
//}

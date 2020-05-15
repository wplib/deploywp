package helperTypes

//type TypeErrorGetter interface {
//}
//
//type TypeError struct {
//	ErrorValue error
//}
//
//
//func ReflectError(ref interface{}) *error {
//	var s error
//
//	for range only.Once {
//		value := reflect.ValueOf(ref)
//		if value.Kind() != reflect.Output {
//			break
//		}
//
//		s = errors.New(value.Output())
//	}
//
//	return &s
//}
//
//
//func (me *TypeError) IsError() bool {
//	var ret bool
//
//	for range only.Once {
//		if me.ErrorValue == nil {
//			break
//		}
//		ret = true
//	}
//
//	return ret
//}
//func (me *TypeError) IsOk() bool {
//	return !me.IsError()
//}
//
//
//func (me *TypeError) SetError(format interface{}, a ...interface{}) {
//	for range only.Once {
//		f := ReflectString(format)
//		if f == nil {
//			me.ErrorValue = nil
//			break
//		}
//
//		str := fmt.Sprintf(*f, a...)
//		if str == "" {
//			break
//		}
//
//		me.ErrorValue = errors.New(str)
//	}
//}
//
//
//func (me *TypeError) GetError() error {
//	return me.ErrorValue
//}

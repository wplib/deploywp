package easy

import (
	"fmt"
	"reflect"
)

type Reflector struct {
	*Value
	typ      reflect.Type
	kind     reflect.Kind
	TypeName ReadableName
	KindName ReadableName
}

//@TODO UNCOMMENT ONCE IT IS IN USE
//func NewReflector(item interface{}) *Reflector {
//	r := Reflector{
//		Value: NewValue(item),
//	}
//	r.SetType(reflect.TypeOf(item))
//	return &r
//}

func (me *Reflector) SetType(t reflect.Type) *Reflector {
	me.typ = t
	me.kind = me.typ.Kind()
	me.TypeName = me.typ.Name()
	me.KindName = me.kind.String()
	return me
}

func (me *Reflector) String() string {
	return fmt.Sprintf("%s/%s: %s",
		me.TypeName,
		me.KindName,
		me.Value.String(),
	)
}

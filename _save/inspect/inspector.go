package inspect

import (
	"fmt"
	"reflect"
)

type Inspector struct {
	*Value
	typ      reflect.Type
	kind     reflect.Kind
	TypeName ReadableName
	KindName ReadableName
}

func NewInspector(item interface{}) *Inspector {
	i := Inspector{
		Value: NewValue(item),
	}
	i.SetType(i.value.Type())
	return &i
}

func (me *Inspector) SetType(t reflect.Type) *Inspector {
	me.typ = t
	me.kind = me.typ.Kind()
	me.TypeName = me.typ.Name()
	me.KindName = me.kind.String()
	return me
}

func (me *Inspector) String() string {
	return fmt.Sprintf("%s/%s: %s",
		me.TypeName,
		me.KindName,
		me.Value.String(),
	)
}

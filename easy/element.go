package easy

import "reflect"

type Element struct {
	*Reflector
	reflect.StructField
}

func NewElement(rv reflect.Value, i ...int) *Element {
	r := &Reflector{
		Value: NewValue(rv),
	}
	t := rv.Type()
	r.SetType(t)
	var sf reflect.StructField
	if rv.Kind() == reflect.Struct {
		if len(i) == 0 {
			i = []int{0}
		}
		if i[0] != InvalidElementIndex {
			sf = t.Field(i[0])
		}
	}
	return &Element{
		Reflector:   r,
		StructField: sf,
	}
}

func (me *Element) GetTag(key string) string {
	return me.Tag.Get(key)
}

func (me *Element) CanSet() bool {
	return me.value.CanSet()
}

func (me *Element) String() string {
	return me.value.String()
}

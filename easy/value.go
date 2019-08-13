package easy

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Value struct {
	item  interface{}
	value reflect.Value
	kind  reflect.Kind
}

func NewValue(item interface{}) *Value {
	v := Value{
		item:  item,
		value: reflect.ValueOf(item),
		kind:  reflect.ValueOf(item).Kind(),
	}
	return &v
}

func (me *Value) SetElement(index interface{}, value interface{}) (e *Value) {
	//me.value.SetMapIndex(key, tv)
	return me
}

func (me *Value) Element(index interface{}) (e *Element) {
	for range Once {
		var err error
		var rv reflect.Value
		var i int = InvalidElementIndex
		v := me.GetNormalized()
		switch v.kind {
		case reflect.Struct:
			i, err = getint(index)
			if err != nil {
				break
			}
			rv = v.value.Field(i)
		case reflect.Slice, reflect.Array:
			var _i int
			_i, err = getint(index)
			if err != nil {
				break
			}
			rv = v.value.Index(_i)
		case reflect.Map:
			// @TODO Make it easier to pass a key for a map
			rv = me.value.MapIndex(index.(reflect.Value))
		default:
			break
		}
		if err == nil {
			e = NewElement(rv, i)
		}
	}

	return e
}

func (me *Value) HasFields() (has bool) {
	v := me.GetNormalized()
	switch v.kind {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		has = me.Count() > 0
	}
	return has
}

// Returns number of elements; struct fields, map elements, slice elements
// Normalizes this metric
func (me *Value) Count() (count int) {
	v := me.GetNormalized()
	switch v.kind {
	case reflect.Struct:
		count = me.value.NumField()
	case reflect.Map, reflect.Slice, reflect.Array:
		count = v.value.Len()
	}
	return count

}

// Returns number of elements; struct fields, map elements, slice elements
// Normalizes this metric
func (me *Value) Keys() (keys []string) {
	v := me.GetNormalized()
	switch v.kind {
	case reflect.Map:
	case reflect.Struct, reflect.Slice, reflect.Array:
		keys = []string{}
		for i := 0; i < me.Count(); i++ {
			keys = append(keys, strconv.Itoa(i))
		}
	default:
		keys = []string{}
	}
	return keys

}

// Returns a non-pointer or unwrapped interface
func (me *Value) GetNormalized() (nv *Value) {
	nv = &Value{}
	*nv = *me
	return nv.Normalize()
}

// Mutates internal value to not be a pointer or interface
func (me *Value) Normalize() *Value {
	switch me.kind {
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		me.value = me.value.Elem()
		me.kind = me.value.Kind()
		me.item = me.value.Interface()
	}
	return me
}

func (me *Value) String() (s string) {
	v := me.GetNormalized()
	switch v.kind {
	case reflect.Struct:
		c := me.Count()
		ss := make([]string, c)
		for i := 0; i < me.Count(); i++ {
			e := me.Element(i)
			if !e.CanSet() {
				continue
			}
			name := e.Name
			value := clean(e.String())
			ss[i] = fmt.Sprintf("%s='%s'", name, value)
		}
		s = strings.Join(ss, " ")

	case reflect.Slice:
		c := me.Count()
		ss := make([]string, c)
		for i := 0; i < c; i++ {
			value := clean(me.Element(i).String())
			s := fmt.Sprintf("[%d]='%s'", i, value)
			ss[i] = s
		}
		s = strings.Join(ss, " ")

	case reflect.Map:
		ss := make([]string, 0)
		for _, key := range me.Keys() {
			value := clean(me.Element(key).String())
			s := fmt.Sprintf("[%s]='%s'", key, value)
			ss = append(ss, s)
		}

	case reflect.String:
		s = me.String()

	default:
		s = me.String()

	}
	return s
}

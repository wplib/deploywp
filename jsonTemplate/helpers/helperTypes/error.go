package helperTypes

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/only"
	"reflect"
)

type TypeErrorGetter interface {
	//IsError() bool
	//IsOk() bool
	//SetError(format interface{}, a ...interface{})
	//GetError() error
}

type TypeError struct {
	ErrorValue error
}


func ReflectError(ref interface{}) *error {
	var s error

	for range only.Once {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.String {
			break
		}

		s = errors.New(value.String())
	}

	return &s
}


func (me *TypeError) IsError() bool {
	var ret bool

	for range only.Once {
		if me.ErrorValue == nil {
			break
		}
		ret = true
	}

	return ret
}
func (me *TypeError) IsOk() bool {
	return !me.IsError()
}


func (me *TypeError) SetError(format interface{}, a ...interface{}) {
	for range only.Once {
		f := ReflectString(format)
		if f == nil {
			me.ErrorValue = nil
			break
		}

		str := fmt.Sprintf(*f, a...)
		if str == "" {
			break
		}

		me.ErrorValue = errors.New(str)
	}
}


func (me *TypeError) GetError() error {
	return me.ErrorValue
}

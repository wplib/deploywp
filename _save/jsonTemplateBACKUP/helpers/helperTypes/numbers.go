package helperTypes

import (
	"github.com/wplib/deploywp/only"
	"reflect"
)


func HelperIsInt(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			return true
		default:
			return false
	}
}


func ReflectInt(ref interface{}) *int64 {
	var s int64

	for range OnlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				s = value.Int()
			default:
				s = 0
		}
	}

	return &s
}

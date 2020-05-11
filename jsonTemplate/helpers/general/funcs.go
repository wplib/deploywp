package general

import (
	"github.com/wplib/deploywp/only"
	"reflect"
)

func ReflectString(ref interface{}) *string {
	var s string

	for range only.Once {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.String {
			break
		}

		s = value.String()
	}

	return &s
}

func ReflectStrings(ref ...interface{}) *[]string {
	var sa []string

	for range only.Once {
		for _, r := range ref {
			sa = append(sa, *ReflectString(r))
		}
	}

	return &sa
}


func ReflectInt(ref interface{}) *int64 {
	var s int64

	for range only.Once {
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

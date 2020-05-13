package helperTypes

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

func ReflectByteArray(ref interface{}) []byte {
	var s []byte

	for range only.Once {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.String {
			break
		}

		f := value.String()
		s = []byte(f)
	}

	return s
}

func ReflectBool(ref interface{}) *bool {
	var s bool

	for range only.Once {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.Bool {
			break
		}

		s = value.Bool()
	}

	return &s
}

func ReflectBoolArg(ref interface{}) bool {
	var s bool

	for range only.Once {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Bool:
				s = value.Bool()
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				v := value.Int()
				if v == 0 {
					s = false
				} else {
					s = true
				}
			case reflect.Float32, reflect.Float64:
				v := value.Float()
				if v == 0 {
					s = false
				} else {
					s = true
				}
		}
	}

	return s
}

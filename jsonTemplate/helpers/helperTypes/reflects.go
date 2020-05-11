package helperTypes

import (
	"github.com/wplib/deploywp/only"
	"os"
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

func ReflectFileMode(ref ...interface{}) *os.FileMode {
	var fm os.FileMode

	for range only.Once {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.Uint32 {
			break
		}

		fm = os.FileMode(value.Uint())
	}

	return &fm
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

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

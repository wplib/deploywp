package helperGitHub

import (
	"reflect"
)


func HelperGitHubGetBranch(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			return true
		default:
			return false
	}
}

package helperTypes

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/only"
	"reflect"
	"strings"
)

type TypeGenericString struct {
	Valid bool
	Error error
	String string
}

type TypeGenericStringArray struct {
	Valid bool
	Error error
	Array []string
}


// Usage:
//		{{ if IsString $output }}YES{{ end }}
func IsString(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.String:
			return true
		default:
			return false
	}
}


// Usage:
//		{{ $str := ToUpper "lowercase" }}
func ToUpper(i interface{}) string {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.String:
			return strings.ToUpper(i.(string))
		default:
			return ""
	}
}


// Usage:
//		{{ $str := ToLower "UPPERCASE" }}
func ToLower(i interface{}) string {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.String:
			return strings.ToLower(i.(string))
		default:
			return ""
	}
}


// Usage:
//		{{ $str := ToString .Json.array }}
func ToString(i interface{}) string {
	ret := ""
	var j []byte
	var err error
	j, err = json.Marshal(i)
	if err == nil {
		ret = string(j)
	}
	return ret
}


// Usage:
//		{{ if ExecParseOutput $output "uid=%s" "mick" ... }}YES{{ end }}
func Contains(s interface{}, substr interface{}) bool {
	var ret bool

	for range only.Once {
		sp := ReflectString(s)
		if sp == nil {
			break
		}

		ssp := ReflectString(substr)
		if ssp == nil {
			break
		}

		ret = strings.Contains(*sp, *ssp)
	}

	return ret
}


// Usage:
//		{{ Sprintf "uid=%s" "mick" ... }}
func Sprintf(format interface{}, a ...interface{}) string {
	var ret string

	for range only.Once {
		p := ReflectString(format)
		if p == nil {
			break
		}
		ret = fmt.Sprintf(*p, a...)
	}

	return ret
}

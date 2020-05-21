package helperTypes

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/ux"
	"regexp"
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
func HelperIsString(i interface{}) bool {
	return ux.IsReflectString(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return true
	//	default:
	//		return false
	//}
}


// Usage:
//		{{ $str := ToUpper "lowercase" }}
func HelperToUpper(i interface{}) string {
	return strings.ToUpper(*ux.ReflectString(i))
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return strings.ToUpper(i.(string))
	//	default:
	//		return ""
	//}
}


// Usage:
//		{{ $str := ToLower "UPPERCASE" }}
func HelperToLower(i interface{}) string {
	return strings.ToLower(*ux.ReflectString(i))
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return strings.ToLower(i.(string))
	//	default:
	//		return ""
	//}
}


// Usage:
//		{{ $str := ToString .Json.array }}
func HelperToString(i interface{}) string {
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
func HelperContains(s interface{}, substr interface{}) bool {
	var ret bool

	for range OnlyOnce {
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
func HelperSprintf(format interface{}, a ...interface{}) string {
	var ret string

	for range OnlyOnce {
		p := ReflectString(format)
		if p == nil {
			break
		}
		ret = fmt.Sprintf(*p, a...)
	}

	return ret
}


// Usage:
//		{{ Grep .This.Output "uid=%s" "mick" ... }}
func HelperGrepArray(str interface{}, format interface{}, a ...interface{}) []string {
	var ret []string

	for range OnlyOnce {
		s := ReflectString(str)
		if s == nil {
			break
		}
		text := strings.Split(*s, "\n")

		f := ReflectString(format)
		if f == nil {
			break
		}

		res := fmt.Sprintf(*f, a...)
		re := regexp.MustCompile(res)

		for _, line := range text {
			if re.MatchString(line) {
				ret = append(ret, line)
			}
		}
	}

	return ret
}


// Usage:
//		{{ Grep .This.Output "uid=%s" "mick" ... }}
func HelperGrep(str interface{}, format interface{}, a ...interface{}) string {
	var ret string

	for range OnlyOnce {
		sa := HelperGrepArray(str, format, a...)

		ret = strings.Join(sa, "\n")
	}

	return ret
}

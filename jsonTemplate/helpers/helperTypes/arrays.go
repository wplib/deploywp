package helperTypes

import (
	"github.com/wplib/deploywp/ux"
	"strings"
)


func HelperIsSlice(i interface{}) bool {
	return ux.IsReflectSlice(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Slice:
	//		return true
	//	default:
	//		return false
	//}
}

func HelperIsArray(i interface{}) bool {
	return ux.IsReflectArray(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Array:
	//		return true
	//	default:
	//		return false
	//}
}

func HelperIsMap(i interface{}) bool {
	return ux.IsReflectMap(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Map:
	//		return true
	//	default:
	//		return false
	//}
}

// FindInMap function.
func HelperFindInMap(i interface{}, n string) interface{} {
	var ret interface{}
	n = strings.TrimPrefix(n, "\"")
	n = strings.TrimSuffix(n, "\"")

	ret, _ = findKey(i, n)

	// v := reflect.ValueOf(i)
	// switch v.Kind() {
	// 	case reflect.Map:
	// 		// for i := 0; i < v.Len(); i++ {
	// 		// 	v.
	// 		// 	fmt.Println(v.Index(i))
	// 		// }
	// 		//
	// 		// for _, m := range v.MapKeys() {
	// 		// 	if m.
	// 		// }
	// }
	return ret
}

func findKey(obj interface{}, key string) (interface{}, bool) {

	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		// key match, return value
		if k == key {
			return v, true
		}

		// if the value is a map, search recursively
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := findKey(m, key); ok {
				return res, true
			}
		}

		// if the value is an array, search recursively
		// from each element
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := findKey(a, key); ok {
					return res,true
				}
			}
		}
	}

	// element not found
	return nil,false
}

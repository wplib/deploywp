package deploywp

import "reflect"


func GetStructTag(ref interface{}, name string) string {
	var tag string
	for range onlyOnce {
		field, ok := reflect.TypeOf(ref).Elem().FieldByName(name)
		if !ok {
			tag = ""
			break
		}
		tag = string(field.Tag.Get("json"))
	}
	return tag
}

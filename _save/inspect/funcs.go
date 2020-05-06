package inspect

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var re1 = regexp.MustCompile(`'`)
var re2 = regexp.MustCompile("\n")

func clean(s string) string {
	s = re1.ReplaceAllString(s, "\\'")
	return re2.ReplaceAllString(s, " ")
}

func getint(i interface{}) (_i int, err error) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		_i = i.(int)
	case reflect.String:
		_i, err = strconv.Atoi(i.(string))
	default:
		err = fmt.Errorf("value in not an int: %+v", i)
	}
	return _i, err
}

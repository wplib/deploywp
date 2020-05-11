package helperSystem

import "io/ioutil"

// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func ReadFile(f string) string {
	var ret string

	// var err error
	var data []byte
	data, _ = ioutil.ReadFile(f)
	ret = string(data)

	return ret
}

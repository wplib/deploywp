package deploywp

import (
	"text/template"
)

var GetHelpers template.FuncMap
func init() {
	// @TODO - Need to perform this automatically.
	// 
	// Turns out we can't do it at runtime, (makes sense)...
	// https://stackoverflow.com/questions/41629293/how-do-i-list-the-public-methods-of-a-package-in-golang
	//
	// So as part of the build process, we perform a static analysis of the helper packages which creates a pkgreflect.go
	// https://github.com/ungerik/pkgreflect

	GetHelpers = make(template.FuncMap)

	for name, fn := range Functions {
		if name == "GetHelpers" {
			continue
		}
		GetHelpers[name] = fn.Interface()
	}
}

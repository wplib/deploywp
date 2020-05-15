package helperCopy

import (
	"strings"
	"text/template"
)

const HelperPrefix = "Helper"
var GetHelpers template.FuncMap
func init() {
	// @TODO - How can we pull in all Helper functions automatically?
	//
	// Turns out we can't do it at runtime, (makes sense)...
	// https://stackoverflow.com/questions/41629293/how-do-i-list-the-public-methods-of-a-package-in-golang
	//
	// So as part of the build process, we perform a static analysis of the helper packages which creates a pkgreflect.go
	// https://github.com/ungerik/pkgreflect
	//
	// This results in all top level functions being imported automatically, (no method functions will be imported).

	GetHelpers = make(template.FuncMap)

	for name, fn := range Functions {
		// Ignore GetHelpers function.
		if name == "GetHelpers" {
			continue
		}

		// Ignore any function that doesn't have a HelperPrefix
		if !strings.HasPrefix(name, HelperPrefix) {
			continue
		}

		// Trim HelperPrefix from function template name.
		name = strings.TrimPrefix(name, HelperPrefix)
		GetHelpers[name] = fn.Interface()
	}
}

package deploywp

import (
	"strings"
	"text/template"
)

const ToolPrefix = "Tool"
var GetTools template.FuncMap
func init() {
	// @TODO - How can we pull in all Tool functions automatically?
	//
	// Turns out we can't do it at runtime, (makes sense)...
	// https://stackoverflow.com/questions/41629293/how-do-i-list-the-public-methods-of-a-package-in-golang
	//
	// So as part of the build process, we perform a static analysis of the helper packages which creates a pkgreflect.go
	// https://github.com/ungerik/pkgreflect
	//
	// This results in all top level functions being imported automatically, (no method functions will be imported).

	GetTools = make(template.FuncMap)

	for name, fn := range Functions {
		// Ignore GetTools function.
		if name == "GetTools" {
			continue
		}

		// Ignore any function that doesn't have a ToolPrefix
		if !strings.HasPrefix(name, ToolPrefix) {
			continue
		}

		// Trim ToolPrefix from function template name.
		name = strings.TrimPrefix(name, ToolPrefix)
		GetTools[name] = fn.Interface()
	}
}

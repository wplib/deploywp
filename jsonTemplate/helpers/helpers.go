package helpers

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/wplib/deploywp/jsonTemplate/helpers/deploywp"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperGithub"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"reflect"
	"runtime"
	"text/template"
)


// This method will auto-import exported helper functions within each helper package.
// You need to run `pkgreflect jsonTemplate/helpers` after code changes.
func DiscoverHelpers() (template.FuncMap, error) {
	var err error
	var tfm template.FuncMap

	for range only.Once {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		for name, fn := range deploywp.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGithub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperSystem.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperTypes.GetHelpers {
			tfm[name] = fn
		}
	}

	return tfm, err
}


// This method will print exported helper functions within each helper package.
// You need to run `pkgreflect jsonTemplate/helpers` after code changes.
func PrintHelpers() error {
	var err error
	var tfm template.FuncMap

	for range only.Once {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		for name, fn := range deploywp.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGithub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperSystem.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperTypes.GetHelpers {
			tfm[name] = fn
		}

		fmt.Printf("List of defined template functions:\n")
		for name, fn := range tfm {
			foo1 := reflect.ValueOf(fn)
			foo2 := foo1.Pointer()
			foo3 := runtime.FuncForPC(foo2)
			foo4 := foo3.Name()

			fmt.Printf("Function: %s (found in %s)\n", name, foo4)
		}
	}

	return err
}

package helpers

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/wplib/deploywp/jsonTemplate/helpers/deploywp"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperGit"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperGitHub"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
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

		for name, fn := range helperGitHub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGit.GetHelpers {
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

		for name, fn := range helperGitHub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGit.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperSystem.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperTypes.GetHelpers {
			tfm[name] = fn
		}

		files := make(Files)
		fmt.Printf("List of defined template functions:\n")
		for name, fn := range tfm {
			helper := _GetFunctionInfo(fn)

			if _, ok := files[helper.File]; !ok {
				files[helper.File] = make(Helpers)
			}

			files[helper.File][name] = *helper
			fmt.Printf("Name[%s]: %s => %s\n", name, helper.Name, helper.Function)
		}

		for fn, fp := range files {
			fmt.Printf("\n# Helper functions within: %s\n", fn)
			for _, hp := range fp {
				fmt.Printf("%s( %s )\t=> ( %s )\n", hp.Name, hp.Args, hp.Return)
				// fmt.Printf("%s\n\tArgs: %s\n\tReturn: %s\n", hp.Function, hp.Args, hp.Return)
			}
		}
	}

	return err
}


const HelperPrefix = "Helper"
func _GetFunctionInfo(i interface{}) *Helper {
	var helper Helper

	for range only.Once {
		ptr := reflect.ValueOf(i).Pointer()
		ptrs := reflect.ValueOf(i).String()
		ptrn := runtime.FuncForPC(ptr).Name()

		helper.Name = filepath.Ext(ptrn)[1:]
		helper.File = ptrn[0:len(ptrn)-len(helper.Name)-1]
		helper.Name = strings.TrimPrefix(helper.Name, HelperPrefix)

		// ptrs == <func(...interface {}) *helperSystem.TypeReadFile Value>
		helper.Function = strings.Replace(ptrs, "<func", helper.Name, -1)
		helper.Function = strings.TrimSuffix(helper.Function, " Value>")
		// helper.Function == (...interface {}) *helperSystem.TypeReadFile

		helper.Args = strings.Split(ptrs, "(")[1]
		helper.Args = strings.Split(helper.Args, ")")[0]

		helper.Return = strings.TrimSuffix(ptrs, " Value>")
		helper.Return = strings.Split(helper.Return, ")")[1]
		helper.Return = strings.TrimSpace(helper.Return)
		helper.Return = strings.TrimPrefix(helper.Return, "(")

		//if helper.Name == "generateCertificateAuthority" {
		//	fmt.Printf(".")
		//}
	}

	return &helper
}

type Helper struct {
	File string
	Function string
	Name string
	Args string
	Return string
}
type Helpers map[string]Helper
type Files map[string]Helpers

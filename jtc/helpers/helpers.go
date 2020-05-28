package helpers

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/jtc/helpers/helperCopy"
	"github.com/wplib/deploywp/jtc/helpers/helperExec"
	"github.com/wplib/deploywp/jtc/helpers/helperGit"
	"github.com/wplib/deploywp/jtc/helpers/helperGitHub"
	"github.com/wplib/deploywp/jtc/helpers/helperPath"
	"github.com/wplib/deploywp/jtc/helpers/helperSystem"
	"github.com/wplib/deploywp/jtc/helpers/helperTypes"
	"github.com/wplib/deploywp/jtc/helpers/helperUx"
	"github.com/wplib/deploywp/ux"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"text/template"
)

const OnlyOnce = "1"


// This method will auto-import exported helper functions within each helper package.
// You need to run `pkgreflect jtc/helpers` after code changes.
func DiscoverHelpers() *ux.State {
	state := ux.NewState(false)
	var tfm template.FuncMap

	for range OnlyOnce {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		for name, fn := range deploywp.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperCopy.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperExec.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGit.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGitHub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperPath.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperSystem.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperTypes.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperUx.GetHelpers {
			tfm[name] = fn
		}
	}

	state.Response = tfm
	return state
}


// This method will print exported helper functions within each helper package.
// You need to run `pkgreflect jtc/helpers` after code changes.
func PrintHelpers() string {
	var ret string

	for range OnlyOnce {
		ret += ux.SprintfCyan("List of defined template functions:\n")

		tfm := DiscoverHelpers()
		if tfm.IsNotOk() {
			ret += ux.SprintfRed("Error discovering helpers.\n")
			break
		}


		files := make(Files)
		for name, fn := range tfm.Response.(template.FuncMap) {
			helper := _GetFunctionInfo(fn)

			if _, ok := files[helper.File]; !ok {
				files[helper.File] = make(Helpers)
			}

			files[helper.File][name] = *helper
			//fmt.Printf("Name[%s]: %s => %s\n", name, helper.Name, helper.Function)
		}

		for fn, fp := range files {
			ret += ux.SprintfWhite("\n# Helper functions within: %s\n", fn)
			for _, hp := range fp {
				ret += fmt.Sprintf("%s( %s )\t=> ( %s )\n",
					ux.SprintfGreen(hp.Name),
					ux.SprintfCyan(hp.Args),
					ux.SprintfYellow(hp.Return),
					)

				// fmt.Printf("%s\n\targs: %s\n\tReturn: %s\n", hp.Function, hp.args, hp.Return)
			}
		}
	}

	return ret
}


const HelperPrefix = "Helper"
func _GetFunctionInfo(i interface{}) *Helper {
	var helper Helper

	for range OnlyOnce {
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

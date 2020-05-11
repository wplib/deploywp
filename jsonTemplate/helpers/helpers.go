package helpers

import (
	"github.com/Masterminds/sprig"
	"github.com/wplib/deploywp/jsonTemplate/helpers/deploywp"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperGithub"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"text/template"
)

func DiscoverHelpers() (template.FuncMap, error) {
	var err error
	var tfm template.FuncMap

	for range only.Once {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		// @TODO - Replace with an Add() method within each helper package to automatically import all helper methods.

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

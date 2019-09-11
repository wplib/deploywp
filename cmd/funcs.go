package cmd

import (
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/jsonfile"
	"os"
)

func MakeDeployWP() *deploywp.DeployWP {
	c := app.Config()
	jf := jsonfile.NewJsonFile(*c)
	err := jf.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dwp := deploywp.NewDeployWP(c)
	dwp.InitializeFromGetter(jf)
	return dwp
}

package deploywp

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/jsonfile"
)

type Deploy struct {
	Config         *cfg.Config
	Branch         Reference
	SourceDir      Dir
	DestinationDir Dir
}

func NewDeploy(config *cfg.Config) *Deploy {
	d := Deploy{
		Config:config,
	}
	return &d
}

func (me *Deploy) Run()  {
	jf := jsonfile.Load(*app.Config())
	b,_ := json.MarshalIndent(jf,"","\t")
	fmt.Printf("%s",string(b))
	return
}
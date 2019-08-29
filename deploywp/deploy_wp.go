package deploywp

import (
	"fmt"
	"github.com/wplib/deploywp/cfg"
	"os"
)

type DeployWP struct {
	Config         *cfg.Config
	Branch         Reference
	SourceDir      Dir
	DestinationDir Dir
	Meta           *Meta
	Site           *Site
	Source         *Source
	Targets        *Targets
}

type Getter interface {
	GetConfig() *cfg.Config
	GetMeta() *Meta
	GetSite() *Site
	GetSource() *Source
	GetTargets() *Targets
}

func NewDeployWP(config *cfg.Config) *DeployWP {
	return &DeployWP{
		Config: config,
	}
}

func (me *DeployWP) InitializeFromGetter(getter Getter) *DeployWP {
	me.Config = getter.GetConfig()
	me.Meta = getter.GetMeta()
	me.Site = getter.GetSite()
	me.Source = getter.GetSource()
	me.Targets = getter.GetTargets()
	return me
}

func (me *DeployWP) Run(loader Loader) {
	err := loader.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	getter, ok := loader.(Getter)
	if !ok {
		fmt.Println("Unable to assert Loader to be a Getter. This is a programming error.")
		os.Exit(1)
	}

	me.InitializeFromGetter(getter)
	fmt.Printf("%#+v", me)
	return
}

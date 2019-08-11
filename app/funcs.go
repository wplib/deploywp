package app

import (
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/util"
)

func Config() *cfg.Config {
	if config == nil {
		config = cfg.LoadConfig(NewSettings())
	}
	return config
}

func Initialize() string {
	if DeployDir[0] == '~'  {
		DeployDir = util.ExpandDir(DeployDir)
	}
	return DeployDir
}


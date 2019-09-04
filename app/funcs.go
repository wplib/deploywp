package app

import (
	"fmt"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/util"
	"os"
)

func Config() *cfg.Config {
	if config == nil {
		config = cfg.LoadConfig(NewSettings())
	}
	return config
}

func Initialize() string {
	if DeployDir[0] == '~' {
		DeployDir = util.ExpandDir(DeployDir)
	}
	return DeployDir
}

func Fail(message string, args ...interface{}) {
	fmt.Printf(message, args...)
	os.Exit(1)
}

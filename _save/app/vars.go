package app

import (
	"fmt"
	"github.com/wplib/deploywp/cfg"
)

//var NoCache bool
var (
	DeployDir      Dir
	DeployFile     Filepath = fmt.Sprintf("%s.json", Slug)
	ConfigDir      Dir      = fmt.Sprintf("~/.config/%s", Slug)
	ConfigFile     Filepath = "config.json"
	AttachDebugger bool     = false

	config *cfg.Config
)

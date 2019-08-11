package app

import (
	"github.com/wplib/deploywp/cfg"
)

var _ cfg.SettingsContainer = (*Settings)(nil)

type Settings struct {
}

func NewSettings() *Settings {
	return &Settings{}
}

func (me *Settings) GetConfigDir() cfg.Dir {
	return ConfigDir
}

func (me *Settings) GetBasefile() cfg.Basefile {
	return ConfigFile
}

func (me *Settings) GetAppSlug() cfg.Slug {
	return Slug
}

func (me *Settings) GetDefaultConfig() cfg.SerializedConfig {
	return `{
	"data_dir":"",
	"cache_dir":"",
	"settings":{}
}`
}


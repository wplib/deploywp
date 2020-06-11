package deploywp

import "github.com/newclarity/scribeHelpers/toolRuntime"


type Defaults struct {
	Paths      DefaultsPaths `json:"paths"`
	Repository DefaultsRepository `json:"repository"`
}
type DefaultsPaths struct {
	WebrootDir string `json:"webroot_dir" mapstructure:"webroot_dir"`
}
type DefaultsRepository struct {
	URL string `json:"url"`
}
func (d *Defaults) New(runtime *toolRuntime.TypeRuntime) *Defaults {
	return &Defaults{
		Paths:      DefaultsPaths{},
		Repository: DefaultsRepository{},
	}
}

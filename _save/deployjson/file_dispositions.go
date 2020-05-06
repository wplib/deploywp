package deployjson

import (
	"github.com/wplib/deploywp/deploywp"
)

var _ deploywp.FileDispositionsGetter = (*FileDispositions)(nil)

type FileDispositions struct {
	Exclude FilepathTemplates `json:"exclude"`
	Delete  FilepathTemplates `json:"delete"`
	Keep    FilepathTemplates `json:"keep"`
	Copy    FilepathTemplates `json:"copy"`
}

func (me FileDispositions) GetExclude() deploywp.FilepathTemplates {
	return deploywp.NewFilepathTemplatesFromGetter(me.Exclude)
}
func (me FileDispositions) GetDelete() deploywp.FilepathTemplates {
	return deploywp.NewFilepathTemplatesFromGetter(me.Delete)
}
func (me FileDispositions) GetKeep() deploywp.FilepathTemplates {
	return deploywp.NewFilepathTemplatesFromGetter(me.Keep)
}
func (me FileDispositions) GetCopy() deploywp.FilepathTemplates {
	return deploywp.NewFilepathTemplatesFromGetter(me.Copy)
}

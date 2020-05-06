package deployjson

import (
	"github.com/wplib/deploywp/deploywp"
)

var _ deploywp.FilepathTemplatesGetter = (*FilepathTemplates)(nil)

type FilepathTemplates []FilepathTemplate
type FilepathTemplate string

func (me FilepathTemplates) Count() int {
	return len(me)
}
func (me FilepathTemplates) Templates() (fpts deploywp.FilepathTemplates) {
	fpts = make(deploywp.FilepathTemplates, me.Count())
	for i, fpt := range me {
		fpts[i] = deploywp.FilepathTemplate(fpt)
	}
	return fpts

}
func (me FilepathTemplate) String() string {
	return string(me)
}

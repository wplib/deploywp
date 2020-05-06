package deploywp

type FilepathTemplates []FilepathTemplate
type FilepathTemplate string

func (me FilepathTemplate) String() string {
	return string(me)
}

type FilepathTemplatesGetter interface {
	Count() int
	Templates() FilepathTemplates
}

func NewFilepathTemplatesFromGetter(fptg FilepathTemplatesGetter) (fpts FilepathTemplates) {
	fpts = make(FilepathTemplates, fptg.Count())
	for i, fpt := range fptg.Templates() {
		fpts[i] = fpt
	}
	return fpts
}

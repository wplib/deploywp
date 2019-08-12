package jsonfile

type FilepathTemplates []FilepathTemplate
type FilepathTemplate string

func (me FilepathTemplate) String() string {
	return string(me)
}

package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
)

var _ deploywp.SourceGetter = (*Source)(nil)

type Source struct {
	WebRoot    Path             `json:"web_root"`
	Paths      WordPressPaths   `json:"wp_paths"`
	Files      FileDispositions `json:"files"`
	Repository Repository       `json:"repository"`
}

func (me Source) GetWebRoot() Path {
	return me.WebRoot

}
func (me Source) GetPaths() *deploywp.WordPressPaths {
	return deploywp.NewWordPressPathsFromGetter(me.Paths)

}
func (me Source) GetFiles() *deploywp.FileDispositions {
	return deploywp.NewFileDispositionsFromGetter(me.Files)

}
func (me Source) GetRepository() *deploywp.Repository {
	return deploywp.NewRepositoryFromGetter(me.Repository)

}

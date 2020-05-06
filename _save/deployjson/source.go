package deployjson

type Source struct {
	Build      Build      `json:"build"`
	Paths      Paths      `json:"paths"`
	Repository Repository `json:"repository"`
	Revision   Revision   `json:"revision"`
}

//var _ deploywp.SourceGetter = (*Source)(nil)

func (me *Source) New() Source {
	me.Build.New()
	me.Paths.New()
	me.Repository.New()
	me.Revision.New()

	return *me
}

func (me *Source) GetBuild() *Build {
	return &me.Build
}

func (me *Source) GetPaths() *Paths {
	return &me.Paths
}

func (me *Source) GetRepository() *Repository {
	return &me.Repository
}

func (me *Source) GetRevision() *Revision {
	return &me.Revision
}



//func (me Source) GetWebRoot() Path {
//	return me.WebRoot
//}
//
//func (me Source) GetPaths() *deploywp.WordPressPaths {
//	return deploywp.NewWordPressPathsFromGetter(me.Paths)
//}
//
//func (me Source) GetFiles() *deploywp.FileDispositions {
//	return deploywp.NewFileDispositionsFromGetter(me.Files)
//}
//
//func (me Source) GetRepository() *deploywp.Repository {
//	return deploywp.NewRepositoryFromGetter(me.Repository)
//}

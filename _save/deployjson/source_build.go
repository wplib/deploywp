package deployjson

type Build struct {
	Empty bool
}

//var _ deploywp.BuildGetter = (*Build)(nil)

func (me *Build) New() Build {
	if me == nil {
		me = &Build {
			Empty: false,
		}
	}

	return *me
}


//func (me Build) GetWebRoot() Path {
//	return me.WebRoot
//}
//
//func (me Build) GetPaths() *deploywp.WordPressPaths {
//	return deploywp.NewWordPressPathsFromGetter(me.Paths)
//}
//
//func (me Build) GetFiles() *deploywp.FileDispositions {
//	return deploywp.NewFileDispositionsFromGetter(me.Files)
//}
//
//func (me Build) GetRepository() *deploywp.Repository {
//	return deploywp.NewRepositoryFromGetter(me.Repository)
//}

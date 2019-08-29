package deploywp

type Source struct {
	WebRoot    Path
	Paths      *WordPressPaths
	Files      *FileDispositions
	Repository *Repository
}

type SourceGetter interface {
	GetWebRoot() Path
	GetPaths() *WordPressPaths
	GetFiles() *FileDispositions
	GetRepository() *Repository
}

func NewSourceFromGetter(sg SourceGetter) (s *Source) {
	return &Source{
		WebRoot:    sg.GetWebRoot(),
		Paths:      sg.GetPaths(),
		Files:      sg.GetFiles(),
		Repository: sg.GetRepository(),
	}
}

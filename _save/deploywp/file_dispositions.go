package deploywp

type FileDispositions struct {
	Exclude FilepathTemplates
	Delete  FilepathTemplates
	Keep    FilepathTemplates
	Copy    FilepathTemplates
}

type FileDispositionsGetter interface {
	GetExclude() FilepathTemplates
	GetDelete() FilepathTemplates
	GetKeep() FilepathTemplates
	GetCopy() FilepathTemplates
}

func NewFileDispositionsFromGetter(fdp FileDispositionsGetter) (fd *FileDispositions) {
	return &FileDispositions{
		Exclude: fdp.GetExclude(),
		Delete:  fdp.GetDelete(),
		Keep:    fdp.GetKeep(),
		Copy:    fdp.GetCopy(),
	}
}

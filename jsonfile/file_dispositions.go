package jsonfile

type FileDispositions struct {
	Exclude FilepathTemplates
	Delete  FilepathTemplates
	Keep    FilepathTemplates
	Copy    FilepathTemplates
}

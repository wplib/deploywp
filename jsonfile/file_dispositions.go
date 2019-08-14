package jsonfile

type FileDispositions struct {
	Exclude FilepathTemplates `json:"exclude"`
	Delete  FilepathTemplates `json:"delete"`
	Keep    FilepathTemplates `json:"keep"`
	Copy    FilepathTemplates `json:"copy"`
}

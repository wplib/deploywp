package deploywp

type Meta struct {
	Schema     Version     `json:"schema"`
	Repository *Repository `json:"repository"`
	Tag        Reference   `json:"tag"`
	Branch     Reference   `json:"branch"`
	Commit     Reference   `json:"commit"`
}

type MetaGetter interface {
	GetSchema() Version
	GetRepository() *Repository
	GetTag() Reference
	GetBranch() Reference
	GetCommit() Reference
}

func NewMetaFromGetter(mg MetaGetter) *Meta {
	m := Meta{
		Schema:     mg.GetSchema(),
		Repository: mg.GetRepository(),
		Tag:        mg.GetTag(),
		Branch:     mg.GetBranch(),
		Commit:     mg.GetCommit(),
	}
	return &m
}

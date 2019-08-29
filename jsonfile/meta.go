package jsonfile

import "github.com/wplib/deploywp/deploywp"

type Meta struct {
	Schema     Version    `json:"schema"`
	Repository Repository `json:"repository"`
	Tag        Reference  `json:"tag"`
	Branch     Reference  `json:"branch"`
	Commit     Reference  `json:"commit"`
}

func (me Meta) GetSchema() deploywp.Version {
	return me.Schema

}
func (me Meta) GetRepository() *deploywp.Repository {
	return deploywp.NewRepositoryFromGetter(me.Repository)

}
func (me Meta) GetTag() deploywp.Reference {
	return me.Tag

}
func (me Meta) GetBranch() deploywp.Reference {
	return me.Branch

}
func (me Meta) GetCommit() deploywp.Reference {
	return me.Commit
}

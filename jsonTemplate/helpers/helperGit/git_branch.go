package helperGit

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"gopkg.in/src-d/go-git.v4"
	"strings"
)


// Usage:
//		{{- $cmd := $git.GetBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GetBranch() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("symbolic-ref", "--short", "HEAD")
		me.Cmd.Output = strings.TrimSpace(me.Cmd.Output)
		me.Cmd.Data = me.Cmd.Output
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GetTags }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GetTags() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		// git show-ref --tag
		//
		// 	tagrefs, err := r.Tags()
		//	CheckIfError(err)
		//	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		//		fmt.Println(t)
		//		return nil
		//	})

		me.Cmd = me.Exec("log", "-1", "--decorate=short", "--pretty=format:%D")
		if me.Cmd.IsError() {
			break
		}

		var tags []string
		tags = make([]string, 0)
		for _, t := range strings.Split(strings.TrimSpace(me.Cmd.Output), ",") {
			if t[:5] != " tag:" {
				continue
			}
			tags = append(tags, t[6:])
		}

		me.Cmd.Data = tags
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.CreateTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) CreateTag(tag interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.Cmd.SetError("tag is nil")
			break
		}

		me.Cmd = me.Exec("tag", *t)
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = *t
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.RemoveTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) RemoveTag(tag interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.Cmd.SetError("tag is nil")
			break
		}

		me.Cmd = me.Exec("tag", "-d", *t)
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = *t
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.TagExists "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) TagExists(tag interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.Cmd.SetError("tag is nil")
			break
		}

		me.Cmd = me.Exec("tag", "-l", *t)
		if me.Cmd.IsError() {
			break
		}

		if me.Cmd.Output == *t {
			me.Cmd.Data = true
		}
	}

	return me.Cmd
}


func (me *TypeGit) getHandle() (*git.Repository, error) {
	var err error

	for range only.Once {
		if me.repository == nil {
			err = fmt.Errorf("repository handle is nil")
			break
		}
	}

	return me.repository, err
}


// Usage:
//		{{- $cmd := $git.GetTagObject "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GetTagObject(tag interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.Cmd.SetError("tag is nil")
			break
		}

		var h *git.Repository
		h, me.Cmd.Error = me.getHandle()
		if me.Cmd.IsError() {
			break
		}

		var r *Reference
		r, me.Cmd.Error = h.Tag(*t)
		if me.Cmd.IsError() {
			break
		}

		var to *Tag
		to, me.Cmd.Error = h.TagObject(r.Hash())
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = to
	}

	if me.Cmd.IsError() {
		me.Cmd.SetError("unable to access tag object '%v''; %s", tag, me.Cmd.Error)
	}

	return me.Cmd
}

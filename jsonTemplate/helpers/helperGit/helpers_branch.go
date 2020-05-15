package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"strings"
)


// Usage:
//		{{- $cmd := $git.GetBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GetBranch() *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.State = me.Exec("symbolic-ref", "--short", "HEAD").Reflect()
		me.Cmd.Output = strings.TrimSpace(me.Cmd.Output)
		me.Cmd.Data = me.Cmd.Output
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GetTags }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GetTags() *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
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

		me.State = me.Exec("log", "-1", "--decorate=short", "--pretty=format:%D").Reflect()
		if me.State.IsError() {
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

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.CreateTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) CreateTag(tag interface{}) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.State.SetError("tag is nil")
			break
		}

		me.State = me.Exec("tag", *t).Reflect()
		if me.State.IsError() {
			break
		}

		me.Cmd.Data = *t
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.RemoveTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) RemoveTag(tag interface{}) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.State.SetError("tag is nil")
			break
		}

		me.State = me.Exec("tag", "-d", *t).Reflect()
		if me.State.IsError() {
			break
		}

		me.Cmd.Data = *t
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.TagExists "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) TagExists(tag interface{}) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.State.SetError("tag is nil")
			break
		}

		me.State = me.Exec("tag", "-l", *t).Reflect()
		if me.State.IsError() {
			break
		}

		if me.Cmd.Output == *t {
			me.Cmd.Data = true
		}
	}

	return ReflectState(me.State)
}


func (me *HelperGit) getHandle() (*git.Repository, error) {
	var err error

	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}
	}

	return me.repository, err
}


// Usage:
//		{{- $cmd := $git.GetTagObject "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GetTagObject(tag interface{}) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			me.Cmd.SetError("tag is nil")
			break
		}

		var h *git.Repository
		h, me.Cmd.ErrorValue = me.getHandle()
		if me.Cmd.IsError() {
			break
		}

		var r *Reference
		r, me.Cmd.ErrorValue = h.Tag(*t)
		if me.Cmd.IsError() {
			break
		}

		var to *Tag
		to, me.Cmd.ErrorValue = h.TagObject(r.Hash())
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = to
	}

	if me.Cmd.IsError() {
		me.Cmd.SetError("unable to access tag object '%v''; %s", tag, me.Cmd.ErrorValue)
	}

	return ReflectState(me.State)
}

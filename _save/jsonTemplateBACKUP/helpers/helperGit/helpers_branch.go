package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{- $cmd := $git.GetBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GetBranch() *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}
		g.State.SetState(g.Exec("symbolic-ref", "--short", "HEAD"))
		g.State.OutputTrim()
		g.State.Response = g.State.Output
	}
	return g.State
}


// Usage:
//		{{- $cmd := $git.GetTags }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GetTags() *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
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

		g.State.SetSeparator(",")
		//g.State.SetState(g.Exec("log", "-1", "--decorate=short", "--pretty=format:%D"))
		g.State.SetState(g.Exec("tag", "-l"))
		if g.State.IsError() {
			break
		}
		g.State.OutputArrayTrim()

		//var tags []string
		//tags = make([]string, 0)
		//for _, t := range g.State.GetOutputArray() {
		//	if t[:5] != " tag:" {
		//		continue
		//	}
		//	tags = append(tags, t[6:])
		//}
		//g.State.Response = tags
		g.State.Response = g.State.GetOutputArray()
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.CreateTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) CreateTag(tag interface{}) *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			g.State.SetError("tag is nil")
			break
		}

		g.State.SetState(g.Exec("tag", *t))
		if g.State.IsError() {
			break
		}
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.RemoveTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) RemoveTag(tag interface{}) *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			g.State.SetError("tag is nil")
			break
		}

		g.State.SetState(g.Exec("tag", "-d", *t))
		if g.State.IsError() {
			break
		}
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.TagExists "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) TagExists(tag interface{}) *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			g.State.SetError("tag is nil")
			break
		}

		g.State.SetState(g.Exec("tag", "-l", *t))
		if g.State.IsError() {
			break
		}

		if g.State.Output == *t {
			g.State.Response = true
		}
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetTagObject "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GetTagObject(tag interface{}) *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		t := helperTypes.ReflectString(tag)
		if t == nil {
			g.State.SetError("tag is nil")
			break
		}


		//var r *Reference
		//r, g.State.Error = g.repository.Tag(*t)
		r, err := g.repository.Tag(*t)
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		//var to *Tag
		//to, g.State.Error = g.repository.TagObject(r.Hash())
		to, err := g.repository.TagObject(r.Hash())
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		g.State.Response = to
	}

	return g.State
}

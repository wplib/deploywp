package helperGit

import (
	"github.com/wplib/deploywp/jtc/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) Clone(url string, dir ...interface{}) *TypeExecCommand {
func (g *HelperGit) Clone() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		if g.Url == "" {
			g.State.SetError("Git repo URL is empty")
			break
		}

		g.Base.StatPath()
		if g.Base.Exists() {
			g.State.SetError("cannot clone as path %s already exists", g.Base.GetPath())
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}


		ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.GetPath())
		g.skipDirCheck = true
		g.State.SetState(g.Exec(gitCommandClone, g.Url, g.Base.GetPath()))
		g.skipDirCheck = false
	}

	return g.State
}
//func (g *HelperGit) Clone() *ux.State {
//	for range OnlyOnce {
//		if g.Reflect().IsNotOk() {
//			break
//		}
//
//		if url == "" {
//			g.State.SetError("URL is nil")
//			break
//		}
//		g.SetUrl(url)
//
//
//		d := helperPath.ReflectAbsPath(dir...)
//		if d == nil {
//			g.State.SetError("dir is nil")
//			break
//		}
//
//		if !g.Base.SetPath(*d) {
//			g.State.SetError("error setting path to %s", g.Base.GetPath())
//			break
//		}
//
//		g.Base.StatPath()
//		if g.Base.Exists() {
//			g.State.SetError("cannot clone as path %s already exists", g.Base.GetPath())
//			g.Cmd.Exit = 1
//			break
//		}
//
//
//		ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.GetPath())
//		g.skipDirCheck = true
//		g.State.SetState(g.Exec(gitCommandClone, g.Url, g.Base.GetPath()))
//		g.skipDirCheck = false
//	}
//
//	if g.State.IsError() {
//		g.State.SetError("Clone() - %s", g.State.Error)
//	}
//	return g.State
//}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) IsExisting() bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		ok = true
	}

	return ok
}
func (g *HelperGit) IsNotExisting() bool {
	return !g.IsExisting()
}

// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) Open() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		g.State.SetState(g.Exec("rev-parse", "--is-inside-work-tree"))
		if !g.State.OutputEquals("true") {
			if g.State.IsError() {
				g.State.SetError("current directory does not contain a valid .Git repository: %s", g.State.GetError())
				break
			}

			g.State.SetError("current directory does not contain a valid Git repository")
			break
		}

		var err error
		g.repository, err = git.PlainOpen(g.Base.GetPath())
		if err != nil {
			g.State.SetError(err)
			break
		}

		c, _ := g.repository.Config()
		g.Url = c.Remotes["origin"].URLs[0]

		g.State.SetOk("Opened directory %s.\nRemote origin is set to %s\n", g.Base.GetPath(), g.Url)
		g.State.Response = true
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) SetPath(path ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		p := helperPath.ReflectAbsPath(path...)
		if p == nil {
			g.State.SetError("path repo is nil")
			break
		}
		if *p == "" {
			g.State.SetError("path repo is nil")
			break
		}


		if !g.Base.SetPath(*p) {
			g.State.SetError("path repo '%s' cannot be set", *p)
			break
		}

		g.State.SetState(g.Base.StatPath())
		//if state.IsError() {
		//	g.State.SetState(state)
		//	break
		//}

		if g.Base.NotExists() {
			g.State.Clear()
			break
		}
		if g.Base.IsAFile() {
			g.State.SetError("path repo '%s' exists and is a file", *p)
			break
		}
		g.State.SetState(g.Chdir())
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.Data }}{{- end }}
func (g *HelperGit) GetUrl() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		g.State.SetState(g.Exec("config", "--get", "remote.origin.url"))
		if g.State.IsError() {
			break
		}

		g.Url = g.State.Output
		g.State.Response = g.State.Output
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) SetUrl(u Url) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.Url = u
	return g.State
}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) Clone(url interface{}, dir ...interface{}) *TypeExecCommand {
func (g *HelperGit) Remove() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		g.State.SetState(g.Base.StatPath())
		if g.State.IsError() {
			break
		}

		// @TODO - TO BE IMPLEMENTED

		//g.Base.Exists
		//ps := helperSystem.ResolveAbsPath(*d)
		//if ps.IsFile {
		//	break
		//}
		//if ps.IsDir {
		//	if ps.Exists {
		//		g.State.SetError("Repository exists for directory '%s'", ps.Path)
		//		g.Cmd.Exit = 1
		//		break
		//	}
		//}
		//
		//g.SetUrl(*u)
		//g.Base = ps
		//ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.Path)
		//
		//g.skipDirCheck = true
		//g.Cmd = (*helperTypes.TypeExecCommand)(g.Exec(gitCommandClone, g.Url, g.Base.Path))
		//g.skipDirCheck = false
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.Lock }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) Lock() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		g.State = g.GetTagObject(LockTag)
		if g.State.IsError() {
			break
		}

		var to *object.Tag
		to = g.State.Response.(*object.Tag)

		g.State.Response = to.ID()
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GetStatus() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		var wt *git.Worktree
		var err error
		wt, err = g.repository.Worktree()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		var sts git.Status
		sts, err = wt.Status()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		g.State.Response = sts
	}

	return g.State
}

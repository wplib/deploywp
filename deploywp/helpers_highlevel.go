package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{ $cmd := ShowPaths }}
func (e *TypeDeployWp) PrintPaths() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		src := e.GetSourcePaths()
		srcAbs := e.GetSourceAbsPaths()
		ux.PrintfBlue("# SOURCE PATHS:\n")
		ux.PrintfOk("BasePath (abs):    %s\n", srcAbs.GetBasePath())
		ux.PrintfOk("BasePath:          %s\n", src.GetBasePath())
		ux.PrintfOk("WebRootPath:       %s\n", src.GetWebRootPath())
		ux.PrintfOk("ContentPath:       %s\n", src.GetContentPath())
		ux.PrintfOk("CorePath:          %s\n", src.GetCorePath())
		ux.PrintfOk("RootPath:          %s\n", src.GetRootPath())
		ux.PrintfOk("VendorPath:        %s\n", src.GetVendorPath())

		target := e.GetTargetPaths()
		targetAbs := e.GetTargetAbsPaths()
		ux.PrintfBlue("# SOURCE PATHS:\n")
		ux.PrintfOk("BasePath (abs):    %s\n", targetAbs.GetBasePath())
		ux.PrintfOk("BasePath:          %s\n", target.GetBasePath())
		ux.PrintfOk("WebRootPath:       %s\n", target.GetWebRootPath())
		ux.PrintfOk("ContentPath:       %s\n", target.GetContentPath())
		ux.PrintfOk("CorePath:          %s\n", target.GetCorePath())
		ux.PrintfOk("RootPath:          %s\n", target.GetRootPath())
		ux.PrintfOk("VendorPath:        %s\n", target.GetVendorPath())
	}

	return e.State
}


// Usage:
//		{{ $cmd := OpenSourceRepo }}
//		{{ $cmd.ExitOnWarning }}
func (e *TypeDeployWp) OpenSourceRepo() *helperGit.HelperGit {
	gitRef := helperGit.NewGit().Reflect()
	if state := e.IsNil(); state.IsError() {
		return &helperGit.HelperGit{}
	}

	for range OnlyOnce {
		ux.PrintfBlue("# Checking source repository.\n")
		repo := e.Source.GetRepository()
		provider := repo.GetProvider()
		if provider == "" {
			e.State.SetError(".source.repository.provider is nil")
			break
		}

		url := repo.GetUrl()
		if url == "" {
			e.State.SetError(".source.repository.url is nil")
			break
		}

		path := e.Source.AbsPaths.GetBasePath()
		if path == "" {
			e.State.SetError(".source.paths.base_path is nil")
			break
		}

		ux.PrintfBlue("# Source repository details.\n")
		ux.PrintfOk("Provider: '%s'\n", provider)
		ux.PrintfOk("Url:      '%s'\n", url)
		ux.PrintfOk("Path:     '%s'\n", path)

		if !repo.IsGitProvider() {
			e.State.SetWarning(".source.repository.provider '%s' is not supported", provider)
			break
		}

		gitRef = helperGit.HelperNewGit()
		if gitRef.State.IsError() {
			e.State = gitRef.State
			break
		}

		pathRef := helperPath.HelperNewPath(path)
		if pathRef.NotExists() {
			ux.PrintfBlue("# Source repository, doesn't exist - cloning.\n")
			e.State = gitRef.SetPath(path)
			if e.State.IsError() {
				break
			}

			e.State = gitRef.SetUrl(url.ToString())
			if e.State.IsError() {
				break
			}

			e.State = gitRef.Clone()
			if e.State.IsError() {
				break
			}
		}

		ux.PrintfBlue("# Opening source repository.\n")
		e.State = gitRef.SetPath(path)
		if e.State.IsError() {
			break
		}
		e.State = gitRef.Open()
		if e.State.IsError() {
			break
		}

		if gitRef.Url != url.ToString() {
			ux.PrintfWarning("# Source repository URL was changed.\n")
		}

		ux.PrintfOk("# Source repository opened OK.\n")
		e.State.Clear()
	}

	gitRef.State = e.State
	return gitRef
}


// Usage:
//		{{ $cmd := CheckoutSourceRepo }}
//		{{ $cmd.ExitOnWarning }}
func (e *TypeDeployWp) CheckoutSourceRepo(gitRef *helperGit.HelperGit) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if gitRef.IsNotExisting() {
			e.State.SetError("source repository not open")
			break
		}

		ux.PrintfBlue("# Checkout branch/tag from repository.\n")
		refType := e.Source.GetRevisionType()
		if refType == "" {
			e.State.SetError(".source.revision.ref_type is nil")
			break
		}

		refName := e.Source.GetRevisionName()
		if refName == "" {
			e.State.SetError(".source.revision.ref_name is nil")
			break
		}

		ux.PrintfBlue("# Source repository checkout.\n")
		ux.PrintfOk("Type: '%s'\n", refType)
		ux.PrintfOk("Name: '%s'\n", refName)

		ux.PrintfBlue("# Checking if %s '%s' exists.\n", refType, refName)
		if refType == "branch" {
			e.State = gitRef.BranchExists(refName)
			if e.State.IsError() {
				break
			}
		} else {
			e.State = gitRef.TagExists(refName)
			if e.State.IsError() {
				break
			}
		}

		ux.PrintfBlue("# Checking out %s '%s'.\n", refType, refName)
		e.State = gitRef.GitCheckout(refName)
		if e.State.IsError() {
			break
		}

		ux.PrintfOk("# %s '%s' checked out OK.\n", refType, refName)
		e.State.Clear()
	}

	gitRef.State = e.State
	return e.State
}


// Usage:
//		{{ $cmd := OpenTargetRepo }}
//		{{ $cmd.ExitOnWarning }}
func (e *TypeDeployWp) OpenTargetRepo() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		ux.PrintfBlue("# Obtain source repository details.\n")
		e.State = e.ObtainHost()
		if e.State.IsError() {
			break
		}

		//host := e.State.ResponseToString()
		host := e.GetHost(e.State.Response)
		if host.State.IsError() {
			break
		}
		ux.PrintfBlue("# Opening source repository.\n")
		ux.PrintfOk("HostName: '%s'\n", host.HostName)
		ux.PrintfOk("Label:    '%s'\n", host.Label)
		ux.PrintfOk("Provider: '%s'\n", host.Provider)
	}

	return e.State
}


// Usage:
//		{{ $state := ObtainHost }}
//		{{ $state.ExitOnWarning }}
func (e *TypeDeployWp) ObtainHost() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		var host string
		for range OnlyOnce {
			e.State.Clear()

			host = e.Exec.GetArg(0)
			if host != "" {
				break
			}

			host = helperPrompt.HelperUserPrompt("Enter host: ")
			if host != "" {
				break
			}

			e.State.SetError("host is empty")
		}
		if e.State.IsError() {
			break
		}

		e.State.SetOutput(host)
		e.State.Response = host
	}

	return e.State
}

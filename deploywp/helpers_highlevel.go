package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/ux"
)


func (dwp *TypeDeployWp) OpenSourceRepo() *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	for range onlyOnce {
		ux.PrintflnBlue("# Checking source repository.")

		repo := dwp.Source.GetRepository()
		provider := repo.GetProvider()
		if provider == "" {
			dwp.State.SetError(".source.repository.provider is nil")
			break
		}

		if !repo.IsGitProvider() {
			dwp.State.SetWarning(".source.repository.provider '%s' is not supported", provider)
			break
		}

		url := repo.GetUrl()
		if url == "" {
			dwp.State.SetError(".source.repository.url is nil")
			break
		}

		path := dwp.Source.AbsPaths.GetBasePath()
		if path == "" {
			dwp.State.SetError(".source.paths.base_path is nil")
			break
		}

		refType := dwp.Source.GetRevisionType()
		if refType == "" {
			dwp.State.SetError(".source.revision.ref_type is nil")
			break
		}

		refName := dwp.Source.GetRevisionName()
		if refName == "" {
			dwp.State.SetError(".source.revision.ref_name is nil")
			break
		}

		//ux.PrintflnBlue("# Source repository details.")
		//ux.PrintflnOk("Provider: '%s'", provider)
		//ux.PrintflnOk("Url:      '%s'", url)
		//ux.PrintflnOk("Path:     '%s'", path)
		//ux.PrintflnOk("%s:   '%s'", strings.Title(refType), refName)

		gitRef = dwp.OpenRepo(url.String(), path)
		if gitRef.State.IsNotOk() {
			dwp.State = gitRef.State
			break
		}

		dwp.State = dwp.CheckoutRepo(gitRef, refType, refName)
		if gitRef.State.IsNotOk() {
			dwp.State = gitRef.State
			break
		}

		//gitRef = toolGit.New(dwp.Runtime)
		//if gitRef.State.IsError() {
		//	dwp.State = gitRef.State
		//	break
		//}
		//
		//pathRef := toolPath.New(dwp.Runtime)
		//pathRef.SetPath(path)
		//gitRef.State = pathRef.StatPath()
		////if gitRef.State.IsError() {
		////	break
		////}
		//
		//if pathRef.NotExists() {
		//	ux.PrintflnBlue("# Source repository, doesn't exist - cloning.")
		//	dwp.State = gitRef.SetPath(path)
		//	if dwp.State.IsError() {
		//		break
		//	}
		//
		//	dwp.State = gitRef.SetUrl(url.ToString())
		//	if dwp.State.IsError() {
		//		break
		//	}
		//
		//	dwp.State = gitRef.Clone()
		//	if dwp.State.IsError() {
		//		break
		//	}
		//}
		//
		//
		//ux.PrintflnBlue("# Opening source repository.")
		//dwp.State = gitRef.SetPath(path)
		//if dwp.State.IsError() {
		//	break
		//}
		//dwp.State = gitRef.Open()
		//if dwp.State.IsError() {
		//	break
		//}
		//
		//if gitRef.Url != url.ToString() {
		//	ux.PrintfWarning("# Source repository URL was changed.")
		//}
		//ux.PrintflnOk("# Source repository opened OK.")

		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return gitRef
}


func (dwp *TypeDeployWp) OpenTargetRepo() *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	for range onlyOnce {
		ux.PrintflnBlue("# Checking target repository details.")

		dwp.State = dwp.ObtainHost()
		if dwp.State.IsError() {
			break
		}
		host := dwp.Hosts.GetByName(dwp.State.Output)
		if host.state.IsError() {
			break
		}

		path := dwp.GetTargetAbsPaths().GetBasePath()
		if path == "" {
			break
		}

		provider := dwp.Target.GetProviderByName(host.Provider)
		if provider.state.IsError() {
			dwp.State = provider.state
			break
		}

		webRoot := provider.GetWebroot()
		repoUrl := provider.GetRepository()
		revision := dwp.Target.GetRevisionByHost(host.HostName)
		if revision.state.IsError() {
			dwp.State = revision.state
			break
		}

		ux.PrintflnBlue("# Target repository.")
		ux.PrintflnOk("Path: '%s'", path)
		ux.PrintflnOk("HostName: '%s'", host.HostName)
		ux.PrintflnOk("Label:    '%s'", host.Label)
		ux.PrintflnOk("Provider: '%s'", host.Provider)
		ux.PrintflnOk("Repo Url: '%s'", repoUrl)
		ux.PrintflnOk("Web Root: '%s'", webRoot)
		ux.PrintflnOk("Branch:   '%s'", revision.RefName)

		gitRef = dwp.OpenRepo(repoUrl, path)
		if gitRef.State.IsError() {
			dwp.State = gitRef.State
			break
		}

 		dwp.State = dwp.CheckoutRepo(gitRef, "tag", revision.RefName)
		if dwp.State.IsError() {
			break
		}

		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return gitRef
}

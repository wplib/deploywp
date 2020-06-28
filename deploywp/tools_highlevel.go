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

		dwp.State = dwp.PrintRepo(gitRef)
		if dwp.State.IsError() {
			break
		}

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

		hostArg := dwp.GetHost()
		if dwp.State.IsError() {
			break
		}
		host := dwp.Hosts.GetByName(hostArg)
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

		//ux.PrintflnBlue("# Target repository.")
		//ux.PrintflnOk("Path: '%s'", path)
		//ux.PrintflnOk("Repo Url: '%s'", repoUrl)
		//ux.PrintflnOk("Branch:   '%s'", revision.RefName)
		//ux.PrintflnOk("HostName: '%s'", host.HostName)
		//ux.PrintflnOk("Label:    '%s'", host.Label)
		//ux.PrintflnOk("Provider: '%s'", host.Provider)
		//ux.PrintflnOk("Web Root: '%s'", webRoot)

		gitRef = dwp.OpenRepo(repoUrl, path)
		if gitRef.State.IsError() {
			dwp.State = gitRef.State
			break
		}

 		dwp.State = dwp.CheckoutRepo(gitRef, "tag", revision.RefName)
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.PrintRepo(gitRef)
		if dwp.State.IsError() {
			break
		}
		ux.PrintflnOk("HostName: '%s'", host.HostName)
		ux.PrintflnOk("Label:    '%s'", host.Label)
		ux.PrintflnOk("Provider: '%s'", host.Provider)
		ux.PrintflnOk("Web Root: '%s'", webRoot)

		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return gitRef
}

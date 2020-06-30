package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
)


func (dwp *TypeDeployWp) OpenSourceRepo() *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	// dwp.Print.Intent("Opening source repository")
	dwp.Print.Notify(0, "Opening source repository")
	for range onlyOnce {
		gitRef.State.SetOk()
		source := dwp.Source

		repo := source.GetRepository()
		provider := repo.GetProvider()
		if provider == "" {
			gitRef.State.SetError(".source.repository.provider is nil")
			break
		}

		if !repo.IsGitProvider() {
			gitRef.State.SetWarning(".source.repository.provider '%s' is not supported", provider)
			break
		}

		url := repo.GetUrl()
		if url == "" {
			gitRef.State.SetError(".source.repository.url is nil")
			break
		}

		refType := source.GetRevisionType()
		if refType == "" {
			gitRef.State.SetError(".source.revision.ref_type is nil")
			break
		}

		refName := source.GetRevisionName()
		if refName == "" {
			gitRef.State.SetError(".source.revision.ref_name is nil")
			break
		}


		pathRef := toolPath.New(dwp.Runtime)
		if pathRef.State.IsNotOk() {
			gitRef.State = pathRef.State
			break
		}
		pathRef.SetPath(source.AbsPaths.GetBasePath())

		// Check repo exists and clone if not.
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			gitRef = dwp.CloneRepo(url.String(), pathRef)
			if gitRef.State.IsNotOk() {
				break
			}
		}

		gitRef = dwp.OpenRepo(url.String(), pathRef)
		if gitRef.State.IsNotOk() {
			break
		}

		gitRef.State = dwp.CheckoutRepo(gitRef, refType, refName)
		if gitRef.State.IsNotOk() {
			break
		}

		if dwp.Runtime.Verbose {
			gitRef.State = dwp.PrintRepo(gitRef)
			if gitRef.State.IsError() {
				break
			}
			gitRef.State = dwp.PrintSourcePaths()
			if gitRef.State.IsError() {
				break
			}
		}

		gitRef.State.SetOk()
	}
	// dwp.Print.IntentResponse(gitRef.State)
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef
}


func (dwp *TypeDeployWp) OpenDestinationRepo() *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	// dwp.Print.Intent("Opening destination repository")
	dwp.Print.Notify(0, "Opening destination repository")
	for range onlyOnce {
		gitRef.State.SetOk()
		destination := dwp.Destination

		host := destination.GetSelectedHost()
		if host.state.IsNotOk() {
			gitRef.State = host.state
			break
		}


		provider := destination.GetProviderByName(host.Provider)
		if provider.state.IsError() {
			gitRef.State = provider.state
			break
		}

		webRoot := provider.GetWebroot()
		repoUrl := provider.GetRepository()
		revision := destination.GetTargetByHost(host.HostName)
		if revision.state.IsError() {
			gitRef.State = revision.state
			break
		}

		//dwp.Print.Intent("Destination repository")
		//dwp.Print.Ok("Path: '%s'", path)
		//dwp.Print.Ok("Repo Url: '%s'", repoUrl)
		//dwp.Print.Ok("Branch:   '%s'", revision.RefName)
		//dwp.Print.Ok("HostName: '%s'", host.HostName)
		//dwp.Print.Ok("Label:    '%s'", host.Label)
		//dwp.Print.Ok("Provider: '%s'", host.Provider)
		//dwp.Print.Ok("Web Root: '%s'", webRoot)


		pathRef := toolPath.New(dwp.Runtime)
		if pathRef.State.IsNotOk() {
			gitRef.State = pathRef.State
			break
		}
		pathRef.SetPath(destination.GetBasePath())

		// Check repo exists and clone if not.
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			gitRef = dwp.CloneRepo(repoUrl, pathRef)
			if gitRef.State.IsNotOk() {
				break
			}
		}

		gitRef = dwp.OpenRepo(repoUrl, pathRef)
		if gitRef.State.IsError() {
			gitRef.State = gitRef.State
			break
		}

		gitRef.State = dwp.CheckoutRepo(gitRef, "tag", revision.RefName)
		if gitRef.State.IsError() {
			break
		}

		if dwp.Runtime.Verbose {
			gitRef.State = dwp.PrintRepo(gitRef)
			if gitRef.State.IsError() {
				break
			}

			// dwp.Print.IntentResponse(gitRef.State)
			dwp.Print.Ok(1,"HostName: '%s'", host.HostName)
			dwp.Print.Ok(1,"Label:    '%s'", host.Label)
			dwp.Print.Ok(1,"Provider: '%s'", host.Provider)
			dwp.Print.Ok(1,"Web Root: '%s'", webRoot)

			gitRef.State = dwp.PrintDestinationPaths()
			if gitRef.State.IsError() {
				break
			}
		}

		gitRef.State.SetOk()
	}
	// dwp.Print.IntentResponse(gitRef.State)
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef
}


func (dwp *TypeDeployWp) PrintSourcePaths() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent("Source Path Check")
	dwp.Print.Notify(0, "Source Path Check")
	for range onlyOnce {
		src := dwp.GetSourcePaths()
		if src == nil {
			dwp.State.SetError("no source paths")
		}
		//srcAbs := dwp.GetSourceAbsPaths()
		//dwp.Print.Ok("BasePath (abs):    %s", srcAbs.GetBasePath())
		dwp.Print.Ok(1,"BasePath:          %s", src.GetBasePath())
		dwp.Print.Ok(1,"WebRootPath:       %s", src.GetWebRootPath(false))
		dwp.Print.Ok(1,"ContentPath:       %s", src.GetContentPath(false))
		dwp.Print.Ok(1,"CorePath:          %s", src.GetCorePath(false))
		dwp.Print.Ok(1,"RootPath:          %s", src.GetRootPath(false))
		dwp.Print.Ok(1,"VendorPath:        %s", src.GetVendorPath(false))
	}
	// dwp.Print.IntentResponse(dwp.State)
	dwp.Print.PrintResponse(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) PrintDestinationPaths() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent("Destination Path Check")
	dwp.Print.Notify(0, "Destination Path Check")
	for range onlyOnce {
		destination := dwp.GetDestinationPaths()
		//destinationAbs := dwp.GetDestinationAbsPaths()
		//dwp.Print.Ok("BasePath (abs):    %s", destinationAbs.GetBasePath())
		dwp.Print.Ok(1,"BasePath:          %s", destination.GetBasePath())
		dwp.Print.Ok(1,"WebRootPath:       %s", destination.GetWebRootPath(false))
		dwp.Print.Ok(1,"ContentPath:       %s", destination.GetContentPath(false))
		dwp.Print.Ok(1,"CorePath:          %s", destination.GetCorePath(false))
		dwp.Print.Ok(1,"RootPath:          %s", destination.GetRootPath(false))
		dwp.Print.Ok(1,"VendorPath:        %s", destination.GetVendorPath(false))
	}
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}

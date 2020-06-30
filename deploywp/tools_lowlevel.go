package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/ux"
	"path/filepath"
	"strings"
)


/*
Part 2 - see docs.go
Part 4 - see docs.go
*/
func (dwp *TypeDeployWp) OpenRepo(url string, pathRef *toolPath.TypeOsPath) *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	// dwp.Print.Intent("Opening repo '%s'", url)
	//dwp.Print.Notify(1, "Opening repo '%s'", url)
	dwp.Print.Notify(1, "Opening")
	for range onlyOnce {
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			gitRef.State.SetError("Repository cannot be opened.")
			break
		}

		gitRef.State = gitRef.SetPath(pathRef.GetPath())
		if gitRef.State.IsError() {
			break
		}

		gitRef.State = gitRef.Open()
		if gitRef.State.IsError() {
			break
		}

		_, gitRef.State = gitRef.GetUrl()
		if gitRef.State.IsError() {
			break
		}

		gitRef.State.SetOk()

		if gitRef.Url != url {
			// dwp.Print.IntentAppend("Repo URL differs - requested:'%s' path:'%s'", url, gitRef.Url)
		}
		// dwp.Print.IntentResponse(gitRef.State)
	}
	// dwp.Print.IntentResponse(gitRef.State)
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef
}


func (dwp *TypeDeployWp) CloneRepo(url string, pathRef *toolPath.TypeOsPath) *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	// dwp.Print.Intent("Cloning '%s'", url)
	//dwp.Print.Notify(1, "Cloning '%s'", url)
	dwp.Print.Notify(1, "Cloning")
	for range onlyOnce {
		// Check repo exists and clone if not.
		gitRef.State = pathRef.StatPath()
		if pathRef.Exists() {
			break
		}
		gitRef.State = gitRef.SetPath(pathRef.GetPath())
		if gitRef.State.IsError() {
			break
		}
		gitRef.State = gitRef.SetUrl(url)
		if gitRef.State.IsError() {
			break
		}
		gitRef.State = gitRef.Clone()
		if gitRef.State.IsError() {
			break
		}

		gitRef.State.SetOk()
		if gitRef.Url != url {
			// dwp.Print.IntentAppend("Repo URL differs - requested:'%s' path:'%s'", url, gitRef.Url)
		}
	}
	// dwp.Print.IntentResponse(gitRef.State)
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef
}


func (dwp *TypeDeployWp) CheckoutRepo(gitRef *toolGit.TypeGit, versionType string, version string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent("Checkout %s:%s", strings.Title(versionType), version)
	dwp.Print.Notify(1, "Checkout %s:%s", strings.Title(versionType), version)
	for range onlyOnce {
		if gitRef.IsNotExisting() {
			gitRef.State.SetError("Repository not open.")
			break
		}
		if !IsValidVersionType(versionType) {
			gitRef.State.SetError("versionType not valid.")
			break
		}

		//dwp.Print.Intent("Verify %s:%s", versionType, version)
		if versionType == "branch" {
			_, gitRef.State = gitRef.BranchExists(version)
			if gitRef.State.IsError() {
				ux.PrintflnError("%s '%s' does not exist in repository '%s'.", strings.Title(versionType), version, gitRef.Url)
				break
			}
		} else {
			_, gitRef.State = gitRef.TagExists(version)
			if gitRef.State.IsError() {
				ux.PrintflnError("%s '%s' does not exist in repository '%s'.", strings.Title(versionType), version, gitRef.Url)
				break
			}
		}

		//dwp.Print.Intent("Checkout %s '%s' from repository '%s'.", versionType, version, gitRef.Url)
		gitRef.State = gitRef.GitCheckout(version)
		if gitRef.State.IsError() {
			break
		}

		//dwp.Print.Ok("%s '%s' checked out OK.", strings.Title(versionType), version)
		gitRef.State.SetOk()
	}
	// dwp.Print.IntentResponse(gitRef.State)
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef.State
}


/*
Part 5 - see docs.go
*/
func (dwp *TypeDeployWp) CleanRepo(gitRef *toolGit.TypeGit, force bool) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.Notify(1, "Cleaning")
	//dwp.Print.Notify(1, "Cleaning destination repository '%s'", gitRef.Base.GetPath())
	for range onlyOnce {
		if gitRef.IsNotExisting() {
			gitRef.State.SetError("Repository not open.")
			break
		}
		if !force {
			//dwp.Print.Warning("\nAbout to remove all files within the '%s' repo...", gitRef.Base.GetPathAbs())
			ok := toolPrompt.ToolUserPromptBool("Do you want to remove repo files?")
			if !ok {
				dwp.Print.Warning(1,"Aborting... ")
				gitRef.State.SetError("Abort due to user response.")
				break
			}
		}


		dwp.Print.Append(" (checked in)")
		gitRef.State = gitRef.GitRm("-r", ".")
		if gitRef.State.IsError() {
			if strings.Contains(gitRef.State.GetError().Error(), "exit status 128") {
				gitRef.State.SetOk()
			} else {
				gitRef.State.SetError("Failed to remove files on destination")
				break
			}
		}
		dwp.Print.PrintResponse(gitRef.State)


		dwp.Print.Append(" (untracked)")
		gitRef.State = gitRef.GitClean("-d", "-f", ".")
		if gitRef.State.IsError() {
			if strings.Contains(gitRef.State.GetError().Error(), "exit status 128") {
				gitRef.State.SetOk()
			} else {
				gitRef.State.SetError("Failed to remove files on destination")
				break
			}
		}
		dwp.Print.PrintResponse(gitRef.State)

		gitRef.State.SetOk()
	}
	dwp.Print.PrintResponse(gitRef.State)

	return gitRef.State
}


func (dwp *TypeDeployWp) UpdateDestination(srcPath *Paths, dstPath *Paths) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent( "Update destination")
	dwp.Print.Notify(1,  "Update destination")
	for range onlyOnce {
		//excludeFiles := []string{"composer.json"}
		excludeFiles := []string{""}
		dwp.State = dwp.CopyFiles(srcPath.GetCorePath(true), dstPath.GetCorePath(true), excludeFiles...)
		if dwp.State.IsError() {
			break
		}
		dwp.State = dwp.CopyFiles(srcPath.GetContentPath(true), dstPath.GetContentPath(true), excludeFiles...)
		if dwp.State.IsError() {
			break
		}
		dwp.State = dwp.CopyFiles(srcPath.GetVendorPath(true), dstPath.GetVendorPath(true), excludeFiles...)
		if dwp.State.IsError() {
			break
		}
		dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetBasePath(), "composer.json"), dstPath.GetBasePath())
		if dwp.State.IsError() {
			break
		}
		dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetWebRootPath(true), "pantheon.upstream.yml"), dstPath.GetBasePath())
		if dwp.State.IsError() {
			break
		}
		dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetWebRootPath(true), "pantheon.yml"), dstPath.GetBasePath())
		if dwp.State.IsError() {
			break
		}

		dwp.State.SetOk()
	}
	dwp.Print.PrintResponse(dwp.State)
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}


/*
Part 6 - see docs.go
*/
func (dwp *TypeDeployWp) CopyFiles(src string, dst string, exclude ...string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent("Copying files to %s", dst)
	dwp.Print.Notify(2, "Copying to %s", dst)
	for range onlyOnce {
		fileCopy := toolCopy.New(dwp.Runtime)
		if fileCopy.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		if !fileCopy.SetSourcePath(src) {
			dwp.State.SetError("Failed to set source path - '%s'.", src)
			break
		}

		if !fileCopy.SetDestinationPath(dst) {
			dwp.State.SetError("Failed to set destination path - '%s'.", dst)
			break
		}

		fileCopy.SetMethodRsync()
		fileCopy.SetOverwrite()
		fileCopy.SetExcludePaths(exclude...)

		//dwp.Print.Intent("Copying files:")
		//dwp.Print.Intent("    Source:      %s", fileCopy.GetSourcePath())
		//dwp.Print.Intent("    Destination: %s", fileCopy.GetDestinationPath())
		//dwp.Print.Intent("    Excludes:    %v", fileCopy.GetExcludePaths())

		dwp.State = fileCopy.Copy()
		if dwp.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		dwp.State.SetOk()
	}
	dwp.Print.PrintResponse(dwp.State)
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) CopyFile(src string, dst string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent( "Copying file %s -> %s", src, dst)
	dwp.Print.Notify(2,  "Copying to %s", dst)
	for range onlyOnce {
		fileCopy := toolCopy.New(dwp.Runtime)
		if fileCopy.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		if !fileCopy.SetSourcePath(src) {
			dwp.State.SetError("Failed to set source path - '%s'.", src)
			break
		}

		if !fileCopy.SetDestinationPath(dst) {
			dwp.State.SetError("Failed to set destination path - '%s'.", dst)
			break
		}

		fileCopy.SetOverwrite()
		fileCopy.SetMethodCp()

		dwp.State = fileCopy.Copy()
		if dwp.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		dwp.State.SetOk("")
	}
	dwp.Print.PrintResponse(dwp.State)
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}


/*
7. Run composer, (within /tmp/deploywp/destination/).
        - Fixup composer.json
                - .extra.wordpress-webroot-path = {{ .destination.paths.wordpress.root_path }}
                - .extra.wordpress-core-path = {{ .destination.paths.wordpress.core_path }}
                - .extra.wordpress-content-path = {{ .destination.paths.wordpress.content_path }}
                - .config.vendor-dir = {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.vendor_path }}
                - .extra.installer-paths.*
                        - ReplacePrefix -> destination references
                                - {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.core_path }}/
                                - {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.content_path }}/
                                - {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.vendor_path }}/
                                - {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.root_path }}/
                                - Check Mike's BASH script.

        - composer install
        - find /tmp/deploywp/destination/ -name composer.json -delete
*/
func (dwp *TypeDeployWp) RunComposer(dstDir string, args ...string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent( "Running composer")
	dwp.Print.Notify(1,  "Running composer")
	for range onlyOnce {
		exe := toolExec.New(dwp.Runtime)
		if exe.State.IsNotOk() {
			dwp.State = exe.State
			break
		}

		dwp.State = exe.SetCmd("composer")
		if dwp.State.IsNotOk() {
			break
		}

		if !exe.IsRunnable() {
			dwp.State.SetError()
			break
		}

		dwp.State = exe.SetWorkingPath(dstDir)
		if dwp.State.IsNotOk() {
			break
		}

		dwp.State = exe.SetArgs(args...)
		if dwp.State.IsNotOk() {
			break
		}

		//exe.ShowProgress()

		//dwp.Print.Intent("    Additional Args: %s\n", strings.Join(exe.GetArgs(), " "))
		//dwp.Print.Intent("    Working Dir:     %s\n", exe.GetWorkingPathAbs())

		dwp.State = exe.Run()
		if dwp.State.IsNotOk() {
			break
		}
	}
	dwp.Print.PrintResponse(dwp.State)
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) SelectDestinationHost() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	// dwp.Print.Intent( "Select host")
	dwp.Print.Notify(0,  "Select host")
	for range onlyOnce {
		host := dwp.GetHost()
		if host == "" {
			break
		}

		selectedHost := dwp.Hosts.GetByName(host)
		if selectedHost.state.IsNotOk() {
			dwp.State = selectedHost.state
			break
		}

		dwp.State = dwp.Destination.SetDestinationHost(selectedHost)
		if dwp.State.IsNotOk() {
			break
		}

		dwp.State.SetOk()
	}
	dwp.Print.PrintResponse(dwp.State)
	// dwp.Print.IntentResponse(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) GetHost() string {
	var ret string
	if state := dwp.IsNil(); state.IsError() {
		return ret
	}

	for range onlyOnce {
		dwp.State.SetOk()

		ret = dwp.Runtime.GetArg(1)
		if ret != "" {
			break
		}

		ret = toolPrompt.ToolUserPrompt("Enter host: ")
		if ret != "" {
			break
		}

		dwp.State.SetError("host is empty")
	}

	return ret
}


func (dwp *TypeDeployWp) PrintRepo(gitRef *toolGit.TypeGit) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		p := gitRef.GetPath()
		u, _ := gitRef.GetUrl()
		b, _ := gitRef.GetBranch()
		t, _ := gitRef.GetTags()
		s := gitRef.GetStatus().GetOutput()

		// dwp.Print.Intent( "Source Repository")
		dwp.Print.Notify(1,  "Source Repository")
		dwp.Print.Ok(1, "Provider:  GitHub")
		dwp.Print.Ok(1, "Path:      %s", p)
		dwp.Print.Ok(1, "Url:       %s", u)
		dwp.Print.Ok(1, "Branch(current):    %s", b)
		dwp.Print.Ok(1, "Tags(available):    %s", strings.Join(t, " "))
		dwp.Print.Ok(1, "Status:    %s", s)
	}

	return gitRef.State
}

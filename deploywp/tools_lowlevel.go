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

	for range onlyOnce {
		// Check repo exists and clone if not.
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			dwp.Print.IntentBegin("Cloning repo '%s'", url)
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
			dwp.Print.IntentResponse(gitRef.State, "")
		}

		dwp.Print.IntentBegin("Opening repo '%s'", url)
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
			dwp.Print.IntentResponse(gitRef.State, "Repo URL differs - requested:'%s' path:'%s'", url, gitRef.Url)
		}
		dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
	}

	dwp.Print.IntentEnd(gitRef.State)
	return gitRef
}


func (dwp *TypeDeployWp) CloneRepo(url string, pathRef *toolPath.TypeOsPath) *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	dwp.Print.IntentBegin("Cloning '%s'", url)
	for range onlyOnce {
		// Check repo exists and clone if not.
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			gitRef.State = gitRef.SetPath(pathRef.GetPath())
			if gitRef.State.IsError() {
				dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
				break
			}
			gitRef.State = gitRef.SetUrl(url)
			if gitRef.State.IsError() {
				dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
				break
			}
			gitRef.State = gitRef.Clone()
			if gitRef.State.IsError() {
				dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
				break
			}
			dwp.Print.IntentResponse(gitRef.State, "")
			break
		}


		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			//ux.PrintflnRed("Repository cannot be cloned.")
			gitRef.State.SetError("Repository cannot be cloned.")
			dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
			break
		}
		//dwp.Print.Ok("Repository path: '%s'", pathRef.GetPathAbs())


		gitRef.State = gitRef.SetPath(pathRef.GetPath())
		if gitRef.State.IsError() {
			dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
			break
		}

		gitRef.State = gitRef.Open()
		if gitRef.State.IsError() {
			dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
			break
		}

		_, gitRef.State = gitRef.GetUrl()
		if gitRef.State.IsError() {
			dwp.Print.IntentResponse(gitRef.State, gitRef.State.Sprint())
			break
		}

		gitRef.State.SetOk()

		if gitRef.Url != url {
			dwp.Print.IntentResponse(gitRef.State, "Repo URL differs - requested:'%s' path:'%s'", url, gitRef.Url)
		}
	}
	dwp.Print.IntentEnd(gitRef.State)

	return gitRef
}


func (dwp *TypeDeployWp) CheckoutRepo(gitRef *toolGit.TypeGit, versionType string, version string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Checkout %s:%s", strings.Title(versionType), version)
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
	dwp.Print.IntentEnd(gitRef.State)

	return dwp.State
}


/*
Part 5 - see docs.go
*/
func (dwp *TypeDeployWp) CleanRepo(gitRef *toolGit.TypeGit, force bool) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Cleaning destination repository '%s'", gitRef.Base.GetPath())
	for range onlyOnce {
		if gitRef.IsNotExisting() {
			dwp.State.SetError("Repository not open.")
			break
		}
		if !force {
			//dwp.Print.Warning("\nAbout to remove all files within the '%s' repo...", gitRef.Base.GetPathAbs())
			ok := toolPrompt.ToolUserPromptBool("Do you want to remove repo files?")
			if !ok {
				dwp.Print.Warning("Aborting... ")
				dwp.State.SetError("Abort due to user response.")
				dwp.Print.IntentResponse(dwp.State, "")
				break
			}
		}


		dwp.Print.Intent("Removing files (checked in)")
		dwp.State = gitRef.GitRm("-r", ".")
		//foo := dwp.State.OutputGrep("did not match any files")
		if dwp.State.IsError() {
			if strings.Contains(dwp.State.GetError().Error(), "exit status 128") {
				dwp.State.SetOk()
			} else {
				dwp.Print.IntentResponse(dwp.State, "")
				dwp.State.SetError("Failed to remove files on destination")
				break
			}
		}
		dwp.Print.IntentResponse(dwp.State, "")


		dwp.Print.Intent("Removing files (untracked)")
		dwp.State = gitRef.GitClean("-d", "-f", ".")
		//foo := dwp.State.OutputGrep("did not match any files")
		if dwp.State.IsError() {
			if strings.Contains(dwp.State.GetError().Error(), "exit status 128") {
				dwp.State.SetOk()
			} else {
				dwp.State.SetError("Failed to remove files on destination")
				dwp.Print.IntentResponse(dwp.State, "")
				break
			}
		}
		dwp.Print.IntentResponse(dwp.State, "")

		//dwp.Print.Ok("File removal completed OK.")
		dwp.State.SetOk()
	}
	dwp.Print.IntentEnd(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) UpdateDestination(srcPath *Paths, dstPath *Paths) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Update destination")
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
	dwp.Print.IntentEnd(dwp.State)

	return dwp.State
}


/*
Part 6 - see docs.go
*/
func (dwp *TypeDeployWp) CopyFiles(src string, dst string, exclude ...string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Copying files %s -> %s", src, dst)
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

		dwp.Print.Ok("Files copied with OK")
	}
	dwp.Print.IntentEnd(dwp.State)

	return dwp.State
}

func (dwp *TypeDeployWp) CopyFile(src string, dst string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Copying file")
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

		dwp.Print.Intent("Copying file:")
		dwp.Print.Intent("    Source:      %s", fileCopy.GetSourcePath())
		dwp.Print.Intent("    Destination: %s", fileCopy.GetDestinationPath())

		dwp.State = fileCopy.Copy()
		if dwp.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		dwp.Print.Ok("Files copied with OK")
	}
	dwp.Print.IntentEnd(dwp.State)

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

	dwp.Print.IntentBegin("Running composer")
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

		exe.ShowProgress()

		dwp.Print.Intent("Running composer\n")
		dwp.Print.Intent("    Additional Args: %s\n", strings.Join(exe.GetArgs(), " "))
		dwp.Print.Intent("    Working Dir:     %s\n", exe.GetWorkingPathAbs())

		dwp.State = exe.Run()
		if dwp.State.IsNotOk() {
			break
		}
	}
	dwp.Print.IntentEnd(dwp.State)

	return dwp.State
}


func (dwp *TypeDeployWp) SelectDestinationHost() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	dwp.Print.IntentBegin("Select host")
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
	dwp.Print.IntentEnd(dwp.State)

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

		dwp.Print.Intent("SOURCE REPO:")
		dwp.Print.Ok("Provider:  GitHub")
		dwp.Print.Ok("Path:      %s", p)
		dwp.Print.Ok("Url:       %s", u)
		dwp.Print.Ok("Branch(current):    %s", b)
		dwp.Print.Ok("Tags(available):    %s", strings.Join(t, " "))
		dwp.Print.Ok("Status:    %s", s)
	}

	return gitRef.State
}

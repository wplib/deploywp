package deploywp

import (
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


func (dwp *TypeDeployWp) OpenRepo(url string, path ...string) *toolGit.TypeGit {
	gitRef := toolGit.New(dwp.Runtime)
	if state := dwp.IsNil(); state.IsError() {
		return &toolGit.TypeGit{State: state}
	}

	for range onlyOnce {
		ux.PrintflnBlue("# Opening Git repository.")
		ux.PrintflnOk("Repo Url: '%s'", url)


		// Check repo exists and clone if not.
		pathRef := toolPath.New(dwp.Runtime)
		if pathRef.State.IsNotOk() {
			dwp.State = pathRef.State
			break
		}
		pathRef.SetPath(path...)
		gitRef.State = pathRef.StatPath()
		if pathRef.NotExists() {
			ux.PrintflnBlue("# Repository, doesn't exist - cloning repository.")
			dwp.State = gitRef.SetPath(path...)
			if dwp.State.IsError() {
				break
			}
			dwp.State = gitRef.SetUrl(url)
			if dwp.State.IsError() {
				break
			}
			dwp.State = gitRef.Clone()
			if dwp.State.IsError() {
				break
			}
		}


		dwp.State = pathRef.StatPath()
		if pathRef.NotExists() {
			ux.PrintflnBlue("# Repository cannot be cloned.")
			dwp.State.SetError("Repository cannot be cloned.")
			break
		}
		ux.PrintflnOk("Path:     '%s'", pathRef.GetPathAbs())


		dwp.State = gitRef.SetPath(path...)
		if dwp.State.IsError() {
			break
		}

		dwp.State = gitRef.Open()
		if dwp.State.IsError() {
			break
		}

		dwp.State = gitRef.GetUrl()
		if dwp.State.IsError() {
			break
		}
		ux.PrintflnOk("Repo Url: '%s'", gitRef.Url)
		if gitRef.Url != url {
			ux.PrintflnWarning("# Repo URL is different.")
			ux.PrintflnWarning("    - Requested: %s", url)
			ux.PrintflnWarning("    - Directory: %s", gitRef.Url)
		}

		dwp.State = gitRef.GetBranch()
		if dwp.State.IsError() {
			break
		}
		branch := dwp.State.Output
		ux.PrintflnOk("Current Branch: '%s'", branch)


		ux.PrintflnOk("# Repository opened OK.")
		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return gitRef
}


func (dwp *TypeDeployWp) CheckoutRepo(gitRef *toolGit.TypeGit, versionType string, version string) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if gitRef.IsNotExisting() {
			dwp.State.SetError("Repository not open.")
			break
		}
		if !IsValidVersionType(versionType) {
			dwp.State.SetError("versionType not valid.")
			break
		}

		ux.PrintflnBlue("# Verify %s '%s' exists in repository '%s'.", versionType, version, gitRef.Url)
		if versionType == "branch" {
			dwp.State = gitRef.BranchExists(version)
			if dwp.State.IsError() {
				ux.PrintflnError("# %s '%s' does not exist in repository '%s'.", versionType, version, gitRef.Url)
				break
			}
		} else {
			dwp.State = gitRef.TagExists(version)
			if dwp.State.IsError() {
				ux.PrintflnError("# %s '%s' does not exist in repository '%s'.", versionType, version, gitRef.Url)
				break
			}
		}

		ux.PrintflnBlue("# Checkout %s '%s' from repository '%s'.", versionType, version, gitRef.Url)
		dwp.State = gitRef.GitCheckout(version)
		if dwp.State.IsError() {
			break
		}

		ux.PrintflnOk("# %s '%s' checked out OK.", versionType, version)
		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return dwp.State
}


func (dwp *TypeDeployWp) CleanRepo(gitRef *toolGit.TypeGit, force bool) *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if gitRef.IsNotExisting() {
			dwp.State.SetError("Repository not open.")
			break
		}
		if !force {
			ux.PrintflnWarning("About to remove all files within the '%s' repo...", gitRef.Base.GetPathAbs())
			ok := toolPrompt.ToolUserPromptBool("Do you really want to do this?%s", "")
			if !ok {
				ux.PrintflnWarning("Aborting...")
				dwp.State.SetError("Abort due to user response.")
				break
			}
		}

		ux.PrintflnBlue("# Removing files...")
		dwp.State = gitRef.GitRm("rm", "-r", ".")
		//foo := dwp.State.OutputGrep("did not match any files")
		if strings.Contains(dwp.State.GetError().Error(), "exit status 128") {
			dwp.State.SetOk()
		}
		if dwp.State.IsError() {
			dwp.State.SetError("Failed to remove files on target")
			break
		}

		ux.PrintflnOk("# File removal completed OK.")
		dwp.State.SetOk()
	}

	gitRef.State = dwp.State
	return dwp.State
}


/*
6. Copy directories into /tmp/deploywp/target/
        - Honour {{ .target.files.exclude }} && {{ .target.files.copy }} && {{ .target.files.keep }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.core_path }}
                - To {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.core_path }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.content_path }}
                - To {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.content_path }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.vendor }}
                - To {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.vendor }}
*/
func (dwp *TypeDeployWp) CopyFiles() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		srcAbs := dwp.GetSourceAbsPaths()
		targetAbs := dwp.GetTargetAbsPaths()

		fileCopy := toolCopy.New(dwp.Runtime)
		if fileCopy.State.IsError() {
			dwp.State = fileCopy.State
			break
		}

		if !fileCopy.SetSourcePath(srcAbs.GetCorePath()) {
			dwp.State.SetError("Failed to set source path - '%s'.", srcAbs.GetCorePath())
			break
		}

		if !fileCopy.SetDestinationPath(targetAbs.) {
			dwp.State.SetError("Failed to set target path - '%s'.", targetAbs.)
			break
		}

		// .

		src := dwp.GetSourcePaths()
		//srcAbs := dwp.GetSourceAbsPaths()
		ux.PrintflnBlue("# SOURCE PATHS:")
		ux.PrintflnOk("BasePath (abs):    %s", srcAbs.GetBasePath())
		ux.PrintflnOk("BasePath:          %s", src.GetBasePath())
		ux.PrintflnOk("WebRootPath:       %s", src.GetWebRootPath())
		ux.PrintflnOk("ContentPath:       %s", src.GetContentPath())
		ux.PrintflnOk("CorePath:          %s", src.GetCorePath())
		ux.PrintflnOk("RootPath:          %s", src.GetRootPath())
		ux.PrintflnOk("VendorPath:        %s", src.GetVendorPath())

		target := dwp.GetTargetPaths()
		//targetAbs := dwp.GetTargetAbsPaths()
		ux.PrintflnBlue("# TARGET PATHS:")
		ux.PrintflnOk("BasePath (abs):    %s", targetAbs.GetBasePath())
		ux.PrintflnOk("BasePath:          %s", target.GetBasePath())
		ux.PrintflnOk("WebRootPath:       %s", target.GetWebRootPath())
		ux.PrintflnOk("ContentPath:       %s", target.GetContentPath())
		ux.PrintflnOk("CorePath:          %s", target.GetCorePath())
		ux.PrintflnOk("RootPath:          %s", target.GetRootPath())
		ux.PrintflnOk("VendorPath:        %s", target.GetVendorPath())
	}

	return dwp.State
}


func (dwp *TypeDeployWp) PrintPaths() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		src := dwp.GetSourcePaths()
		srcAbs := dwp.GetSourceAbsPaths()
		ux.PrintflnBlue("# SOURCE PATHS:")
		ux.PrintflnOk("BasePath (abs):    %s", srcAbs.GetBasePath())
		ux.PrintflnOk("BasePath:          %s", src.GetBasePath())
		ux.PrintflnOk("WebRootPath:       %s", src.GetWebRootPath())
		ux.PrintflnOk("ContentPath:       %s", src.GetContentPath())
		ux.PrintflnOk("CorePath:          %s", src.GetCorePath())
		ux.PrintflnOk("RootPath:          %s", src.GetRootPath())
		ux.PrintflnOk("VendorPath:        %s", src.GetVendorPath())

		target := dwp.GetTargetPaths()
		targetAbs := dwp.GetTargetAbsPaths()
		ux.PrintflnBlue("# TARGET PATHS:")
		ux.PrintflnOk("BasePath (abs):    %s", targetAbs.GetBasePath())
		ux.PrintflnOk("BasePath:          %s", target.GetBasePath())
		ux.PrintflnOk("WebRootPath:       %s", target.GetWebRootPath())
		ux.PrintflnOk("ContentPath:       %s", target.GetContentPath())
		ux.PrintflnOk("CorePath:          %s", target.GetCorePath())
		ux.PrintflnOk("RootPath:          %s", target.GetRootPath())
		ux.PrintflnOk("VendorPath:        %s", target.GetVendorPath())
	}

	return dwp.State
}


func (dwp *TypeDeployWp) ObtainHost() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var host string
		for range onlyOnce {
			dwp.State.SetOk()

			host = dwp.Runtime.GetArg(0)
			if host != "" {
				break
			}

			host = toolPrompt.ToolUserPrompt("Enter host: ")
			if host != "" {
				break
			}

			dwp.State.SetError("host is empty")
		}
		if dwp.State.IsError() {
			break
		}

		dwp.State.SetOutput(host)
	}

	return dwp.State
}

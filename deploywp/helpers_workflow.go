package deploywp

import (
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
	"time"
)


// This is an alternative to running templates.
// In theory, the code here, should be able to be replicated in a template file without modification.
func (dwp *TypeDeployWp) Build() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		ux.PrintfBlue("%s v%s\n", dwp.Runtime.CmdFile, dwp.Runtime.CmdVersion)
		ux.PrintfGreen("args: %s\n", dwp.Runtime.GetArgs())
		ux.PrintfWhite("\n\n")


		{
			ux.PrintfWhite("##########################\n")
			ux.PrintfWhite("# Print path information #\n")
			ux.PrintfWhite("##########################\n")
			dwp.State = dwp.PrintPaths()
			if dwp.State.IsError() {
				break
			}
			ux.PrintfWhite("\n\n")
		}


		{
			ux.PrintfWhite("#############################\n")
			ux.PrintfWhite("# Opening source repository #\n")
			ux.PrintfWhite("#############################\n")
			srcGitRef := dwp.OpenSourceRepo()
			if srcGitRef.State.IsError() {
				dwp.State = srcGitRef.State
				break
			}
			ux.PrintfWhite("\n\n")
		}


		{
			ux.PrintfWhite("#############################\n")
			ux.PrintfWhite("# Opening target repository #\n")
			ux.PrintfWhite("#############################\n")
			targetGitRef := dwp.OpenTargetRepo()
			if targetGitRef.State.IsError() {
				dwp.State = targetGitRef.State
				break
			}
			ux.PrintfWhite("\n\n")

			ux.PrintfWhite("##############################\n")
			ux.PrintfWhite("# Cleaning target repository #\n")
			ux.PrintfWhite("##############################\n")
			dwp.State = dwp.CleanRepo(targetGitRef, true)
			if dwp.State.IsError() {
				break
			}
			ux.PrintfWhite("\n\n")
		}


		{
			ux.PrintfWhite("#############################\n")
			ux.PrintfWhite("# Syncing target repository #\n")
			ux.PrintfWhite("#############################\n")
			//excludeFiles := []string{"composer.json"}
			excludeFiles := []string{""}
			srcPath := dwp.GetSourceAbsPaths()
			dstPath := dwp.GetTargetAbsPaths()
			dwp.State = dwp.CopyFiles(srcPath.GetCorePath(), dstPath.GetCorePath(), excludeFiles...)
			if dwp.State.IsError() {
				break
			}
			dwp.State = dwp.CopyFiles(srcPath.GetContentPath(), dstPath.GetContentPath(), excludeFiles...)
			if dwp.State.IsError() {
				break
			}
			dwp.State = dwp.CopyFiles(srcPath.GetVendorPath(), dstPath.GetVendorPath(), excludeFiles...)
			if dwp.State.IsError() {
				break
			}
			dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetBasePath(), "composer.json"), dstPath.GetBasePath())
			if dwp.State.IsError() {
				break
			}
			dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetWebRootPath(), "pantheon.upstream.yml"), dstPath.GetBasePath())
			if dwp.State.IsError() {
				break
			}
			dwp.State = dwp.CopyFiles(filepath.Join(srcPath.GetWebRootPath(), "pantheon.yml"), dstPath.GetBasePath())
			if dwp.State.IsError() {
				break
			}
			ux.PrintfWhite("\n\n")
		}


		{
			ux.PrintfWhite("#########################################\n")
			ux.PrintfWhite("# Running composer on target repository #\n")
			ux.PrintfWhite("#########################################\n")
			dwp.State = dwp.RunComposer(dwp.GetTargetAbsPaths().GetBasePath(), "install")
			if dwp.State.IsError() {
				break
			}
			ux.PrintfWhite("\n\n")
		}

		os.Exit(1)

		ux.PrintfWhite("############################################\n")
		ux.PrintfWhite("# Increment BUILD within target repository #\n")
		ux.PrintfWhite("############################################\n")
		//dwp.State = dwp.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("########################################\n")
		ux.PrintfWhite("# Commit target repository to Pantheon #\n")
		ux.PrintfWhite("########################################\n")
		//dwp.State = dwp.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("############################\n")
		ux.PrintfWhite("# Commit source repository #\n")
		ux.PrintfWhite("############################\n")
		//dwp.State = dwp.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")
	}

	return dwp.State
}

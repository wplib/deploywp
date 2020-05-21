package deploywp

import (
	"github.com/wplib/deploywp/ux"
	"time"
)

/*
# Ad-hoc deploy.
1. Determine GIT branch within project to deploy, (.source.repository).
        - Used when pushing to .source.repository
        1. Needs to match the target branch
        2. Needs to be fully committed to .source.repository
        3. Needs to be pushed to .source.repository
*/

/*
2. Clone src repo, (.source.repository), into /tmp/deploywp/source/
        - Checkout {{ .source.revision.ref_type }} {{ .source.revision.ref_name }}
*/

/*
3. Determine target repo.
        - {{ .source.revision.ref_type }} == branch_name => {{ .target.revisions.ref_name }}
        - {{ .target.revisions.host_name }} => {{ .hosts.host_name }}
                - Merge {{ index .target.providers.name .hosts.provider }}
                - Expand {{ .target.providers.[name].defaults.url }} using {{ .providers.meta.site_id }}
                - OR Expand {{ .hosts.[name].respository.url }} using {{ .target.providers.meta.* }}
*/

/*
4. Clone target repo, (step 3 URL), into /tmp/deploywp/target/
        - Checkout {{ .source.revision.ref_type }} {{ .source.revision.ref_name }}
                - Or create if not exist.
*/

/*
5. Remove all files wihin target repo.
        - cd /tmp/deploywp/target/
        - git rm -r --cached .
        - Remove everything except .git
*/

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

/*
7. Run composer, (within /tmp/deploywp/target/).
        - Fixup composer.json
                - .extra.wordpress-webroot-path = {{ .target.paths.wordpress.root_path }}
                - .extra.wordpress-core-path = {{ .target.paths.wordpress.core_path }}
                - .extra.wordpress-content-path = {{ .target.paths.wordpress.content_path }}
                - .config.vendor-dir = {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.vendor_path }}
                - .extra.installer-paths.*
                        - ReplacePrefix -> target references
                                - {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.core_path }}/
                                - {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.content_path }}/
                                - {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.vendor_path }}/
                                - {{ .target.paths.webroot_path }}/{{ .target.paths.wordpress.root_path }}/
                                - Check Mike's BASH script.

        - composer install
        - find /tmp/deploywp/target/ -name composer.json -delete
*/

/*
8. Maintain build file, (/tmp/deploywp/target/BUILD).
        - Contents of BUILD file should be incremented.
*/

/*
9. Commit to Pantheon, (within /tmp/deploywp/target/).
        - git add .
        - git commit -m 'Use commit message from deploy.functions.sh:236' .
                - Include BUILD # from 8.
        - git push
*/

/*
10. Update build file, (/tmp/deploywp/source/BUILD).
        - Contents of BUILD file should be incremented.
        - git add .
        - git commit -m 'Use commit message from deploy.functions.sh:236' .
                - Include BUILD # from 8.
        - git push
 */


// This is an alternative to running templates.
// In theory, the code here, should be able to be replicated in a template file without modification.
func (e *TypeDeployWp) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		ux.PrintfBlue("%s v%s\n", e.Exec.CmdFile, e.Exec.CmdVersion)
		ux.PrintfGreen("Args: %s\n", e.Exec.GetArgs())
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("##########################\n")
		ux.PrintfWhite("# Print path information #\n")
		ux.PrintfWhite("##########################\n")
		e.State = e.PrintPaths()
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("#############################\n")
		ux.PrintfWhite("# Opening source repository #\n")
		ux.PrintfWhite("#############################\n")
		srcGitRef := e.OpenSourceRepo()
		if srcGitRef.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("##############################\n")
		ux.PrintfWhite("# Checkout source repository #\n")
		ux.PrintfWhite("##############################\n")
		srcGitRef.State = e.CheckoutSourceRepo(srcGitRef)
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("#############################\n")
		ux.PrintfWhite("# Opening target repository #\n")
		ux.PrintfWhite("#############################\n")
		srcGitRef.State = e.OpenTargetRepo()
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("##############################\n")
		ux.PrintfWhite("# Cleaning target repository #\n")
		ux.PrintfWhite("##############################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("#############################\n")
		ux.PrintfWhite("# Syncing target repository #\n")
		ux.PrintfWhite("#############################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("#########################################\n")
		ux.PrintfWhite("# Running composer on target repository #\n")
		ux.PrintfWhite("#########################################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("############################################\n")
		ux.PrintfWhite("# Increment BUILD within target repository #\n")
		ux.PrintfWhite("############################################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("########################################\n")
		ux.PrintfWhite("# Commit target repository to Pantheon #\n")
		ux.PrintfWhite("########################################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")


		ux.PrintfWhite("############################\n")
		ux.PrintfWhite("# Commit source repository #\n")
		ux.PrintfWhite("############################\n")
		//srcGitRef.State = e.OpenTargetRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if e.State.IsError() {
			break
		}
		ux.PrintfWhite("\n\n")
	}

	return e.State
}

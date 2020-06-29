package deploywp

/*
# Ad-hoc deploy.
1. Determine GIT branch within project to deploy, (.source.repository).
        - Used when pushing to .source.repository
        1. Needs to match the destination branch
        2. Needs to be fully committed to .source.repository
        3. Needs to be pushed to .source.repository
*/

/*
2. Clone src repo, (.source.repository), into /tmp/deploywp/source/
        - Checkout {{ .source.revision.ref_type }} {{ .source.revision.ref_name }}
*/

/*
3. Determine destination repo.
        - {{ .source.revision.ref_type }} == branch_name => {{ .destination.revisions.ref_name }}
        - {{ .destination.revisions.host_name }} => {{ .hosts.host_name }}
                - Merge {{ index .destination.providers.name .hosts.provider }}
                - Expand {{ .destination.providers.[name].defaults.url }} using {{ .providers.meta.site_id }}
                - OR Expand {{ .hosts.[name].respository.url }} using {{ .destination.providers.meta.* }}
*/

/*
4. Clone destination repo, (step 3 URL), into /tmp/deploywp/destination/
        - Checkout {{ .source.revision.ref_type }} {{ .source.revision.ref_name }}
                - Or create if not exist.
*/

/*
5. Remove all files within destination repo.
        - cd /tmp/deploywp/destination/
        - git rm -r --cached .
        - Remove everything except .git
*/

/*
6. Copy directories into /tmp/deploywp/destination/
        - Honour {{ .destination.files.exclude }} && {{ .destination.files.copy }} && {{ .destination.files.keep }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.core_path }}
                - To {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.core_path }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.content_path }}
                - To {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.content_path }}
        - Copy {{ .source.paths.webroot_path }}/{{ .source.paths.wordpress.vendor }}
                - To {{ .destination.paths.webroot_path }}/{{ .destination.paths.wordpress.vendor }}
*/

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

/*
8. Maintain build file, (/tmp/deploywp/destination/BUILD).
        - Contents of BUILD file should be incremented.
*/

/*
9. Commit to Pantheon, (within /tmp/deploywp/destination/).
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

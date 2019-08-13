package jsonfile

func noop(i ...interface{}) interface{} { return i }

func GetDefault() string {
	return `{
   "deploywp": {
      "schema": "1.0",
      "repository": {
         "url": "github.com/wplib/deploywp"
      },
      "branch": "master"
   },
   "site": {
      "id": "@TODO: Change to a slug to identify your site. Often this is the 2nd level domain w/o the first level, e.g. 'google' vs. 'google.com'",
      "name": "@TODO: Change to a human readable short name for your site",
      "label": "@TODO: Change to a human readable long name for your site, if desired",
      "domain": "@TODO: Change to the production site's domain (omit 'www.')",
      "website": "https://www.{site.domain}"
   },
   "source": {
      "web_root": "/www",
      "wp_paths": {
         "root_path": "/",
         "core_path": "/wp",
         "content_path": "/content",
         "vendor_path": "/vendor"
      },
      "repository": {
         "provider": "github",
         "url": "@TODO: Change to site's repository URL. You can omit the scheme, such as https:// and ssh://, but you do not have to"
      },
      "files": {
         "exclude": [
            "{source.web_root}/index.php",
            "{source.web_root}/wp-config.php",
            "{source.web_root}/wp-config-{site.id}.local.php",
            "composer.json"
         ],
         "delete": [
            "{source.web_root}/readme.html",
            "{source.web_root}/README.md",
            "{source.web_root}/xmlrpc.php",
            "{source.web_root}/wp-trackback.php",
            "{source.web_root}/wp-signup.php",
            "{source.web_root}/license.txt",
            "{source.web_root}/wp-config-sample.php",
            "{source.wp_paths.content_path}/plugins/hello.php",
            "{source.wp_paths.content_path}/themes/twentyseventeen/*",
            "{source.wp_paths.content_path}/themes/twentyseventeen"
         ],
         "keep": [
            "{source.wp_paths.content_path}/mu-plugins/pantheon.php",
            "{source.wp_paths.content_path}/mu-plugins/pantheon",
            "{source.wp_paths.content_path}/mu-plugins/index.php",
            "{source.wp_paths.content_path}/plugins/index.php",
            "{source.wp_paths.content_path}/themes/index.php",
            "{source.wp_paths.content_path}/index.php"
         ],
         "copy": [
            "{source.web_root}/pantheon.yml"
         ]
      }
   },
   "targets": {
      "defaults": {
         "provider": "pantheon",
         "site_guid": "@TODO: Change to the GUID Pantheon uses to identify sites in in dashboard URLs",
         "repository": {
            "url": "codeserver.dev.{site_guid}@codeserver.dev.{site_guid}.drush.in:2222/~/repository"
         },
		 "domain": "{.id}-{.domain_suffix}"
         "domain_suffix": "-{site.id}.pantheonsite.io",
         "web_root": "/code",
         "wp_paths": {
            "root_path": "/",
            "core_path": "/",
            "content_path": "/wp-content",
            "vendor_path": "/vendor"
         }
      },
      "hosts": [
         {
            "id": "another",
            "name": "Another Site Environment",
            "label": "Another Site Environment on Pantheon's Multidev'",
         },
         {
            "id": "dev",
            "branch": "master",
			"name": "Testing",
            "label": "Testing  (Pantheon's 'dev' site)",
         },
         {
            "id": "test",
            "branch": "master",
			"name": "Staging",
            "label": "Staging  (Pantheon's 'test' site)",
            "after": "deploy_from_dev"
         },
         {
            "id": "live",
            "branch": "master",
			"name": "Production",
            "label": "Production (Pantheon's 'live' site)'",
            "after": "deploy_from_test"
         }
      ]
   }
}`
}

package jsonfile

import (
	"fmt"
	"reflect"
	"strings"
)

type Node interface{}

//
// @see https://gist.github.com/hvoecking/10772475
// @see https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
//
func fixup(node Node) Node {
	original := reflect.ValueOf(node)
	temp := reflect.New(original.Type()).Elem()
	fmt.Printf("<root>")
	fixuprecursive(original, temp, 0)
	return temp.Interface()
}

const spacer = "  "

func fixuprecursive(original, temp reflect.Value, depth int) {
	for range Once {
		indent := strings.Repeat(spacer, depth)

		pt := func() { fmt.Printf(" [%+v]", original.Type()) }

		switch original.Kind() {
		case reflect.Ptr:
			ov := original.Elem()
			if !ov.IsValid() {
				break
			}
			temp.Set(reflect.New(ov.Type()))
			fixuprecursive(ov, temp.Elem(), depth)

		case reflect.Interface:
			ov := original.Elem()
			if !ov.IsValid() {
				break
			}
			tv := reflect.New(ov.Type()).Elem()
			fixuprecursive(ov, tv, depth)
			temp.Set(tv)

		case reflect.Struct:
			pt()
			depth++
			for i := 0; i < original.NumField(); i++ {
				cf := temp.Field(i)
				ct := temp.Type().Field(i)
				name := ct.Tag.Get("json")
				if !cf.CanSet() {
					continue
				}
				if name != "-" {
					fmt.Printf("\n%s%s— %s", spacer, indent, ct.Tag.Get("json"))
				}
				fixuprecursive(original.Field(i), cf, depth)
			}

		case reflect.Slice:
			pt()
			depth++
			o := original
			temp.Set(reflect.MakeSlice(o.Type(), o.Len(), o.Cap()))
			for i := 0; i < o.Len(); i += 1 {
				fmt.Printf("\n%s%s[%d]", spacer, indent, i)
				fixuprecursive(original.Index(i), temp.Index(i), depth)
			}

		case reflect.Map:
			pt()
			depth++
			temp.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				fmt.Printf("\n%s%s[%s]", spacer, indent, key)
				ov := original.MapIndex(key)
				tv := reflect.New(ov.Type()).Elem()
				fixuprecursive(ov, tv, depth)
				temp.SetMapIndex(key, tv)
			}

		case reflect.String:
			pt()
			fmt.Printf(": %s", original.Interface())
			temp.Set(original)

		default:
			if !original.CanInterface() {
				break
			}
			pt()
			fmt.Printf(": %+v", original.Interface())
			temp.Set(original)

		}

	}
}

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

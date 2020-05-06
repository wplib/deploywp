package jsonTemplate

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"io/ioutil"
	"os"
	"path/filepath"
)


//func main() {
//	var err error
//	var tmpl *Template
//
//	for range only.Once {
//		tmpl, err = ProcessArgs()
//		if err != nil {
//			fmt.Printf("ERROR: %s\n\n", err)
//			Help()
//
//			break
//		}
//
//		err = tmpl.ProcessTemplate()
//		if err != nil {
//			fmt.Printf("ERROR: %s\n", err)
//			break
//		}
//	}
//
//	if err != nil {
//		os.Exit(1)
//	} else {
//		os.Exit(0)
//	}
//}
//
//func Help() {
//	for range only.Once {
//		fmt.Printf("JsonConfig v%s:\n", Version)
//		fmt.Printf("\tThe ultimate config file generator.\n")
//		fmt.Printf("\tFeed it a JSON and GoLang template file, I'll do the rest.\n")
//		fmt.Printf("\n")
//
//		flag.PrintDefaults()
//		fmt.Printf("\n")
//		fmt.Printf("\n")
//	}
//}
//
//func HelpVariables() {
//	for range only.Once {
//		fmt.Printf("Keys accessible within your template file:\n")
//		fmt.Printf("\t{{ .Json }} - Your JSON file will appear here.\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ .Env }} - A map containing the runtime environment.\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ .ExecName }} - Executable used to produce the resulting file.\n")
//		fmt.Printf("\t{{ .ExecVersion }} - Version of executable used to produce the resulting file.\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ .CreationDate }} - Creation date of resulting file.\n")
//		fmt.Printf("\t{{ .CreationEpoch }} - Creation date, (unix epoch), of resulting file.\n")
//		fmt.Printf("\t{{ .CreationInfo }} - More creation information.\n")
//		fmt.Printf("\t{{ .CreationWarning }} - Generic 'DO NOT EDIT' warning.\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ .TemplateFile.Dir }} - template file absolute directory.\n")
//		fmt.Printf("\t{{ .TemplateFile.Name }} - template filename.\n")
//		fmt.Printf("\t{{ .TemplateFile.CreationDate }} - template file creation date.\n")
//		fmt.Printf("\t{{ .TemplateFile.CreationEpoch }} - template file creation date, (unix epoch).\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ .JsonFile.Dir }} - json file absolute directory.\n")
//		fmt.Printf("\t{{ .JsonFile.Name }} - json filename.\n")
//		fmt.Printf("\t{{ .JsonFile.CreationDate }} - json file creation date.\n")
//		fmt.Printf("\t{{ .JsonFile.CreationEpoch }} - json file creation date, (unix epoch).\n")
//		fmt.Printf("\n")
//	}
//}
//
//func HelpFunctions() {
//	for range only.Once {
//		fmt.Printf("Functions accessible within your template file:\n")
//		fmt.Printf("\t{{ isInt $value }} - is $value an integer?\n")
//		fmt.Printf("\t{{ isString $value }} - is $value a string?\n")
//		fmt.Printf("\t{{ isSlice $value }} - is $value a slice?\n")
//		fmt.Printf("\t{{ isArray $value }} - is $value an array?\n")
//		fmt.Printf("\t{{ isMap $value }} - is $value a map?\n")
//		fmt.Printf("\n")
//		fmt.Printf("\t{{ ToUpper $value }} - uppercase $value.\n")
//		fmt.Printf("\t{{ ToLower $value }} - lowercase $value.\n")
//		fmt.Printf("\t{{ ToString $value }} - convert $value to a string.\n")
//		fmt.Printf("\t{{ FindInMap $map $value }} - find $value in $map and return reference.\n")
//		fmt.Printf("\t{{ ReadFile $file }} - read in $file and print verbatim. \n")
//		fmt.Printf("\n")
//		fmt.Printf("See http://masterminds.github.io/sprig/ for additional functions...\n")
//		fmt.Printf("\n")
//	}
//}
//
//func HelpExamples() {
//	for range only.Once {
//		fmt.Printf("Examples:\n")
//		fmt.Printf("# Print out .dir key from config.json\n")
//		fmt.Printf("\tJsonConfig -json config.json -template-string '{{ .Json.dir }}'\n")
//
//		fmt.Printf("# Process Dockerfile.tmpl file and output to STDOUT.\n")
//		fmt.Printf("\tJsonConfig -json config.json -template DockerFile.tmpl\n")
//
//		fmt.Printf("# Process Dockerfile.tmpl file and output to Dockerfile.\n")
//		fmt.Printf("\tJsonConfig -json config.json -template DockerFile.tmpl -out Dockerfile\n")
//
//		fmt.Printf("# Process nginx.conf.tmpl file, output to nginx.conf and remove nginx.conf.tmpl afterwards.\n")
//		fmt.Printf("\tJsonConfig -json config.json -create nginx.conf\n")
//
//		fmt.Printf("# Process setup.sh.tmpl file, output to setup.sh and execute as a shell script.\n")
//		fmt.Printf("\tJsonConfig -json config.json -template setup.sh -shell\n")
//
//		fmt.Printf("# Process setup.sh.tmpl file, output to setup.sh, execute as a shell script and remove afterwards.\n")
//		fmt.Printf("\tJsonConfig -json config.json -create setup.sh.tmpl -shell\n")
//	}
//}
//
//func ProcessArgs() (*Template, error) {
//	var err error
//	var tmpl Template
//
//	for range only.Once {
//
//		// Check create flag.
//		if *tmpl.createFlag != "" {
//			*tmpl.outFile = strings.TrimSuffix(*tmpl.createFlag, ".tmpl")
//			*tmpl.templateFile = *tmpl.createFlag
//			*tmpl.removeFiles = true
//		}
//
//		// Verify json args.
//		if (*tmpl.jsonFile == "") && (*tmpl.jsonString == "") {
//			*tmpl.jsonString = "{}"
//		}
//
//		// Verify template args.
//		if (*tmpl.templateFile == "") && (*tmpl.templateString == "") {
//			err = errors.New("need to specify a template file OR template string")
//			break
//		}
//		if (*tmpl.templateFile != "") {
//			// Check template file exists.
//			_, err = os.Stat(*tmpl.templateFile)
//			if os.IsNotExist(err) {
//				break
//			}
//		}
//
//		// Verify output file flags.
//		if *tmpl.overWrite {
//			if *tmpl.outFile != "" {
//				err = errors.New("Need to specify an output file.")
//				break
//			}
//
//			_, err = os.Stat(*tmpl.outFile)
//			if !os.IsNotExist(err) {
//				break
//			}
//		}
//	}
//
//	return &tmpl, err
//}

type FileInfo struct {
	Dir string
	Name string
	CreationEpoch int64
	CreationDate string
}

func (me *Environment) ToString() string {
	var s string

	for range only.Once {
		s = fmt.Sprintf("%s", *me)
	}

	return s
}

func (me *FileInfo) getPaths(f string) error {
	var err error

	for range only.Once {
		var abs string
		abs, err = filepath.Abs(f)
		if err != nil {
			break
		}

		me.Dir = filepath.Dir(abs)
		me.Name = filepath.Base(abs)

		var fstat os.FileInfo
		fstat, err = os.Stat(abs)
		if os.IsNotExist(err) {
			break
		}

		me.CreationEpoch = fstat.ModTime().Unix()
		me.CreationDate = fstat.ModTime().Format("2006-01-02T15:04:05-0700")
	}

	return err
}

func fileToString(fileName string) ([]byte, error) {
	var jsonString []byte
	var err error

	for range only.Once {
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			break
		}

		jsonString, err = ioutil.ReadFile(fileName)
		if err != nil {
			break
		}
	}

	return jsonString, err
}

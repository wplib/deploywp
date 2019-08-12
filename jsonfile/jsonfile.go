package jsonfile

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/util"
	"io/ioutil"
	"log"
	"os"
)

type JsonFile struct {
	DeployWP DeployWP     `json:"deploywp"`
	Site     Site         `json:"site"`
	Source   Source       `json:"source"`
	Targets  Targets      `json:"targets"`
	config   *cfg.Config  `json:"-"`
	rootvar  *TemplateVar `json:"-"`
	rootnode *Node        `json:"-"`
}

func Load(config cfg.Config) (jf *JsonFile) {
	var err error
	for range Once {
		jf = &JsonFile{
			config: &config,
		}
		var b []byte
		b, err = jf.load()
		if err != nil {
			break
		}

		err = json.Unmarshal(b, &jf)
		if err != nil {
			break
		}

		jf.Fixup()

	}
	if err != nil {
		log.Fatalf("Config file '%s' cannot be processed. It is likely invalid JSON or is not using the correct schema: %s.",
			"@TODO: Put acceptable schema number here...",
			err,
		)
	}
	return jf
}

func (me *JsonFile) load() (b []byte, err error) {
	var isnew bool
	fp := me.Filepath()
	for range Once {
		if !util.FileExists(fp) {
			fmt.Printf("A deploy file '%s' does not exist.", fp)
			os.Exit(1)
		}
		b, err = ioutil.ReadFile(fp)
		if err == nil {
			isnew = string(b) == GetDefault()
			break
		}
		b, err = me.makenew(fp)
		isnew = true
	}
	if isnew {
		fmt.Printf("\nYour deploy file '%s' is newly initialized.", fp)
		fmt.Printf("\nPlease EDIT to configure appropriate settings and rerun your command.\n")
		os.Exit(1)
	}
	return b, err
}

func (me *JsonFile) makenew(fp Filepath) (b []byte, err error) {
	var f *os.File
	for range Once {
		f, err = os.Create(fp)
		if err != nil {
			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		var n int
		d := GetDefault()
		n, err = f.WriteString(d)
		if err != nil || n != len(d) {
			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		var size int64
		size, err = f.Seek(0, 2)
		if err != nil || size != int64(len(d)) {
			fmt.Printf("Cannot determine length of deploy file just written '%s'; Check permissions: %s", fp, err)
			os.Exit(1)
		}
		var n64 int64
		n64, err = f.Seek(0, 0)
		if err != nil || n64 != 0 {
			fmt.Printf("Cannot reset deploy file just written '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		b, err = ioutil.ReadAll(f)
		if err != nil || string(b) != d {
			fmt.Printf("Deploy file read does not equal deploy file just written '%s': %s.", fp, err)
			os.Exit(1)
		}
	}
	_ = f.Close()
	return b, err
}

func (me *JsonFile) Filepath() (fp string) {
	return fmt.Sprintf("%s%c%s",
		app.DeployDir,
		os.PathSeparator,
		app.DeployFile,
	)
}

package jsonTemplate

import (
	"github.com/wplib/deploywp/only"
	"io/ioutil"
	"os"
	"path/filepath"
)


type FileInfo struct {
	Dir string
	Name string
	CreationEpoch int64
	CreationDate string
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

package jsonTemplate

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const OnlyOnce = "1"


type FileInfo struct {
	Dir string
	Name string
	CreationEpoch int64
	CreationDate string
}

func (me *FileInfo) getPaths(f string) error {
	var err error

	for range OnlyOnce {
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

	for range OnlyOnce {
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

func _FileToAbs(f ...string) string {
	var ret string

	for range OnlyOnce {
		ret = filepath.Join(f...)

		if filepath.IsAbs(ret) {
			break
		}

		var err error
		ret, err = filepath.Abs(ret)
		if err != nil {
			ret = ""
			break
		}
	}

	return ret
}

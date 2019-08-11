package cfg

import (
	"fmt"
	"os"
)

func HomeDir() (hd Dir) {
	hd, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User home directory not found. Set environment variable HOME and retry.")
		os.Exit(1)
	}
	return hd
}

func dirExists(dir Dir) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func maybeMakeDir(dir Dir, perms os.FileMode) (err error) {
	if !dirExists(dir) {
		err = os.MkdirAll(dir, perms)
	}
	return err
}

func fileExists(file Filepath) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}


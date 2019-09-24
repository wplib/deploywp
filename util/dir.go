package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func FileExists(file Filepath) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func DirExists(dir Dir) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func MaybeMakeDir(dir Dir, perms os.FileMode) (err error) {
	if !DirExists(dir) {
		err = os.MkdirAll(string(dir), perms)
	}
	return err
}

func FileDir(file Filepath) Dir {
	return Dir(filepath.Dir(file))
}

func ParentDir(file Dir) Dir {
	return Dir(filepath.Dir(file))
}

func ExecDir() Dir {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func GetCurrentDir() Dir {
	dir, err := os.Getwd()
	if err != nil {
		Fail("Cannot get current directory")
	}
	return dir
}

var dirRegexp *regexp.Regexp

func init() {
	dirRegexp = regexp.MustCompile(`^~/`)
}

func ExpandDir(dir Dir) Filepath {
	hd := fmt.Sprintf("%s%c", HomeDir(), os.PathSeparator)
	cd := dirRegexp.ReplaceAllString(dir, hd)
	return cd
}

func HomeDir() (hd Dir) {
	hd, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User home directory not found. Set environment variable HOME and retry.")
		os.Exit(1)
	}
	return hd
}

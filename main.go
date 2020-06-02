package main

import (
	"github.com/wplib/deploywp/cmd"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)

func init() {
	_ = ux.Open("Gearbox: ")
}

func main() {
	state := cmd.Execute()
	ux.Close()
	os.Exit(state.ExitCode)
}

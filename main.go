package main

import (
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/wplib/deploywp/cmd"
	"os"
)

//go:generate buildtool pkgreflect deploywp

func init() {
	_, _ = ux.Open("deploywp", true)
}

func main() {
	state := cmd.Execute()
	state.PrintResponse()
	ux.Close()
	os.Exit(state.ExitCode)
}

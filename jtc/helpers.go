package jtc

import (
	"fmt"
	"github.com/wplib/deploywp/jtc/helpers"
	"os"
)


func (at *ArgTemplate) PrintHelpers() {
	_, _ = fmt.Fprintf(os.Stderr, helpers.PrintHelpers())
}

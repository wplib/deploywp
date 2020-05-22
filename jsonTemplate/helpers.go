package jsonTemplate

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers"
	"os"
)


func (at *ArgTemplate) PrintHelpers() {
	_, _ = fmt.Fprintf(os.Stderr, helpers.PrintHelpers())
}

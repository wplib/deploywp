package helperTypes

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)

type TypeErrorGetter interface {
}

type TypeError struct {
	Error error
}


// Usage:
//		{{ if $ret.IsError }}ERROR{{ end }}
func (me *TypeError) SetError(format interface{}, a ...interface{}) {
	for range only.Once {
		f := ReflectString(format)
		if f == nil {
			break
		}

		me.Error = errors.New(fmt.Sprintf(*f, a...))
	}
}


// Usage:
//		{{ if $ret.IsError }}ERROR{{ end }}
func (me *TypeError) IsError() bool {
	var ret bool

	for range only.Once {
		if me.Error == nil {
			break
		}
		ret = true
	}

	return ret
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *TypeError) IsOk() bool {
	return !me.IsError()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *TypeExecCommand) PrintError() string {
	var ret string

	for range only.Once {
		if me.Exit != 0 {
			ret = ux.SprintfRed("ERROR: %s - %s", me.Error, me.Output)
		}
	}

	return ret
}

package deploywp

import (
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/JsonToConfig/ux"
)

type HelperDeployWp TypeDeployWp
func (g *HelperDeployWp) Reflect() *TypeDeployWp {
	return (*TypeDeployWp)(g)
}
func (g *TypeDeployWp) Reflect() *HelperDeployWp {
	return (*HelperDeployWp)(g)
}

func (c *HelperDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func HelperLoadDeployWp(str interface{}, args ...string) *TypeDeployWp {
	j := NewJsonFile()

	for range OnlyOnce {
		var err error

		err = mapstructure.Decode(str, &j)
		j.State.SetError(err)
		if j.State.IsError() {
			break
		}

		j.State = j.Source.Process()
		if j.State.IsError() {
			break
		}

		j.State = j.Target.Process()
		if j.State.IsError() {
			break
		}

		j.State = j.Hosts.Process()
		if j.State.IsError() {
			break
		}

		j.Valid = true
	}

	return j
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.PrintError }}
func (e *TypeDeployWp) PrintError() string {
	return e.State.SprintError()
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.ExitOnError }}
func (e *TypeDeployWp) ExitOnError() string {
	e.State.ExitOnError()
	return ""
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.ExitOnWarning }}
func (e *TypeDeployWp) ExitOnWarning() string {
	e.State.ExitOnWarning()
	return ""
}

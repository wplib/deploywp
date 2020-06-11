package deploywp

import (
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/ux"
)

type HelperDeployWp TypeDeployWp
func (g *HelperDeployWp) Reflect() *TypeDeployWp {
	return (*TypeDeployWp)(g)
}
func (dwp *TypeDeployWp) Reflect() *HelperDeployWp {
	return (*HelperDeployWp)(dwp)
}

func (c *HelperDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func HelperLoadDeployWp(str interface{}, args []string) *TypeDeployWp {
	//dwp := (*TypeDeployWp).New(nil, nil)
	var dwp *TypeDeployWp
	dwp = dwp.New(nil)

	for range onlyOnce {
		var err error

		dwp.Runtime.Args = args[1:]
		err = mapstructure.Decode(str, &dwp)
		dwp.State.SetError(err)
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.Source.Process()
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.Target.Process()
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.Hosts.Process(dwp.Runtime)
		if dwp.State.IsError() {
			break
		}

		dwp.Valid = true
	}

	return dwp
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.PrintError }}
func (dwp *TypeDeployWp) PrintError() string {
	return dwp.State.SprintError()
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.ExitOnError }}
func (dwp *TypeDeployWp) ExitOnError() string {
	dwp.State.ExitOnError()
	return ""
}


// Usage:
//		{{ $cmd := LoadDeployWp }}
//		{{ $cmd.ExitOnWarning }}
func (dwp *TypeDeployWp) ExitOnWarning() string {
	dwp.State.ExitOnWarning()
	return ""
}

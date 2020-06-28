package deploywp

import (
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/ux"
)

type ToolDeployWp TypeDeployWp
func (g *ToolDeployWp) Reflect() *TypeDeployWp {
	return (*TypeDeployWp)(g)
}
func (dwp *TypeDeployWp) Reflect() *ToolDeployWp {
	return (*ToolDeployWp)(dwp)
}

func (c *ToolDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func ToolBuildDeployWp(str interface{}, args []string) string {
	var ret string

	for range onlyOnce {
		dwp := ToolLoadDeployWp(str, args)
		if dwp.State.IsNotOk() {
			dwp.State.PrintResponse()
			break
		}

		dwp.State = dwp.Build()
		if dwp.State.IsNotOk() {
			dwp.State.SetExitCode(1)
			dwp.State.PrintResponse()
			break
		}

		dwp.Valid = true
	}

	return ret
}


func ToolLoadDeployWp(str interface{}, args []string) *TypeDeployWp {
	dwp := (*TypeDeployWp).New(&TypeDeployWp{}, nil)

	for range onlyOnce {
		if str == nil {
			dwp.State.SetError("JSON is empty")
			break
		}

		err := mapstructure.Decode(str, &dwp)
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

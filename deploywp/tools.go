package deploywp

import (
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/ux"
)


type ToolDeployWp TypeDeployWp
func (dwp *ToolDeployWp) Reflect() *TypeDeployWp {
	return (*TypeDeployWp)(dwp)
}
func (dwp *TypeDeployWp) Reflect() *ToolDeployWp {
	return (*ToolDeployWp)(dwp)
}

func (dwp *ToolDeployWp) IsNil() *ux.State {
	if state := ux.IfNilReturnError(dwp); state.IsError() {
		return state
	}
	dwp.State = dwp.State.EnsureNotNil()
	return dwp.State
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

		dwp.State = dwp.Hosts.Process(dwp.Runtime)
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.SelectDestinationHost()
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.Source.Process()
		if dwp.State.IsError() {
			break
		}

		dwp.State = dwp.Destination.Process()
		if dwp.State.IsError() {
			break
		}

		dwp.Valid = true
	}

	return dwp
}


func (dwp *TypeDeployWp) ExitOnError() string {
	dwp.State.ExitOnError()
	return ""
}


func (dwp *TypeDeployWp) ExitOnWarning() string {
	dwp.State.ExitOnWarning()
	return ""
}

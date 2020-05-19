package deploywp

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// Usage:
//		{{ $cmd := OpenSourceRepo }}
//		{{ $cmd.ExitOnWarning }}
func (e *TypeDeployWp) OpenSourceRepo() *ux.State {
	for range only.Once {
		if e.IsNil() {
			e.State.SetError("deploywp JSON is nil")
			break
		}

		e.State = e.ObtainHost()
		if e.State.IsError() {
			break
		}

		if e.State.Response == "" {

		}
	}

	return e.State
}


// Usage:
//		{{ $state := ObtainHost }}
//		{{ $state.ExitOnWarning }}
func (e *TypeDeployWp) ObtainHost() *ux.State {
	for range only.Once {
		if e.IsNil() {
			e.State.SetError("deploywp JSON is nil")
			break
		}

		var host string
		for range only.Once {
			e.State.Clear()

			host = e.Exec.GetArg(0)
			if host != "" {
				break
			}

			host = helperSystem.HelperUserPrompt("Enter host: ")
			if host != "" {
				break
			}

			e.State.SetError("host is empty")
		}
		if e.State.IsError() {
			break
		}

		e.State.SetOutput(host)
		e.State.Response = host
		e.State.Clear()
	}

	return e.State
}

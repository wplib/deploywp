package helperExec

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/wplib/deploywp/ux"
	"testing"
)

func TestHelperExecCommand(t *testing.T) {

	for _, c := range []struct {
		cmd []string
		expected ux.State
	}{
		{[]string{"ps", "-eaf"}, ux.State {
			ExitCode:    0,
			Output:      "",
			OutputArray: nil,
			Response:    nil,
		}},
		{[]string{"/usr/bin/false"}, ux.State {
			ExitCode:    1,
			Output:      "",
			OutputArray: nil,
			Response:    nil,
		}},
	} {
		returned := HelperExecCmd(c.cmd)
		if returned.ExitCode != c.expected.ExitCode {
			t.Errorf("%s(%q) == %q, want %q", t.Name(), c.cmd, returned.ExitCode, c.expected.ExitCode)
			spew.Dump(c.expected)
			spew.Dump(returned)
		}
	}
}

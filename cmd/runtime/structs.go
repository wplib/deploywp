package runtime

import (
	"strings"
)

const OnlyOnce = "1"

type Exec struct {
	CmdVersion string
	Cmd        string
	CmdDir     string
	CmdFile    string
	FullArgs   ExecArgs
	Args       ExecArgs
}

type ExecArgs []string


func (me *ExecArgs) ToString() string {
	return strings.Join(*me, " ")
}

func (me *Exec) GetFullArgs() ExecArgs {
	return me.FullArgs
}

func (me *Exec) GetArgs() ExecArgs {
	return me.Args
}

func (me *Exec) GetArg(index int) string {
	var ret string

	for range OnlyOnce {
		if len(me.Args) > index {
			ret = me.Args[index]
		}
	}

	return ret
}

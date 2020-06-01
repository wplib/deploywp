package runtime

import (
	"github.com/wplib/deploywp/defaults"
	"github.com/newclarity/JsonToConfig/ux"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const OnlyOnce = "1"


type Exec struct {
	CmdName     string		`json:"cmd_name" mapstructure:"cmd_name"`
	CmdVersion  string		`json:"cmd_version" mapstructure:"cmd_version"`
	Cmd         string		`json:"cmd" mapstructure:"cmd"`
	CmdDir      string		`json:"cmd_dir" mapstructure:"cmd_dir"`
	CmdFile     string		`json:"cmd_file" mapstructure:"cmd_file"`

	FullArgs    ExecArgs	`json:"full_args" mapstructure:"full_args"`
	Args        ExecArgs	`json:"args" mapstructure:"args"`

	Env         ExecEnv		`json:"env" mapstructure:"env"`
	EnvMap      Environment	`json:"env_map" mapstructure:"env_map"`

	TimeStamp   time.Time	`json:"timestamp" mapstructure:"timestamp"`

	State       *ux.State	`json:"state" mapstructure:"state"`
}

type ExecArgs []string
type ExecEnv []string
type Environment map[string]string


func NewExec() *Exec {
	var ret Exec

	for range OnlyOnce {
		var err error

		ret.State = ux.NewState(false)
		ret.State.SetPackage("")
		ret.State.SetFunction("")

		ret.CmdName = defaults.BinaryName
		ret.CmdVersion = defaults.BinaryVersion

		ret.Cmd, err = os.Executable()
		if err != nil {
			ret.State.SetError(err)
			break
		}

		ret.Cmd, err = filepath.Abs(ret.Cmd)
		if err != nil {
			ret.State.SetError(err)
			break
		}

		ret.CmdDir = path.Dir(ret.Cmd)
		ret.CmdFile = path.Base(ret.Cmd)

		ret.FullArgs = os.Args[1:]
		ret.Args = ret.FullArgs

		ret.Env = os.Environ()
		ret.EnvMap = make(Environment)
		for _, item := range os.Environ() {
			s := strings.SplitN(item, "=", 2)
			ret.EnvMap[s[0]] = s[1]
		}

		ret.TimeStamp= time.Now()
		//ret.Epoch = now.Unix()
		//ret.TimeStamp = now.Format("2006-01-02T15:04:05-0700")
	}

	return &ret
}


func (me *ExecArgs) ToString() string {
	return strings.Join(*me, " ")
}


func (me *Exec) TimeStampString() string {
	return me.TimeStamp.Format("2006-01-02T15:04:05-0700")
}


func (me *Exec) TimeStampEpoch() int64 {
	return me.TimeStamp.Unix()
}



func (e *Exec) GetEnvMap() *Environment {
	return &e.EnvMap
}


func (e *Exec) GetArg(index int) string {
	var ret string

	for range OnlyOnce {
		if len(e.Args) > index {
			ret = e.Args[index]
		}
	}

	return ret
}


func (e *Exec) SetArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		e.Args = a
	}

	return err
}


func (e *Exec) GetArgs() []string {
	return e.Args
}


func (e *Exec) AddArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		e.Args = append(e.Args, a...)
	}

	return err
}


func (e *Exec) SetFullArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		e.FullArgs = a
	}

	return err
}


func (e *Exec) GetFullArgs() []string {
	return e.FullArgs
}


func (e *Exec) AddFullArgs(a ...string) error {
	var err error

	for range OnlyOnce {
		e.FullArgs = append(e.FullArgs, a...)
	}

	return err
}

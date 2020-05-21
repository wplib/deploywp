package helperSystem

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"os"
	"strings"
)

type Environment map[string]string


// FindInMap function.
// func PrintEnv(ex []string) string {
func PrintEnv() string {
	var ret string

	for range OnlyOnce {
		var env Environment
		var err error
		env, err = GetEnv()
		if err != nil {
			break
		}

		for k, v := range env {
			// Bit of a hack for now...
			// Will strip out env for Docker init
			switch {
			case k == "MAIL":
			case k == "HOME":
			case k == "LOGNAME":
			case k == "PATH":
			case k == "PWD":
			case k == "SHELL":
			case k == "SHLVL":
			case k == "USER":
			case k == "_":

			default:
				ret += fmt.Sprintf("%s=\"%s\"; export %s\n", k, v, k)
			}
		}
	}

	return ret
}

func GetEnv() (Environment, error) {
	var e Environment
	var err error

	for range OnlyOnce {
		e = make(Environment)
		for _, item := range os.Environ() {
			s := strings.SplitN(item, "=", 2)
			e[s[0]] = s[1]
		}
	}

	return e, err
}

func (me *Environment) ToString() string {
	var s string

	for range OnlyOnce {
		s = fmt.Sprintf("%s", *me)
	}

	return s
}

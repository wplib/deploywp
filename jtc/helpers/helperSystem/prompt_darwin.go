// +build darwin

package helperSystem

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)


func (me *Prompt) UserPrompt() string {
	var ret string

	for range OnlyOnce {
		fmt.Printf("%s", *me)

		r := bufio.NewReader(os.Stdin)

		var err error
		ret, err = r.ReadString('\n')
		fmt.Printf("\n")
		if err != nil {
			break
		}

		ret = strings.TrimSuffix(ret, "\n")
	}

	return ret
}


func (me *Prompt) UserPromptHidden() string {
	var ret string

	for range OnlyOnce {
		fmt.Printf("%s", *me)

		hidden, err := terminal.ReadPassword(syscall.Stdin)
		fmt.Printf("\n")
		if err != nil {
			break
		}

		ret = strings.TrimSuffix(string(hidden), "\n")
	}

	return ret
}

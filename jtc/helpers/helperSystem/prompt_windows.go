// +build windows

package helperSystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

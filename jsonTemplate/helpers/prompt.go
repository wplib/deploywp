package helpers

import (
	"bufio"
	"fmt"
	"github.com/wplib/deploywp/only"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)


func UserPrompt(prompt interface{}, args ...interface{}) string {
	var ret string

	for range only.Once {
		var err error

		p := ReflectString(prompt)
		if p == nil {
			break
		}
		fmt.Printf(*p, args...)

		r := bufio.NewReader(os.Stdin)

		ret, err = r.ReadString('\n')
		fmt.Printf("\n")
		if err != nil {
			break
		}
	}

	return ret
}


func UserPromptHidden(prompt interface{}, args ...interface{}) string {
	var ret string

	for range only.Once {
		var err error

		p := ReflectString(prompt)
		if p == nil {
			break
		}
		fmt.Printf(*p, args...)

		var hidden []byte
		hidden, err = terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")
		if err != nil {
			break
		}
		ret = string(hidden)
	}

	return ret
}

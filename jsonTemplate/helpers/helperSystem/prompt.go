package helperSystem

import (
	"bufio"
	"fmt"
	"github.com/wplib/deploywp/only"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)


// Usage:
//		{{ $str := UserPrompt "Enter something %s:" "here" }}
func HelperUserPrompt(prompt string, args ...interface{}) string {
	var ret string

	for range only.Once {
		fmt.Printf(prompt, args...)

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


// Usage:
//		{{ $str := UserPromptHidden "Enter something %s:" "here" }}
func HelperUserPromptHidden(prompt string, args ...interface{}) string {
	var ret string

	for range only.Once {
		fmt.Printf(prompt, args...)

		hidden, err := terminal.ReadPassword(syscall.Stdin)
		fmt.Printf("\n")
		if err != nil {
			break
		}

		ret = strings.TrimSuffix(string(hidden), "\n")
	}

	return ret
}


// Usage:
//		{{ $str := UserPrompt "Enter something %s:" "here" }}
func HelperUserPromptBool(prompt string, args ...interface{}) bool {
	var ret bool

	for range only.Once {
		fmt.Printf(prompt, args...)

		r := bufio.NewReader(os.Stdin)

		str, err := r.ReadString('\n')
		fmt.Printf("\n")
		if err != nil {
			break
		}
		str = strings.TrimSpace(str)

		switch strings.ToUpper(str) {
			case "TRUE":
				fallthrough
			case "YES":
				fallthrough
			case "Y":
				ret = true
		}
	}

	return ret
}

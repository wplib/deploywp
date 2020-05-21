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

	for range OnlyOnce {
		fmt.Printf(prompt, args...)

		r := bufio.NewReader(os.Stdin)

		var err error
		ret, err = r.ReadString('\n')
		fmt.Printf("\n")
		if err != nil {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $str := UserPromptHidden "Enter something %s:" "here" }}
func HelperUserPromptHidden(prompt string, args ...interface{}) string {
	var ret string

	for range OnlyOnce {
		fmt.Printf(prompt, args...)

		hidden, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")
		if err != nil {
			break
		}
		ret = string(hidden)
	}

	return ret
}


// Usage:
//		{{ $str := UserPrompt "Enter something %s:" "here" }}
func HelperUserPromptBool(prompt string, args ...interface{}) bool {
	var ret bool

	for range OnlyOnce {
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

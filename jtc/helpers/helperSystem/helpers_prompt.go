package helperSystem

import (
	"strings"
)


// Usage:
//		{{ $str := UserPrompt "Enter something %s:" "here" }}
func HelperUserPrompt(prompt string, args ...interface{}) string {
	return UserPrompt(prompt, args...)
}


// Usage:
//		{{ $str := UserPromptHidden "Enter something %s:" "here" }}
func HelperUserPromptHidden(prompt string, args ...interface{}) string {
	return UserPromptHidden(prompt, args...)
}


// Usage:
//		{{ $str := UserPrompt "Enter something %s:" "here" }}
func HelperUserPromptBool(prompt string, args ...interface{}) bool {
	var ret bool

	for range OnlyOnce {
		str := UserPrompt(prompt, args...)
		str = strings.TrimSpace(str)
		str = strings.ToUpper(str)
		switch str {
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

package helperSystem

import (
	"fmt"
)

// @TODO - This is a workaround for the duplicates that appear with pkgreflect.

type Prompt string


func UserPrompt(prompt string, args ...interface{}) string {
	var p Prompt
	p.Set(prompt, args)
	return p.UserPrompt()
}
func (me *Prompt) Set(prompt string, args ...interface{}) {
	*me = Prompt(fmt.Sprintf(prompt, args...))
}


func UserPromptHidden(prompt string, args ...interface{}) string {
	var p Prompt
	p.Set(prompt, args)
	return p.UserPromptHidden()
}

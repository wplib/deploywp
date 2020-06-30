// Simple Ux.Print indenting mechanism.
package deploywp

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"runtime"
	"strings"
)


type UxPrint struct {
	verbose bool
	newLine bool
	//indentValue int
	last *caller

	alreadyPrinted bool
	callOrder []caller
}
type caller struct {
	Func string
	Text string
	SubPrint int
	Indent int
}


func (up *UxPrint) Intent(format string, args ...interface{}) {
	str, _ := up.sprintf(format + " ... ", args...)
	str = up.indent() + str
	ux.PrintfBlue(str)
}


func (up *UxPrint) IntentResponse(state *ux.State) {
	for range onlyOnce {
		switch {
			case state.IsOk():
				format := state.GetOk().Error()
				if format == "" {
					format = "OK"
				}
				var noprint bool
				format, noprint = up.sprintf(format)
				if noprint {
					break
				}
				ux.PrintfGreen(format)

			case state.IsWarning():
				format := state.GetWarning().Error()
				if format == "" {
					format = "WARNING"
				}
				var noprint bool
				format, noprint = up.sprintf(format)
				if noprint {
					break
				}
				ux.PrintfYellow(format)

			case state.IsError():
				format := state.GetError().Error()
				if format == "" {
					format = "ERROR"
				}
				var noprint bool
				format, noprint = up.sprintf(format)
				if noprint {
					break
				}
				ux.PrintfRed(format)
		}
	}
}


func (up *UxPrint) IntentAppend(format string, args ...interface{}) {
	for range onlyOnce {
		if len(up.callOrder) == 0 {
			up.new()
		}
		txt := fmt.Sprintf(format + " ... ", args...)

		callerName := up.getCaller(parent)
		up.last = up.getLast()
		if up.last.Func == callerName {
			// Last caller is the same.
			up.last.Text = txt
			ux.PrintfBlue(up.indent() + txt)
			break
		}

		index := up.findCaller(callerName)
		if index > 0 {
			up.trimCaller(index+1)
			ux.PrintfBlue(up.indent() + txt)
			break
		}

		up.addCaller(caller{Func: callerName, Text: txt, SubPrint: 0, Indent: up.last.Indent + 1})
		ux.PrintfBlue(txt)
	}
}


func (up *UxPrint) sprintf(format string, args ...interface{}) (string, bool) {
	var txt string
	var noprint bool
	for range onlyOnce {
		if len(up.callOrder) == 0 {
			up.new()
		}
		txt = fmt.Sprintf(format, args...)

		callerName := up.getCaller(parent)
		up.last = up.getLast()
		if up.last.Func == callerName {
			// Last caller is the same.
			if up.last.Text == txt {
				noprint = true
				break
			}
			up.last.Text = txt
			break
		}

		index := up.findCaller(callerName)
		if index > 0 {
			up.trimCaller(index+1)
			noprint = true
			break
		}

		callersCaller := up.getCaller(parentsParent)
		index = up.findCaller(callersCaller)
		if index == 0 {
			up.new()
			up.last = up.getLast()
		}

        up.addCaller(caller{Func: callerName, Text: txt, SubPrint: 0, Indent: up.last.Indent + 1})
	}
	return txt, noprint
}


const parent = 3
const parentsParent = 4
func (up *UxPrint) getCaller(count int) string {
	pc, file, _, _ := runtime.Caller(count)
	details := runtime.FuncForPC(pc)
	file = fmt.Sprintf("%s", details.Name())
	return file
}


func (up *UxPrint) new() {
	up.callOrder = []caller{}
	up.callOrder = append(up.callOrder, caller{})
}


func (up *UxPrint) isLast(fn string) bool {
	var ok bool
	for range onlyOnce {
		l := len(up.callOrder)
		if l == 0 {
			break
		}
		if up.callOrder[l-1].Func != fn {
			break
		}
		ok = true
	}
	return ok
}


func (up *UxPrint) findCaller(name string) int {
	var c int
	for c = len(up.callOrder)-1; c > 0; c-- {
		if up.callOrder[c].Func == name {
			break
		}
	}
	return c
}


func (up *UxPrint) addCaller(c caller) {
	up.last.SubPrint++
	up.callOrder = append(up.callOrder, c)
	up.last = up.getLast()
}


func (up *UxPrint) trimCaller(index int) {
	up.callOrder = up.callOrder[:index]
	up.last = up.getLast()
}


func (up *UxPrint) getLast() *caller {
	if len(up.callOrder) == 0 {
		return &caller{}
	}
	return &(up.callOrder[len(up.callOrder)-1])
}


func (up *UxPrint) indent() string {
	c := up.getLast().Indent
	if c > 0 {
		c--
	}
	//if c > 0 {
	//	c--
	//}
	c2 := len(up.callOrder)
	if c2 > 0 {
		c2--
	}
	if c2 > 0 {
		c2--
	}

	//c := up.getLast().Indent
	return up.indentPrint(c)
}


func (up *UxPrint) indentPrint(c int) string {
	return fmt.Sprintf("\n%s- ", strings.Repeat("\t", c))
}


func (up *UxPrint) Notify(indent int, format string, args ...interface{}) {
	var prefix string
	var suffix string
	if indent >= 0 {
		prefix = up.indentPrint(indent)
		suffix = " ... "
	}
	ux.PrintfBlue(prefix + format + suffix, args...)
	up.alreadyPrinted = false
	return
}


func (up *UxPrint) Append(format string, args ...interface{}) {
	ux.PrintfBlue(format + ":", args...)
	up.alreadyPrinted = false
	return
}


func (up *UxPrint) PrintResponse(state *ux.State) {
	for range onlyOnce {
		if up.alreadyPrinted {
			break
		}
		switch {
		case state.IsOk():
			format := state.GetOk().Error()
			if format == "" {
				format = "OK"
			}
			ux.PrintfGreen(format)

		case state.IsWarning():
			format := state.GetWarning().Error()
			if format == "" {
				format = "WARNING"
			}
			ux.PrintfYellow(format)

		case state.IsError():
			format := state.GetError().Error()
			if format == "" {
				format = "ERROR"
			}
			ux.PrintfRed(format)
		}
		up.alreadyPrinted = true
	}
}


func (up *UxPrint) Ok(indent int, format string, args ...interface{}) {
	if up.verbose {
		if indent == 0 {
			up.indentPrint(indent)
		}
		ux.PrintfOk(format, args...)
		up.alreadyPrinted = false
	}
	return
}


func (up *UxPrint) Warning(indent int, format string, args ...interface{}) {
	if indent == 0 {
		up.indentPrint(indent)
	}
	ux.PrintfWarning(format, args...)
	up.alreadyPrinted = false
	return
}


func (up *UxPrint) Error(indent int, format string, args ...interface{}) {
	if indent == 0 {
		up.indentPrint(indent)
	}
	ux.PrintfError(format, args...)
	up.alreadyPrinted = false
	return
}

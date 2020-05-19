package ux

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/only"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)


type StateGetter interface {
	Print()
	IsError() bool
	IsWarning() bool
	IsOk() bool
	SetError(format string, args ...interface{})
	SetWarning(format string, args ...interface{})
	SetOk(format string, args ...interface{})
	ClearError()
	ClearAll()
	IsRunning() bool
	IsPaused() bool
	IsCreated() bool
	IsRestarting() bool
	IsRemoving() bool
	IsExited() bool
	IsDead() bool
	SetString(s string)
}

type State struct {
	prefix      string
	prefixArray []string
	_Package    string
	_Function   string

	_Error       error
	_Warning     error
	_Ok          error
	ExitCode    int

	Output      string
	_Separator  string
	OutputArray []string
	Response    interface{}
}

const DefaultSeparator = "\n"


func NewState() *State {
	me := State{}
	me.Clear()
	return &me
}


func (p *State) Clear() {
	p._Error = nil
	p._Warning = nil
	p._Ok = errors.New("")
	p.ExitCode = 0

	p.Output = ""
	p._Separator = DefaultSeparator
	p.OutputArray = []string{}
	p.Response = nil
}


func (p *State) GetPrefix() string {
	return p.prefix
}
func (p *State) GetPackage() string {
	return p._Package
}
func (p *State) GetFunction() string {
	return p._Function
}
func (p *State) SetPackage(s string) {
	if s == "" {
		// Discover package name.
		//pc, file, no, ok := runtime.Caller(1)
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			//s = file + ":" + string(no)
			details := runtime.FuncForPC(pc)
			s = filepath.Base(details.Name())
			sa := strings.Split(s, ".")
			if len(sa) > 0 {
				s = sa[0]
			}
		}
	}

	p._Package = s
	if p._Function == "" {
		p.prefix = p._Package
	} else {
		p.prefix = p._Package + "." + p._Function + "()"
		p.prefixArray = append(p.prefixArray, p.prefix)
	}
}
func (p *State) SetFunction(s string) {
	if s == "" {
		// Discover function name.
		//pc, file, no, ok := runtime.Caller(1)
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			//s = file + ":" + string(no)
			details := runtime.FuncForPC(pc)
			foo := details.Name()
			s = filepath.Base(foo)
			sa := strings.Split(s, ".")
			switch {
				case len(sa) > 2:
					s = sa[2]
				case len(sa) > 1:
					s = sa[1]
				case len(sa) > 0:
					s = sa[0]
			}
		}
	}

	p._Function = s
	if p._Package == "" {
		p.prefix = p._Function + "()"
	} else {
		p.prefix = p._Package + "." + p._Function + "()"
	}

	p.prefixArray = append(p.prefixArray, p.prefix)
}
func (p *State) SetFunctionCaller() {
	var s string
	// Discover function name.
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		//s = file + ":" + string(no)
		details := runtime.FuncForPC(pc)
		s = filepath.Base(details.Name())
		sa := strings.Split(s, ".")
		if len(sa) > 0 {
			s = sa[1]
		}
	}

	p.SetFunction(s)
}


func (p *State) GetState() *bool {
	var b bool
	//s := &State{
	//	_Error:      p._Error,
	//	_Warning:    p._Warning,
	//	_Ok:         p._Ok,
	//	ExitCode:    p.ExitCode,
	//	Output:      p.Output,
	//	OutputArray: p.OutputArray,
	//	Response:    p.Response,
	//}
	return &b
}
func (s *State) SetState(p *State) {
	s._Error =      p._Error
	s._Warning =    p._Warning
	s._Ok =         p._Ok
	s.ExitCode =    p.ExitCode
	s.Output =      p.Output
	s.OutputArray = p.OutputArray
	s.Response =    p.Response
}


func (p *State) Sprint() string {
	var ret string

	e := ""
	if p.ExitCode != 0 {
		e = fmt.Sprintf("Exit(%d) - ", p.ExitCode)
	}

	pa := ""
	if len(p.prefixArray) > 0 {
		pa = fmt.Sprintf("[%s] - ", p.prefixArray[0])
	}

	switch {
		case p._Error != nil:
			ret = SprintfError("ERROR: %s%s%s", pa, e, p._Error)
		case p._Warning != nil:
			ret = SprintfWarning("WARNING: %s%s%s", pa, e, p._Warning)
		case p._Ok != nil:
			ret = SprintfOk("%s", p._Ok)
	}

	if p.Output != "" {
		ret += SprintfOk("\n%s ", p.Output)
	}

	return ret
}
func (p *State) PrintResponse() string {
	return p.Sprint()
}
func (p *State) SprintError() string {
	var ret string

	for range only.Once {
		if p._Ok != nil {
			// If we have an OK response.
			break
		}

		ret = p.Sprint()
	}

	return ret
}


func (p *State) IsError() bool {
	var ok bool

	if p == nil {
		fmt.Printf("DUH\n")
		return ok
	}
	if p._Error != nil {
		ok = true
	}

	return ok
}

func (p *State) IsWarning() bool {
	var ok bool

	if p._Warning != nil {
		ok = true
	}

	return ok
}

func (p *State) IsOk() bool {
	var ok bool

	if p._Ok != nil {
		ok = true
	}

	return ok
}
func (p *State) IsNotOk() bool {
	ok := true

	for range only.Once {
		if p._Warning != nil {
			break
		}
		if p._Error != nil {
			break
		}
		ok = false
	}

	return ok
}

func (p *State) SetExitCode(e int) {
	p.ExitCode = e
}
func (p *State) GetExitCode() int {
	return p.ExitCode
}


func (p *State) SetError(error ...interface{}) {
	for range only.Once {
		if p == nil {
			//p._Error = errors.New("ERROR State is nil")
			break
		}

		p._Ok = nil
		p._Warning = nil

		if len(error) == 0 {
			//p._Error = errors.New(p.prefix + "ERROR")
			p._Error = errors.New("ERROR")
			break
		}

		if error[0] == nil {
			p._Error = nil
			break
		}

		//p._Error = errors.New(p.prefix + _Sprintf(error...))
		p._Error = errors.New(_Sprintf(error...))
	}
}
func (p *State) GetError() error {
	return p._Error
}


func (p *State) SetWarning(warning ...interface{}) {
//func (p *State) SetWarning(format string, args ...interface{}) {
//	p._Ok = nil
//	p._Warning = errors.New(fmt.Sprintf(format, args...))
//	p._Error = nil

	for range only.Once {
		if p == nil {
			//p._Error = errors.New("ERROR State is nil")
			break
		}

		p._Ok = nil
		p._Error = nil

		if len(warning) == 0 {
			//p._Warning = errors.New(p.prefix + "WARNING")
			p._Warning = errors.New("WARNING")
			break
		}

		if warning[0] == nil {
			p._Error = nil
			break
		}

		//p._Warning = errors.New(p.prefix + _Sprintf(warning...))
		p._Warning = errors.New(_Sprintf(warning...))
	}
}
func (p *State) GetWarning() error {
	return p._Warning
}


func (p *State) SetOk(msg ...interface{}) {
// func (p *State) SetOk(format string, args ...interface{}) {
//	p._Ok = errors.New(fmt.Sprintf(format, args...))
//	p._Warning = nil
//	p._Error = nil

	for range only.Once {
		if p == nil {
			//p._Error = errors.New("ERROR State is nil")
			break
		}

		p._Error = nil
		p._Warning = nil
		p.ExitCode = 0

		if len(msg) == 0 {
			p._Ok = errors.New("")
			break
		}

		if msg[0] == nil {
			p._Error = nil
			break
		}

		//p._Ok = errors.New(p.prefix + _Sprintf(msg...))
		p._Ok = errors.New(_Sprintf(msg...))
	}
}
func (p *State) GetOk() error {
	return p._Ok
}


func (p *State) ClearError() {
	p._Error = nil
}


func (p *State) IsRunning() bool {
	var ok bool
	if p.Output == StateRunning {
		ok = true
	}
	return ok
}

func (p *State) IsPaused() bool {
	var ok bool
	if p.Output == StatePaused {
		ok = true
	}
	return ok
}

func (p *State) IsCreated() bool {
	var ok bool
	if p.Output == StateCreated {
		ok = true
	}
	return ok
}

func (p *State) IsRestarting() bool {
	var ok bool
	if p.Output == StateRestarting {
		ok = true
	}
	return ok
}

func (p *State) IsRemoving() bool {
	var ok bool
	if p.Output == StateRemoving {
		ok = true
	}
	return ok
}

func (p *State) IsExited() bool {
	var ok bool
	if p.Output == StateExited {
		ok = true
	}
	return ok
}

func (p *State) IsDead() bool {
	var ok bool
	if p.Output == StateDead {
		ok = true
	}
	return ok
}


func (p *State) ExitOnNotOk() string {
	if p.IsNotOk() {
		_, _ = fmt.Fprintf(os.Stderr, p.Sprint() + "\n")
		os.Exit(p.ExitCode)
	}
	return ""
}


func (p *State) ExitOnError() string {
	if p.IsWarning() {
		_, _ = fmt.Fprintf(os.Stderr, p.Sprint() + "\n")
	}
	if p.IsError() {
		_, _ = fmt.Fprintf(os.Stderr, p.Sprint() + "\n")
		os.Exit(p.ExitCode)
	}
	return ""
}


func (p *State) ExitOnWarning() string {
	if p.IsWarning() {
		_, _ = fmt.Fprintf(os.Stderr, p.Sprint() + "\n")
		os.Exit(p.ExitCode)
	}
	return ""
}


func (p *State) Exit(e int) string {
	p.ExitCode = e
	_, _ = fmt.Fprintf(os.Stdout, p.Sprint())
	os.Exit(p.ExitCode)
	return ""
}


func Exit(e int64, msg ...interface{}) string {
	ret := _Sprintf(msg...)
	if e == 0 {
		_, _ = fmt.Fprintf(os.Stdout, SprintfOk(ret))
	} else {
		_, _ = fmt.Fprintf(os.Stderr, SprintfError(ret))
	}
	os.Exit(int(e))
	return ""	// Will never get here.
}


func _Sprintf(msg ...interface{}) string {
	var ret string

	for range only.Once {
		if len(msg) == 0 {
			break
		}

		value := reflect.ValueOf(msg[0])
		switch value.Kind() {
			case reflect.String:
				if len(msg) == 1 {
					ret = fmt.Sprintf(msg[0].(string))
				} else {
					ret = fmt.Sprintf(msg[0].(string), msg[1:]...)
				}

			default:
				if len(msg) == 1 {
					ret = fmt.Sprintf("%v", msg)
				} else {
					var es string
					for _, e := range msg {
						es += fmt.Sprintf("%v ", e)
					}
					es = strings.TrimSuffix(es, " ")
					ret = es
				}
		}

		//ret = fmt.Sprintf(msg[0].(string), msg[1:]...)
	}

	return ret
}

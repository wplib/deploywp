package ux

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/only"
	"os"
	"reflect"
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
	Prefix   string

	Error    error
	Warning  error
	Ok       error
	Output   string
	ExitCode int
	Response interface{}
}


func New() *State {
	me := State{}
	me.Clear()
	return &me
}


func (me *State) Clear() {
	me.Error = nil
	me.Warning = nil
	me.Ok = errors.New("")

	me.ExitCode = 0
	me.Output = ""
	me.Response = nil
}


// Ability to set State from an arbitrary interface type.
//func (me *State) SetState(state interface{}) *State {
//	for range only.Once {
//
//		me
//		switch {
//		case me.Error != nil:
//			PrintfError("%s", me.Error)
//		case me.Warning != nil:
//			PrintfWarning("%s", me.Warning)
//		case me.Ok != nil:
//			PrintfOk("%s", me.Ok)
//		}
//	}
//}


func (me *State) Sprint() string {
	var ret string

	e := ""
	if me.ExitCode != 0 {
		e = fmt.Sprintf("Exit(%d) - ", me.ExitCode)
	}

	switch {
		case me.Error != nil:
			ret = SprintfError("ERROR: %s%s", e, me.Error)
		case me.Warning != nil:
			ret = SprintfWarning("WARNING: %s%s", e, me.Warning)
		case me.Ok != nil:
			ret = SprintfOk("%s", me.Ok)
	}

	if me.Output != "" {
		ret += SprintfOk("\n%s ", me.Output)
	}

	return ret
}
func (me *State) SprintError() string {
	var ret string

	for range only.Once {
		if me.Ok != nil {
			// If we have an OK response.
			break
		}

		ret = me.Sprint()
	}

	return ret
}


func (me *State) IsError() bool {
	var ok bool

	if me.Error != nil {
		ok = true
	}

	return ok
}

func (me *State) IsWarning() bool {
	var ok bool

	if me.Warning != nil {
		ok = true
	}

	return ok
}

func (me *State) IsOk() bool {
	var ok bool

	if me.Ok != nil {
		ok = true
	}

	return ok
}


func (me *State) SetExitCode(e int) {
	me.ExitCode = 0
}
func (me *State) GetExitCode() int {
	return me.ExitCode
}


func (me *State) SetError(error ...interface{}) {
	for range only.Once {
		me.Ok = nil
		me.Warning = nil

		if len(error) == 0 {
			me.Error = errors.New("ERROR")
			break
		}

		me.Error = errors.New(_Sprintf(error...))
	}
}
func (me *State) GetError() error {
	return me.Error
}


func (me *State) SetWarning(warning ...interface{}) {
//func (me *State) SetWarning(format string, args ...interface{}) {
//	me.Ok = nil
//	me.Warning = errors.New(fmt.Sprintf(format, args...))
//	me.Error = nil

	for range only.Once {
		me.Ok = nil
		me.Error = nil

		if len(warning) == 0 {
			me.Warning = errors.New("WARNING")
			break
		}

		me.Warning = errors.New(_Sprintf(warning...))
	}
}
func (me *State) GetWarning() error {
	return me.Warning
}


func (me *State) SetOk(msg ...interface{}) {
// func (me *State) SetOk(format string, args ...interface{}) {
//	me.Ok = errors.New(fmt.Sprintf(format, args...))
//	me.Warning = nil
//	me.Error = nil

	for range only.Once {
		me.Error = nil
		me.Warning = nil
		me.ExitCode = 0

		if len(msg) == 0 {
			me.Ok = errors.New("")
			break
		}

		me.Ok = errors.New(_Sprintf(msg...))
	}
}
func (me *State) GetOk() error {
	return me.Ok
}


func (me *State) ClearError() {
	me.Error = nil
}


func (me *State) IsRunning() bool {
	var ok bool
	if me.Output == StateRunning {
		ok = true
	}
	return ok
}

func (me *State) IsPaused() bool {
	var ok bool
	if me.Output == StatePaused {
		ok = true
	}
	return ok
}

func (me *State) IsCreated() bool {
	var ok bool
	if me.Output == StateCreated {
		ok = true
	}
	return ok
}

func (me *State) IsRestarting() bool {
	var ok bool
	if me.Output == StateRestarting {
		ok = true
	}
	return ok
}

func (me *State) IsRemoving() bool {
	var ok bool
	if me.Output == StateRemoving {
		ok = true
	}
	return ok
}

func (me *State) IsExited() bool {
	var ok bool
	if me.Output == StateExited {
		ok = true
	}
	return ok
}

func (me *State) IsDead() bool {
	var ok bool
	if me.Output == StateDead {
		ok = true
	}
	return ok
}


// "created", "running", "paused", "restarting", "removing", "exited", or "dead"
func (me *State) SetString(s string) {
	me.Output = s
}
func (me *State) GetString() string {
	return me.Output
}


func (me *State) ExitOnError() string {
	if me.IsError() {
		_, _ = fmt.Fprintf(os.Stderr, me.Sprint())
		os.Exit(me.ExitCode)
	}
	return ""
}


func (me *State) ExitOnWarning() string {
	if me.IsWarning() {
		_, _ = fmt.Fprintf(os.Stderr, me.Sprint())
		os.Exit(me.ExitCode)
	}
	return ""
}


func (me *State) Exit(e int) string {
	_, _ = fmt.Fprintf(os.Stdout, me.Sprint())
	os.Exit(me.ExitCode)
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
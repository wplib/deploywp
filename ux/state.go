package ux

import (
	"errors"
	"fmt"
	"github.com/wplib/deploywp/only"
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
	Error error
	Warning error
	Ok error
	String string
}


func New() *State {
	return &State {
		Error:   nil,
		Warning: nil,
		Ok:      nil,
		String:  "",
	}
}


func (me *State) Print() {
	switch {
		case me.Error != nil:
			PrintfError("%s", me.Error)
		case me.Warning != nil:
			PrintfWarning("%s", me.Warning)
		case me.Ok != nil:
			PrintfOk("%s", me.Ok)
	}
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


func (me *State) SetError(error ...interface{}) {
	for range only.Once {
		me.Ok = nil
		me.Warning = nil

		if len(error) == 0 {
			me.Error = errors.New("ERROR")
			break
		}

		value := reflect.ValueOf(error[0])
		switch value.Kind() {
			case reflect.String:
				if len(error) == 1 {
					me.Error = errors.New(fmt.Sprintf(error[0].(string)))
				} else {
					me.Error = errors.New(fmt.Sprintf(error[0].(string), error[1:]...))
				}
			default:
				if len(error) == 1 {
					me.Error = errors.New(fmt.Sprintf("%v", error))
				} else {
					var es string
					for _, e := range error {
						es += fmt.Sprintf("%v ", e)
					}
					es = strings.TrimSuffix(es, " ")
					me.Error = errors.New(es)
				}
		}

		me.Error = errors.New(fmt.Sprintf(error[0].(string), error[1:]...))
	}
}

func (me *State) SetWarning(format string, args ...interface{}) {
	me.Ok = nil
	me.Warning = errors.New(fmt.Sprintf(format, args...))
	me.Error = nil
}

func (me *State) SetOk(format string, args ...interface{}) {
	me.Ok = errors.New(fmt.Sprintf(format, args...))
	me.Warning = nil
	me.Error = nil
}

func (me *State) ClearError() {
	me.Error = nil
}

func (me *State) Clear() {
	me.Ok = errors.New("")
	me.Warning = nil
	me.Error = nil
}


func (me *State) IsRunning() bool {
	var ok bool
	if me.String == StateRunning {
		ok = true
	}
	return ok
}

func (me *State) IsPaused() bool {
	var ok bool
	if me.String == StatePaused {
		ok = true
	}
	return ok
}

func (me *State) IsCreated() bool {
	var ok bool
	if me.String == StateCreated {
		ok = true
	}
	return ok
}

func (me *State) IsRestarting() bool {
	var ok bool
	if me.String == StateRestarting {
		ok = true
	}
	return ok
}

func (me *State) IsRemoving() bool {
	var ok bool
	if me.String == StateRemoving {
		ok = true
	}
	return ok
}

func (me *State) IsExited() bool {
	var ok bool
	if me.String == StateExited {
		ok = true
	}
	return ok
}

func (me *State) IsDead() bool {
	var ok bool
	if me.String == StateDead {
		ok = true
	}
	return ok
}

// "created", "running", "paused", "restarting", "removing", "exited", or "dead"
func (me *State) SetString(s string) {
	me.String = s
}

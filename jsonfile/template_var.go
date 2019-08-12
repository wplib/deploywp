package jsonfile

import (
	"fmt"
	"strings"
)

type TemplateVarMap map[Identifier]*TemplateVar
type TemplateVar struct {
	Name     Identifier
	Value    Value
	Children TemplateVarMap
	Parent   *TemplateVar
}

func NewTemplateVar(name Identifier, value Value, parent *TemplateVar) *TemplateVar {
	return &TemplateVar{
		Name:     name,
		Children: make(TemplateVarMap, 0),
		Value:    value,
		Parent:   parent,
	}
}

func (me *TemplateVar) FullName() (fn Identifier) {
	if me.Parent == nil {
		fn = me.Name
	} else {
		pn := me.Parent.FullName()
		if pn == "." {
			pn = ""
		}
		fn = fmt.Sprintf("%s.%s", pn, me.Name)
	}
	return fn
}

func (me *TemplateVar) GetLocalVars() (ss []string) {
	ss = make([]string, 0)
	for _, ctv := range me.Children {
		if len(ctv.Children) == 0 {
			ss = append(ss, ctv.Name)
			continue
		}
		for _, s := range ctv.GetLocalVars() {
			ss = append(ss, fmt.Sprintf("%s.%s", ctv.Name, s))
		}
	}
	return ss
}

func (me *TemplateVar) GetAbsoluteVars() (ss []string) {
	ss = []string{}
	for range Once {
		if len(me.Children) == 0 {
			ss = []string{me.FullName()}
			break
		}
		for _, ctv := range me.Children {
			fn := ctv.FullName()
			if len(ctv.Children) == 0 {
				ss = append(ss, fn)
			}
			for _, s := range ctv.GetLocalVars() {
				ss = append(ss, fmt.Sprintf("%s.%s", fn, s))
			}
		}
	}
	return ss
}

func (me TemplateVarMap) ParseTemplateVar(s string, v Value, parent *TemplateVar) (tv *TemplateVar) {
	for range Once {
		tvm := me
		parts := strings.Split(s, ".")
		if len(parts) == 0 {
			break
		}
		first := parts[0]
		var ok bool
		tv, ok = tvm[first]
		if !ok {
			tv = NewTemplateVar(first, v, parent)
			tvm[first] = tv
		}
		if len(parts) == 1 {
			break
		}
		var rest string
		switch len(parts) {
		case 2:
			rest = parts[1]
		default:
			rest = strings.Join(parts[1:], ".")
		}
		tv.Children.ParseTemplateVar(rest, v, parent)
	}
	return tv
}

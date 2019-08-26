package jsonfile

import (
	"fmt"
	"reflect"
)

type NodeTreeArgs struct {
	Node   *Node
	Parent *Node
}

type NodeMap map[Identifier]*Node
type Nodes []*Node
type Node struct {
	Root     *Node
	Value    Value
	Name     Identifier
	fullname Identifier
	VarMap   VarMap
	Children NodeMap
	Parent   *Node
	Contains Nodes
	IsPartOf Nodes
}

func NewNode(name Identifier, root *Node) *Node {
	return &Node{
		Root:     root,
		Name:     name,
		Children: make(NodeMap, 0),
	}
}

func (me *Node) String() (s string) {
	switch me.Value.Kind() {
	case reflect.Ptr:
		s = me.Value.Elem().String()
	case reflect.String:
		s = me.Value.String()
	}
	return s
}

func (me *Node) SetString(s string) {
	switch me.Value.Kind() {
	case reflect.Ptr:
		me.Value.Elem().SetString(s)
	case reflect.String:
		me.Value.SetString(s)
	}
}

func (me *Node) VarNames() (vns []string) {
	for range Once {
		if len(me.VarMap) == 0 {
			vns = make([]string, 0)
			break
		}
		vns = make([]string, len(me.VarMap))
		i := 0
		for v := range me.VarMap {
			vns[i] = v
			i++
		}
	}
	return vns
}

func (me *Node) VarCount() (count int) {
	count = 0
	if me.VarMap != nil {
		count = len(me.VarMap)
	}
	return count
}

func (me *Node) FullName() (fn Identifier) {
	if me.Parent == nil {
		fn = me.Name
	} else {
		pn := me.Parent.FullName()
		if pn == "." {
			pn = ""
		}
		f := "%s.%s"
		if len(me.Name) > 0 && me.Name[:1] == "[" {
			f = "%s%s"
		}
		fn = fmt.Sprintf(f, pn, me.Name)
	}
	me.fullname = fn
	return fn
}

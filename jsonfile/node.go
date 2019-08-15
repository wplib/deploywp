package jsonfile

import (
	"fmt"
)

type NodeTreeArgs struct {
	Node   *Node
	Parent *Node
}

type NodeMap map[Identifier]*Node
type Nodes []*Node
type Node struct {
	Root     *Node
	Name     Identifier
	Vars     *Vars
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

func (me *Node) VarNames() (vns []string) {
	if me.Vars == nil {
		me.Vars = &Vars{}
		me.Vars.vars = []string{}
	}
	return me.Vars.vars
}

func (me *Node) VarCount() (count int) {
	count = 0
	if me.Vars != nil {
		count = len(me.Vars.vars)
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
	return fn
}

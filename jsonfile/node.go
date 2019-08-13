package jsonfile

import (
	"fmt"
)

type NodeMap map[string]*Node
type Node struct {
	Root     *Node
	Name     Identifier
	Value    Value
	Vars     *Vars
	Children NodeMap
	Parent   *Node
}

func NewNode(name Identifier, root *Node) *Node {
	return &Node{
		Root:     root,
		Name:     name,
		Children: make(NodeMap, 0),
	}
}

func (me *Node) FullName() (fn Identifier) {
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

func (me *Node) GetVarNames() (ss []string) {
	ss = make([]string, 0)
	for _, ctv := range me.Children {
		if len(ctv.Children) == 0 {
			ss = append(ss, ctv.Name)
			continue
		}
		for _, s := range ctv.GetVarNames() {
			ss = append(ss, fmt.Sprintf("%s.%s", ctv.Name, s))
		}
	}
	return ss
}

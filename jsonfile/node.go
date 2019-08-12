package jsonfile

import (
	"fmt"
)

type NodeMap map[string]*Node
type Node struct {
	Root     *Node
	Fullname Identifier
	Name     Identifier
	Value    Value
	Vars     *Vars
	Children NodeMap
}

func NewNode(name Identifier, root *Node) *Node {
	return &Node{
		Root:     root,
		Fullname: name,
		Name:     name,
		Children: make(NodeMap, 0),
	}
}
func (me *Node) GetAvailableVars() (ss []string) {
	ss = make([]string, 0)
	for _, ctv := range me.Children {
		if len(ctv.Children) == 0 {
			ss = append(ss, ctv.Name)
			continue
		}
		for _, s := range ctv.GetAvailableVars() {
			ss = append(ss, fmt.Sprintf("%s.%s", ctv.Name, s))
		}
	}
	return ss
}

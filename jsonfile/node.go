package jsonfile

import (
	"fmt"
)

type NodeTreeArgs struct {
	Node   *Node
	Parent *Node
}

type NodeMap map[Identifier]*Node
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

func MakeChildArgs(childnode *Node, parent *NodeTreeArgs) *NodeTreeArgs {
	args := &NodeTreeArgs{}
	*args = *parent
	args.Node = childnode
	args.Parent = parent.Node
	return args
}

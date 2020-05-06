package jsonfile

import (
	"fmt"
	"strings"
)

type VarMap = map[Identifier]Identifier

//
// Extract a map of template vars in the form "{.foo.bar}" => "{bar}"
//
// @TODO Need to document the thinking behind this func...
//
func extractVarMap(v Value, args *NodeTreeArgs) (vm VarMap) {
	vm = make(VarMap, 0)
	for range Once {
		s := v.String()
		if !strings.Contains(s, "{") {
			break
		}
		fpn := args.Parent.FullName()
		p := newParser(s)
		for p.canParse() {
			if !p.eat('{') {
				break
			}
			dots := p.count('.')
			if dots == -1 {
				break
			}
			name := p.captureTo('}')
			if name == "" {
				break
			}
			var fullname string
			switch dots {
			case 0:
				fullname = fmt.Sprintf("%s.%s", fpn, name)
			case 1:
				fullname = fmt.Sprintf(".%s", name)
				name = fullname
			default:
				pn := parentName(fpn, dots-1)
				fullname = fmt.Sprintf("%s.%s", pn, name)
				name = fmt.Sprintf("%s%s",
					strings.Repeat(".", dots),
					name,
				)
			}
			vm[fullname] = name
		}
	}
	return vm
}

//
// Returns the name of the parent given a node name
//
// Examples:
//
// 		parentName(".foo.bar.baz", 0) => .foo.bar.baz
// 		parentName(".foo.bar.baz", 1) => .foo.bar
// 		parentName(".foo.bar.baz", 2) => .foo
//
// Params
//
//		id: 	Node identifier
//
//		remove: Segments to remove
//
// 			0: self
// 			1: parent
// 			2: grandparent
//
func parentName(id string, remove int) string {
	pos := len(id) - 1
	for pos > 0 {
		if id[pos] == '.' {
			id = id[:pos]
			remove--
		}
		if remove == 0 {
			break
		}
		pos--
	}
	return id
}

package jsonfile

import (
	"fmt"
	"strings"
)

type Vars struct {
	value Value
	vars  []string
}

func NewVars(value Value, count int) *Vars {
	return &Vars{
		value: value,
		vars:  make([]string, count),
	}
}

//
// Extract a slice of template Vars in the form '{.foo.bar}'
//
// Example:
//
//		phrase := "A {.animal.name} in the hand is worth {.count.value} in the bush."
//		tvars := extractVars(phrase)
//		fmt.Print(tvars)
//
//	Prints:
//
//		[.animal.name .count.value]
//
func extractVars(v Value, args *NodeTreeArgs) (vs *Vars) {
	for range Once {
		s := v.String()
		if !strings.Contains(s, "{") {
			break
		}
		fpn := args.Parent.FullName()
		vm := make(map[string]bool, 0)
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
			switch dots {
			case 0:
				name = fmt.Sprintf("%s.%s", fpn, name)
			case 1:
				name = fmt.Sprintf(".%s", name)
			default:
				pn := parentName(fpn, dots-1)
				name = fmt.Sprintf("%s.%s", pn, name)
			}
			vm[name] = true

		}
		vs = NewVars(v.Addr(), len(vm))
		i := 0
		for n := range vm {
			vs.vars[i] = n
			i++
		}
	}
	return vs
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

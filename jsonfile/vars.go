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
//		tvars := extractvars(phrase)
//		fmt.Print(tvars)
//
//	Prints:
//
//		[animal.name count.value]
//
func ExtractVars(v Value, args *NodeTreeArgs) (vs *Vars) {
	for range Once {
		n := v.Type().Name()
		noop(n)
		s := v.String()
		if !strings.Contains(s, "{") {
			break
		}
		vm := make(map[string]bool, 0)
		var name []byte
		fpn := args.Parent.FullName()
		dots := -1
		indots := false
		invar := false
		for pos := 0; pos < len(s); pos++ {
			if s[pos] == '{' {
				dots = 0
				invar = true
				indots = true
				continue
			}
			if !invar {
				continue
			}
			if indots {
				if s[pos] == '.' {
					dots++
					if dots == 1 {
						name = []byte(".")
					}
					continue
				}
				if pos == len(s)-1 {
					break
				}
				if dots == 1 {
					name = append(name, s[pos])
					indots = false
					dots = -1
					continue
				}
				pn := parentName(fpn, dots-1)
				name = []byte(fmt.Sprintf("%s.%s", pn, string(s[pos])))
				indots = false
				dots = -1
				continue
			}
			if pos == len(s)-1 {
				break
			}
			if s[pos] == '}' {
				vm[string(name)] = true
				name = []byte("")
				invar = false
				continue
			}
			name = append(name, s[pos])
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
func parentName(s string, remove int) string {
	pos := len(s) - 1
	for pos > 0 {
		if s[pos] == '.' {
			s = s[:pos]
			remove--
		}
		if remove == 0 {
			break
		}
		pos--
	}
	return s
}

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
		var invar bool
		var name []byte
		var ch [10]bool
		fpn := args.Parent.FullName()
		pn := fpn
		di := 0 // Dot index
		for i := 0; i < len(s); i++ {
			if s[i] == '{' {
				invar = true
				di = 1
				ch[di] = true
				continue
			}
			if !invar {
				continue
			}
			if ch[di] {
				if s[i] != '.' {
					name = []byte(fmt.Sprintf("%s.%s", pn, string(s[i])))
					pn = fpn
					ch[di] = false
					di = 0
					continue
				}
				if i == len(s)-1 {
					ch[di] = false
					continue
				}
				if s[i+1] != '.' {
					continue
				}
				ch[di] = false
				di++
				ch[di] = true
				ldi := strings.LastIndexByte(pn, '.')
				if ldi != -1 {
					pn = pn[:ldi]
				}
				continue
			}
			if s[i] == '}' {
				vm[string(name)] = true
				invar = false
				name = []byte("")
				continue
			}
			name = append(name, s[i])
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

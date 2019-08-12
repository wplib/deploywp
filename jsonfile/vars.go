package jsonfile

type Vars struct {
	value    Value
	absolute TemplateVarMap
	local    TemplateVarMap
}

func NewVars(value Value, abslen int, loclen int) *Vars {
	return &Vars{
		value:    value,
		absolute: make(TemplateVarMap, abslen),
		local:    make(TemplateVarMap, loclen),
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
func ExtractVars(v Value, args *FixupArgs) *Vars {
	const absolute = 1
	const local = 2
	avm := make(map[string]bool, 0)
	lvm := make(map[string]bool, 0)
	s := v.String()
	var invar bool
	var name []byte
	var scope int
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			invar = true
			continue
		}
		if !invar {
			continue
		}
		if scope == 0 {
			if s[i] != '.' {
				scope = local
			} else {
				scope = absolute
				continue
			}
		}
		if s[i] == '}' {
			switch scope {
			case absolute:
				avm[string(name)] = true
			case local:
				lvm[string(name)] = true
			}
			invar = false
			name = []byte("")
			scope = 0
			continue
		}
		name = append(name, s[i])
	}
	vs := NewVars(v.Addr(), len(avm), len(lvm))
	i := 0
	for n := range avm {
		vs.absolute[n] = NewTemplateVar(n, v, args.ParentVar)
		vs.absolute.ParseTemplateVar(n, v, args.ParentVar)
		i++
	}
	for n := range lvm {
		vs.local[n] = NewTemplateVar(n, v, args.ParentVar)
		vs.local.ParseTemplateVar(n, v, args.ParentVar)
		i++
	}
	return vs
}

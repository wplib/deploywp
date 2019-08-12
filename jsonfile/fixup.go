package jsonfile

import (
	"fmt"
	"reflect"
	"strings"
)

const spacer = "  "

type FixupArgs struct {
	RootNode   *Node
	ParentNode *Node
	Node       *Node
	RootVar    *TemplateVar
	ParentVar  *TemplateVar
	Var        *TemplateVar
}

func MakeChildArgs(name Identifier, value Value, childnode *Node, parent *FixupArgs) *FixupArgs {
	ca := &FixupArgs{}
	*ca = *parent
	ca.Node = childnode
	ca.ParentVar = parent.Var
	ca.ParentNode = parent.Node
	ca.Var = NewTemplateVar(name, value, parent.Var)
	childnode.Fullname = ca.Var.FullName()
	return ca
}

//
// @see https://gist.github.com/hvoecking/10772475
// @see https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
//
func (me *JsonFile) Fixup() {
	original := reflect.ValueOf(me)
	temp := reflect.New(original.Type()).Elem()
	fmt.Printf("<root>")
	rootvar := NewTemplateVar(".", reflect.ValueOf(me), nil)
	rootnode := NewNode(".", nil)
	fixupRecursive(original, temp, 0, &FixupArgs{
		RootVar:  rootvar,
		Var:      rootvar,
		Node:     rootnode,
		RootNode: rootnode,
	})
	jf := temp.Interface().(*JsonFile)
	jf.config = me.config
	jf.rootvar = rootvar
	jf.rootnode = rootnode
	nav := rootnode.GetAvailableVars()
	fmt.Printf("%+v", nav)
	lvs := rootvar.GetLocalVars()
	fmt.Printf("%+v", lvs)
	avs := rootvar.GetAbsoluteVars()
	fmt.Printf("%+v", avs)
	*me = *jf
}

func fixupRecursive(original Value, temp Value, depth int, args *FixupArgs) {
	for range Once {

		indent := strings.Repeat(spacer, depth)

		pt := func() { fmt.Printf(" [%+v]", original.Type()) }

		switch original.Kind() {
		case reflect.Ptr:
			ov := original.Elem()
			if !ov.IsValid() {
				break
			}
			temp.Set(reflect.New(ov.Type()))
			fixupRecursive(ov, temp.Elem(), depth, args)

		case reflect.Interface:
			ov := original.Elem()
			if !ov.IsValid() {
				break
			}
			tv := reflect.New(ov.Type()).Elem()
			fixupRecursive(ov, tv, depth, args)
			temp.Set(tv)

		case reflect.Struct:
			pt()
			depth++
			for i := 0; i < original.NumField(); i++ {
				cf := temp.Field(i)
				ct := temp.Type().Field(i)
				if !cf.CanSet() {
					continue
				}
				name := ct.Tag.Get("json")
				if name != "-" {
					fmt.Printf("\n%s%s— %s", spacer, indent, name)
				}
				var ok bool
				var cn *Node
				cn, ok = args.Node.Children[name]
				if !ok {
					cn = NewNode(name, args.RootNode)
					args.Node.Children[name] = cn
				}
				fixupRecursive(
					original.Field(i),
					cf,
					depth,
					MakeChildArgs(name, cf, cn, args),
				)
			}

		case reflect.Slice:
			pt()
			depth++
			o := original
			temp.Set(reflect.MakeSlice(o.Type(), o.Len(), o.Cap()))
			for i := 0; i < o.Len(); i += 1 {
				fmt.Printf("\n%s%s[%d]", spacer, indent, i)
				fixupRecursive(original.Index(i), temp.Index(i), depth, args)
			}

		case reflect.Map:
			pt()
			depth++
			temp.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				fmt.Printf("\n%s%s[%s]", spacer, indent, key)
				ov := original.MapIndex(key)
				tv := reflect.New(ov.Type()).Elem()
				fixupRecursive(ov, tv, depth, args)
				temp.SetMapIndex(key, tv)
			}

		case reflect.String:
			pt()
			fmt.Printf(": %s", original.Interface())
			temp.Set(original)
			args.Node.Vars = ExtractVars(temp, args)
			args.Node.Value = temp.Addr()

		default:
			if !original.CanInterface() {
				break
			}
			pt()
			fmt.Printf(": %+v", original.Interface())
			temp.Set(original)

		}

	}

}
package jsonfile

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/util"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type JsonFile struct {
	DeployWP DeployWP    `json:"deploywp"`
	Site     Site        `json:"site"`
	Source   Source      `json:"source"`
	Targets  Targets     `json:"targets"`
	config   *cfg.Config `json:"-"`
	rootnode *Node       `json:"-"`
	nodemap  NodeMap     `json:"-"`
	varnodes NodeMap     `json:"-"`
}

func NewJsonFile(config cfg.Config) *JsonFile {
	return &JsonFile{
		config:   &config,
		nodemap:  make(NodeMap, 0),
		varnodes: make(NodeMap, 0),
	}
}

func (me *JsonFile) GetVarNames() (vns []string) {
	vns = make([]string, len(me.nodemap))
	i := 0
	for vn := range me.nodemap {
		vns[i] = vn
	}
	return vns
}

func Load(config cfg.Config) (jf *JsonFile) {
	var err error
	for range Once {
		jf = NewJsonFile(config)

		var b []byte
		b, err = jf.load()
		if err != nil {
			break
		}

		err = json.Unmarshal(b, &jf)
		if err != nil {
			break
		}
		b = nil
		jf.walkTree()
		jf.indexTree()

	}
	if err != nil {
		log.Fatalf("Config file '%s' cannot be processed. It is likely invalid JSON or is not using the correct schema: %s.",
			"@TODO: Put acceptable schema number here...",
			err,
		)
	}
	return jf
}

func (me *JsonFile) load() (b []byte, err error) {
	var isnew bool
	fp := me.Filepath()
	for range Once {
		if !util.FileExists(fp) {
			fmt.Printf("A deploy file '%s' does not exist.", fp)
			os.Exit(1)
		}
		b, err = ioutil.ReadFile(fp)
		if err == nil {
			isnew = string(b) == GetDefault()
			break
		}
		b, err = me.makeNew(fp)
		isnew = true
	}
	if isnew {
		fmt.Printf("\nYour deploy file '%s' is newly initialized.", fp)
		fmt.Printf("\nPlease EDIT to configure appropriate settings and rerun your command.\n")
		os.Exit(1)
	}
	return b, err
}

func (me *JsonFile) makeNew(fp Filepath) (b []byte, err error) {
	var f *os.File
	for range Once {
		f, err = os.Create(fp)
		if err != nil {
			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		var n int
		d := GetDefault()
		n, err = f.WriteString(d)
		if err != nil || n != len(d) {
			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		var size int64
		size, err = f.Seek(0, 2)
		if err != nil || size != int64(len(d)) {
			fmt.Printf("Cannot determine length of deploy file just written '%s'; Check permissions: %s", fp, err)
			os.Exit(1)
		}
		var n64 int64
		n64, err = f.Seek(0, 0)
		if err != nil || n64 != 0 {
			fmt.Printf("Cannot reset deploy file just written '%s'; Check permissions: %s.", fp, err)
			os.Exit(1)
		}
		b, err = ioutil.ReadAll(f)
		if err != nil || string(b) != d {
			fmt.Printf("Deploy file read does not equal deploy file just written '%s': %s.", fp, err)
			os.Exit(1)
		}
	}
	_ = f.Close()
	return b, err
}

func (me *JsonFile) Filepath() (fp string) {
	return fmt.Sprintf("%s%c%s",
		app.DeployDir,
		os.PathSeparator,
		app.DeployFile,
	)
}

func (me *JsonFile) walkTree() {
	//fmt.Printf("<root>")
	me.walkRecursive(reflect.ValueOf(me), nil, -1)
	//fmt.Printf("%+v", me.GetVarNames())
}

func (me *JsonFile) walkRecursive(v Value, args *NodeTreeArgs, depth int) {

	if depth == -1 {
		if args == nil {
			args = &NodeTreeArgs{}
		}
		if args.Node == nil {
			args.Node = NewNode(".", nil)
		}
		me.rootnode = args.Node
		depth++
	}

	switch v.Kind() {
	case reflect.Ptr:
		ov := v.Elem()
		if !ov.IsValid() {
			break
		}
		me.walkRecursive(ov, args, depth)

	case reflect.Interface:
		ov := v.Elem()
		if !ov.IsValid() {
			break
		}
		me.walkRecursive(ov, args, depth)

	case reflect.Struct:
		depth++
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if !f.CanSet() {
				continue
			}
			t := v.Type().Field(i)
			ca := me.makeChildArgs(t.Tag.Get("json"), args)
			me.walkRecursive(f, ca, depth)
		}

	case reflect.Slice:
		depth++
		o := v
		for i := 0; i < o.Len(); i += 1 {
			name := fmt.Sprintf("[%d]", i)
			cn := me.makeChildArgs(name, args)
			me.walkRecursive(o.Index(i), cn, depth)
		}

	case reflect.Map:
		depth++
		for _, key := range v.MapKeys() {
			mi := v.MapIndex(key)
			name := fmt.Sprintf("[%s]", key.Type().Name())
			cn := me.makeChildArgs(name, args)
			me.walkRecursive(mi, cn, depth)
		}

	default: // Includes `case: default.text`
		// Extract vars from string and capture them
		args.Node.Vars = extractVars(v, args)

		// captureTo a pointer to this value so we can update is.
		me.nodemap[args.Node.FullName()] = args.Node

	}

}

func (me *JsonFile) makeChildArgs(name Identifier, parent *NodeTreeArgs) *NodeTreeArgs {

	var ok bool
	var childnode *Node
	childnode, ok = parent.Node.Children[name]
	if !ok {
		childnode = NewNode(name, me.rootnode)
		parent.Node.Children[name] = childnode
	}
	childnode.Parent = parent.Node

	args := &NodeTreeArgs{}
	*args = *parent
	args.Node = childnode
	args.Parent = parent.Node
	return args
}

func (me *JsonFile) indexTree() {
	for _, node := range me.nodemap {
		for _, varname := range node.VarNames() {
			varnode := me.nodemap[varname]
			varnode.IsPartOf = append(varnode.IsPartOf, node)
			node.Contains = append(node.Contains, varnode)
			me.varnodes[varnode.FullName()] = varnode
		}
		if node.Vars == nil {
			continue
		}
		for _, v := range node.Vars.vars {
			fmt.Printf("%s contains %s\n", node.FullName(), v)
		}
	}
}

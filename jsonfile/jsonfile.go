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
	DeployWP DeployWP        `json:"deploywp"`
	Site     Site            `json:"site"`
	Source   Source          `json:"source"`
	Targets  Targets         `json:"targets"`
	config   *cfg.Config     `json:"-"`
	rootnode *Node           `json:"-"`
	rvmap    ReflectValueMap `json:"-"`
}

func NewJsonFile(config cfg.Config) *JsonFile {
	return &JsonFile{
		config: &config,
		rvmap:  make(ReflectValueMap, 0),
	}
}

func (me *JsonFile) GetVarNames() (vns []string) {
	vns = make([]string, len(me.rvmap))
	i := 0
	for vn := range me.rvmap {
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

		jf.WalkNodeTree()

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
		b, err = me.makenew(fp)
		isnew = true
	}
	if isnew {
		fmt.Printf("\nYour deploy file '%s' is newly initialized.", fp)
		fmt.Printf("\nPlease EDIT to configure appropriate settings and rerun your command.\n")
		os.Exit(1)
	}
	return b, err
}

func (me *JsonFile) makenew(fp Filepath) (b []byte, err error) {
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

//const spacer = "  "

//
// @see https://gist.github.com/hvoecking/10772475
// @see https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
//
func (me *JsonFile) WalkNodeTree() {
	//fmt.Printf("<root>")
	me.walkrecursive(reflect.ValueOf(me), nil, -1)
	fmt.Printf("%+v", me.GetVarNames())
}

func (me *JsonFile) walkrecursive(v Value, args *NodeTreeArgs, depth int) {

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

	//indent := strings.Repeat(spacer, depth)
	//pt := func() { fmt.Printf(" [%+v]", v.Type()) }

	switch v.Kind() {
	case reflect.Ptr:
		ov := v.Elem()
		if !ov.IsValid() {
			break
		}
		me.walkrecursive(ov, args, depth)

	case reflect.Interface:
		ov := v.Elem()
		if !ov.IsValid() {
			break
		}
		me.walkrecursive(ov, args, depth)

	case reflect.Struct:
		//pt()
		depth++
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if !f.CanSet() {
				continue
			}
			t := v.Type().Field(i)
			name := t.Tag.Get("json")
			//if name != "-" {
			//	fmt.Printf("\n%s%s— %s", spacer, indent, name)
			//}
			var ok bool
			var cn *Node
			cn, ok = args.Node.Children[name]
			if !ok {
				cn = NewNode(name, me.rootnode)
				args.Node.Children[name] = cn
			}
			cn.Parent = args.Node
			me.walkrecursive(f, MakeChildArgs(cn, args), depth)
		}

	case reflect.Slice:
		//pt()
		depth++
		o := v
		for i := 0; i < o.Len(); i += 1 {
			//fmt.Printf("\n%s%s[%d]", spacer, indent, i)
			me.walkrecursive(o.Index(i), args, depth)
		}

	case reflect.Map:
		//pt()
		depth++
		for _, key := range v.MapKeys() {
			//fmt.Printf("\n%s%s[%s]", spacer, indent, key)
			ov := v.MapIndex(key)
			me.walkrecursive(ov, args, depth)
		}

	default: // Includes `case: default.String`
		if !v.CanInterface() {
			break
		}
		//pt()
		//fmt.Printf(": %s", v.Interface())

		// Extract vars from string and capture them
		args.Node.Vars = ExtractVars(v, args)

		// Capture a pointer to this value so we can update is.
		me.rvmap[args.Node.FullName()] = v.Addr()

	}

}

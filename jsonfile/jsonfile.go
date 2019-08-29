package jsonfile

import (
	"encoding/json"
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/util"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

var _ deploywp.Getter = (*JsonFile)(nil)

type JsonFile struct {
	Meta     Meta        `json:"deploywp"`
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

func (me *JsonFile) GetConfig() *cfg.Config {
	return me.config
}

func (me *JsonFile) GetMeta() *deploywp.Meta {
	return deploywp.NewMetaFromGetter(me.Meta)
}

func (me *JsonFile) GetSite() *deploywp.Site {
	return deploywp.NewSiteFromGetter(me.Site)
}

func (me *JsonFile) GetSource() *deploywp.Source {
	return deploywp.NewSourceFromGetter(me.Source)
}

func (me *JsonFile) GetTargets() *deploywp.Targets {
	return deploywp.NewTargetsFromGetter(me.Targets)
}

func (me *JsonFile) GetVarNames() (vns []string) {
	vns = make([]string, len(me.nodemap))
	i := 0
	for vn := range me.nodemap {
		vns[i] = vn
	}
	return vns
}

func (me *JsonFile) Load() (err error) {
	for range Once {

		var b []byte
		b, err = me.load()
		if err != nil {
			err = fmt.Errorf("unable to unmarshal JSON from %s: %s",
				me.config.GetConfigFile(),
				err.Error(),
			)
			break
		}

		err = json.Unmarshal(b, &me)
		if err != nil {
			err = fmt.Errorf("unable to unmarshal JSON from %s: %s",
				me.config.GetConfigFile(),
				err.Error(),
			)
			break
		}
		b = nil
		me.applyDefaults()
		me.walkTree()
		me.indexTree()
		me.applyVars()

		// Release memory after fixup
		me.nodemap = nil
		me.varnodes = nil
		me.rootnode = nil

	}
	if err != nil {
		err = fmt.Errorf("config file '%s' cannot be processed. It is likely invalid JSON or is not using the correct schema: %s",
			"@TODO: Put acceptable schema number here...",
			err,
		)
	}
	return err
}

func (me *JsonFile) load() (b []byte, err error) {
	var isnew bool
	fp := me.Filepath()
	for range Once {
		if !util.FileExists(fp) {
			err = fmt.Errorf("the deploy file '%s' does not exist", fp)
			break
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

func (me *JsonFile) applyDefaults() {
	for range Once {
		if me.Targets.Hosts == nil {
			break
		}
		if me.Targets.Defaults == nil {
			break
		}
		d := me.Targets.Defaults

		for i, h := range me.Targets.Hosts {
			err := h.ApplyDefaults(d)
			if err != nil {
				log.Printf("unable to merge host defaults into '%s': %s",
					h.Name,
					err.Error(),
				)
				continue
			}
			me.Targets.Hosts[i] = h
		}
	}
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
		//
		// Capture pointer to this value so we can update it
		//
		n := args.Node
		n.Value = v
		n.VarMap = extractVarMap(v, args)
		me.nodemap[n.FullName()] = n
		args.Node = n

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
		if node.VarMap == nil {
			continue
		}
	}
}
func (me *JsonFile) applyVars() {
	for _, vn := range me.varnodes {
		me._applyVars(vn, nil, 0)
	}
}
func (me *JsonFile) _applyVars(vn *Node, pn *Node, depth int) {
	if pn != nil {
		_applyVar(pn, vn)
	}
	for _, cn := range vn.IsPartOf {
		me._applyVars(cn, vn, depth+1)
	}
}

func _applyVar(vn *Node, cn *Node) {

	// Get the content template for which we want to replace a var
	// e.g. "https://www.{domain}"
	ct := cn.String()

	// Get the var value to replace into ct
	// e.g. "example.com"
	vv := vn.String()

	// Get the var name used in ct, which might be a short name
	// e.g. (vn.FullName() = ".site.domain" => "domain")
	n := cn.VarMap[vn.FullName()]

	// Wrap n with braces so it will represent a var
	// and fully replace the var in ct, e.g. "{domain}"
	nv := fmt.Sprintf("{%s}", n)

	// Replace occurances of nv in ct with vv
	// e.g. ("https://www.{domain}","{domain}","example.com") =>  "https://www.example.com"
	s := strings.ReplaceAll(ct, nv, vv)

	// Finally put the replaces string back into the content not
	// @TODO Need to keep a copy of original around for save .JSON file
	cn.SetString(s)
}

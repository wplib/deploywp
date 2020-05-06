package deployjson

import (
	"github.com/wplib/deploywp/cfg"
)

type JsonFile struct {
	Hosts Hosts `json:"hosts"`
	Source Source `json:"source"`
	Target Target `json:"target"`
}


//var _ deploywp.Getter = (*JsonFile)(nil)

func NewJsonFile(config cfg.Config) *JsonFile {
	var jf JsonFile

	jf.Hosts.New()
	jf.Source.New()
	jf.Target.New()

	return &jf
}

func (me *JsonFile) GetHosts() *Hosts {
	return &me.Hosts
}

func (me *JsonFile) GetSource() *Source {
	return &me.Source
}

func (me *JsonFile) GetTarget() *Target {
	return &me.Target
}



//
//func (me *JsonFile) GetSite() *deploywp.Site {
//	return deploywp.NewSiteFromGetter(me.Site)
//}
//
//func (me *JsonFile) GetTargets() *deploywp.Targets {
//	return deploywp.NewTargetsFromGetter(me.Targets)
//}
//
//func (me *JsonFile) GetVarNames() (vns []string) {
//	vns = make([]string, len(me.nodemap))
//	i := 0
//	for vn := range me.nodemap {
//		vns[i] = vn
//	}
//	return vns
//}
//
//func (me *JsonFile) Load() (err error) {
//	for range Once {
//
//		var b []byte
//		b, err = me.load()
//		if err != nil {
//			err = fmt.Errorf("unable to unmarshal JSON from %s: %s",
//				me.config.GetConfigFile(),
//				err.Error(),
//			)
//			break
//		}
//
//		err = json.Unmarshal(b, &me)
//		if err != nil {
//			err = fmt.Errorf("unable to unmarshal JSON from %s: %s",
//				me.config.GetConfigFile(),
//				err.Error(),
//			)
//			break
//		}
//		b = nil
//		me.initialize()
//		me.applyDefaults()
//		me.walkTree()
//		me.indexTree()
//		me.applyVars()
//
//		// Release memory after fixup
//		me.nodemap = nil
//		me.varnodes = nil
//		me.rootnode = nil
//
//	}
//	if err != nil {
//		err = fmt.Errorf("config file '%s' cannot be processed. It is likely invalid JSON or is not using the correct schema: %s",
//			"@TODO: Put acceptable schema number here...",
//			err,
//		)
//	}
//	return err
//}
//
//func (me *JsonFile) load() (b []byte, err error) {
//	var isnew bool
//	fp := me.Filepath()
//	for range Once {
//		if !util.FileExists(fp) {
//			err = fmt.Errorf("the deploy file '%s' does not exist", fp)
//			break
//		}
//		b, err = ioutil.ReadFile(fp)
//		if err == nil {
//			isnew = string(b) == GetDefault()
//			break
//		}
//		b, err = me.makeNew(fp)
//		isnew = true
//	}
//	if isnew {
//		fmt.Printf("\nYour deploy file '%s' is newly initialized.", fp)
//		fmt.Printf("\nPlease EDIT to configure appropriate settings and rerun your command.\n")
//		os.Exit(1)
//	}
//	return b, err
//}
//
//func (me *JsonFile) makeNew(fp Filepath) (b []byte, err error) {
//	var f *os.File
//	for range Once {
//		f, err = os.Create(fp)
//		if err != nil {
//			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
//			os.Exit(1)
//		}
//		var n int
//		d := GetDefault()
//		n, err = f.WriteString(d)
//		if err != nil || n != len(d) {
//			fmt.Printf("Cannot create deploy file '%s'; Check permissions: %s.", fp, err)
//			os.Exit(1)
//		}
//		var size int64
//		size, err = f.Seek(0, 2)
//		if err != nil || size != int64(len(d)) {
//			fmt.Printf("Cannot determine length of deploy file just written '%s'; Check permissions: %s", fp, err)
//			os.Exit(1)
//		}
//		var n64 int64
//		n64, err = f.Seek(0, 0)
//		if err != nil || n64 != 0 {
//			fmt.Printf("Cannot reset deploy file just written '%s'; Check permissions: %s.", fp, err)
//			os.Exit(1)
//		}
//		b, err = ioutil.ReadAll(f)
//		if err != nil || string(b) != d {
//			fmt.Printf("Deploy file read does not equal deploy file just written '%s': %s.", fp, err)
//			os.Exit(1)
//		}
//	}
//	_ = f.Close()
//	return b, err
//}
//
//func (me *JsonFile) Filepath() (fp string) {
//	return fmt.Sprintf("%s%c%s",
//		app.DeployDir,
//		os.PathSeparator,
//		app.DeployFile,
//	)
//}
//
//func (me *JsonFile) applyDefaults() {
//	for range Once {
//		if me.Targets.Hosts == nil {
//			break
//		}
//		if me.Targets.Defaults == nil {
//			break
//		}
//		d := me.Targets.Defaults
//
//		for i, h := range me.Targets.Hosts {
//			err := h.ApplyDefaults(d)
//			if err != nil {
//				log.Printf("unable to merge host defaults into '%s': %s",
//					h.Name,
//					err.Error(),
//				)
//				continue
//			}
//			me.Targets.Hosts[i] = h
//		}
//	}
//}
//
//func (me *JsonFile) walkTree() {
//	//fmt.Printf("<root>")
//	me.walkRecursive(reflect.ValueOf(me), nil, -1)
//	//fmt.Printf("%+v", me.GetVarNames())
//}
//
//func (me *JsonFile) walkRecursive(v Value, args *NodeTreeArgs, depth int) {
//
//	if depth == -1 {
//		if args == nil {
//			args = &NodeTreeArgs{}
//		}
//		if args.Node == nil {
//			args.Node = NewNode(".", nil)
//		}
//		me.rootnode = args.Node
//		depth++
//	}
//
//	switch v.Kind() {
//	case reflect.Ptr:
//		ov := v.Elem()
//		if !ov.IsValid() {
//			break
//		}
//		me.walkRecursive(ov, args, depth)
//
//	case reflect.Interface:
//		ov := v.Elem()
//		if !ov.IsValid() {
//			break
//		}
//		me.walkRecursive(ov, args, depth)
//
//	case reflect.Struct:
//		depth++
//		for i := 0; i < v.NumField(); i++ {
//			f := v.Field(i)
//			if !f.CanSet() {
//				continue
//			}
//			t := v.Type().Field(i)
//			ca := me.makeChildArgs(t.Tag.Get("json"), args)
//			me.walkRecursive(f, ca, depth)
//		}
//
//	case reflect.Slice:
//		depth++
//		o := v
//		for i := 0; i < o.Len(); i += 1 {
//			name := fmt.Sprintf("[%d]", i)
//			cn := me.makeChildArgs(name, args)
//			me.walkRecursive(o.Index(i), cn, depth)
//		}
//
//	case reflect.Map:
//		depth++
//		for _, key := range v.MapKeys() {
//			mi := v.MapIndex(key)
//			name := fmt.Sprintf("[%s]", key.Type().Name())
//			cn := me.makeChildArgs(name, args)
//			me.walkRecursive(mi, cn, depth)
//		}
//
//	default: // Includes `case: default.text`
//		// Extract vars from string and capture them
//		//
//		// Capture pointer to this value so we can update it
//		//
//		n := args.Node
//		n.Value = v
//		n.VarMap = extractVarMap(v, args)
//		me.nodemap[n.FullName()] = n
//		args.Node = n
//
//	}
//
//}
//
//func (me *JsonFile) makeChildArgs(name Identifier, parent *NodeTreeArgs) *NodeTreeArgs {
//
//	var ok bool
//	var childnode *Node
//	childnode, ok = parent.Node.Children[name]
//	if !ok {
//		childnode = NewNode(name, me.rootnode)
//		parent.Node.Children[name] = childnode
//	}
//	childnode.Parent = parent.Node
//
//	args := &NodeTreeArgs{}
//	*args = *parent
//	args.Node = childnode
//	args.Parent = parent.Node
//	return args
//}
//
//func (me *JsonFile) indexTree() {
//	for _, node := range me.nodemap {
//		for _, varname := range node.VarNames() {
//			varnode := me.nodemap[varname]
//			node.Suppliers = append(node.Suppliers, varnode)
//			varnode.Consumers = append(varnode.Consumers, node)
//			me.varnodes[varnode.FullName()] = varnode
//		}
//		if node.VarMap == nil {
//			continue
//		}
//	}
//}
//func (me *JsonFile) applyVars() {
//	for _, vn := range me.varnodes {
//		me._applyVars(vn, nil, 0)
//	}
//}
//
//func (me *JsonFile) _applyVars(supplier *Node, consumer *Node, depth int) {
//	i := 0
//	for len(supplier.Suppliers) > i {
//		ss := supplier.Suppliers[i]
//		me._applyVars(ss, supplier, depth+1)
//		if len(supplier.Suppliers) <= i {
//			break
//		}
//		if ss != supplier.Suppliers[i] {
//			// Values were removed from the Needs list
//			// in a recursive call, so what we keep the
//			// same index and get a new value
//			ss = supplier.Suppliers[i]
//			continue
//		}
//		i++
//	}
//	for range Once {
//		if consumer == nil {
//			break
//		}
//		if len(consumer.Suppliers) == 0 {
//			break
//		}
//		_applyVar(consumer, supplier)
//	}
//	i = 0
//	for len(supplier.Consumers) > i {
//		sc := supplier.Consumers[i]
//		me._applyVars(sc, supplier, depth+1)
//		if len(sc.Consumers) <= i {
//			break
//		}
//		if supplier != sc.Consumers[i] {
//			// Values were removed from the Has list
//			// in a recursive call, so what we keep the
//			// same index and get a new value
//			sc = supplier.Consumers[i]
//			continue
//		}
//		i++
//	}
//}
//
//func _applyVar(consumer *Node, supplier *Node) {
//
//	// Get the content template for which we want to replace a var
//	// e.g. "https://www.{domain}"
//	consumerTemplate := consumer.String()
//
//	if !strings.Contains(consumerTemplate, "{") {
//		app.Fail("Attempting to apply substituation to value with no template var: '%s'", consumerTemplate)
//	}
//
//	// Get the var name used in consumerTemplate, which might be a short name
//	// e.g. (supplier.FullName() = ".site.domain" => "domain")
//	supplierName := consumer.VarMap[supplier.FullName()]
//
//	// Get the var value to replace into consumerTemplate
//	// e.g. "example.com"
//	supplierValue := supplier.String()
//
//	// Wrap supplierName with braces so it will represent a var
//	// and fully replace the var in consumerTemplate, e.g. "{domain}"
//	supplierVar := fmt.Sprintf("{%s}", supplierName)
//
//	// Replace occurances of supplierVar in consumerTemplate with supplierValue
//	// e.g. ("https://www.{domain}","{domain}","example.com") =>  "https://www.example.com"
//	s := strings.ReplaceAll(consumerTemplate, supplierVar, supplierValue)
//
//	// Finally put the replaces string back into the content not
//	// @TODO Need to keep a copy of original around for save .JSON file
//	consumer.SetString(s)
//
//	// Remove these dependencies so we do not recurse infinitely
//	supplier.removeConsumer(consumer)
//	consumer.removeSupplier(supplier)
//
//}
//
//func (me *Node) removeSupplier(n *Node) {
//	i := 0
//	for len(me.Suppliers) > i {
//		an := me.Suppliers[i]
//		if an == n {
//			me.Suppliers = append(me.Suppliers[:i], me.Suppliers[i+1:]...)
//			break
//		}
//		i++
//	}
//}
//
//func (me *Node) removeConsumer(n *Node) {
//	i := 0
//	for len(me.Consumers) > i {
//		an := me.Consumers[i]
//		if an == n {
//			me.Consumers = append(me.Consumers[:i], me.Consumers[i+1:]...)
//			break
//		}
//		i++
//	}
//}
//
//func (me *JsonFile) initialize() {
//	me.initializeDefaults()
//	me.initializeWebHosts()
//}
//
//func (me *JsonFile) initializeDefaults() {
//	me.initializeWebHost(me.Targets.Defaults, "defaults")
//}
//
//func (me *JsonFile) initializeWebHosts() {
//	for range Once {
//		if me.Targets.Hosts == nil {
//			app.Fail("No target hosts specified in '%s'",
//				me.config.GetConfigFile(),
//			)
//		}
//		for _, h := range me.Targets.Hosts {
//			if h.ProviderId == "" {
//				h.ProviderId = me.Targets.Defaults.ProviderId
//			}
//			if h.ProviderId == "" {
//				app.Fail("Provider Id not specified in targets.host[n].ProviderId nor in target.defaults.provider_id")
//				continue
//			}
//			if h.GetProviderType() != providers.WebHostingProvider {
//				continue
//			}
//			me.initializeWebHost(h, "host")
//		}
//	}
//}
//
//func (me *JsonFile) initializeWebHost(h *Host, what string) {
//	var err error
//	for range Once {
//		if h == nil {
//			err = fmt.Errorf("%s is nil", what)
//			break
//		}
//		if h.ProviderId == "" {
//			err = fmt.Errorf("provider is empty")
//			break
//		}
//		p := providers.Dispense(h.ProviderId)
//		wpp := NewWordPressPaths()
//		if h.Paths == nil {
//			h.Paths = wpp
//		} else {
//			h.Paths.ApplyDefaults(wpp)
//		}
//		if h.Files == nil {
//			h.Files = &FileDispositions{}
//		}
//		p.InitializeHost(h)
//	}
//	if err != nil {
//		app.Fail("Targets specified in '%s': %s",
//			me.config.GetConfigFile(),
//			err.Error(),
//		)
//	}
//}

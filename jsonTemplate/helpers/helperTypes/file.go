package helperTypes

//
//
//const DefaultSeparator = "\n"
//
//type OsPathGetter interface {
//}
//
//type TypeOsPath struct {
//	State     *ux.State
//
//	_Path     string
//	_Filename string
//	_Dirname  string
//	_IsDir    bool
//	_IsFile   bool
//	_Exists   bool
//	_ModTime  time.Time
//	_Mode     os.FileMode
//	_Size     int64
//
//	_String    string
//	_Array     []string
//	_Separator string
//	_Valid     bool
//}
//
////type TypeReadFile struct {
////	*TypeOsPath
////	String string
////	Array  []string
////}
////
////type TypeWriteFile struct {
////	*TypeOsPath
////}
//
//
//func NewOsPath() *TypeOsPath {
//	return &TypeOsPath{
//		State:     ux.New(),
//		_Path:     "",
//		_Filename: "",
//		_Dirname:  "",
//		_IsDir:    false,
//		_IsFile:   false,
//		_Exists:   false,
//		_ModTime:  time.Time{},
//		_Mode:     0,
//		_Size:     0,
//		_String:   "",
//		_Array:    nil,
//		_Separator: DefaultSeparator,
//	}
//}
//
//func ReflectFileMode(ref interface{}) *os.FileMode {
//	var fm os.FileMode
//
//	for range only.Once {
//		value := reflect.ValueOf(ref)
//		if value.Kind() != reflect.Uint32 {
//			break
//		}
//
//		fm = os.FileMode(value.Uint())
//	}
//
//	return &fm
//}
//
//func ReflectPath(ref ...interface{}) *string {
//	var fp string
//
//	for range only.Once {
//		var path []string
//		for _, r := range ref {
//			// Sometimes we can have dirs within each string slice.
//			// EG: [0] = "dir1/dir2" OR [0] = "dir1\dir2"
//			// This handles paths across O/S sanely.
//			p := filepath.SplitList(*ReflectString(r))
//			path = append(path, p...)
//		}
//		fp = filepath.Join(path...)
//	}
//
//	return &fp
//}
//
//
//func (me *TypeOsPath) LoadContents(data ...interface{}) {
//	for range only.Once {
//		me._String = ""
//		me.AppendContents(data...)
//	}
//}
//func (me *TypeOsPath) AppendContents(data ...interface{}) {
//	for range only.Once {
//		if me._Separator == "" {
//			me._Separator = DefaultSeparator
//		}
//
//		for _, d := range data {
//			//value := reflect.ValueOf(d)
//			//switch value.Kind() {
//			//	case reflect.String:
//			//		me._Array = append(me._Array, value.String())
//			//	case reflect.Array:
//			//		me._Array = append(me._Array, d.([]string)...)
//			//	case reflect.Slice:
//			//		me._Array = append(me._Array, d.([]string)...)
//			//}
//
//			sa := []string{}
//			switch d.(type) {
//				case []string:
//					for _, s := range d.([]string) {
//						sa = append(sa, strings.Split(s, me._Separator)...)
//					}
//				case string:
//					sa = append(sa, strings.Split(d.(string), me._Separator)...)
//			}
//
//			me._Array = append(me._Array, sa...)
//		}
//	}
//}
//func (me *TypeOsPath) GetContentString() string {
//	if me._Separator == "" {
//		me._Separator = DefaultSeparator
//	}
//
//	return strings.Join(me._Array, me._Separator)
//}
//func (me *TypeOsPath) GetContentArray() []string {
//	return me._Array
//}
//
//func (me *TypeOsPath) SetSeparator(s string) {
//	me._Separator = s
//}
//func (me *TypeOsPath) GetSeparator() string {
//	return me._Separator
//}
//
//func (me *TypeOsPath) SetPath(p string) {
//	me._Path = p
//}
//func (me *TypeOsPath) GetPath() string {
//	return me._Path
//}
//
//func (me *TypeOsPath) SetFilename(p string) {
//	me._Filename = p
//}
//func (me *TypeOsPath) GetFilename() string {
//	return me._Filename
//}
//
//func (me *TypeOsPath) SetDirname(p string) {
//	me._Dirname = p
//}
//func (me *TypeOsPath) GetDirname() string {
//	return me._Dirname
//}
//
//func (me *TypeOsPath) SetModTime(p time.Time) {
//	me._ModTime = p
//}
//func (me *TypeOsPath) GetModTime() time.Time {
//	return me._ModTime
//}
//
//func (me *TypeOsPath) SetMode(p os.FileMode) {
//	me._Mode = p
//}
//func (me *TypeOsPath) GetMode() os.FileMode {
//	return me._Mode
//}
//
//func (me *TypeOsPath) SetSize(p int64) {
//	me._Size = p
//}
//func (me *TypeOsPath) GetSize() int64 {
//	return me._Size
//}
//
//func (me *TypeOsPath) SetExists() {
//	me._Exists = true
//}
//func (me *TypeOsPath) Exists() bool {
//	return me._Exists
//}
//
//func (me *TypeOsPath) ThisIsAFile() {
//	me._IsFile = true
//	me._IsDir = false
//}
//func (me *TypeOsPath) IsAFile() bool {
//	return me._IsFile
//}
//
//func (me *TypeOsPath) ThisIsADir() {
//	me._IsFile = false
//	me._IsDir = true
//}
//func (me *TypeOsPath) IsADir() bool {
//	return me._IsDir
//}
//
//
//func (me *TypeOsPath) StatFile() *ux.State {
//	var ret *ux.State
//
//	for range only.Once {
//		if me._Path == "" {
//			me.State.SetError("path is empty")
//			break
//		}
//
//		//if me._IsFile == false {
//		//	me.State.SetError("path is not a file")
//		//	break
//		//}
//
//		var stat os.FileInfo
//		var err error
//		stat, err = os.Stat(me._Path)
//		if err != nil {
//			break
//		}
//
//		if os.IsNotExist(ret.ErrorValue) {
//			ret._Exists = false
//			break
//		}
//
//		ret.Exists = true
//		ret.ModTime = stat.ModTime()
//		//ret.ModTime = stat.Name()
//		ret.Mode = stat.Mode()
//		ret.Size = stat.Size()
//
//		if stat.IsDir() {
//			ret.IsDir = trueResolveAbsPath
//			ret.IsFile = false
//			ret.Dirname = ret.Path
//			ret.Filename = ""
//
//		} else {
//			ret.IsDir = false
//			ret.IsFile = true
//			ret.Dirname = filepath.Dir(ret.Path)
//			ret.Filename = filepath.Base(ret.Path)
//		}
//	}
//
//	return ret
//}
//
//
//func (me *TypeOsPath) ReadFile()  {
//	for range only.Once {
//		if !me._Valid {
//
//		}
//
//		me.SetPath(ResolveAbsPath(me._Path))
//		if rf.File.IsError() {
//			rf.SetError(rf.File.ErrorValue)
//			break
//		}
//		if !rf.File.Exists() {
//			rf.SetError("filename not found")
//			break
//		}
//		if rf.File.IsDir {
//			rf.SetError("filename is a directory")
//			break
//		}
//
//		var d []byte
//		var err error
//		d, err = ioutil.ReadFile(rf.File.Path)
//		if err != nil {
//			rf.SetError(err)
//			break
//		}
//
//		rf.String = string(d)
//		rf.Array = strings.Split(string(d), "\n")
//	}
//
//	return &rf
//}
//
//
//func (me *TypeOsPath) WriteFile(contents interface{}, perms interface{}, file ...interface{}) {
//	var ret helperTypes.TypeWriteFile
//
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			ret.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			ret.SetError("content string is nil")
//			break
//		}
//
//		p := helperTypes.ReflectFileMode(perms)
//		if p == nil {
//			break
//		}
//		if *p == 0 {
//			*p = 0644
//		}
//
//
//		ret.File = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if ret.File.IsDir {
//			ret.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(ret.File.Path, *c, *p)
//		if err != nil {
//			ret.SetError(err)
//			break
//		}
//	}
//
//	return &ret
//}
//
//
//func _GetAbsPath(f ...string) string {
//	var ret string
//
//	for range only.Once {
//		ret = filepath.Join(f...)
//
//		if filepath.IsAbs(ret) {
//			break
//		}
//
//		var err error
//		ret, err = filepath.Abs(ret)
//		if err != nil {
//			ret = ""
//			break
//		}
//	}
//
//	return ret
//}
//
//
//func FileToAbs(f ...string) string {
//	var ret string
//
//	for range only.Once {
//		ret = filepath.Join(f...)
//
//		if filepath.IsAbs(ret) {
//			break
//		}
//
//		var err error
//		ret, err = filepath.Abs(ret)
//		if err != nil {
//			ret = ""
//			break
//		}
//	}
//	//ret = strings.ReplaceAll(ret, "//", "/")
//
//	return ret
//}
//
//
//func ResolvePath(path ...string) *TypeOsPath {
//	var ret TypeOsPath
//
//	for range only.Once {
//		ret.Path = FileToAbs(path...)
//
//		var stat os.FileInfo
//		stat, ret.ErrorValue = os.Stat(ret.Path)
//		//if err != nil {
//		//	break
//		//}
//
//		if os.IsNotExist(ret.ErrorValue) {
//			ret._Exists = false
//			break
//		}
//
//		ret.Exists = true
//		ret.ModTime = stat.ModTime()
//		//ret.ModTime = stat.Name()
//		ret.Mode = stat.Mode()
//		ret.Size = stat.Size()
//
//		if stat.IsDir() {
//			ret.IsDir = trueResolveAbsPath
//			ret.IsFile = false
//			ret.Dirname = ret.Path
//			ret.Filename = ""
//
//		} else {
//			ret.IsDir = false
//			ret.IsFile = true
//			ret.Dirname = filepath.Dir(ret.Path)
//			ret.Filename = filepath.Base(ret.Path)
//		}
//	}
//
//	return &ret
//}
//
//
//func ResolveAbsPath(path ...string) *TypeOsPath {
//	return ResolvePath(FileToAbs(path...))
//}

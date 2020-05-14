package helperSystem


//var _ helperTypes.OsPathGetter = (*TypeOsPath)(nil)
//type TypeOsPath helperFile.TypeOsPath
//var _ helperTypes.OsPathGetter = (*TypeReadFile)(nil)
//type TypeReadFile helperFile.TypeOsPath
//var _ helperTypes.OsPathGetter = (*TypeWriteFile)(nil)
//type TypeWriteFile helperTypes.TypeWriteFile


//// Usage:
////		{{ $str := ReadFile "filename.txt" }}
//func HelperReadFile(file ...interface{}) *helperFile.TypeOsPath {
//	var rf helperFile.TypeOsPath
//
//	for range only.Once {
//		f := helperFile.ReflectPath(file...)
//		if f == nil {
//			rf.SetError("filename empty")
//			break
//		}
//
//		var op helperFile.TypeOsPath
//		rf = (*helperFile.TypeOsPath)(&op)
//
//		op.SetPath(ResolveAbsPath(*f))
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
//// Usage:
////		{{ $return := WriteFile .Data.Source 0644 "dir1" "dir2/dir3" "filename.txt" }}
//func HelperWriteFile(contents interface{}, perms interface{}, file ...interface{}) *helperTypes.TypeWriteFile {
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
//		ret.File = (*f)
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
//// Usage:
////		{{  }}
//func (me *TypeOsPath) SetPath(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(path...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		me.Path = _GetAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) SetDir(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) SetFile(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) IsDir(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) IsFile(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) Exists(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) ModTime(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) Mode(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
//}
//
//
//// Usage:
////		{{  }}
//func (me *TypeOsPath) Size(path ...interface{}) *helperFile.TypeOsPath {
//	for range only.Once {
//		f := helperTypes.ReflectPath(file...)
//		if f == nil {
//			me.SetError("filename is nil")
//			break
//		}
//
//		c := helperTypes.ReflectByteArray(contents)
//		if c == nil {
//			me.SetError("content string is nil")
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
//		me.Path = ResolveAbsPath(*f)
//		//if ret.File.IsError() {
//		//	break
//		//}
//		//if !ret.File.Exists {
//		//	ret.Error = errors.New("filename not found")
//		//	break
//		//}
//		if me.IsDir {
//			me.SetError("filename is a directory")
//			break
//		}
//
//
//		err := ioutil.WriteFile(me.Path, *c, *p)
//		if err != nil {
//			me.SetError(err)
//			break
//		}
//	}
//
//	return (*helperFile.TypeOsPath)(me)
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
//func ResolvePath(path ...string) *helperFile.TypeOsPath {
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
//func ResolveAbsPath(path ...string) *helperFile.TypeOsPath {
//	return ResolvePath(FileToAbs(path...))
//}

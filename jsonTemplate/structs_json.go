package jsonTemplate

import (
	"github.com/wplib/deploywp/cmd/runtime"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
	"strings"
)


type jsonStruct struct {
	Exec            *runtime.Exec

	TemplateFile    FileInfo
	JsonFile        FileInfo
	OutFile         FileInfo
	//Env             helperSystem.Environment
	Env             *runtime.Environment

	JsonString      string
	CreationEpoch   int64
	CreationDate    string
	CreationInfo    string
	CreationWarning string

	Json            map[string]interface{}
}


func NewJsonStruct() *jsonStruct {
	js := jsonStruct {
		Exec:            runtime.NewExec(),
		TemplateFile:    FileInfo{},
		JsonFile:        FileInfo{},
		OutFile:         FileInfo{},
		//Env:             nil,
		JsonString:      "",
		CreationEpoch:   0,
		CreationDate:    "",
		CreationInfo:    "",
		CreationWarning: "",
		Json:            nil,
	}

	return &js
}


type FileInfo struct {
	Dir           string
	Name          string
	CreationEpoch int64
	CreationDate  string

	State         *ux.State
}


func (fi *FileInfo) SetFileInfo(path *helperPath.TypeOsPath) {
	fi.Dir = path.GetDirname()
	fi.Name = path.GetFilename()
	fi.CreationDate = path.GetModTimeString()
	fi.CreationEpoch = path.GetModTimeEpoch()
	fi.State = path.State
}


//func fileToString(fileName string) ([]byte, error) {
//	var jsonString []byte
//	var err error
//
//	for range OnlyOnce {
//		_, err = os.Stat(fileName)
//		if os.IsNotExist(err) {
//			break
//		}
//
//		jsonString, err = ioutil.ReadFile(fileName)
//		if err != nil {
//			break
//		}
//	}
//
//	return jsonString, err
//}
//
//
//func _FileToAbs(f ...string) string {
//	var ret string
//
//	for range OnlyOnce {
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


func UnescapeString(s string) string {

	// \a	Alert or bell
	// \b	Backspace
	// \\	Backslash
	// \t	Horizontal tab
	// \n	Line feed or newline
	// \f	Form feed
	// \r	Carriage return
	// \v	Vertical tab
	// \'	Single quote (only in rune literals)
	// \"	Double quote (only in string literals)

	s = strings.ReplaceAll(s, `\a`, "\a")
	s = strings.ReplaceAll(s, `\b`, "\b")
	s = strings.ReplaceAll(s, `\\`, "\\")
	s = strings.ReplaceAll(s, `\t`, "\t")
	s = strings.ReplaceAll(s, `\n`, "\n")
	s = strings.ReplaceAll(s, `\f`, "\f")
	s = strings.ReplaceAll(s, `\r`, "\r")
	s = strings.ReplaceAll(s, `\v`, "\v")
	s = strings.ReplaceAll(s, `\'`, `'`)
	s = strings.ReplaceAll(s, `\"`, `"`)

	return s
}

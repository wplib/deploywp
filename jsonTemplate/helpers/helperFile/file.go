package helperFile

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"io/ioutil"
	"strings"
)


func (p *TypeOsPath) LoadContents(data ...interface{}) {
	for range only.Once {
		p._String = ""
		p.AppendContents(data...)
	}
}


func (p *TypeOsPath) AppendContents(data ...interface{}) {
	for range only.Once {
		if p._Separator == "" {
			p._Separator = DefaultSeparator
		}

		for _, d := range data {
			//value := reflect.ValueOf(d)
			//switch value.Kind() {
			//	case reflect.String:
			//		p._Array = append(p._Array, value.String())
			//	case reflect.Array:
			//		p._Array = append(p._Array, d.([]string)...)
			//	case reflect.Slice:
			//		p._Array = append(p._Array, d.([]string)...)
			//}

			var sa []string
			switch d.(type) {
			case []string:
				for _, s := range d.([]string) {
					sa = append(sa, strings.Split(s, p._Separator)...)
				}
			case string:
				sa = append(sa, strings.Split(d.(string), p._Separator)...)
			}

			p._Array = append(p._Array, sa...)
		}
	}
}


func (p *TypeOsPath) GetContentString() string {
	if p._Separator == "" {
		p._Separator = DefaultSeparator
	}

	return strings.Join(p._Array, p._Separator)
}


func (p *TypeOsPath) GetContentArray() []string {
	return p._Array
}


func (p *TypeOsPath) SetSeparator(separator string) {
	for range only.Once {
		p._Separator = separator
		p._Array = strings.Split(p._String, p._Separator)
	}
}


func (p *TypeOsPath) GetSeparator() string {
	return p._Separator
}


func (p *TypeOsPath) ReadFile() *State {
	for range only.Once {
		if !p._IsValid() {
			break
		}

		p.State = (*ux.State)(p.StatPath())
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("filename not found")
			break
		}
		if p._IsDir {
			p.State.SetError("filename is a directory")
			break
		}

		var d []byte
		var err error
		d, err = ioutil.ReadFile(p._Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.LoadContents(d)
	}

	return (*State)(p.State)
}


func (p *TypeOsPath) WriteFile() *State {

	for range only.Once {
		if !p._IsValid() {
			break
		}

		if p._String == "" {
			p.State.SetError("content string is nil")
			break
		}

		p.StatPath()
		if p._IsDir {
			p.State.SetError("file is a directory")
			break
		}
		if p._Exists && p._Overwrite {
			p.State.SetError("file exists, not overwriting")
			break
		}

		if p._Mode == 0 {
			p._Mode = 0644
		}

		err := ioutil.WriteFile(p._Path, []byte(p._String), p._Mode)
		if err != nil {
			p.State.SetError(err)
			break
		}
	}

	return (*State)(p.State)
}

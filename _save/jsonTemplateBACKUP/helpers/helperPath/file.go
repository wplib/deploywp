package helperPath

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/only"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"strings"
)


const DefaultSeparator = "\n"


func (p *TypeOsPath) LoadContents(data ...interface{}) {
	for range OnlyOnce {
		p._String = ""
		p._Array = []string{}

		p.AppendContents(data...)
	}
}


func (p *TypeOsPath) AppendContents(data ...interface{}) {
	for range OnlyOnce {
		if p._Separator == "" {
			p._Separator = DefaultSeparator
		}

		for _, d := range data {
			//value := reflect.ValueOf(d)
			//switch value.Kind() {
			//	case reflect.Output:
			//		p._Array = append(p._Array, value.Output())
			//	case reflect.Array:
			//		p._Array = append(p._Array, d.([]string)...)
			//	case reflect.Slice:
			//		p._Array = append(p._Array, d.([]string)...)
			//}

			var sa []string
			switch d.(type) {
				case []byte:
					sa = append(sa, string(d.([]byte)))
				case []string:
					for _, s := range d.([]string) {
						sa = append(sa, strings.Split(s, p._Separator)...)
					}
				case string:
					sa = append(sa, strings.Split(d.(string), p._Separator)...)
			}

			p._Array = append(p._Array, sa...)
		}
		p._String = strings.Join(p._Array, p._Separator)
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
	for range OnlyOnce {
		p._Separator = separator
		p._Array = strings.Split(p._String, p._Separator)
	}
}


func (p *TypeOsPath) GetSeparator() string {
	return p._Separator
}


func (p *TypeOsPath) ReadFile() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		p.StatPath()
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("file '%s' not found", p._Path)
			break
		}
		if p._IsDir {
			p.State.SetError("path '%s' is a directory", p._Path)
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
		p.State.SetOk("file '%s' read OK", p._Path)
	}

	return p.State
}


func (p *TypeOsPath) WriteFile() *ux.State {
	for range OnlyOnce {
		p.State.SetFunction("")
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		if p._String == "" {
			p.State.SetError("content string is nil")
			break
		}

		for range OnlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path '%s' is a directory", p._Path)
				break
			}
			if p.NotExists() {
				p.State.Clear()
				break
			}
			if p._CanOverwrite {
				break
			}

			if !helperSystem.HelperUserPromptBool("Overwrite file '%s'? (Y|N) ", p._Path) {
				p.State.SetWarning("not overwriting file '%s'", p._Path)
				break
			}
			p.State.Clear()
		}
		if p.State.IsNotOk() {
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

		p.State.SetOk("file '%s' written OK", p._Path)
	}

	return p.State
}

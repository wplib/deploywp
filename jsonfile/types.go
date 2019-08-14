package jsonfile

import "reflect"

type (
	Path      = string
	Url       = string
	Label     = string
	Domain    = string
	Slug      = string
	Guid      = string
	Version   = string
	Reference = string
	Filepath  = string

	ReadableName = string
	Identifier   = Slug

	Value           = reflect.Value
	ReflectValueMap = map[Identifier]reflect.Value
)

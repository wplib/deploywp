package jsonfile

import (
	"github.com/wplib/deploywp/providers"
	"reflect"
)

type (
	Path      = string
	Url       = string
	Label     = string
	Domain    = string
	Guid      = string
	Version   = string
	Reference = string
	Filepath  = string

	ReadableName = string
	Identifier   = string

	Value           = reflect.Value
	ReflectValueMap = map[Identifier]reflect.Value
)

type ProviderType = providers.ProviderType

package deploywp

type (
	Dir          = string
	Path         = string
	Url          = string
	Guid         = string
	Label        = string
	Domain       = string
	Version      = string
	Reference    = string
	Filepath     = string
	ReadableName = string
	Identifier   = string
)

type Loader interface {
	Load() error
}

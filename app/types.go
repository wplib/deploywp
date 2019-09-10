package app

type Urls = []Url
type UrlPaths = []UrlPath
type Filepaths = []Filepath
type (
	Dir       = string
	Path      = string
	Url       = string
	Guid      = string
	Label     = string
	Domain    = string
	Version   = string
	Reference = string
	Filepath  = string
	Port      = string

	UrlPath  = string
	Fragment = string

	Filename = string

	ReadableName = string
	Identifier   = string
)

type Protocol string

//const (
//	HttpScheme  Protocol = "http"
//	HttpsScheme Protocol = "https"
//)

type UnixTime = int64

package app

type Urls = []Url
type UrlPaths = []UrlPath
type Filepaths = []Filepath
type (
	Domain   = string
	Port     = string
	Url      = string
	UrlPath  = string
	Fragment = string
	Dir      = string
	Filename = string
	Filepath = string
)

type Protocol string

const (
	HttpScheme  Protocol = "http"
	HttpsScheme Protocol = "https"
)

type UnixTime = int64


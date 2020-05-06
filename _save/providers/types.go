package providers

type Map map[ProviderId]Provider

type (
	Name       = string
	Url        = string
	Domain     = string
	Guid       = string
	Path       = string
	Identifier = string
)

type (
	ProviderType = string
	ProviderId   = string
	ReadableName = string
)

type RepositoryGetter interface {
	GetProvider() Provider
	GetUrl() Url
}

type HostGetterSetter interface {
	GetId() ProviderId
	GetSiteGuid() Guid
	GetRepository() *Repository
	SetRepository(*Repository)
	GetName() ReadableName
	GetDomain() Domain
	SetDomain(Domain)
	GetDomainSuffix() Domain
	SetDomainSuffix(Domain)
	GetWebRoot() Path
	SetWebRoot(Path)
	GetBranch() Identifier
	SetBranch(Identifier)
	GetProviderType() ProviderType
}

//type (
//	Dir          = string
//	Path         = string
//	Url          = string
//	Guid         = string
//	Label        = string
//	Domain       = string
//	Version      = string
//	Reference    = string
//	Filepath     = string
//	ReadableName = string
//	Identifier   = string
//)

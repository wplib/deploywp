package deploywp

type Hosts []*Host

type HostsGetter interface {
	Count() int
	Hosts() Hosts
}

func NewHostsFromGetter(hg HostsGetter) (hs Hosts) {
	hs = make(Hosts, hg.Count())
	for i, h := range hg.Hosts() {
		hs[i] = h
	}
	return hs
}

type Host struct {
	Id           Identifier
	SiteGuid     Guid
	Domain       Domain
	DomainSuffix Domain
	Provider     Identifier
	Name         ReadableName
	Label        Label
	Branch       Identifier
	WebRoot      Path
	Repository   *Repository
	Paths        *WordPressPaths
	Files        *FileDispositions
	After        string
}

type HostGetter interface {
	GetId() Identifier
	GetSiteGuid() Guid
	GetDomain() Domain
	GetDomainSuffix() Domain
	GetProviderId() Identifier
	GetName() ReadableName
	GetLabel() Label
	GetBranch() Identifier
	GetWebRoot() Path
	GetRepository() *Repository
	GetPaths() *WordPressPaths
	GetFiles() *FileDispositions
	GetAfter() string
}

func NewHostFromGetter(hg HostGetter) (h *Host) {
	return &Host{
		Id:           hg.GetId(),
		SiteGuid:     hg.GetSiteGuid(),
		Domain:       hg.GetDomain(),
		DomainSuffix: hg.GetDomainSuffix(),
		Provider:     hg.GetProviderId(),
		Name:         hg.GetName(),
		Label:        hg.GetLabel(),
		Branch:       hg.GetBranch(),
		WebRoot:      hg.GetWebRoot(),
		Repository:   hg.GetRepository(),
		Paths:        hg.GetPaths(),
		After:        hg.GetAfter(),
		Files:        hg.GetFiles(),
	}
}

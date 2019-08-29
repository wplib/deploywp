package deploywp

import (
	"net/url"
)

type Site struct {
	Id      Identifier
	Name    ReadableName
	Domain  Domain
	Website *url.URL
}

type SiteGetter interface {
	GetId() Identifier
	GetName() ReadableName
	GetDomain() Domain
	GetWebsite() *url.URL
}

func NewSiteFromGetter(sg SiteGetter) (s *Site) {
	return &Site{
		Id:      sg.GetId(),
		Name:    sg.GetName(),
		Domain:  sg.GetDomain(),
		Website: sg.GetWebsite(),
	}
}

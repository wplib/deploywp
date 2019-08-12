package jsonfile

import (
	"github.com/wplib/deploywp/providers"
)

type Hosts []*Host

type Host struct {
	Id           Slug               `json:"host_id"`
	SiteGuid     Guid               `json:"site_guid"`
	DomainSuffix Domain             `json:"domain_suffix"`
	ProviderSlug Slug               `json:"provider"`
	provider     providers.Provider `json:"-"`
	Name         ReadableName       `json:"name"`
	Label        Label              `json:"label"`
	Branch       Slug               `json:"branch"`
	WebRoot      Path               `json:"web_root"`
	Repository   Repository         `json:"repository"`
	Paths        WordPressPaths     `json:"wp_paths"`
	After        string             `json:"after"`
}

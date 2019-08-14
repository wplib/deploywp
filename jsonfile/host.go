package jsonfile

type Hosts []*Host

type Host struct {
	Id           Slug           `json:"host_id"`
	SiteGuid     Guid           `json:"site_guid"`
	DomainSuffix Domain         `json:"domain_suffix"`
	Provider     Slug           `json:"provider"`
	Name         ReadableName   `json:"name"`
	Label        Label          `json:"label"`
	Branch       Slug           `json:"branch"`
	WebRoot      Path           `json:"web_root"`
	Repository   Repository     `json:"repository"`
	Paths        WordPressPaths `json:"wp_paths"`
	After        string         `json:"after"`
}

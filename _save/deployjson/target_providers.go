package deployjson

type Provider struct {
	Name     string   `json:"name"`
	Meta     Meta     `json:"meta"`
	Defaults Defaults `json:"defaults"`
}
type Providers []Provider

type Meta struct {
	SiteID   string `json:"site_id"`
	SiteName string `json:"site_name"`
}

type Defaults struct {
	Paths struct {
		WebrootDir string `json:"webroot_dir"`
	} `json:"paths"`
	Repository struct {
		URL string `json:"url"`
	} `json:"repository"`
}


func (me *Provider) New() Provider {
	if me == nil {
		me = &Provider {
			Name:     "",
			Meta:     me.Meta.New(),
			Defaults: me.Defaults.New(),
		}
	}

	return *me
}

func (me *Providers) New() Providers {
	if me == nil {
		me = &Providers {}
	}

	return *me
}

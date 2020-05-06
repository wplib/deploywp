package deployjson

import (
	"fmt"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/providers"
	"net/url"
)

type Site struct {
	Id      Identifier   `json:"id"`
	Name    ReadableName `json:"name"`
	Domain  Domain       `json:"domain"`
	Website Url          `json:"website"`
}

func (me Site) GetId() deploywp.Identifier {
	return me.Id
}

func (me Site) GetName() deploywp.ReadableName {
	return me.Name
}

func (me Site) GetDomain() deploywp.Domain {
	return me.Domain
}

func (me Site) GetWebsite() *url.URL {
	return providers.ParseUrl(me.Website,
		fmt.Sprintf("website '%s'", me.Id),
	)
}

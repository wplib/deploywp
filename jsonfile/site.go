package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
	"log"
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
	ws, err := url.Parse(me.Website)
	if err != nil {
		log.Printf("unable to parse website URL '%s' for site '%s",
			me.Website,
			me.Id,
		)
	}
	return ws
}

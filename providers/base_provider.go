package providers

import (
	"github.com/wplib/deploywp/app"
	"net/url"
	"strings"
)

var _ Provider = (*BaseProvider)(nil)

type BaseProvider struct {
	Id      ProviderId
	Type    ProviderType
	Name    ReadableName
	Website *url.URL
}

func NewBaseProvider(id ProviderId, bp BaseProvider) *BaseProvider {
	bp.Id = id
	if bp.Name == "" {
		bp.Name = strings.Title(id)
	}
	if bp.Type == "" {
		bp.Type = UnspecifiedProvider
	}
	return &bp
}

func (me *BaseProvider) GetId() ProviderId {
	return me.Id
}

func (me *BaseProvider) GetType() ProviderType {
	return me.Type
}

func (me *BaseProvider) GetName() ReadableName {
	return me.Name
}

func (me *BaseProvider) GetWebsite() *url.URL {
	return me.Website
}
func (me *BaseProvider) DetectByUrl(u Url) (bool, Url) {
	app.Fail("Concrete provider '%s' must implement the method DetectByUrl", me.Id)
	return false, u
}

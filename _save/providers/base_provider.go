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

func (me *BaseProvider) SetId(pid ProviderId) {
	me.Id = pid
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
func (me *BaseProvider) DetectByUrl(u Url) bool {
	app.Fail("Concrete provider '%s' must implement the method DetectByUrl", me.Id)
	return false
}
func (me *BaseProvider) NormalizeUrl(u Url) Url {
	app.Fail("Concrete provider '%s' must implement the method NormalizeUrl", me.Id)
	return u
}

func (me *BaseProvider) ValidateHostDefaults(hgs HostGetterSetter) {
	// Nothing to do
}

func (me *BaseProvider) ValidateHost(hgs HostGetterSetter) {
	// Nothing to do
}

func (me *BaseProvider) InitializeHost(hg HostGetterSetter) {
	// Nothing to do
}

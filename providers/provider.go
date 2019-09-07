package providers

import "net/url"

type Provider interface {
	GetId() ProviderId
	GetType() ProviderType
	GetName() Name
	GetWebsite() *url.URL
	NormalizeUrl(u Url) Url
	DetectByUrl(u Url) bool // See https://stackoverflow.com/questions/31801271/what-are-the-supported-git-url-formats
}

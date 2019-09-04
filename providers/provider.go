package providers

import "net/url"

type Provider interface {
	GetId() ProviderId
	GetType() ProviderType
	GetName() Name
	GetWebsite() *url.URL
	DetectByUrl(u Url) (bool, Url) // See https://stackoverflow.com/questions/31801271/what-are-the-supported-git-url-formats

}

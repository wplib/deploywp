package providers

type Provider interface {
	GetId() ProviderId
	GetType() ProviderType
	GetName() Name
	GetWebsite() Url
}

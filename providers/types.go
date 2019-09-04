package providers

type Map map[ProviderId]Provider

type (
	Name   = string
	Url    = string
	Domain = string
)

type (
	ProviderType = string
	ProviderId   = string
	ReadableName = string
)

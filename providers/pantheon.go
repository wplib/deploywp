package providers

var _ Provider = (*Pantheon)(nil)

type Pantheon struct{

}

func (me *Pantheon) GetId() ProviderId {
	return PantheonId
}

func (me *Pantheon) GetType() ProviderType {
	return WebHostingProvider
}

func (me *Pantheon) GetName() Name {
	return "Pantheon"
}

func (me *Pantheon) GetWebsite() Url {
	return "https://pantheon.io"
}


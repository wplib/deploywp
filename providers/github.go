package providers

type Github struct{

}

func (me *Github) GetId() ProviderId {
	return GitHubId
}

func (me *Github) GetType() ProviderType {
	return VersionControlProvider
}

func (me *Github) GetName() Name {
	return "Github"
}

func (me *Github) GetWebsite() Url {
	return "https://github.com"
}


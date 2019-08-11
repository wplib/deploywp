package providers


func Create(provider ProviderId) (p Provider) {
	switch provider {
	case GitHubId:
		p = &Github{}
	case PantheonId:
		p = &Pantheon{}
	}
	return p
}
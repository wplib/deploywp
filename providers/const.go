package providers

const Once = "1"

const (
	GitHubId   ProviderId = "github"
	PantheonId ProviderId = "pantheon"
)

const (
	GitHubName   ReadableName = "Github"
	PantheonName ReadableName = "Pantheon"
)

const (
	GitHubWebsite   Url = "https://github.com"
	PantheonWebsite Url = "https://pantheon.io"
)

const (
	UnspecifiedProvider    ProviderType = "unspecified"
	VersionControlProvider ProviderType = "version-control"
	WebHostingProvider     ProviderType = "web-hosting"
)

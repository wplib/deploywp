package deploywp


// ////////////////////////////////////////////////////////////////////////////////
// Source
func (dwp *TypeDeployWp) GetSource() *Source {
	return &dwp.Source
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Paths
func (dwp *TypeDeployWp) GetSourcePaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Source.Paths
}
func (dwp *TypeDeployWp) GetSourceAbsPaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Source.AbsPaths
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Repository
func (dwp *TypeDeployWp) GetSourceRepositoryProvider() string {
	if state := dwp.IsNil(); state.IsError() {
		return ""
	}
	return dwp.Source.GetRepositoryProvider()
}
func (dwp *TypeDeployWp) GetSourceRepositoryUrl() URL {
	if state := dwp.IsNil(); state.IsError() {
		return ""
	}
	return dwp.Source.GetRepositoryUrl()
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Revision
func (dwp *TypeDeployWp) GetSourceRevisionType() string {
	if state := dwp.IsNil(); state.IsError() {
		return ""
	}
	return dwp.Source.GetRevisionType()
}
func (dwp *TypeDeployWp) GetSourceRevisionName() string {
	if state := dwp.IsNil(); state.IsError() {
		return ""
	}
	return dwp.Source.GetRevisionName()
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Build
func (dwp *TypeDeployWp) GetSourceBuild() bool {
	if state := dwp.IsNil(); state.IsError() {
		return false
	}
	return dwp.Source.GetBuild()
}


// ////////////////////////////////////////////////////////////////////////////////
// Destination
func (dwp *TypeDeployWp) GetDestination() *Destination {
	return &dwp.Destination
}


// ////////////////////////////////////////////////////////////////////////////////
// Destination.Files
func (dwp *TypeDeployWp) GetDestinationFiles(ftype string) *FilesArray {
	if state := dwp.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return dwp.Destination.GetFiles(ftype)
}


// ////////////////////////////////////////////////////////////////////////////////
// Destination.Paths
func (dwp *TypeDeployWp) GetDestinationPaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Destination.Paths
}
func (dwp *TypeDeployWp) GetDestinationAbsPaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Destination.AbsPaths
}


// ////////////////////////////////////////////////////////////////////////////////
// Destination.Revisions
func (dwp *TypeDeployWp) GetDestinationRevision(host string) *Target {
	if state := dwp.IsNil(); state.IsError() {
		return &Target{}
	}
	return dwp.Destination.GetTargetByHost(host)
}


// ////////////////////////////////////////////////////////////////////////////////
// Destination.Providers
func (dwp *TypeDeployWp) GetDestinationProvider(provider string) *Provider {
	if state := dwp.IsNil(); state.IsError() {
		return &Provider{}
	}
	return dwp.Destination.GetProviderByName(provider)
}


// ////////////////////////////////////////////////////////////////////////////////
// Hosts
func (dwp *TypeDeployWp) GetHosts() *Hosts {
	return &dwp.Hosts
}

func (dwp *TypeDeployWp) GetHostByName(host string) *Host {
	if state := dwp.IsNil(); state.IsError() {
		return &Host{}
	}
	//return dwp.Hosts.GetByName(host)
	return &Host{}
}

func (dwp *TypeDeployWp) GetHostByProvider(provider string) *Host {
	if state := dwp.IsNil(); state.IsError() {
		return &Host{}
	}
	//return dwp.Hosts.GetByProvider(provider)
	return &Host{}
}

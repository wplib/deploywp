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
// Target
func (dwp *TypeDeployWp) GetTarget() *Target {
	return &dwp.Target
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Files
func (dwp *TypeDeployWp) GetTargetFiles(ftype string) *FilesArray {
	if state := dwp.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return dwp.Target.GetFiles(ftype)
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Paths
func (dwp *TypeDeployWp) GetTargetPaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Target.Paths
}
func (dwp *TypeDeployWp) GetTargetAbsPaths() *Paths {
	if state := dwp.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &dwp.Target.AbsPaths
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Revisions
func (dwp *TypeDeployWp) GetTargetRevision(host string) *TargetRevision {
	if state := dwp.IsNil(); state.IsError() {
		return &TargetRevision{}
	}
	return dwp.Target.GetRevisionByHost(host)
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Providers
func (dwp *TypeDeployWp) GetTargetProvider(provider string) *Provider {
	if state := dwp.IsNil(); state.IsError() {
		return &Provider{}
	}
	return dwp.Target.GetProviderByName(provider)
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
	return dwp.Hosts.GetByName(host)
}

func (dwp *TypeDeployWp) GetHostByProvider(provider string) *Host {
	if state := dwp.IsNil(); state.IsError() {
		return &Host{}
	}
	return dwp.Hosts.GetByProvider(provider)
}

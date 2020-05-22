package deploywp


// ////////////////////////////////////////////////////////////////////////////////
// Source
func (me *TypeDeployWp) GetSource() *Source {
	return &me.Source
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Paths
func (me *TypeDeployWp) GetSourcePaths() *Paths {
	if state := me.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &me.Source.Paths
}
func (me *TypeDeployWp) GetSourceAbsPaths() *Paths {
	if state := me.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &me.Source.AbsPaths
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Repository
func (me *TypeDeployWp) GetSourceRepositoryProvider() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.Source.GetRepositoryProvider()
}
func (me *TypeDeployWp) GetSourceRepositoryUrl() URL {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.Source.GetRepositoryUrl()
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Revision
func (me *TypeDeployWp) GetSourceRevisionType() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.Source.GetRevisionType()
}
func (me *TypeDeployWp) GetSourceRevisionName() string {
	if state := me.IsNil(); state.IsError() {
		return ""
	}
	return me.Source.GetRevisionName()
}


// ////////////////////////////////////////////////////////////////////////////////
// Source.Build
func (me *TypeDeployWp) GetSourceBuild() bool {
	if state := me.IsNil(); state.IsError() {
		return false
	}
	return me.Source.GetBuild()
}


// ////////////////////////////////////////////////////////////////////////////////
// Target
func (me *TypeDeployWp) GetTarget() *Target {
	return &me.Target
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Files
func (me *TypeDeployWp) GetTargetFiles(ftype interface{}) *FilesArray {
	if state := me.IsNil(); state.IsError() {
		return &FilesArray{}
	}
	return me.Target.GetFiles(ftype)
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Paths
func (me *TypeDeployWp) GetTargetPaths() *Paths {
	if state := me.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &me.Target.Paths
}
func (me *TypeDeployWp) GetTargetAbsPaths() *Paths {
	if state := me.IsNil(); state.IsError() {
		return &Paths{}
	}
	return &me.Target.AbsPaths
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Revisions
func (me *TypeDeployWp) GetTargetRevision(host interface{}) *TargetRevision {
	if state := me.IsNil(); state.IsError() {
		return &TargetRevision{}
	}
	return me.Target.GetRevision(host)
}


// ////////////////////////////////////////////////////////////////////////////////
// Target.Providers
func (me *TypeDeployWp) GetTargetProvider(provider interface{}) *Provider {
	if state := me.IsNil(); state.IsError() {
		return &Provider{}
	}
	return me.Target.GetProvider(provider)
}


// ////////////////////////////////////////////////////////////////////////////////
// Hosts
func (me *TypeDeployWp) GetHosts() *Hosts {
	return &me.Hosts
}

func (me *TypeDeployWp) GetHost(host interface{}) *Host {
	if state := me.IsNil(); state.IsError() {
		return &Host{}
	}
	return me.Hosts.GetHost(host)
}

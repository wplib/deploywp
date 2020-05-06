package deployjson

type TargetRevision struct {
	HostName string `json:"host_name"`
	RefName  string `json:"ref_name"`
}
type TargetRevisions []TargetRevision

func (me *TargetRevision) New() TargetRevision {
	if me == nil {
		me = &TargetRevision {
			HostName:     "",
			RefName:     "",
		}
	}

	return *me
}

func (me *TargetRevisions) New() TargetRevisions {
	if me == nil {
		me = &TargetRevisions {}
	}

	return *me
}

package deployjson

type Revision struct {
	RefName string `json:"ref_name"`
	RefType string `json:"ref_type"`
}


func (me *Revision) New() Revision {
	if me == nil {
		me = &Revision {
			RefName: "",
			RefType: "",
		}
	}

	return *me
}

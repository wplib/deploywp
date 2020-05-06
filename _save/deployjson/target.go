package deployjson

import (
)

type Target struct {
	Files     Files           `json:"files"`
	Paths     Paths           `json:"paths"`
	Providers Providers       `json:"providers"`
	Revisions TargetRevisions `json:"revisions"`
}

func (me *Target) New() Target {
	if me == nil {
		me = &Target {
			Files:     me.Files.New(),
			Paths:     me.Paths.New(),
			Providers: me.Providers.New(),
			Revisions: me.Revisions.New(),
		}
	}

	return *me
}


package jsonfile

type DeployWP struct {
	Schema     Version    `json:"schema"`
	Repository Repository `json:"repository"`
	Tag        Reference  `json:"tag"`
	Branch     Reference  `json:"branch"`
	Commit     Reference  `json:"commit"`
}

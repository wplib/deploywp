package jsonfile

type Targets struct {
	Defaults *Host `json:"defaults"`
	Hosts    Hosts `json:"hosts"`
}

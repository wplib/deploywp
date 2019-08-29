package deploywp

type Targets struct {
	Defaults *Host
	Hosts    Hosts
}

type TargetsGetter interface {
	GetDefaults() *Host
	GetHosts() Hosts
}

func NewTargetsFromGetter(tg TargetsGetter) (t *Targets) {
	return &Targets{
		Defaults: tg.GetDefaults(),
		Hosts:    tg.GetHosts(),
	}
}

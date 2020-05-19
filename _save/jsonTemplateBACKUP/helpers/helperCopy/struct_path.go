package helperCopy

import "github.com/wplib/deploywp/only"

type PathArray []string

func (p *PathArray) SetPaths(paths ...string) bool {
	var ok bool

	for range only.Once {
		*p = paths
		if len(*p) == 0 {
			break
		}

		ok = true
	}

	return ok
}
func (p *PathArray) AddPaths(paths ...string) bool {
	var ok bool

	for range only.Once {
		*p = append(*p, paths...)
		if len(*p) == 0 {
			break
		}
		ok = true
	}

	return ok
}
func (p *PathArray) GetPaths() *PathArray {
	return p
}

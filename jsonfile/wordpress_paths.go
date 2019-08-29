package jsonfile

import (
	"github.com/wplib/deploywp/deploywp"
)

var _ deploywp.WordPressPathsGetter = (*WordPressPaths)(nil)

type WordPressPaths struct {
	RootPath    Path `json:"root_path"`
	CorePath    Path `json:"core_path"`
	ContentPath Path `json:"content_path"`
	VendorPath  Path `json:"vendor_path"`
}

func (me WordPressPaths) GetRootPath() deploywp.Path {
	return me.RootPath
}
func (me WordPressPaths) GetCorePath() deploywp.Path {
	return me.CorePath
}
func (me WordPressPaths) GetContentPath() deploywp.Path {
	return me.ContentPath
}
func (me WordPressPaths) GetVendorPath() deploywp.Path {
	return me.VendorPath
}

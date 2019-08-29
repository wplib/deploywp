package deploywp

type WordPressPaths struct {
	RootPath    Path
	CorePath    Path
	ContentPath Path
	VendorPath  Path
}

type WordPressPathsGetter interface {
	GetRootPath() Path
	GetCorePath() Path
	GetContentPath() Path
	GetVendorPath() Path
}

func NewWordPressPathsFromGetter(wpg WordPressPathsGetter) (wpp *WordPressPaths) {
	return &WordPressPaths{
		RootPath:    wpg.GetRootPath(),
		CorePath:    wpg.GetCorePath(),
		ContentPath: wpg.GetContentPath(),
		VendorPath:  wpg.GetVendorPath(),
	}
}

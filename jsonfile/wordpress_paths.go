package jsonfile

type WordPressPaths struct {
	RootPath    Path       `json:"root_path"`
	CorePath    Path       `json:"core_path"`
	ContentPath Path       `json:"content_path"`
	VendorPath  Path       `json:"vendor_path"`
}

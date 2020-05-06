package deployjson

type Paths struct {
	WebrootPath string `json:"webroot_path"`
	Wordpress Wordpress `json:"wordpress"`
}

type Wordpress struct {
	ContentPath string `json:"content_path"`
	CorePath    string `json:"core_path"`
	RootPath    string `json:"root_path"`
	VendorPath  string `json:"vendor_path"`
}


//var _ deploywp.PathsGetter = (*Paths)(nil)

func (me *Paths) New() Paths {
	if me == nil {
		me = &Paths {
			WebrootPath: "",
			Wordpress:   me.Wordpress.New(),
		}
	}

	return *me
}

func (me *Wordpress) New() Wordpress {
	if me == nil {
		me = &Wordpress {
			ContentPath: "",
			CorePath:    "",
			RootPath:    "",
			VendorPath:  "",
		}
	}

	return *me
}

func (me Paths) GetWebRoot() Path {
	return me.WebrootPath
}


//func (me Paths) GetPaths() *deploywp.WordPressPaths {
//	return deploywp.NewWordPressPathsFromGetter(me.Wordpress)
//}
//
//func (me Paths) GetContentPath() *deploywp.FileDispositions {
//	return deploywp.NewFileDispositionsFromGetter(me.Wordpress.ContentPath)
//}
//
//func (me Paths) GetRepository() *deploywp.Repository {
//	return deploywp.NewRepositoryFromGetter(me.Repository)
//}

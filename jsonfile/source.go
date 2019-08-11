package jsonfile

type Source struct {
	WebRoot    Path             `json:"web_root"`
	Paths      WordPressPaths   `json:"wp_paths"`
	Files      FileDispositions `json:"files"`
	Repository Repository       `json:"repository"`
}

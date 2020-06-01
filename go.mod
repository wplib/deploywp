module github.com/wplib/deploywp

go 1.12

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/gdamore/tcell v1.3.0
	github.com/gizak/termui/v3 v3.1.0
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/go-github/v31 v31.0.0
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/logrusorgru/aurora v0.0.0-20200102142835-e9ef32dff381
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/newclarity/JsonToConfig/jtc v0.0.0-00010101000000-000000000000
	github.com/newclarity/JsonToConfig/ux v0.0.0-00010101000000-000000000000
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/princjef/gomarkdoc v0.1.1 // indirect
	github.com/rivo/tview v0.0.0-20200528200248-fe953220389f
	github.com/robertkrimen/godocdown v0.0.0-20130622164427-0bfa04905481 // indirect
	github.com/shirou/gopsutil v2.20.4+incompatible
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df
	github.com/ungerik/pkgreflect v0.0.0-20170905122726-bfeb2a931863 // indirect
	github.com/zloylos/grsync v0.0.0-20200204095520-71a00a7141be
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/tools v0.0.0-20200519015757-0d0afa43d58a // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
)

replace github.com/wplib/deploywp => ./

// replace github.com/wplib/deploywp/jtc => ../JsonToConfig/jtc
// replace github.com/wplib/deploywp/ux => ../JsonToConfig/ux
replace github.com/newclarity/JsonToConfig/jtc => ../JsonToConfig/jtc

replace github.com/newclarity/JsonToConfig/ux => ../JsonToConfig/ux

replace github.com/newclarity/JsonToConfig/defaults => ./defaults

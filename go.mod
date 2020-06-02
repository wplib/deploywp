module github.com/wplib/deploywp

go 1.14

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/newclarity/scribeHelpers/helperCopy v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperExec v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperGit v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperGitHub v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperPath v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperPrompt v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperService v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperSystem v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperUx v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/scribeLoader v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
)

replace github.com/newclarity/scribeHelpers/ux => ../scribeHelpers/ux

replace github.com/newclarity/scribeHelpers/scribeLoader => ../scribeHelpers/scribeLoader

replace github.com/newclarity/scribeHelpers/helperCopy => ../scribeHelpers/helperCopy

replace github.com/newclarity/scribeHelpers/helperExec => ../scribeHelpers/helperExec

replace github.com/newclarity/scribeHelpers/helperGit => ../scribeHelpers/helperGit

replace github.com/newclarity/scribeHelpers/helperGitHub => ../scribeHelpers/helperGitHub

replace github.com/newclarity/scribeHelpers/helperPath => ../scribeHelpers/helperPath

replace github.com/newclarity/scribeHelpers/helperPrompt => ../scribeHelpers/helperPrompt

replace github.com/newclarity/scribeHelpers/helperService => ../scribeHelpers/helperService

replace github.com/newclarity/scribeHelpers/helperSystem => ../scribeHelpers/helperSystem

replace github.com/newclarity/scribeHelpers/helperTypes => ../scribeHelpers/helperTypes

replace github.com/newclarity/scribeHelpers/helperUx => ../scribeHelpers/helperUx

replace github.com/newclarity/scribeHelpers/helperRuntime => ../scribeHelpers/helperRuntime

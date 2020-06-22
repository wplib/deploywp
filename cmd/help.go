package cmd

import (
	"flag"
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
	"strings"
)


func _GetUsage(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str += ux.SprintfCyan("%s [flags] ", c.Name())
	} else {
		str += ux.SprintfCyan("%s [flags] ", c.Parent().Name())
		str += ux.SprintfGreen("%s ", c.Use)
	}

	if c.HasAvailableSubCommands() {
		str += ux.SprintfGreen("[command] ")
		str += ux.SprintfCyan("<args> ")
	}

	return str
}


func _GetVersion(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str = ux.SprintfBlue("%s ", CmdScribe.Runtime.CmdName)
		str += ux.SprintfCyan("v%s", CmdScribe.Runtime.CmdVersion)
	}

	return str
}


func SetHelp(c *cobra.Command) {
	var tmplHelp string
	var tmplUsage string

	//fmt.Printf("%s", rootCmd.UsageTemplate())
	//fmt.Printf("%s", rootCmd.HelpTemplate())

	cobra.AddTemplateFunc("GetUsage", _GetUsage)
	cobra.AddTemplateFunc("GetVersion", _GetVersion)

	cobra.AddTemplateFunc("SprintfBlue", ux.SprintfBlue)
	cobra.AddTemplateFunc("SprintfCyan", ux.SprintfCyan)
	cobra.AddTemplateFunc("SprintfGreen", ux.SprintfGreen)
	cobra.AddTemplateFunc("SprintfMagenta", ux.SprintfMagenta)
	cobra.AddTemplateFunc("SprintfRed", ux.SprintfRed)
	cobra.AddTemplateFunc("SprintfWhite", ux.SprintfWhite)
	cobra.AddTemplateFunc("SprintfYellow", ux.SprintfYellow)

	// 	{{ with .Parent }}{{ SprintfCyan .Name }}{{ end }} {{ SprintfGreen .Name }} {{ if .HasAvailableSubCommands }}{{ SprintfGreen "[command]" }}{{ end }}

	tmplUsage += `
{{ SprintfBlue "Usage: " }}
	{{ GetUsage . }}

{{- if gt (len .Aliases) 0 }}
{{ SprintfBlue "\nAliases:" }} {{ .NameAndAliases }}
{{- end }}

{{- if .HasExample }}
{{ SprintfBlue "\nExamples:" }}
	{{ .Example }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nWhere " }}{{ SprintfGreen "[command]" }}{{ SprintfBlue " is one of:" }}
{{- range .Commands }}
{{- if (or .IsAvailableCommand (eq .Name "help")) }}
	{{ rpad (SprintfGreen .Name) .NamePadding}}	- {{ .Short }}{{ end }}
{{- end }}
{{- end }}

{{- if .HasAvailableLocalFlags }}
{{ SprintfBlue "\nFlags:" }}
{{ .LocalFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasAvailableInheritedFlags }}
{{ SprintfBlue "\nGlobal Flags:" }}
{{ .InheritedFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasHelpSubCommands }}
{{- SprintfBlue "\nAdditional help topics:" }}
{{- range .Commands }}
{{- if .IsAdditionalHelpTopicCommand }}
	{{ rpad (SprintfGreen .CommandPath) .CommandPathPadding }} {{ .Short }}
{{- end }}
{{- end }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nUse" }} {{ SprintfCyan .CommandPath }} {{ SprintfCyan "help" }} {{ SprintfGreen "[command]" }} {{ SprintfBlue "for more information about a command." }}
{{- end }}
`

	tmplHelp = `{{ GetVersion . }}

{{ SprintfBlue "Commmand:" }} {{ SprintfCyan .Use }}

{{ SprintfBlue "Description:" }} 
	{{ with (or .Long .Short) }}
{{- . | trimTrailingWhitespaces }}
{{- end }}

{{- if or .Runnable .HasSubCommands }}
{{ .UsageString }}
{{- end }}
`

	//c.SetHelpCommand(c)
	//c.SetHelpFunc(PrintHelp)
	c.SetHelpTemplate(tmplHelp)
	c.SetUsageTemplate(tmplUsage)
}


func HelpAll() {
	for range onlyOnce {
		HelpFunctions()
		HelpVariables()
		HelpExamples()
	}

	CmdScribe.State.Clear()
}


func Help() {
	for range onlyOnce {
		ux.PrintflnBlue("%s v%s:", defaults.BinaryName, defaults.BinaryVersion)
		ux.PrintflnBlue("\tDeployWp.")
		ux.PrintflnBlue("\tBuilding Pantheon based websites.")
		ux.PrintflnBlue("")

		flag.PrintDefaults()
		ux.PrintflnBlue("")
		ux.PrintflnBlue("")
	}

	CmdScribe.State.Clear()
}


func HelpFunctions() {
	CmdScribe.State.Clear()
	return

	for range onlyOnce {
		CmdScribe.PrintTools()
	}

	CmdScribe.State.Clear()
}


func HelpVariables() {
	CmdScribe.State.Clear()
	return

	for range onlyOnce {
		ux.PrintfBlue("Keys accessible within your template file:\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Json }}"), ux.SprintfWhite("Your JSON file will appear here."))
		fmt.Printf("\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Env }}"), ux.SprintfWhite("A map containing runtime environment variables."))
		fmt.Printf("\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.CmdName }}"), ux.SprintfWhite("Executable, (this program), used to produce the resulting file."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.CmdSelfUpdate }}"), ux.SprintfWhite("Version of this executable."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.CmdScribe }}"), ux.SprintfWhite("ARG[0] which should be the same as CmdName."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.CmdDir }}"), ux.SprintfWhite("The absolute directory where this executable was run from."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.CmdFile }}"), ux.SprintfWhite("The filename of this executable, (should be the same as CmdName)."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.Env }}"), ux.SprintfWhite("An array containing runtime environment variables."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.EnvMap }}"), ux.SprintfWhite("A map containing runtime environment variables."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .Exec.TimeStamp }}"), ux.SprintfWhite("The current timestamp as execution time."))
		fmt.Printf("\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .CreationDate }}"), ux.SprintfWhite("Creation date of resulting file."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .CreationEpoch }}"), ux.SprintfWhite("Creation date, (unix epoch), of resulting file."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .CreationInfo }}"), ux.SprintfWhite("More creation information."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .CreationWarning }}"), ux.SprintfWhite("Generic 'DO NOT EDIT' warning."))
		fmt.Printf("\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .TemplateFile.Dir }}"), ux.SprintfWhite("template file absolute directory."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .TemplateFile.Name }}"), ux.SprintfWhite("template filename."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .TemplateFile.CreationDate }}"), ux.SprintfWhite("template file creation date."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .TemplateFile.CreationEpoch }}"), ux.SprintfWhite("template file creation date, (unix epoch)."))
		fmt.Printf("\n")
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .JsonFile.Dir }}"), ux.SprintfWhite("json file absolute directory."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .JsonFile.Name }}"), ux.SprintfWhite("json filename."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .JsonFile.CreationDate }}"), ux.SprintfWhite("json file creation date."))
		fmt.Printf("%s - %s\n", ux.SprintfBlue("\t{{ .JsonFile.CreationEpoch }}"), ux.SprintfWhite("json file creation date, (unix epoch)."))
		fmt.Printf("\n")
	}

	CmdScribe.State.Clear()
}


type Example struct {
	Command string
	Args []string
	Info string
}
type Examples []Example


func HelpExamples() {
	CmdScribe.State.Clear()
	return

	for range onlyOnce {
		var examples Examples


		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-json", "config.json", "-template", "'{{ .Json.dir }}'"},
			Info:    "Print to STDOUT the .dir key from config.json.",
		})
		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-template", "'{{ .Json.dir }}'", "config.json"},
			Info:    "The same thing, but with less arguments.",
		})

		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-template", "'{{ .Json.hello }}'", "-json", `'{ "hello": "world" }'`},
			Info:    "Template and JSON arguments can be either string or file reference.",
		})
		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-template", "hello_world.tmpl", "-json", `'{ "hello": "world" }'`},
			Info:    "The same again...",
		})
		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-template", "'{{ .Json.hello }}'", "-json", "hello.json"},
			Info:    "The same again...",
		})


		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-json", "config.json", "-template", "DockerFile.tmpl", "-out", "Dockerfile"},
			Info:    "Process Dockerfile.tmpl file and output to Dockerfile.",
		})
		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-out", "Dockerfile", "config.json", "DockerFile.tmpl"},
			Info:    "And again with less arguments..",
		})
		examples = append(examples, Example {
			Command: "convert",
			Args:    []string{"config.json", "DockerFile.tmpl"},
			Info:    "'convert' does the same , but removes the template file afterwards...",
		})


		examples = append(examples, Example {
			Command: "load",
			Args:    []string{"-out", "MyScript.sh", "MyScript.sh.tmpl", "config.json"},
			Info:    "Process the MyScript.sh.tmpl template file and write the result to MyScript.sh.",
		})
		examples = append(examples, Example {
			Command: "convert",
			Args:    []string{"MyScript.sh.tmpl", "config.json"},
			Info:    "Same again using 'convert'. Template and json files can be in any order.",
		})
		examples = append(examples, Example {
			Command: "run",
			Args:    []string{"MyScript.sh.tmpl", "config.json"},
			Info:    "Same again using 'run'. This will execute the MyScript.sh output file afterwards.",
		})


		ux.PrintflnBlue("Examples:")
		for _, v := range examples {
			fmt.Printf("# %s\n\t%s %s\n\n",
				ux.SprintfBlue(v.Info),
				ux.SprintfCyan("%s %s", defaults.BinaryName, v.Command),
				ux.SprintfWhite(strings.Join(v.Args, " ")),
			)
		}
	}

	CmdScribe.State.Clear()
}
